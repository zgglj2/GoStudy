package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	ESConf    `ini:"es"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
	Nums        int    `ini:"nums"`
}
