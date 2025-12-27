package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type TimeStamp interface {
	GetTimeStamp() uint64
	GetString() string
}

func Format[T any, PT interface {
	*T
	TimeStamp
}](inCh chan []byte, outCh chan<- string) {
	for bts := range inCh {
		var data T
		ptr := PT(&data)
		reader := bytes.NewReader(bts)
		err := binary.Read(reader, binary.BigEndian, ptr)
		if err == nil {
			fmtdTime := time.UnixMilli(int64(ptr.GetTimeStamp() / 1000))
			if fmtdTime.Year() == time.Now().Year() {
				outCh <- fmt.Sprintf("%s, %s", fmtdTime.Format("2006-01-02 15:04:05.000"), ptr.GetString())
			}
		}
	}
}
