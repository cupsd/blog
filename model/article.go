package model

import "time"

type ArticleInfo struct {
	ArticleId    int       `db:"id"`            // 文章ID
	CategoryId   int       `db:"category_id"`   // 分类ID
	Summary      string    `db:"summary"`       // 摘要
	CommentCount int       `db:"comment_count"` // 评论次数
	Title        string    `db:"title"`         // 标题
	CreateTime   time.Time `db:"create_time"`   // 创建时间
	UpdateTime   time.Time `db:"update_time"`   // 更新时间
	Status       int       `db:"status"`        // 1 显示 2 不显示
	ViewCount    int64     `db:"view_count"`    // 阅读次数
}

type ArticleDetail struct {
	ArticleInfo
	Category
	Content string `db:"content"`
}
