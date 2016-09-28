package config

import (
	"io/ioutil"
	"encoding/json"
)

type SimpleServerConfig struct {
	Port          int    `json:"port"`
	PprofPort     int    `json:"profport"`
	RetryTime     int    `json:"retry"`
	Timeout       int    `json:"timeout"`
}

var Cfg SimpleServerConfig

func ParseConf(file string) (err error) {
	cnt, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(cnt, &Cfg);
	return
}