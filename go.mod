module GoStudy

go 1.15

require (
	github.com/DataDog/zstd v1.4.8 // indirect
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/PaulXu-cn/goeval v0.1.1
	github.com/Shopify/sarama v1.30.0
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/blanchonvincent/ctxarg v0.0.0-20190726074905-a05d037a0c36
	github.com/bwmarrin/snowflake v0.3.0
	github.com/cavaliercoder/go-rpm v0.0.0-20200122174316-8cb9fd9c31a8
	github.com/cheggaaa/pb/v3 v3.0.8
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/coreos/etcd v3.3.27+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-ini/ini v1.63.2
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gops v0.3.22
	github.com/google/uuid v1.3.0 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/isbm/go-deb v0.0.0-20200606113352-45f79b074aa5
	github.com/jmoiron/sqlx v1.3.4
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lxn/walk v0.0.0-20210112085537-c389da54e794
	github.com/lxn/win v0.0.0-20210218163916-a377121e959e // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-sqlite3 v1.14.8
	github.com/mitchellh/go-ps v1.0.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/nsqio/go-nsq v1.1.0
	github.com/panjf2000/ants v1.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/sassoftware/go-rpmutils v0.2.0
	github.com/schollz/progressbar/v3 v3.8.3
	github.com/shirou/gopsutil v3.21.10+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/sony/sonyflake v1.0.0
	github.com/ugorji/go v1.2.6 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9
	go.etcd.io/etcd v3.3.27+incompatible // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871 // indirect
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/tools v0.1.7
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
	google.golang.org/grpc v1.42.0 // indirect
	google.golang.org/protobuf v1.27.1
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/cavaliercoder/go-rpm => ./go-rpm

replace github.com/isbm/go-deb => ./go-deb

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.6
