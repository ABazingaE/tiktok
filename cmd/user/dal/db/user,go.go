package db

import (
	"context"
	"tiktok/pkg/consts"
)

type User struct {
	UserId       int    `gorm:"column:user_id" json:"user_id"`        //type:*int     comment:ID          version:2023-01-02 15:43
	UserName     string `gorm:"column:user_name" json:"username"`     //type:string   comment:用户名      version:2023-01-02 15:43
	UserPassword string `gorm:"column:user_password" json:"password"` //type:string   comment:用户密码    version:2023-01-02 15:43
}

// TableName 表名:tiktok_user，tiktok_user。
// 说明:
func (u *User) TableName() string {
	return consts.UserTableName
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

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	/*先插入user表，再查询id，再将id和name插入到info表*/

	//插入user表
	if err := DB.WithContext(ctx).Create(users).Error; err != nil {
		return err
	}

	//查询id
	var id int
	var name string = users[0].UserName
	res := make([]*User, 0)
	res, _ = QueryUser(ctx, name)
	id = res[0].UserId

	//插入info表
	return DB.WithContext(ctx).Create([]*UserInfo{
		{
			UserId: id,
			Name:   name,
		},
	}).Error
	//return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserInfoById(ctx context.Context, userId int64) (*UserInfo, error) {
	res := &UserInfo{}
	if err := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// IsFollow check if user follow another user
func IsFollow(ctx context.Context, userId int64, requestId int64) (bool, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&Follow{}).Where("followed_user_id = ? and follower_id = ?", userId, requestId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
