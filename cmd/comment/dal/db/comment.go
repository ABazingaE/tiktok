package db

import (
	"context"
	"gorm.io/gorm"
	"time"
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

type Comment struct {
	Id         int    `gorm:"column:id" json:"id"`                   //type:*int         comment:    version:2023-01-08 22:24
	VideoId    int    `gorm:"column:video_id" json:"video_id"`       //type:string       comment:    version:2023-01-08 22:24
	UserId     int    `gorm:"column:user_id" json:"user_id"`         //type:*int         comment:    version:2023-01-08 22:24
	Content    string `gorm:"column:content" json:"content"`         //type:string       comment:    version:2023-01-08 22:24
	CreateDate string `gorm:"column:create_date" json:"create_date"` //type:*time.Time   comment:    version:2023-01-08 22:24
}

// TableName 表名:comment，。
// 说明:
func (c *Comment) TableName() string {
	return "comment"
}

// 添加评论
func AddComment(ctx context.Context, video_id int, user_id int, content string) (*Comment, error) {
	comment := &Comment{
		VideoId:    video_id,
		UserId:     user_id,
		Content:    content,
		CreateDate: time.Now().String(),
	}
	if err := DB.WithContext(ctx).Create(comment).Error; err != nil {
		return nil, err
	}
	//更新视频评论数
	if err := DB.WithContext(ctx).Model(&Video{}).Where("id = ?", video_id).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

// 根据评论id删除评论
func DeleteCommentById(ctx context.Context, comment_id int) error {
	if err := DB.WithContext(ctx).Where("id = ?", comment_id).Delete(&Comment{}).Error; err != nil {
		return err
	}
	return nil
}

// 根据视频id查询评论,按时间倒序排列
func GetCommentListByVideoId(ctx context.Context, video_id int) ([]*Comment, error) {
	var comments []*Comment
	if err := DB.WithContext(ctx).Where("video_id = ?", video_id).Order("create_date desc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
