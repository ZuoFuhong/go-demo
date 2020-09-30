module go-demo

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet v2.1.2+incompatible // indirect
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/boltdb/bolt v1.3.1
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/emirpasic/gods v1.12.0
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/googollee/go-engine.io v1.4.3-0.20200220091802-9b2ab104b298
	github.com/googollee/go-socket.io v1.4.3
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.1
	github.com/jinzhu/gorm v1.9.16
	github.com/juju/ratelimit v1.0.1
	github.com/julienschmidt/httprouter v1.3.0
	github.com/justinas/alice v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prashantv/gostub v1.0.0
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/xormplus/builder v0.0.0-20200331055651-240ff40009be // indirect
	github.com/xormplus/xorm v0.0.0-20200731130200-6811f3bde592
	go.mongodb.org/mongo-driver v1.3.3
	go.uber.org/atomic v1.5.1 // indirect
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/sys v0.0.0-20200806125547-5acd03effb82 // indirect
	golang.org/x/tools v0.0.0-20200913032122-97363e29fc9b // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/olivere/elastic.v5 v5.0.86
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

// 解决etcd依赖包冲突
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
