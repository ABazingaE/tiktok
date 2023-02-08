package service

import (
	"context"
	"tiktok/cmd/like/dal/db"
	"tiktok/cmd/like/rpc"
	"tiktok/kitex_gen/like"
	"tiktok/kitex_gen/video"
)

type LikeListService struct {
	ctx context.Context
}

// NewLikeActionService new LikeActionService
func NewLikeListService(ctx context.Context) *LikeListService {
	return &LikeListService{ctx: ctx}
}

// Like List
func (s *LikeListService) LikeList(req *like.LikeListReq) (resp *like.LikeListResp, err error) {
	videoIds, err := db.GetLikeVideoIdByUserId(s.ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	//将videoIds转化为int数组
	var videoIdList []int64
	for _, v := range videoIds {
		videoIdList = append(videoIdList, int64(v))
	}

	videos, err := rpc.VideoInfoListById(s.ctx, &video.VideoInfoListByIdReq{
		VideoIds: videoIdList,
		Token:    req.Token,
	})

	if err != nil {
		return nil, err
	}

	//将videos转化为like.Video数组
	var videoList []*like.Video
	for _, v := range videos {
		videoList = append(videoList, &like.Video{
			Id:            v.Id,
			Author:        (*like.Author)(v.Author),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}

	return &like.LikeListResp{
		VideoList: videoList,
	}, nil

}
