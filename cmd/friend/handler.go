package main

import (
	"context"
	friend "tiktok/kitex_gen/friend"
)

// FriendServiceImpl implements the last service interface defined in the IDL.
type FriendServiceImpl struct{}

// FriendList implements the FriendServiceImpl interface.
func (s *FriendServiceImpl) FriendList(ctx context.Context, req *friend.FriendListReq) (resp *friend.FriendListResp, err error) {
	//  Your code here...
	return
}

// MessageChat implements the FriendServiceImpl interface.
func (s *FriendServiceImpl) MessageChat(ctx context.Context, req *friend.MessageChatReq) (resp *friend.MessageChatResp, err error) {
	//  Your code here...
	return
}

// MessageAction implements the FriendServiceImpl interface.
func (s *FriendServiceImpl) MessageAction(ctx context.Context, req *friend.MessageActionReq) (resp *friend.MessageActionResp, err error) {
	//  Your code here...
	return
}
