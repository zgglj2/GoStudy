package main

import (
	"GoStudy/logagent/conf"
	"GoStudy/logagent/etcd"
	"GoStudy/logagent/kafka"
	"GoStudy/logagent/taillog"
	"GoStudy/logagent/utils"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	err = kafka.Init([]string{appConf.KafkaConf.Address}, appConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init kafka failed, err: ", err)
		return
	}
	defer kafka.Finish()
	fmt.Println("init kafka success")

	err = etcd.Init([]string{appConf.EtcdConf.Address}, time.Duration(appConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("init etcd failed, err: ", err)
		return
	}
	defer etcd.Finish()
	fmt.Println("init etcd success")
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		fmt.Println("get out bound ip failed, err: ", err)
		return
	}
	fmt.Println("out bound ip: ", ipStr)
	etcdConfKey := fmt.Sprintf(appConf.Key, ipStr)
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Println("etcd get conf failed, err: ", err)
		return
	}
	fmt.Println("etcd get conf success")

	confChan := etcd.ConfChan()
	taillog.InitTailTaskMgr(logEntryConf, confChan)

	go etcd.WatchConf(etcdConfKey)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
}
