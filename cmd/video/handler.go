package main

import (
	"context"
	video "tiktok/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// VideoStream implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoStream(ctx context.Context, req *video.VideoStreamReq) (resp *video.VideoStreamResp, err error) {
	//  Your code here...
	return
}

// VideoUpload implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoUpload(ctx context.Context, req *video.VideoUploadReq) (resp *video.VideoUploadResp, err error) {
	//  Your code here...
	return
}

// VideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoList(ctx context.Context, req *video.VideoListReq) (resp *video.VideoListResp, err error) {
	//  Your code here...
	return
}