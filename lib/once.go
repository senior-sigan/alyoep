package lib

type Once struct {
	called bool
}

func NewOnce() *Once {
	return &Once{called: false}
}

func (o *Once) Invoke() bool {
	if !o.called {
		o.called = true
		return true
	}
	return false
}

func (o *Once) Reset() {
	o.called = false
}
