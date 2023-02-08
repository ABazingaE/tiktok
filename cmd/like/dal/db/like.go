package db

import (
	"context"
	"gorm.io/gorm"
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

type Like struct {
	Id      int `gorm:"column:id" json:"id"`             //type:*int   comment:          version:2023-01-07 18:05
	UserId  int `gorm:"column:user_id" json:"user_id"`   //type:*int   comment:          version:2023-01-07 18:05
	VideoId int `gorm:"column:video_id" json:"video_id"` //type:*int   comment:视频id    version:2023-01-07 18:05
}

// TableName 表名:like，。
// 说明:
func (l *Like) TableName() string {
	return "like"
}

// 更新点赞信息——点赞或取消点赞
func LikeAction(ctx context.Context, user_id int, video_id int, action_type int) error {
	/*
		判断action_type,如果是1则点赞，如果是0则取消点赞
		点赞即为向like表插入一条数据，取消点赞即删除一条数据
	*/
	if action_type == 1 {
		like := &Like{
			UserId:  user_id,
			VideoId: video_id,
		}
		//向like表插入数据
		if err := DB.WithContext(ctx).Create(like).Error; err != nil {
			return err
		}

		//更新video表中的点赞数,favorite_count值加1
		if err := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
			return err
		}

	} else {
		if err := DB.WithContext(ctx).Where("user_id = ? and video_id = ?", user_id, video_id).Delete(&Like{}).Error; err != nil {
			return err
		}
		//更新video表中的点赞数,favorite_count值减1
		if err := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count- ?", 1)).Error; err != nil {
			return err
		}
	}
	return nil
}

// 根据用户id查询点赞视频id
func GetLikeVideoIdByUserId(ctx context.Context, user_id int) ([]int, error) {
	var like []Like
	var video_ids []int
	if err := DB.WithContext(ctx).Where("user_id = ?", user_id).Find(&like).Error; err != nil {
		return nil, err
	}
	for _, v := range like {
		video_ids = append(video_ids, v.VideoId)
	}
	return video_ids, nil
}
