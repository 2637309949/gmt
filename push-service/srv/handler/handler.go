package handler

import (
	"push/srv/publisher"
)

type Handler struct {
	Publisher publisher.Publisher
}
