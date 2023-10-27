package main

import (
	go_cache "github.com/wk331100/go-cache"
	"github.com/wk331100/go-cache-server/handlers"
)

const (
	Get    = "GET"
	Set    = "SET"
	SetEx  = "SETEX"
	Incr   = "INCR"
	Decr   = "DECR"
	IncrBy = "INCRBY"
	DecrBy = "DECRBY"
)

var (
	commandList = make(map[string]commandHandler)
)

func init() {
	commandList = map[string]commandHandler{
		Get:    &handlers.GetHandler{},
		Set:    &handlers.SetHandler{},
		SetEx:  &handlers.SetExHandler{},
		Incr:   &handlers.IncrHandler{},
		Decr:   &handlers.DecrHandler{},
		IncrBy: &handlers.IncrByHandler{},
		DecrBy: &handlers.DecrByHandler{},
	}
}

type commandHandler interface {
	Handle(server *go_cache.Cache, args []string) (string, error)
}
