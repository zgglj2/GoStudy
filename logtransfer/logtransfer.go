package main

import (
	"GoStudy/logtransfer/conf"
	"GoStudy/logtransfer/es"
	"GoStudy/logtransfer/kafka"
	"fmt"

	"github.com/go-ini/ini"
)

var (
	appConf conf.AppConf
)

func main() {
	err := ini.MapTo(&appConf, "./conf/config.ini")
	if err != nil {
		fmt.Println("load config failed, err: ", err)
		return
	}
	fmt.Println(appConf)
	fmt.Printf("%#v\n", appConf)

	err = es.Init([]string{appConf.ESConf.Address}, appConf.ChanMaxSize, appConf.Nums)
	if err != nil {
		fmt.Println("init elasticsearch failed, err: ", err)
		return
	}
	defer es.Finish()
	fmt.Println("init elasticsearch success")

	err = kafka.Init([]string{appConf.KafkaConf.Address}, appConf.Topic)
	if err != nil {
		fmt.Println("init kafka failed, err: ", err)
		return
	}
	defer kafka.Finish()
	fmt.Println("init kafka success")

}
