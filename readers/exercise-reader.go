package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 为 MyReader 添加一个 Read([]byte) (int, error) 方法。
// Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
