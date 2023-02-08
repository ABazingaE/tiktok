package service

import (
	"context"
	"tiktok/cmd/video/dal/db"
	"tiktok/kitex_gen/video"
)

type VideoInfoService struct {
	ctx context.Context
}

// NewVideoInfoService NewVideoInfoService new VideoInfoService
func NewVideoInfoService(ctx context.Context) *VideoInfoService {
	return &VideoInfoService{
		ctx: ctx,
	}
}

// VideoInfoListById VideoInfoListById
func (s *VideoInfoService) VideoInfoListById(req *video.VideoInfoListByIdReq) (videoInfo []*db.Video, err error) {
	videoInfo, err = db.GetVideoByIds(s.ctx, req.VideoIds)
	if err != nil {
		return nil, err
	}
	return
}
