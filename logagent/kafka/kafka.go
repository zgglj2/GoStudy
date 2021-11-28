package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
	ctx         context.Context
	cancel      context.CancelFunc
)

type logData struct {
	topic string
	data  string
}

func Init(addrs []string, maxChanSize int) (err error) {
	// 构建 生产者
	// 生成 生产者配置文件
	config := sarama.NewConfig()
	// 设置生产者 消息 回复等级 0 1 all
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 设置生产者 成功 发送消息 将在什么 通道返回
	config.Producer.Return.Successes = true
	// 设置生产者 发送的分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// 连接 kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("sarama.NewSyncProducer failed, err: ", err)
		return
	}
	logDataChan = make(chan *logData, maxChanSize)
	ctx, cancel = context.WithCancel(context.Background())
	go sendToKafka()
	return
}

func Finish() {
	cancel()
	client.Close()
}

func sendToKafka() {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("kafka send is stop...\n")
			return

		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)

			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("pid: %v, offset: %v\n", pid, offset)
		}
	}

}

func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}
