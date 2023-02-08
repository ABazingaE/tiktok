package service

import (
	"context"
	"strconv"
	"tiktok/cmd/comment/dal/db"
	"tiktok/cmd/comment/rpc"
	"tiktok/kitex_gen/comment"
	"tiktok/kitex_gen/user"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *comment.CommentActionReq) (resp *comment.CommentActionResp, err error) {
	actionType := int(req.ActionType)
	resp = &comment.CommentActionResp{}
	switch actionType {
	case 1:
		userId, err := strconv.Atoi(req.Token)
		if err != nil {
			return nil, err
		}
		commentInfo, err := db.AddComment(s.ctx, int(req.VideoId), userId, req.CommentText)
		if err != nil {
			return nil, err
		}
		userInfo, err := rpc.GetUserInfo(s.ctx, &user.UserInfoReq{
			UserId: int64(userId),
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		cmt := &comment.Comment{
			Id:         int64(commentInfo.Id),
			User:       (*comment.User)(userInfo.User),
			Content:    commentInfo.Content,
			CreateDate: commentInfo.CreateDate,
		}
		resp.Comment = cmt

	case 2:
		db.DeleteCommentById(s.ctx, int(req.CommentId))

	}

	return resp, nil
}
