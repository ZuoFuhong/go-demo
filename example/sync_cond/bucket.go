package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

// DataBucket 数据 bucket
type DataBucket struct {
	buffer *bytes.Buffer
	mutex  *sync.RWMutex
	cond   *sync.Cond
}

func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer: bytes.NewBuffer(buf),
		mutex:  new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())
	return db
}

// Read 读取器
func (db *DataBucket) Read(id int) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	var data []byte
	var d byte
	var err error
	for {
		if d, err = db.buffer.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					// data 不为空，则打印它
					fmt.Printf("reader-%d: %s\n", id, data)
				}
				// 缓冲区为空，通过 Wait 方法等待通知，进入阻塞状态
				// 其中：使用condition必须加锁，Wait()会自动释放c.L，并挂起调用者的goroutine。之后恢复执行，Wait()会在返回时对c.L加锁。
				// 除非被Signal或者Broadcast唤醒，否则Wait()不会返回。
				db.cond.Wait()
				// 将 data 清空
				data = data[:0]
				continue
			}
		}
		data = append(data, d)
	}
}

// Put 写入器
func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	n, err := db.buffer.Write(d)
	// 通过 Signal 通知处于阻塞状态的读取器（随机挑一个）
	db.cond.Signal()
	// 通过 Broadcast 唤醒所有等待的goroutine
	//db.cond.Broadcast()
	return n, err
}

func main() {
	db := NewDataBucket()
	// 开启读取器协程
	for i := 0; i < 2; i++ {
		go db.Read(i)
	}
	// 开启写入器协程
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			d := fmt.Sprintf("data-%d", i)
			_, _ = db.Put([]byte(d))
		}
	}()
	time.Sleep(15 * time.Second)
}
