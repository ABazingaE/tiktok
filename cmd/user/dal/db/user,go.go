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
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
