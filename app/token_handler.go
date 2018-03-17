package app

import "github.com/Leondroids/go-ethereum-rpc/rpc"

type TokenHandler struct {
	client *rpc.Client
}

func NewTokenHandler(appContext *Context) *TokenHandler {
	return &TokenHandler{
		client: appContext.Client,
	}
}
