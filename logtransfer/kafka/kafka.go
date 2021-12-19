package kafka

import (
	"GoStudy/logtransfer/es"
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	client sarama.Consumer
	// ctx    context.Context
	// cancel context.CancelFunc
)

func Init(addrs []string, topic string) (err error) {
	// 生成消费者 实例
	client, err = sarama.NewConsumer(addrs, nil)
	if err != nil {
		log.Print(err)
		return
	}
	// 拿到 对应主题下所有分区
	partitionList, err := client.Partitions(topic)
	if err != nil {
		log.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	// 遍历所有分区
	for partition := range partitionList {
		//消费者 消费 对应主题的 具体 分区 指定 主题 分区 offset  return 对应分区的对象
		pc, err := client.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Println(err)
			return err
		}

		// 运行完毕记得关闭
		defer pc.AsyncClose()

		// 去出对应的 消息
		// 通过异步 拿到 消息
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				ld := &es.LogData{
					Topic: topic,
					Data:  string(msg.Value),
				}
				fmt.Printf("LogData: %v\n", ld)
				es.SendToESChan(ld)
			}
		}(pc)
	}
	wg.Wait()
	return
}

func Finish() {
	// cancel()
	client.Close()
}
