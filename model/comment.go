package model

import "time"

// 评论实体
type Comment struct {
	Id         int       `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	Status     int       `db:"status"`
	ArticleId  int       `db:"article_id"`
	CreateTime time.Time `db:"create_time"`
}

// 留言实体
type LeaveMessage struct {
	Id         int       `db:"id"`
	Username   string    `db:"username"`
	CreateTime time.Time `db:"create_time"`
	Content    string    `db:"content"`
}
