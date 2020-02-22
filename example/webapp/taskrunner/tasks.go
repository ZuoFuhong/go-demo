package taskrunner

import (
	"errors"
	"log"
)

// 生产者
func ClearDispatcher(dc dataChan) error {
	for i := 0; i < 30; i++ {
		dc <- i
		log.Printf("Dispatcher sent: %v", i)
	}
	return nil
}

// 消费者
func ClearExecutor(dc dataChan) error {
forloop:
	for {
		select {
		case d := <-dc:
			log.Printf("Executor received: %v", d)
		default:
			break forloop
		}
	}
	return errors.New("executor error")
}
