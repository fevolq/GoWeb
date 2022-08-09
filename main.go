package main

import (
	"fmt"

	"go-web/config"
	"go-web/dao"
	"go-web/logic"
)

func init() {
	dao.MysqlConn()
}

func main() {

	logic.RegisRouter()
	app := logic.Router

	app.Run(fmt.Sprintf(":%d", config.PORT))
}
