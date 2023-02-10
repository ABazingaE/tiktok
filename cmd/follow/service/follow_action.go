package service

import (
	"context"
	"strconv"
	"tiktok/cmd/follow/dal/db"
	"tiktok/kitex_gen/follow"
)

type FollowActionService struct {
	ctx context.Context
}

func NewFollowActionService(ctx context.Context) *FollowActionService {
	return &FollowActionService{ctx: ctx}
}

func (s *FollowActionService) FollowAction(req *follow.FollowActionReq) (resp *follow.FollowActionResp, err error) {
	actionType := int(req.ActionType)
	followerId, err := strconv.Atoi(req.Token)
	if err != nil {
		return nil, err
	}
	resp = &follow.FollowActionResp{}
	switch actionType {
	case 1:
		db.AddFollow(s.ctx, int(req.ToUserId), followerId)

	case 2:
		db.DeleteFollow(s.ctx, int(req.ToUserId), followerId)
	}
	return resp, nil
}
