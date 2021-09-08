package handler

import (
	"proto/helloworld"
)

type Handler struct {
	HelloworldClient helloworld.HelloworldService
}
