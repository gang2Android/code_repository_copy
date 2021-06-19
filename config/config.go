package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var fileLocker sync.Mutex

type Config struct {
	SourceRepository string `json:"source_repository"`
	TargetRepository string `json:"target_repository"`
}

func LoadConfig(tasks *[]Config) {
	fileLocker.Lock()
	data, err := ioutil.ReadFile("./config.json")
	fileLocker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return
	}
	dataJson := []byte(data)
	err = json.Unmarshal(dataJson, tasks)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return
	}
	return
}
