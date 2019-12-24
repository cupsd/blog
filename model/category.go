package model

import "time"

type Category struct {
	CategoryId   int       `db:"id"`
	CategoryName string    `db:"category_name"`
	CreateTime   time.Time `db:"create_time"`
}
