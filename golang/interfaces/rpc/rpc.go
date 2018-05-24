package rpc

type Msg struct {
}

type Reader interface {
	Read() (*Msg, error)
}

type Writer interface {
	Write(m Msg) error
}

type ReadWriter interface {
	Reader
	Writer
}

type Conn interface {
	ReadWriter
	Close() error
}

type Listener interface {
	Accept() (Conn, error)
	Close() error
}

type ServeMux struct {
	c Conn
}

func NewServeMux(c Conn) *ServeMux {
	return &ServeMux{c: c}
}

type Handler interface {
	Serve(Writer, *Msg)
}

func (mux *ServeMux) Handle(uid string, classMethod string, rpctype int, handler Handler) {
}

func (mux *ServeMux) HandleFunc(uid string, classMethod string, rpctype int, handler func(Writer, *Msg)) {
}
