package proxy

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"golang.org/x/net/proxy"
)

func HealthP(PROXY_ADDR string) (net.Conn, error) {

	net1, _ := proxy.SOCKS5("tcp", PROXY_ADDR,
		// &proxy.Auth{User: "username", Password: "password"},
		nil,
		&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
		},
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var conn net.Conn
	ch := make(chan int, 1)
	var err error

	go func() {
		select {
		default:
			conn, err = net1.Dial("tcp", "ptt.cc:23")
			fmt.Println("net.Dial", PROXY_ADDR, err)
			ch <- 1
		case <-ctx.Done():
			fmt.Println("Canceled by timeout")
			return
		}
	}()

	select {
	case <-ch:
		fmt.Println("Read from ch")
	case <-time.After(10 * time.Second):
		err = errors.New("Timed out")
		fmt.Println("Timed out")
	}

	return conn, err
}
