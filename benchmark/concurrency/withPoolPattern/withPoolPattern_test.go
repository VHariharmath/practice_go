package main

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	deamonStarted := startNetworkDeamon()
	deamonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("Cannot dial connection %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("Cannot read conn %v", err)
		}
		conn.Close()
	}
}
