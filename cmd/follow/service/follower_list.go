package service

import (
	"context"
	"tiktok/cmd/follow/dal/db"
	"tiktok/cmd/follow/rpc"
	"tiktok/kitex_gen/follow"
	"tiktok/kitex_gen/user"
)

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

func (s *FollowerListService) FollowerList(req *follow.FollowerListReq) (resp *follow.FollowerListResp, err error) {
	userId := int(req.UserId)
	followerId, err := db.GetFollowerIdList(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	var followerUser []*follow.User
	for _, id := range followerId {
		resp, err := rpc.GetUserInfo(s.ctx, &user.UserInfoReq{
			UserId: int64(id),
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		followerUser = append(followerUser, (*follow.User)(resp.User))
	}
	resp = &follow.FollowerListResp{
		UserList: followerUser,
	}
	return resp, nil
}
