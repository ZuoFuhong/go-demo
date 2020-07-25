module go-demo

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/boltdb/bolt v1.3.1
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-redis/redis/v7 v7.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1 // indirect
	github.com/googollee/go-engine.io v1.4.3-0.20200220091802-9b2ab104b298
	github.com/googollee/go-socket.io v1.4.3
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.1
	github.com/julienschmidt/httprouter v1.3.0
	github.com/justinas/alice v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.5.1
	go.etcd.io/etcd v3.3.22+incompatible
	go.mongodb.org/mongo-driver v1.3.3
	go.uber.org/atomic v1.5.1 // indirect
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190530122614-20be4c3c3ed5
	golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/olivere/elastic.v5 v5.0.86
	sigs.k8s.io/yaml v1.2.0 // indirect
)

// 解决etcd依赖包冲突
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
