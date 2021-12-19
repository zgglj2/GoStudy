package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type LogData struct {
	Topic string
	Data  string
}

type esData struct {
	Data string
}

var (
	client      *elastic.Client
	logDataChan chan *LogData
	ctx         context.Context
	cancel      context.CancelFunc
)

func Init(addrs []string, maxChanSize, nums int) (err error) {
	client, err = elastic.NewClient(elastic.SetURL(addrs...))
	if err != nil {
		panic(err)
	}
	for _, host := range addrs {
		info, code, err := client.Ping(host).Do(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	}
	logDataChan = make(chan *LogData, maxChanSize)
	ctx, cancel = context.WithCancel(context.Background())
	for i := 0; i < nums; i++ {
		go sendToES()
	}
	return
}

func Finish() {
	cancel()
}

func SendToESChan(ld *LogData) {
	logDataChan <- ld
}

func sendToES() {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("es send is stop...\n")
			return

		case ld := <-logDataChan:
			msg := esData{Data: ld.Data}
			put1, err := client.Index().
				Index(ld.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Printf("index failed, err:%v\n", err)
				return
			}
			fmt.Printf("Indexed data %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		}
	}

}
