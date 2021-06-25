package sample

type A struct {
	content string
}

func NewA(content string) *A {
	return &A{content}
}
