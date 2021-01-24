package main

import (
	"io/ioutil"
	"log"
	"time"
	"transwitt/transwitt"

	"gopkg.in/yaml.v2"
)

func main() {
	// 임시 코드. 커밋 시 API 비공개를 위해 yaml에 저장하여 개발 진행.
	yamlFile, err := ioutil.ReadFile("/tmp/api.yaml")
	var conf transwitt.APIConfig
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Println("Fail to load api.ymal", err)
		return
	}
	log.Println("conf", conf)

	err = transwitt.Run(transwitt.OperateConfig{
		Messanger: transwitt.MessagnerConfig{
			Telegram: conf.Telegram,
		},
	})

	if err != nil {
		log.Println("Fail to run transwitt", err)
		return
	}

	for {
		time.Sleep(time.Millisecond * 50)
	}
	log.Println("End of program. never come here")

}
