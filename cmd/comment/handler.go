package main

import (
	"context"
	"tiktok/cmd/comment/rpc"
	"tiktok/cmd/comment/service"
	comment "tiktok/kitex_gen/comment"
	"tiktok/kitex_gen/user"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionReq) (resp *comment.CommentActionResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	resp, err = service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	if err = req.IsValid(); err != nil {
		return nil, err
	}
	commentInfo, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		return nil, err
	}
	var comments []*comment.Comment
	//删除comments中的空元素

	for _, v := range commentInfo {
		userId := int64(v.UserId)
		userInfo, err := rpc.GetUserInfo(ctx, &user.UserInfoReq{
			UserId: userId,
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		comment := &comment.Comment{
			Id:         int64(v.Id),
			User:       (*comment.User)(userInfo.User),
			Content:    v.Content,
			CreateDate: v.CreateDate,
		}
		comments = append(comments, comment)
	}
	return &comment.CommentListResp{
		CommentList: comments,
	}, nil
}
