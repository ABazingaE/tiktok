package db

import (
	"context"
	"gorm.io/gorm"
)

type Follow struct {
	Id             int `gorm:"column:id" json:"id"`                             //type:*int   comment:    version:2023-01-09 19:53
	FollowedUserId int `gorm:"column:followed_user_id" json:"followed_user_id"` //type:*int   comment:    version:2023-01-09 19:53
	FollowerId     int `gorm:"column:follower_id" json:"follower_id"`           //type:*int   comment:    version:2023-01-09 19:53
}

// TableName 表名:follow，。
// 说明:
func (f *Follow) TableName() string {
	return "follow"
}

type LatestMessage struct {
	Id         int    `gorm:"column:id" json:"id"`                   //type:*int     comment:                                                             version:2023-01-10 18:26
	UserId     int    `gorm:"column:user_id" json:"user_id"`         //type:*int     comment:                                                             version:2023-01-10 18:26
	FollowerId int    `gorm:"column:follower_id" json:"follower_id"` //type:*int     comment:                                                             version:2023-01-10 18:26
	Message    string `gorm:"column:message" json:"message"`         //type:string   comment:                                                             version:2023-01-10 18:26
	MsgType    int    `gorm:"column:msg_type" json:"msg_type"`       //type:*int     comment:0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息    version:2023-01-10 18:26
}

// TableName 表名:latest_message，。
// 说明:
func (lm *LatestMessage) TableName() string {
	return "latest_message"
}

type Message struct {
	Id         int    `gorm:"column:id" json:"id"`                     //type:*int     comment:          version:2023-01-10 18:36
	FromUserId int    `gorm:"column:from_user_id" json:"from_user_id"` //type:*int     comment:          version:2023-01-10 18:36
	ToUserId   int    `gorm:"column:to_user_id" json:"to_user_id"`     //type:*int     comment:          version:2023-01-10 18:36
	Content    string `gorm:"column:content" json:"content"`           //type:string   comment:          version:2023-01-10 18:36
	CreateTime int    `gorm:"column:create_time" json:"create_time"`   //type:*int     comment:时间戳    version:2023-01-10 18:36
}

// TableName 表名:message，。
// 说明:
func (m *Message) TableName() string {
	return "message"
}

/*
 好友列表：
	好友的定义：双方互相关注，即为好友，返回好友列表及最新的消息
*/

/*
查询好友，传入用户id，此id关注用户中同样也关注了该用户的用户id列表
*/
func GetFriendIdList(ctx context.Context, userId int) ([]int, error) {
	var followedIdList []int
	var followerIdList []int
	//查询登录用户的粉丝列表
	err := DB.WithContext(ctx).Table("follow").Where("followed_user_id = ?", userId).Pluck("follower_id", &followerIdList).Error
	if err != nil {
		return nil, err
	}

	//查询登录用户的关注列表
	err = DB.WithContext(ctx).Table("follow").Where("follower_id = ?", userId).Pluck("followed_user_id", &followedIdList).Error
	if err != nil {
		return nil, err
	}
	var friendIdList []int
	//比较两个列表，取交集，依次放入map中，遇到重复的key即为好友
	friendMap := make(map[int]int)
	for _, id := range followerIdList {
		friendMap[id] = id
	}
	for _, id := range followedIdList {
		if _, ok := friendMap[id]; ok {
			friendIdList = append(friendIdList, id)
		}
	}
	return friendIdList, nil
}

/*
查询最新消息，返回消息以及消息类型
*/
func GetLatestMessage(ctx context.Context, userId int, friendId int) (message string, msgType int, error error) {
	var latestMessage LatestMessage
	//不确定latest_message表中的user_id是谁
	err := DB.WithContext(ctx).Table("latest_message").Where("user_id = ? and follower_id = ?", userId, friendId).First(&latestMessage).Error
	if err == nil {
		//若查询到，msgType代表的即为当前请求用户接收的消息的状态
		return latestMessage.Message, latestMessage.MsgType, nil
	} else if err == gorm.ErrRecordNotFound {
		err = DB.WithContext(ctx).Table("latest_message").Where("user_id = ? and follower_id = ?", friendId, userId).First(&latestMessage).Error
		if err == nil {
			//颠倒msgType的值，代表当前请求用户发送的消息的状态（msgTyoe代表user_id接收的消息状态）
			latestMessage.MsgType = 1 - latestMessage.MsgType
			return latestMessage.Message, latestMessage.MsgType, nil
		} else {
			return "", 0, err
		}
	}
	return "", 0, err
}
