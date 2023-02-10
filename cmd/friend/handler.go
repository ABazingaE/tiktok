package main

import (
	"context"
	"tiktok/cmd/friend/service"
	friend "tiktok/kitex_gen/friend"
)

// FriendServiceImpl implements the last service interface defined in the IDL.
type FriendServiceImpl struct{}

// FriendList implements the FriendServiceImpl interface.
func (s *FriendServiceImpl) FriendList(ctx context.Context, req *friend.FriendListReq) (resp *friend.FriendListResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	resp, err = service.NewFriendListService(ctx).FriendList(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MessageChat implements the FriendServiceImpl interface.
func (s *FriendServiceImpl) MessageChat(ctx context.Context, req *friend.MessageChatReq) (resp *friend.MessageChatResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	resp, err = service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MessageAction implements the FriendServiceImpl interface.
func (s *FriendServiceImpl) MessageAction(ctx context.Context, req *friend.MessageActionReq) (resp *friend.MessageActionResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	resp, err = service.NewMessageActionService(ctx).SendMessage(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
