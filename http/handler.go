package http

import (
	"log"

	"git.code.oa.com/CSIG_EDU_OA_SERVICE/go_http_demo/dao"
	"git.code.oa.com/CSIG_EDU_OA_SERVICE/go_http_demo/g"
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	responseOk(c, "hello world")
}

func getUsernameHandler(c *gin.Context) {
	kv, err := dao.GetValByUsername(c.Query("username"))
	if err != nil {
		log.Printf("db query err: %s", err.Error())
		responseErr(c, g.EC_MysqlQueryError, err)
	} else {
		responseOk(c, kv)
	}

}
