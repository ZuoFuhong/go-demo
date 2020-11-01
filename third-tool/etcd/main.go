package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

var client *clientv3.Client

func init() {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"47.98.199.80:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	//CURDKey()
	//grantLease()
	//autoGrantLease()
	//setKeyTTL()
	//watchKey()
	//txn()
	for i := 0; i < 3; i++ {
		go distributeLock()
	}
	time.Sleep(time.Hour)
}

// 操作Etcd
func CURDKey() {
	kv := clientv3.NewKV(client)
	_, err := kv.Put(context.Background(), "name", "dazuo")
	if err != nil {
		log.Panic(err)
	}
	resp, err := kv.Get(context.Background(), "name")
	if err != nil {
		log.Panic(err)
	}
	for _, v := range resp.Kvs {
		fmt.Printf("%s : %s version: %x\n", v.Key, v.Value, v.Version)
	}
	// op操作
	putOp := clientv3.OpDelete("name")
	_, err = kv.Do(context.Background(), putOp)
	if err != nil {
		log.Panic(err)
	}
}

// etcd需要先创建 lease，然后 put 命令加上参数 –lease= 来设置
func grantLease() {
	lease := clientv3.NewLease(client)
	grantResp, err := lease.Grant(context.Background(), 30)
	if err != nil {
		log.Panic(err)
	}
	id := grantResp.ID
	fmt.Printf("LeaseID：%x\n", grantResp.ID)

	// 撤销租约会使当前租约的所关联的key-value失效
	_, err = lease.Revoke(context.Background(), id)
	if err != nil {
		log.Panic(err)
	}
	liveResp, err := lease.TimeToLive(context.Background(), id)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("ttl: ", liveResp.TTL)
}

// 自动续租
func autoGrantLease() {
	lease := clientv3.NewLease(client)
	grantResp, err := lease.Grant(context.Background(), 10)
	if err != nil {
		log.Panic(err)
	}
	keepRespChan, err := lease.KeepAlive(context.Background(), grantResp.ID)
	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else {
					fmt.Printf("收到自动续租应答: %x\n", keepResp.ID)
				}
			}
		}
	END:
	}()
	time.Sleep(time.Hour)
}

// 通过租约put
func setKeyTTL() {
	grantResp, err := client.Grant(context.Background(), 10)
	if err != nil {
		log.Panic(err)
	}
	_, err = client.Put(context.Background(), "name", "dazuo", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Panic(err)
	}
}

// 监听Key
func watchKey() {
	wc := client.Watch(context.Background(), "name", clientv3.WithPrevKV())
	for v := range wc {
		for _, e := range v.Events {
			fmt.Printf("type:%v key:%v value:%v\n", e.Type, string(e.Kv.Key), string(e.Kv.Value))
		}
	}
}

// 事务
func txn() {
	kv := clientv3.NewKV(client)
	// 创建事务
	txn := kv.Txn(context.Background())
	txn.If(clientv3.Compare(clientv3.Value("name"), "=", "dazuo")).
		Then(clientv3.OpPut("name", "dazuo2")).
		Else(clientv3.OpPut("name", "dazuo3"))
	txnResp, err := txn.Commit()
	if err != nil {
		log.Panic(err)
	}
	if txnResp.Succeeded {
		fmt.Println("succeeded")
	}
}

// 实现分布式锁的基础
// Lease机制：Etcd 可以为存储的 key-value 对设置租约，当租约到期，key-value 将失效删除；同时也支持续约，通过客户端可以在租约
// 到期之前续约，以避免 key-value 对过期失效。Lease 机制可以保证分布式锁的安全性，为锁对应的 key 配置租约，即使锁的持有者因故障
// 而不能主动释放锁，锁也会因租约到期而自动释放。
//
// Revision机制：每个key带有一个 Revision号，每进行一次事务加一，因此它是全局唯一的，如初始值为 0，进行一次 put(key, value)，
// key的Revision 变为1；同样的操作，再进行一次，Revision 变为 2；换成key1进行put(key1, value)操作，Revision 将变为 3。
// 这种机制有一个作用：通过 Revision 的大小就可以知道进行写操作的顺序。在实现分布式锁时，多个客户端同时抢锁，根据 Revision号大小
// 依次获得锁，可以避免 “羊群效应” （也称 “惊群效应”），实现公平锁。

// Prefix机制：即前缀机制，也称目录机制。例如，一个名为 /mylock 的锁，两个争抢它的客户端进行写操作，实际写入的 key 分别为：
// key1="/mylock/UUID1"，key2="/mylock/UUID2"，其中，UUID 表示全局唯一的 ID，确保两个 key 的唯一性。很显然，写操作都会成
// 功，但返回的Revision不一样，那么，如何判断谁获得了锁呢？通过前缀 /mylock 查询，返回包含两个 key-value 对的的KeyValue列表，
// 同时也包含它们的 Revision，通过 Revision 大小，客户端可以判断自己是否获得锁，如果抢锁失败，则等待锁释放（对应的 key 被删除或
// 者租约过期），然后再判断自己是否可以获得锁；

// Watch机制：即监听机制，Watch 机制支持 Watch 某个固定的key，也支持Watch一个范围（前缀机制），当被Watch的key或范围发生变化，
// 客户端将收到通知；在实现分布式锁时，如果抢锁失败，可通过 Prefix 机制返回的 KeyValue 列表获得 Revision 比自己小且相差最小的
// key（称为 pre-key），对 pre-key 进行监听，因为只有它释放锁，自己才能获得锁，如果 Watch 到 pre-key 的 DELETE事件，则说明
// pre-key 已经释放，自己已经持有锁。
func distributeLock() {
	// 创建租约
	lease := clientv3.NewLease(client)
	grantResp, err := lease.Grant(context.Background(), 5)
	if err != nil {
		log.Panic(err)
	}
	leaseId := grantResp.ID
	ctx, cancelFunc := context.WithCancel(context.TODO())
	keepRespChan, err := lease.KeepAlive(ctx, leaseId)
	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else {
					fmt.Printf("收到自动续租应答: %x\n", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 尝试抢锁
	tryLock("mutex1", leaseId)
	// 释放锁
	unLock(cancelFunc, lease, leaseId)
}

func tryLock(lockKey string, leaseId clientv3.LeaseID) {
	kv := clientv3.NewKV(client)
	txn := kv.Txn(context.Background())
	// 事务抢锁：reversion + 1，同时授权租约
	txn.If(clientv3.Compare(clientv3.CreateRevision(lockKey), "=", 0)).
		Then(clientv3.OpPut(lockKey, "", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet(lockKey))

	txnRes, err := txn.Commit()
	if err != nil {
		log.Panic(err)
	}
	if txnRes.Succeeded {
		fmt.Println("抢锁成功，把持3秒")
		time.Sleep(time.Second * 3)
	} else {
		fmt.Println("锁被占用")
	}
}

func unLock(cancelFunc context.CancelFunc, lease clientv3.Lease, leaseId clientv3.LeaseID) {
	// 取消自动续租协程
	cancelFunc()
	_, err := lease.Revoke(context.TODO(), leaseId)
	if err != nil {
		log.Panic(err)
	}
}
