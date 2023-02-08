package main

import (
	"context"
	"tiktok/cmd/video/rpc"
	"tiktok/cmd/video/service"
	"tiktok/kitex_gen/user"
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

// VideoInfoListById implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoInfoListById(ctx context.Context, req *video.VideoInfoListByIdReq) (resp *video.VideoInfoListByIdResp, err error) {
	//1.由视频id获取视频基本信息
	videos, err := service.NewVideoInfoService(ctx).VideoInfoListById(req)
	if err != nil {
		return resp, err
	}

	//2.构造返回值，遍历videos，获取authorId，查询作者信息
	var videoInfoList []*video.Video

	for _, v := range videos {
		authorId := v.AuthorId
		userInfoReq := &user.UserInfoReq{
			UserId: authorId,
			Token:  req.Token,
		}
		userInfoResp, err := rpc.GetUserInfo(ctx, userInfoReq)
		if err != nil {
			return nil, err
		}

		video := &video.Video{
			Id:            int64(v.Id),
			Author:        (*video.Author)(userInfoResp.User),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: int64(v.FavoriteCount),
			CommentCount:  int64(v.CommentCount),
			//此方法仅在喜欢列表调用，故此字段写死
			IsFavorite: true,
			Title:      v.Title,
		}
		videoInfoList = append(videoInfoList, video)

	}

	return &video.VideoInfoListByIdResp{
		VideoList: videoInfoList,
	}, nil
}
