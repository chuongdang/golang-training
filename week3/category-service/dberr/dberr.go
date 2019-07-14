package dberr

var ErrChannel chan bool

func init() {
	ErrChannel = make(chan bool)
}