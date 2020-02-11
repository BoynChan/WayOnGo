package Reader

type StrReader struct {
	s     string //包含的字符串
	index int    //当前读到的位置
}

func NewStrReader(s string) *StrReader {
	return &StrReader{s: s}
}

//实现一个简单的Reader接口
func (s *StrReader) Read(p []byte) (n int, err error) {
	if s.index >= len(s.s) {
		return 0, nil
	}
	n = copy(p, s.s[s.index:])
	s.index += n
	return n, nil
}
