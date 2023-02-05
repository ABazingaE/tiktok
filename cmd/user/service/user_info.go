package service

import (
	"context"
	"strconv"
	"tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/user"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService NewCheckUserService new CheckUserService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
	}
}

// UserInfo CheckUser check user info
func (s *UserInfoService) UserInfo(req *user.UserInfoReq) (*user.UserInfoResp, error) {
	resp := &user.UserInfoResp{}
	/*
		request的token已解析为请求者id，需根据此id与userId判断二者的关注关系，即请求者id是否关注了userId，最后得到is_follow字段值
	*/
	var isFollow = false
	if req.Token != "" {

		requestId, err := strconv.ParseInt(req.Token, 10, 64)
		if err != nil {
			return nil, err
		}
		isFollow, err = db.IsFollow(s.ctx, req.UserId, requestId)
		if err != nil {
			return nil, err
		}
	}

	baseInfo, errno := db.QueryUserInfoById(s.ctx, req.UserId)
	if errno != nil {
		return nil, errno
	}

	userInfo := &user.User{}
	userInfo.Id = int64(baseInfo.UserId)
	userInfo.Name = baseInfo.Name
	userInfo.FollowCount = int64(baseInfo.FollowCount)
	userInfo.FollowerCount = int64(baseInfo.FollowerCount)
	userInfo.IsFollow = isFollow

	resp.User = userInfo

	return resp, nil
}
