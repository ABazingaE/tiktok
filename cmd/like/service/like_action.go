package service

import (
	"context"
	"strconv"
	"tiktok/cmd/like/dal/db"
	"tiktok/kitex_gen/like"
)

type LikeActionService struct {
	ctx context.Context
}

// NewLikeActionService new LikeActionService
func NewLikeActionService(ctx context.Context) *LikeActionService {
	return &LikeActionService{ctx: ctx}
}

// LikeAction like action.
func (s *LikeActionService) LikeAction(req *like.LikeActionReq) (resp *like.LikeActionResp, err error) {
	//request中的token在api层已解析为user_id
	userId, err := strconv.Atoi(req.Token)
	if err != nil {
		return nil, err
	}

	//调用dal层
	err = db.LikeAction(s.ctx, userId, int(req.VideoId), int(req.ActionType))
	if err != nil {
		return nil, err
	}
	return nil, nil
}
