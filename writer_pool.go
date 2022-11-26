package syslog

import (
	"fmt"
	"math/rand"
	"time"
)

type WriterPool struct {
	pool []*Writer
	rand *rand.Rand
}

func NewWriterPool(poolSize int, addr string) (*WriterPool, error) {
	if poolSize <= 0 {
		return nil, fmt.Errorf("Pool size must be greater than 0")
	}
	pool := make([]*Writer, poolSize)
	for i := 0; i < poolSize; i++ {
		w, err := Dial(addr)
		if err != nil {
			return nil, err
		}
		pool[i] = w
	}
	return &WriterPool{
		pool: pool,
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}, nil
}

func (w *WriterPool) Write(b []byte) (int, error) {
	conn := w.pool[w.rand.Intn(len(w.pool))]
	return conn.Write(b)
}

// Close closes a connection to the syslog daemon.
func (w *WriterPool) Close() error {
	for _, w := range w.pool {
		err := w.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
