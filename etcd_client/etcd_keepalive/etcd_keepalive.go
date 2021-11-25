package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		fmt.Printf("Grant failed, err:%v\n", err)
		return
	}
	_, err = cli.Put(context.TODO(), "gyc", "1", clientv3.WithLease(resp.ID))
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		fmt.Printf("KeepAlive failed, err:%v\n", err)
		return
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}
