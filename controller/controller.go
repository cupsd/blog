package controller

import (
	"blog/dal"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	articleInfoList, err := dal.GetArticleInfoList()
	if err != nil {
		log.Printf("err:%v", err)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", articleInfoList)
}
