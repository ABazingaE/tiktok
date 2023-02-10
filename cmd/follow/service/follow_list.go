package service

import (
	"context"
	"tiktok/cmd/follow/dal/db"
	"tiktok/cmd/follow/rpc"
	"tiktok/kitex_gen/follow"
	"tiktok/kitex_gen/user"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

func (s *FollowListService) FollowList(req *follow.FollowListReq) (resp *follow.FollowListResp, err error) {
	userId := int(req.UserId)
	resp = &follow.FollowListResp{}
	followedId, err := db.GetFollowedUserIdList(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	var followedUser []*follow.User
	for _, id := range followedId {
		resp, err := rpc.GetUserInfo(s.ctx, &user.UserInfoReq{
			UserId: int64(id),
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		followedUser = append(followedUser, (*follow.User)(resp.User))
	}
	resp.UserList = followedUser
	return resp, nil
}
