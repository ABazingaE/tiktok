package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.RegisterReq) (id int64, error error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if len(users) != 0 {
		return -1, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return -1, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	err = db.CreateUser(s.ctx, []*db.User{{
		UserName:     req.Username,
		UserPassword: password,
	}})
	if err != nil {
		return -1, err
	}

	//查询刚刚注册用户的id
	users, err = db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return -1, err
	}
	if len(users) == 0 {
		return -1, errno.UserNotExistErr
	}
	return int64(users[0].UserId), nil

}
