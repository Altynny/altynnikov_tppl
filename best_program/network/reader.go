package network

import (
	"fmt"
	"io"
	"net"
	"time"
)

type NetworkWorker struct {
	Address     string
	PackageSize int
}

func (nw *NetworkWorker) connect() (net.Conn, error) {
	connection, err := net.Dial("tcp", nw.Address)
	if err != nil {
		fmt.Println("Couldn't connect: ", err)
		return connection, err
	}
	secretKey := "isu_pt"
	_, err = connection.Write([]byte(secretKey))
	if err != nil {
		fmt.Println("Couldn't send secret key: ", err)
		return connection, err
	}
	buffer := make([]byte, nw.PackageSize)
	_, err = connection.Read(buffer)
	if err != nil {
		fmt.Println("Couldn't authorize: ", err)
		return connection, err
	}
	return connection, nil
}

func (nw *NetworkWorker) Run(outCh chan<- []byte) {
	buffer := make([]byte, nw.PackageSize)
	for {
		connection, err := nw.connect()
		if err != nil {
			fmt.Println("Couldn't connect: ", err)
			fmt.Println("Trying to reconnect in 0.5s...")
			time.Sleep(500 * time.Millisecond)
			continue
		}
		for {
			_, err := connection.Write([]byte("get"))
			if err != nil {
				fmt.Println("Couldn't send command: ", err)
				break
			}
			if _, err := io.ReadFull(connection, buffer); err != nil {
				fmt.Println("Couldn't get answer: ", err)
				continue
			}
			cs := buffer[nw.PackageSize-1]
			var s byte
			for _, b := range buffer[:nw.PackageSize-1] {
				s += b
			}

			if s == cs {
				outCh <- buffer
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}
