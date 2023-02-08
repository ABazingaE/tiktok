package db

import (
	"context"
)

type Video struct {
	Id            int    `gorm:"column:id" json:"id"`                         //type:*int     comment:                    version:2023-01-05 10:20
	AuthorId      int64  `gorm:"column:author_id" json:"author_id"`           //type:*int     comment:                    version:2023-01-05 10:20
	PlayUrl       string `gorm:"column:play_url" json:"play_url"`             //type:string   comment:                    version:2023-01-05 10:20
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url"`           //type:string   comment:                    version:2023-01-05 10:20
	FavoriteCount int    `gorm:"column:favorite_count" json:"favorite_count"` //type:*int     comment:                    version:2023-01-05 10:20
	CommentCount  int    `gorm:"column:comment_count" json:"comment_count"`   //type:*int     comment:                    version:2023-01-05 10:20
	Title         string `gorm:"column:title" json:"title"`                   //type:string   comment:                    version:2023-01-05 10:20
	PublishTime   int    `gorm:"column:publish_time" json:"publish-time"`     //type:*int     comment:投稿发布的时间戳    version:2023-01-05 10:20
}

// TableName 表名:video，。
// 说明:
func (v *Video) TableName() string {
	return "video"
}

// 插入视频信息
func CreateVideo(ctx context.Context, video []*Video) error {
	if err := DB.WithContext(ctx).Create(video).Error; err != nil {
		return err
	}
	return nil
}

// 视频流
func GetVideoStream(ctx context.Context, latestTIme int64) (videoInfo []*Video, nextTime int64, error error) {
	//查询出所有小于latestTime的最新的最多30条视频信息
	var videos []*Video
	if err := DB.WithContext(ctx).Where("publish_time < ?", latestTIme).Order("publish_time desc").Limit(30).Find(&videos).Error; err != nil {
		return nil, 0, err
	}
	nextTime = int64(videos[0].PublishTime)
	return videos, nextTime, nil
}

// 查询指定用户发布的视频,按照从新到旧的顺序
func GetVideoByAuthorId(ctx context.Context, authorId int64) (videoInfo []*Video, error error) {
	var videos []*Video
	if err := DB.WithContext(ctx).Where("author_id = ?", authorId).Order("publish_time desc").Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

/*
* 查询指定视频id的视频信息,按照从新到旧的顺序
 */
func GetVideoByIds(ctx context.Context, videoId []int64) (videoInfo []*Video, error error) {
	var videos []*Video
	if err := DB.WithContext(ctx).Where("id in (?)", videoId).Order("publish_time desc").Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
