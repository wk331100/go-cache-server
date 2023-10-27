package handlers

import (
	jsoniter "github.com/json-iterator/go"
	go_cache "github.com/wk331100/go-cache"
	"github.com/wk331100/go-cache-server/types"
	"strconv"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type GetHandler struct{}

func (h *GetHandler) Handle(srv *go_cache.Cache, args []string) (string, error) {
	if len(args) == 0 {
		return "", types.ErrArgs
	}
	data, err := srv.Get(args[0])
	if err != nil {
		return "", err
	}
	switch data.(type) {
	case string:
		return data.(string), nil
	case int64:
		return strconv.FormatInt(data.(int64), 10), nil
	}
	return " ", nil
}

type SetHandler struct{}

func (h *SetHandler) Handle(srv *go_cache.Cache, args []string) (string, error) {
	if len(args) != 2 {
		return "", types.ErrArgs
	}
	srv.Set(args[0], args[1])
	return "", nil
}

type SetExHandler struct{}

func (h *SetExHandler) Handle(srv *go_cache.Cache, args []string) (string, error) {
	if len(args) != 3 {
		return "", types.ErrArgs
	}
	seconds, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		return "", types.ErrExpiration
	}
	srv.SetEx(args[0], args[1], time.Second*time.Duration(seconds))
	return "", nil
}

type IncrHandler struct{}

func (h *IncrHandler) Handle(srv *go_cache.Cache, args []string) (string, error) {
	if len(args) != 1 {
		return "", types.ErrArgs
	}
	srv.Incr(args[0])
	return "", nil
}

type DecrHandler struct{}

func (h *DecrHandler) Handle(srv *go_cache.Cache, args []string) (string, error) {
	if len(args) != 1 {
		return "", types.ErrArgs
	}
	srv.Decr(args[0])
	return "", nil
}

type IncrByHandler struct{}

func (h *IncrByHandler) Handle(srv *go_cache.Cache, args []string) (string, error) {
	if len(args) != 2 {
		return "", types.ErrArgs
	}
	num, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return "", types.ErrExpiration
	}
	srv.IncrBy(args[0], num)
	return "", nil
}

type DecrByHandler struct{}

func (h *DecrByHandler) Handle(srv *go_cache.Cache, args []string) (string, error) {
	if len(args) != 2 {
		return "", types.ErrArgs
	}
	num, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return "", types.ErrExpiration
	}
	srv.DecrBy(args[0], num)
	return "", nil
}
