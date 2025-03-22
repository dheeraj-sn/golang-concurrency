package primitives

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func sumulateConnection() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)

	startServer := func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
			}
			sumulateConnection()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}
	go startServer()
	return &wg
}

func TestConnectiond() {
	daemon := startNetworkDaemon()
	daemon.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
