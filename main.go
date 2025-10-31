package main

type ReadWriter interface {
	~string | ~[]rune | ~[]byte

	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type StringReadWriter string

func (s StringReadWriter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (s StringReadWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

type BytesReadWriter []byte

func (s BytesReadWriter) Read(p []byte) (n int, err error) {

	return 0, nil
}

func (s BytesReadWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func main() {

}
