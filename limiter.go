package swagin

import (
	"errors"
	"net"
	"sync"
)

var limitReached = errors.New("connection limit reached")

type LimitListener struct {
	net.Listener
	counter SyncedLimitedInt
	once    sync.Once
}

// NewLimitListener returns a Listener that is limited to given number of connections
func NewLimitListener(listener net.Listener, limit int) *LimitListener {
	return &LimitListener{
		Listener: listener,
		counter:  NewLimitedInt(limit),
	}
}

func (l *LimitListener) Accept() (net.Conn, error) {
	if l.counter.Incr() {
		c, err := l.Listener.Accept()
		if err != nil {
			l.counter.Decr()
			return nil, err
		}
		return c, nil
	} else {
		return nil, limitReached
	}
}

func (l *LimitListener) Close() error {
	var err error
	l.once.Do(func() {
		err = l.Listener.Close()
		l.counter.Decr()
	})
	return err
}

func (l *LimitListener) Addr() net.Addr {
	return l.Listener.Addr()
}

func (l *LimitListener) SetLimit(limit int) {
	l.counter.SetLimit(limit)
}
