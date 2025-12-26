package main

import (
	"bestProgram/network"
	"fmt"
	"os"
	"time"
)

type SensorData struct {
	TimeStamp uint64
	Temp      float32
	Pressure  int16
}

func (s SensorData) GetTimeStamp() uint64 { return s.TimeStamp }

type CoordsData struct {
	TimeStamp uint64
	X, Y, Z   int32
}

func (s CoordsData) GetTimeStamp() uint64 { return s.TimeStamp }

func main() {
	SensorNw := network.NetworkWorker{Address: "95.163.237.76:5123", PackageSize: 15}
	CoordsNw := network.NetworkWorker{Address: "95.163.237.76:5124", PackageSize: 21}
	SensorByteChan := make(chan []byte, 100)
	CoordsByteChan := make(chan []byte, 100)
	TextChan := make(chan string, 200)

	go SensorNw.Run(SensorByteChan)
	go CoordsNw.Run(CoordsByteChan)
	go network.Format[SensorData](SensorByteChan, TextChan)
	go network.Format[CoordsData](CoordsByteChan, TextChan)

	f, err := os.OpenFile(fmt.Sprintf("%dlog.txt", time.Now().UnixNano()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer f.Close()

	for msg := range TextChan {
		f.WriteString(msg + "\n")
		f.Sync()
	}
}
