package gophertime

type Framer interface {
	Frame() int64
}

type Frame int64

func (f Frame) Frame() int64 {
	return int64(f)
}

func (f Frame) Equals(other Framer) bool {
	return other.Frame() == f.Frame()
}

func (f Frame) Add(other Framer) Frame {
	return Frame(int64(f) + other.Frame())
}
