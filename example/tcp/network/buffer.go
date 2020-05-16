package network

import (
	"errors"
	"io"
	"log"
)

type Buffer struct {
	reader io.Reader
	buf    []byte
	start  int
	end    int
}

func newBuffer(reader io.Reader, len int) Buffer {
	return Buffer{
		reader: reader,
		buf:    make([]byte, len),
		start:  0,
		end:    0,
	}
}

func (b *Buffer) len() int {
	return b.end - b.start
}

// 将有效的字节前移
func (b *Buffer) grow() {
	if b.start == 0 {
		return
	}
	copy(b.buf, b.buf[b.start:b.end])
	b.end -= b.start
	b.start = 0
}

// 从reader中读取字节，如果reader阻塞，发生阻塞
func (b *Buffer) readFromReader() error {
	b.grow()
	n, err := b.reader.Read(b.buf[b.end:])
	if err != nil {
		log.Println(err)
		return err
	}
	b.end += n
	return nil
}

// seek返回n个字节，而不产生移位，如果没有足够的字节，则返回错误
func (b *Buffer) seek(offset, limit int) ([]byte, error) {
	if b.end-b.start < offset+limit {
		return nil, errors.New("not enough")
	}
	return b.buf[b.start+offset : b.start+offset+limit], nil
}

// 移动offset个字节，返回limit个字节，如果没有足够的字节，则返回错误
func (b *Buffer) read(offset, limit int) ([]byte, error) {
	if b.len() < offset+limit {
		return nil, errors.New("not enough")
	}
	b.start += offset
	buf := b.buf[b.start : b.start+limit]
	b.start += limit
	return buf, nil
}
