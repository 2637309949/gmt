package handler

import (
	"proto/{{.name}}"
)

type Handler struct {
	{{toTitle .name}}Service {{.name}}.{{toTitle .name}}Service
}
