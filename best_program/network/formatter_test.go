package network

import (
	"fmt"
	"testing"
)

type TestData struct {
	TimeStamp uint64
	Temp      float32
	Pressure  int16
}

func (s TestData) GetTimeStamp() uint64 { return s.TimeStamp }
func (s TestData) GetString() string    { return fmt.Sprintf("%f, %d", s.Temp, s.Pressure) }
func TestFormatBytes(t *testing.T) {
	TestByteChan := make(chan []byte)
	TextChan := make(chan string)
	go Format[TestData](TestByteChan, TextChan)
	TestByteChan <- []byte("\x00\x06F\xd6\xbcN:\x9eA\x93\\)\x03\xe9")
	expectedString := "2025-12-26 16:32:36.453, 18.420000, 1001"
	if fmtdString := <-TextChan; fmtdString != expectedString {
		t.Errorf("Got unexpected string %s instead of %s", fmtdString, expectedString)
	}
}
