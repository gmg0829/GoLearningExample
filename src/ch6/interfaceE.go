package main

import (
    "fmt"
)

// 定义一个数据写入器
type DataWriter interface {
    WriteData(data string) error
}

type DataWriterA struct {
}

type DataWriterB struct {
}

func (d *DataWriterA) WriteData(data string) error {
    fmt.Println("DataWriterA:", data)
    return nil
}

func (d *DataWriterB) WriteData(data string) error {
    fmt.Println("DataWriterB:", data)
    return nil
}

func main() {
	var a DataWriter = new(DataWriterA)
	a.WriteData("gmg")

	var b DataWriter = new(DataWriterB)
	b.WriteData("gmg")
}