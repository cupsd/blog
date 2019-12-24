package dal

import (
	"blog/model"
	"fmt"
	"testing"
)

func init() {
	dns := "root:123@tcp(192.168.150.128:3306)/blog?parseTime=true"
	InitMysql(dns)
}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.Summary = "this is test summary"
	article.Title = "test title"
	article.Content = "this is a test content"
	article.ArticleInfo.CategoryId = 1

	articleId, err := InsertArticle(article)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(articleId)
}

func TestGetArticleInfoList(t *testing.T) {

	articleInfoList, err := GetArticleInfoList()
	if err != nil {
		t.Error(err)
	}

	for _, article := range articleInfoList {
		t.Logf("%#v", article)
	}
}
