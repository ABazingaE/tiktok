package main

import (
	"context"
	comment "tiktok/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionReq) (resp *comment.CommentActionResp, err error) {
	//  Your code here...
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListReq) (resp *comment.CommentListResp, err error) {
	//  Your code here...
	return
}
