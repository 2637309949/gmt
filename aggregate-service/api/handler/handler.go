package handler

import (
	"proto/aggregate"
)

type Handler struct {
	AggregateService aggregate.AggregateService
}
