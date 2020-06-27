package http

import (
	"log"

	"git.code.oa.com/CSIG_EDU_OA_SERVICE/go_http_demo/g"
	"github.com/gin-gonic/gin"
)

var s *gin.Engine

func Init() {
	gin.SetMode(gin.ReleaseMode)
	s = gin.Default()

	setRoutes()

	log.Println("Server listening at", g.Config().Http.Listen)
	s.Run(g.Config().Http.Listen)
}

func setRoutes() {
	s.GET("/", indexHandler)
	s.GET("/get_username", getUsernameHandler)
}

func responseOk(c *gin.Context, data interface{}) {
	js := ResponseJSON{Code: 0, Data: data}
	c.JSON(200, &js)
}

func responseErr(c *gin.Context, code uint32, err error) {
	js := ResponseJSON{Code: code, Message: err.Error()}
	c.JSON(200, &js)
}

type ResponseJSON struct {
	Code    uint32      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}
