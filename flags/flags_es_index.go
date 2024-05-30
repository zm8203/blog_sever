package flags

import (
	"gvb_server/models"
	"gvb_server/service/es_ser/indexs"
)

func ESIndex() {
	indexs.CreateIndex(models.FullTextModel{})
	indexs.CreateIndex(models.ArticleModel{})
}
