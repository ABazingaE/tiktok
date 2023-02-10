package service

import (
	"context"
	"strconv"
	"tiktok/cmd/friend/dal/db"
	"tiktok/kitex_gen/friend"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (m *MessageActionService) SendMessage(req *friend.MessageActionReq) (resp *friend.MessageActionResp, err error) {
	fromUserId, err := strconv.Atoi(req.Token)
	if err != nil {
		return nil, err
	}
	err = db.SendMessage(m.ctx, fromUserId, int(req.ToUserId), req.Content)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
