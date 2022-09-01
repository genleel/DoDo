package model

import (
	"fmt"
	"followService/config"
	pb "followService/proto"
	"github.com/gogf/gf/util/gconv"
	"log"
	"strconv"
	"strings"
	"sync"
	userModel "userService/model"
)

// FindRelation 给定当前用户和目标用户id，查询follow表中相应的记录。
func FindRelation(userId int64, targetId int64) (*Follow, error) {
	// follow变量用于后续存储数据库查出来的用户关系。
	follow := Follow{}
	//当查询出现错误时，日志打印err msg，并return err.
	if err := Db.
		Where("user_id = ?", targetId).
		Where("follower_id = ?", userId).
		Where("cancel = ?", 0).
		Take(&follow).Error; nil != err {
		// 当没查到数据时，gorm也会报错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	//正常情况，返回取到的值和空err.
	return &follow, nil
}

// GetFollowersIds 给定用户id，查询他关注了哪些人的id。
func GetFollowersIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := Db.Model(Follow{}).
		Where("user_id = ?", userId).
		Where("cancel = ?", 0).
		Pluck("follower_id", &ids).Error; nil != err {
		// 没有粉丝，但是不能算错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		// 查询出错。
		log.Println(err.Error())
		return nil, err
	}
	// 查询成功。
	return ids, nil
}

// GetFollowingIds 给定用户id，查询他关注了哪些人的id。
func GetFollowingIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := Db.Model(Follow{}).
		Where("follower_id = ?", userId).
		Pluck("user_id", &ids).Error; nil != err {
		// 没有关注任何人，但是不能算错。
		if "record not found" == err.Error() {
			return nil, nil
		}
		// 查询出错。
		log.Println(err.Error())
		return nil, err
	}
	// 查询成功。
	return ids, nil
}

func IsFollowing(userId int64, targetId int64) (bool, error) {
	// 先查Redis里面是否有此关系。
	if flag, _ := RdbFollowingPart.SIsMember(Ctx, strconv.Itoa(int(userId)), targetId).Result(); flag {
		// 重现设置过期时间。
		RdbFollowingPart.Expire(Ctx, strconv.Itoa(int(userId)), config.ExpireTime)
		return true, nil
	}
	// SQL 查询。
	relation, err := FindRelation(userId, targetId)

	if nil != err {
		return false, err
	}
	if nil == relation {
		return false, nil
	}
	// 存在此关系，将其注入Redis中。
	go addRelationToRedis(int(userId), int(targetId))
	return true, nil
}

func GetFollowerCnt(userId int64) (int64, error) {
	// 查Redis中是否已经存在。
	if cnt, _ := RdbFollowers.SCard(Ctx, strconv.Itoa(int(userId))).Result(); cnt > 0 {
		// 更新过期时间。
		RdbFollowers.Expire(Ctx, strconv.Itoa(int(userId)), config.ExpireTime)
		return cnt - 1, nil
	}
	// SQL中查询。
	ids, err := GetFollowersIds(userId)
	if nil != err {
		return 0, err
	}
	// 将数据存入Redis.
	// 更新followers 和 followingPart
	go addFollowersToRedis(int(userId), ids)

	return int64(len(ids)), nil
}

func GetFollowingCnt(userId int64) (int64, error) {
	// 查看Redis中是否有关注数。
	if cnt, _ := RdbFollowing.SCard(Ctx, strconv.Itoa(int(userId))).Result(); cnt > 0 {
		// 更新过期时间。
		RdbFollowing.Expire(Ctx, strconv.Itoa(int(userId)), config.ExpireTime)
		return cnt - 1, nil

	}
	// 用SQL查询。
	ids, err := GetFollowingIds(userId)

	if nil != err {
		return 0, err
	}
	// 更新Redis中的followers和followPart
	go addFollowingToRedis(int(userId), ids)

	return int64(len(ids)), nil
}

func AddFollowRelation(userId int64, targetId int64) (bool, error) {
	// 加信息打入消息队列。
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(userId)))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(int(targetId)))
	RmqFollowAdd.Publish(sb.String())
	// 记录日志
	// 更新redis信息。
	updateRedisWithAdd(userId, targetId)

	return true, nil
}

func DeleteFollowRelation(userId int64, targetId int64) (bool, error) {
	// 加信息打入消息队列。
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(int(userId)))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(int(targetId)))
	RmqFollowDel.Publish(sb.String())
	// 记录日志
	// 更新redis信息。
	updateRedisWithDelete(userId, targetId)

	return true, nil
}

func GetFollowing(userId int64) ([]*pb.FeedUser, error) {
	// 获取关注对象的id数组。
	ids, err := GetFollowingIds(userId)
	// 查询出错
	if nil != err {
		return nil, err
	}
	// 没得关注者
	if nil == ids {
		return nil, nil
	}
	// 根据每个id来查询用户信息。
	followingNum := len(ids)

	users := make([]pb.FeedUser, followingNum)
	for i := 0; i < followingNum; i++ {

		user, err := userModel.GetFeedUserByIdWithCurId(ids[i], userId)
		if err != nil {
			fmt.Println("userModel.GetFeedUserByIdWithCurId err:", err)
		}

		var tmpUser *pb.FeedUser
		gconv.Struct(user, &tmpUser)

		users[i] = *tmpUser

	}
	// 返回关注对象列表
	var followUser []*pb.FeedUser
	gconv.Struct(users, &followUser)

	return followUser, nil
}

func GetFollowers(userId int64) ([]*pb.FeedUser, error) {
	// 获取粉丝的id数组。
	ids, err := GetFollowersIds(userId)
	// 查询出错
	if nil != err {
		return nil, err
	}
	// 没得粉丝
	if nil == ids {
		return nil, nil
	}
	// 根据每个id来查询用户信息。
	len := len(ids)
	if len > 0 {
		len -= 1
	}
	users := make([]userModel.FeedUser, len)
	var wg sync.WaitGroup
	wg.Add(len)
	i, j := 0, 0
	for ; i < len; j++ {
		// 越过-1
		if ids[j] == -1 {
			continue
		}
		//开启协程来查。
		go func(i int, idx int64) {
			defer wg.Done()

			// 调用微服务的方法
			user, _ := userModel.GetFeedUserByIdWithCurId(idx, userId)
			var tmpUser userModel.FeedUser
			gconv.Struct(user, &tmpUser)

			users[i] = tmpUser

		}(i, ids[i])
		i++
	}
	wg.Wait()
	// 返回粉丝列表。

	var followerUser []*pb.FeedUser
	gconv.Struct(users, &followerUser)

	return followerUser, nil
}
