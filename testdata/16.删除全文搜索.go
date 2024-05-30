package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/es_ser"
)

func main() {
	core.InitConf()
	core.InitLogger()
	global.ESClient = core.EsConnect()
	es_ser.DeleteFullTextByArticleID("MI4aeYYB6uoytGZAtrHU")

}
