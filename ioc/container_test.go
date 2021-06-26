package ioc

import (
	"ioc-demo/sample"
	"testing"
)

func TestIoc(t *testing.T) {
	container := NewContainer()
	target := sample.NewTarget()
	container.BindInstance("a", sample.NewA("ABC"))
	if err := container.Injection(target); err != nil {
		t.Error("error: test not pass")
	}
	if target.ContentOfA() != "ABC" {
		t.Error("error: test not pass")
	}
}
