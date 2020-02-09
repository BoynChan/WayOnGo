package logger

/*
定义一个日志写入器
*/
//声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writerList []LogWriter
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		//将日志输入到每一个写入器中
		writer.Write(data)
	}
}
