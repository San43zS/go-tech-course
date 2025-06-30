package limitRead

import (
	"io"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("limitReader")

type LimitRead struct {
	r io.Reader
	n int64
}

func New(r io.Reader, n int64) IReaderLimit {
	return &LimitRead{
		r: r,
		n: n,
	}
}

func (l *LimitRead) Read(buffer []byte) (int, error) {
	if l.n <= 0 {
		log.Criticalf("EOF")
		return 0, io.EOF
	}

	if int64(len(buffer)) > l.n {
		buffer = buffer[:l.n]
	}

	n, err := l.r.Read(buffer)

	l.n -= int64(n)

	return n, err
}
