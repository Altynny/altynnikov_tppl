package network

import (
	"io"
	"net"
	"testing"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	authBuffer := make([]byte, 6)
	io.ReadFull(conn, authBuffer)
	conn.Write([]byte("granted"))
	responseData := []byte("\x00\x06F\xd8\xb9\xa3\xf6\x9eB9z\xe1\x04G5")
	cmdBuffer := make([]byte, 3)
	for {
		conn.Read(cmdBuffer)
		conn.Write(responseData)
	}
}

const ServerAddress = "localhost:8080"

func TestNetworkWorkerOnMockServer(t *testing.T) {
	listener, _ := net.Listen("tcp", ServerAddress)
	defer listener.Close()

	SensorNw := NetworkWorker{Address: ServerAddress, PackageSize: 15}
	SensorByteChan := make(chan []byte)
	go SensorNw.Run(SensorByteChan)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
		break
	}
	expected := []byte("\x00\x06F\xd8\xb9\xa3\xf6\x9eB9z\xe1\x04G5")
	if res := <-SensorByteChan; string(res) != string(expected) {
		t.Errorf("Got unexpected bytes %d\n instead of\n %d", res, expected)
	}
}
