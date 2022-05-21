package terminal

type Output struct {
	data []byte
}

func NewOutput(b []byte) *Output {
	return &Output{data: b}
}

func (o *Output) String() string {
	return string(o.data)
}
