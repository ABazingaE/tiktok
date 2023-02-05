package main

import (
	"context"
	"tiktok/cmd/video/service"
	video "tiktok/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// VideoStream implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoStream(ctx context.Context, req *video.VideoStreamReq) (resp *video.VideoStreamResp, err error) {
	resp, err = service.NewVideoFeedService(ctx).VideoFeed(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// VideoUpload implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoUpload(ctx context.Context, req *video.VideoUploadReq) (resp *video.VideoUploadResp, err error) {
	_, err = service.NewVideoUploadService(ctx).VideoUpload(req)
	if err != nil {
		return resp, err
	}
	return
}

// VideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoList(ctx context.Context, req *video.VideoListReq) (resp *video.VideoListResp, err error) {
	resp, err = service.NewVideoListService(ctx).VideoList(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
