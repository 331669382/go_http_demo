package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"git.code.oa.com/CSIG_EDU_OA_SERVICE/go_http_demo/dao"
	"git.code.oa.com/CSIG_EDU_OA_SERVICE/go_http_demo/g"
	"git.code.oa.com/CSIG_EDU_OA_SERVICE/go_http_demo/http"
)

func main() {
	cfg := flag.String("c", "cfg.yml", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	g.ParseConfig(*cfg)
	dao.Init(g.Config().Mysql, g.Config().Debug)

	go http.Init()

	select {}
}
