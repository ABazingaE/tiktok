package service

import (
	"context"
	"tiktok/cmd/comment/dal/db"
	"tiktok/kitex_gen/comment"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentListService returns a new CommentListService.
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *comment.CommentListReq) (resp []*db.Comment, err error) {
	videoId := int(req.VideoId)
	commentsInfo, err := db.GetCommentListByVideoId(s.ctx, videoId)
	if err != nil {
		return nil, err
	}
	return commentsInfo, nil
}
