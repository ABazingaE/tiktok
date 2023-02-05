package main

import (
	"context"
	"tiktok/cmd/user/pack"
	"tiktok/cmd/user/service"
	user "tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	//  Your code here...
	resp = new(user.RegisterResp)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	var userid int64
	_ = userid
	userid, err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = userid
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	//  Your code here...
	resp = new(user.LoginResp)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	resp = new(user.UserInfoResp)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	resp, err = service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
