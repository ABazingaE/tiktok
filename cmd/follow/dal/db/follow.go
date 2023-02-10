package db

import (
	"context"
	"gorm.io/gorm"
	"tiktok/pkg/consts"
)

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

type UserInfo struct {
	UserId        int    `gorm:"column:user_id" json:"user_id"`               //type:*int     comment:id      version:2023-01-03 20:08
	Name          string `gorm:"column:name" json:"name"`                     //type:string   comment:name    version:2023-01-03 20:08
	FollowCount   int    `gorm:"column:follow_count" json:"follow_count"`     //type:*int     comment:        version:2023-01-03 20:08
	FollowerCount int    `gorm:"column:follower_count" json:"follower_count"` //type:*int     comment:        version:2023-01-03 20:08
	Avatar        string `gorm:"column:avatar" json:"avatar"`                 //type:string   comment:        version:2023-01-03 20:08
}

// TableName 表名:user_info，user_info。
// 说明:
func (ui *UserInfo) TableName() string {
	return consts.UserInfoTableName
}

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

// 添加关注
func AddFollow(ctx context.Context, followed_user_id int, follower_id int) (*Follow, error) {
	follow := &Follow{
		FollowedUserId: followed_user_id,
		FollowerId:     follower_id,
	}
	if err := DB.WithContext(ctx).Create(follow).Error; err != nil {
		return nil, err
	}
	// 更新关注数,粉丝数
	if err := DB.WithContext(ctx).Model(&UserInfo{}).Where("user_id = ?", followed_user_id).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		return nil, err

	}
	if err := DB.WithContext(ctx).Model(&UserInfo{}).Where("user_id = ?", follower_id).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		return nil, err
	}

	//判断是否成为好友关系，如果是则更新最新消息表，插入一条“你们已成为好友”的信息
	if err := DB.WithContext(ctx).Model(&Follow{}).Where("followed_user_id = ? AND follower_id = ?", follower_id, followed_user_id).First(&Follow{}).Error; err == nil {
		//添加一条最新消息
		if err := DB.WithContext(ctx).Create(&LatestMessage{
			UserId:     followed_user_id,
			FollowerId: follower_id,
			Message:    "你们已成为好友",
			MsgType:    0,
		}).Error; err != nil {
			return nil, err
		}
	}
	return follow, nil
}

// 删除关注
func DeleteFollow(ctx context.Context, followed_user_id int, follower_id int) error {
	//判断二者是否为好友关系，若是则删除最新消息表中记录二者最新消息的数据
	if err := DB.WithContext(ctx).Model(&Follow{}).Where("followed_user_id = ? AND follower_id = ?", follower_id, followed_user_id).First(&Follow{}).Error; err == nil {
		//删除最新消息表中记录二者最新消息的数据
		//不确定latest_message表中的user_id和follower_id具体是哪个，所以先查询出来，再删除
		var latestMessage LatestMessage
		if err := DB.WithContext(ctx).Model(&LatestMessage{}).Where("user_id = ? AND follower_id = ?", followed_user_id, follower_id).First(&latestMessage).Error; err != nil {
			//若err不为空，则说明user_id和follower_id的值是反的，再次查询
			if err := DB.WithContext(ctx).Model(&LatestMessage{}).Where("user_id = ? AND follower_id = ?", follower_id, followed_user_id).First(&latestMessage).Error; err != nil {
				return err
			}
			//查询成功后删除
			if err := DB.WithContext(ctx).Where("id = ?", latestMessage.Id).Delete(&LatestMessage{}).Error; err != nil {
				return err
			}
		}
	}

	//删除数据
	if err := DB.WithContext(ctx).Where("followed_user_id = ? AND follower_id = ?", followed_user_id, follower_id).Delete(&Follow{}).Error; err != nil {
		return err
	}
	// 更新关注数,粉丝数
	if err := DB.WithContext(ctx).Model(&UserInfo{}).Where("user_id = ?", followed_user_id).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		return err
	}
	if err := DB.WithContext(ctx).Model(&UserInfo{}).Where("user_id = ?", follower_id).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户id查询此id关注的用户id列表
func GetFollowedUserIdList(ctx context.Context, follower_id int) ([]int, error) {
	var followedUserIdList []int
	if err := DB.WithContext(ctx).Model(&Follow{}).Where("follower_id = ?", follower_id).Pluck("followed_user_id", &followedUserIdList).Error; err != nil {
		return nil, err
	}
	return followedUserIdList, nil
}

// 根据用户id查询关注此id的用户id列表
func GetFollowerIdList(ctx context.Context, followed_user_id int) ([]int, error) {
	var followerIdList []int
	if err := DB.WithContext(ctx).Model(&Follow{}).Where("followed_user_id = ?", followed_user_id).Pluck("follower_id", &followerIdList).Error; err != nil {
		return nil, err
	}
	return followerIdList, nil
}
