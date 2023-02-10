package service

import (
	"context"
	"strconv"
	"tiktok/cmd/friend/dal/db"
	"tiktok/kitex_gen/friend"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

// 定义全局变量，用于记录聊天记录的最后id,初始化为0，第一次查询所有聊天记录
var lastId = 0

func (m *MessageChatService) MessageChat(req *friend.MessageChatReq) (resp *friend.MessageChatResp, err error) {
	//1.查询聊天记录
	userId, err := strconv.Atoi(req.Token)
	if err != nil {
		return nil, err
	}
	chatList, err := db.GetChatRecord(m.ctx, userId, int(req.ToUserId), lastId)
	if err != nil {
		return nil, err
	}
	//2.返回聊天记录
	var chatRespList []*friend.Message
	for _, chat := range chatList {
		chatRespList = append(chatRespList, &friend.Message{
			FromUserId: int64(chat.FromUserId),
			ToUserId:   int64(chat.ToUserId),
			Content:    chat.Content,
			CreateTime: int64(chat.CreateTime),
		})
	}
	//3.更新lastId
	if len(chatList) > 0 {
		lastId = chatList[len(chatList)-1].Id
	}
	resp = &friend.MessageChatResp{
		MessageList: chatRespList,
	}
	return resp, nil
}
