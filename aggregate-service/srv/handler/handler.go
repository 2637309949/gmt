package handler

import (
	"github.com/olivere/elastic"
)

type Handler struct {
	EsClient *elastic.Client
}
