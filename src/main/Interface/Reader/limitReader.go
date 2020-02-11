package Reader

import "io"

// 实现一个LimitReader
// 接收一个io.Reader接口类型的r和字节数n
// 当读完了n个字节后,就表示收到了文件结束,不会再读
type LimitReader struct {
	r     io.Reader
	limit int
}

func (l *LimitReader) Read(p []byte) (n int, err error) {
	var read int
	//如果n为0,则直接不进行后面的过程返回
	if l.limit == 0 {
		return 0, nil
	}
	//如果接收数组的长度大于n,那么就先创建一个长度为n的数组接收,然后再将其复制
	if len(p) > l.limit {
		dst := make([]byte, l.limit)
		read, _ = l.r.Read(dst)
		copy(p, dst)
	} else {
		read, _ = l.r.Read(p)
	}
	l.limit -= read
	return read, nil
}

func NewLimitReader(r io.Reader, n int) *LimitReader {
	return &LimitReader{r: r, limit: n}
}
