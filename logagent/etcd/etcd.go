package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	client    *clientv3.Client
	confChan  chan []*LogEntry
	watchChan clientv3.WatchChan
)

func Init(addrs []string, timeout time.Duration) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	confChan = make(chan []*LogEntry)
	return
}

func Finish() {
	client.Close()
}

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := client.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Println("unmarchal log conf failed, err: ", err)
			continue
		}
	}
	return
}

func WatchConf(key string) {
	watchChan = client.Watch(context.Background(), key)
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			var newConf []*LogEntry
			if ev.Type == clientv3.EventTypeDelete {
				confChan <- newConf
				continue
			}
			err := json.Unmarshal(ev.Kv.Value, &newConf)
			if err != nil {
				fmt.Println("unmarchal log conf failed, err: ", err)
				continue
			}
			confChan <- newConf
			fmt.Println("send new conf to confChan success")
		}
	}
}

func ConfChan() <-chan []*LogEntry {
	return confChan
}
