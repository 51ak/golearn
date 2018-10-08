package day2

import (
	"code.google.com/p/go-tour/reader"
)

type MyReader struct{}

/*实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。*/
// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	i := 0
	for ; i < len(b); i++ {
		b[i] = 'A'
	}
	return i, nil

}

func Test07() {
	reader.Validate(MyReader{})
}
