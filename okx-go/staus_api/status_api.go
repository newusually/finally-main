package staus_api

import (
	"finally-main/okx-go/client"
	"finally-main/okx-go/consts"
)

type StatusAPI struct {
	*client.Client
}

func NewStatusAPI(apiKey, apiSecretKey, passphrase string, useServerTime bool, flag string) *StatusAPI {
	return &StatusAPI{
		Client: client.NewClient(apiKey, apiSecretKey, passphrase, useServerTime, flag),
	}
}

func (s *StatusAPI) Status(state string) (map[string]interface{}, error) {
	params := make(map[string]string)
	params["state"] = state
	return s.Request(consts.GET, consts.STATUS, params)
}
