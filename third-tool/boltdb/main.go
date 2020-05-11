package main

import (
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

func main() {
	baseUse()
}

func baseUse() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin(true)
	if err != nil {
		panic(err)
	}
	b := tx.Bucket([]byte("myBuckets"))
	err = b.Put([]byte("age"), []byte("23"))
	if err != nil {
		panic(err)
	}

	// 请注意，Get()返回的值只在事务打开时有效。如果需要在事务外部使用值，则必须使用copy()将其复制到另一个字节片。
	value := b.Get([]byte("age"))
	fmt.Println(string(value))

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

// 事务辅助函数
func transactionsWrapper() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	// 读写事务
	err = db.Update(func(tx *bolt.Tx) error {
		return nil
	})

	// 只读事务
	err = db.View(func(tx *bolt.Tx) error {
		return nil
	})

	// 批处理读写事务
	err = db.Batch(func(tx *bolt.Tx) error {
		return nil
	})
	defer db.Close()
}

// 手动管理事务
func manuallyTransaction() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(true)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	_, err = tx.CreateBucket([]byte("MyBucket"))
	if err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}

	defer db.Close()
}

//
func useBucket() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		//e := tx.DeleteBucket([]byte("myBuckets"))
		//if e != nil {
		//	panic(e)
		//}
		b, e := tx.CreateBucketIfNotExists([]byte("myBuckets"))
		if e != nil {
			panic(e)
		}
		e = b.Put([]byte("name"), []byte("dazuo"))
		e = b.Put([]byte("age"), []byte("24"))
		e = b.Put([]byte("gender"), []byte("1"))
		if e != nil {
			panic(e)
		}
		value := b.Get([]byte("name"))
		fmt.Println(string(value))

		e = b.Delete([]byte("name"))
		value = b.Get([]byte("name"))
		if e != nil {
			panic(e)
		}
		fmt.Println("value: ", value)

		// 遍历
		c := b.Cursor()
		for k,v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("k=%s, v=%s\n", k, v)
		}

		fmt.Println("prefix ---------------")

		// 前缀匹配
		prefix := []byte("gen")
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		fmt.Println("ForEach --------------")

		// 遍历
		_ = b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})

		return nil
	})

	defer db.Close()
}