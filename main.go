package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
	"transwitt/transwitt"

	"gopkg.in/yaml.v2"
)

func main() {
	// Logging
	f, err := os.OpenFile("/tmp/transwitt.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(f)

	// 임시 코드. 커밋 시 API 비공개를 위해 yaml에 저장하여 개발 진행.
	yamlFile, err := ioutil.ReadFile("/tmp/api.yaml")
	var conf transwitt.OperateConfig
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Println("Fail to load api.ymal", err)
		return
	}
	err = transwitt.Run(conf)

	if err != nil {
		log.Println("Fail to run transwitt", err)
		return
	}

	for {
		time.Sleep(time.Millisecond * 50)
	}
	log.Println("End of program. never come here")

}
