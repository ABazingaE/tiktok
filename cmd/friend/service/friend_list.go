package service

import (
	"context"
	"tiktok/cmd/friend/dal/db"
	"tiktok/cmd/friend/rpc"
	"tiktok/kitex_gen/friend"
	"tiktok/kitex_gen/user"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s *FriendListService) FriendList(req *friend.FriendListReq) (resp *friend.FriendListResp, err error) {
	//1.查询好友id
	friendIds, err := db.GetFriendIdList(s.ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	//2.查询好友信息
	var friendList []*user.User
	for _, id := range friendIds {
		resp, err := rpc.GetUserInfo(s.ctx, &user.UserInfoReq{
			UserId: int64(id),
		})
		if err != nil {
			return nil, err
		}
		friendList = append(friendList, resp.User)
	}
	//3.查询最新消息
	var userList []*friend.FriendUser
	for i, friendId := range friendIds {
		content, msgType, err := db.GetLatestMessage(s.ctx, int(req.UserId), friendId)
		if err != nil {
			return nil, err
		}
		userList = append(userList, &friend.FriendUser{
			Id:            int64(friendId),
			Name:          friendList[i].Name,
			FollowCount:   friendList[i].FollowCount,
			FollowerCount: friendList[i].FollowerCount,
			IsFollow:      true,
			Message:       &content,
			MsgType:       int64(msgType),
		})
	}
	resp = &friend.FriendListResp{
		UserList: userList,
	}
	return resp, nil
}
