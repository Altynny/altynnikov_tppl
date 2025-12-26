package network

import (
	"testing"
)

type TestData struct {
	TimeStamp uint64
	Temp      float32
	Pressure  int16
}

func (s TestData) GetTimeStamp() uint64 { return s.TimeStamp }
func TestFormatBytes(t *testing.T) {
	TestByteChan := make(chan []byte)
	TextChan := make(chan string)
	go Format[TestData](TestByteChan, TextChan)
	TestByteChan <- []byte("\x00\x06F\xd6\xbcN:\x9eA\x93\\)\x03\xe9")
	expectedString := "[2025-12-26 16:32:36.453] Data: {TimeStamp:1766737956453022 Temp:18.42 Pressure:1001}"
	if fmtdString := <-TextChan; fmtdString != expectedString {
		t.Errorf("Got unexpected string %s instead of %s", fmtdString, expectedString)
	}
}
