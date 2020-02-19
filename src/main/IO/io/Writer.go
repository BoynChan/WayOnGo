package io

import "io"

// Author:Boyn
// Date:2020/2/19

func WriteAt(writer io.WriterAt, offset int, content []byte) (at int, err error) {
	at, err = writer.WriteAt(content, int64(offset))
	return
}

func WriteTo(reader io.WriterTo, writer io.Writer) (int, error) {
	// 传入一个实现了WriteTo方法的接口和writer
	// WriteTo接口可以方便地
	to, err := reader.WriteTo(writer)
	return int(to), err
}
