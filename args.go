package main

import (
	"encoding/json"
	"os"
)


type ServerArgs struct {
	Host string `json:host`
	Port int `json:port`

	DbCfg struct {
		Host string `json:host`
		Port int `json:port`
	}`json:dbCfg`
}

func newServerArgs() (ServerArgs,error) {
	raw := os.Args[0]
	res := ServerArgs{}
	err := json.Unmarshal([]byte(raw),&res)
	return res,err
}