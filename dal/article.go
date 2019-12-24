package dal

import (
	"blog/model"
	"database/sql"
)

// 插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	var result sql.Result

	sqlStr := `INSERT INTO article(category_id,content,title,summary)VALUES(?,?,?,?)`

	result, err = Db.Exec(sqlStr, article.ArticleInfo.CategoryId, article.Content, article.Title, article.Summary)

	if err != nil {
		return
	}

	return result.LastInsertId()
}

// 获取所有文章
func GetArticleInfoList() (articleInfoList []*model.ArticleInfo, err error) {
	sqlStr := `SELECT id,title,create_time,summary,comment_count,view_count FROM article`

	err = Db.Select(&articleInfoList, sqlStr)
	if err != nil {
		return
	}

	return
}
