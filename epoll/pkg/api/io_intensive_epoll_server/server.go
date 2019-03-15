package io_intensive_epoll_server

import (
	"flag"
	"io"
	"log"
	"net"
	_ "net/http/pprof"
	"os"
	"syscall"
	"time"

	"github.com/libp2p/go-reuseport"
	"github.com/rcrowley/go-metrics"
)

var (
	c      = flag.Int("c", 10, "concurrency")
	iotime = flag.Duration("io", time.Duration(10*time.Millisecond), "sleep time")
)

var (
	opsRate = metrics.NewRegisteredMeter("ops", nil)
)

func Run() {
	flag.Parse()

	setLimit()
	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	for i := 0; i < *c; i++ {
		go startEpoll()
	}

	select {}
}

func startEpoll() {
	ln, err := reuseport.Listen("tcp", ":37137")
	if err != nil {
		panic(err)
	}

	epoller, err := MkEpoll()
	if err != nil {
		panic(err)
	}

	go start(epoller)

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		if err := epoller.Add(conn); err != nil {
			log.Printf("failed to add connection %v", err)
			conn.Close()
		}
	}
}

func start(epoller *epoll) {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			log.Printf("failed to epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}

			time.Sleep(*iotime)
			io.CopyN(conn, conn, 8)
			if err != nil {
				if err := epoller.Remove(conn); err != nil {
					log.Printf("failed to remove %v", err)
				}
				conn.Close()
			}

			opsRate.Mark(1)
		}
	}
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}
