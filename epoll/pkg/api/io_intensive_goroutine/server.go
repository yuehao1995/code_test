package io_intensive_goroutine

import (
	"flag"
	"io"
	"log"
	"net"
	_ "net/http/pprof"
	"os"
	"syscall"
	"time"

	"github.com/rcrowley/go-metrics"
)

var (
	iotime = flag.Duration("io", time.Duration(10*time.Millisecond), "sleep time")
)
var (
	opsRate = metrics.NewRegisteredMeter("ops", nil)
)

func Run() {
	flag.Parse()

	setLimit()
	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	ln, err := net.Listen("tcp", ":37137")
	if err != nil {
		panic(err)
	}

	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()

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

		go handleConn(conn)
		connections = append(connections, conn)
		if len(connections)%100 == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}
}

func handleConn(conn net.Conn) {
	for {
		time.Sleep(*iotime)
		_, err := io.CopyN(conn, conn, 8)
		if err != nil {
			log.Printf("failed to copy: %v", err)
			conn.Close()
			return
		}
		opsRate.Mark(1)
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