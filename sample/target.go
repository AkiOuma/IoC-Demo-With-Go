package sample

type Target struct {
	A *A `ioc:"a"`
}

func NewTarget() *Target {
	return &Target{}
}

func (t *Target) ContentOfA() string {
	return t.A.content
}
