package limitRead

type IReaderLimit interface {
	Read(buffer []byte) (n int, err error)
}
