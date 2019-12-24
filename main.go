package main

import (
	"blog/dal"
	"blog/routers"
)

func init() {

	dal.InitMysql("root:123@tcp(192.168.150.128:3306)/blog?parseTime=true")

}

func main() {
	r := routers.InitRouter()
	r.Run(":8080")
}
