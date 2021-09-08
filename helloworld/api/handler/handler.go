package handler

import (
	"proto/helloworld"

	"github.com/micro/go-micro/v2"
)

type Handler struct {
	HelloworldClient helloworld.HelloworldService
	HelloworldEvent  micro.Event
}
