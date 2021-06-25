package ioc

import (
	"ioc-demo/sample"
	"testing"
)

func TestIoc(t *testing.T) {
	container := NewContainer()
	target := sample.NewTarget()
	container.BindInstance("a", sample.NewA("ABC"))
	container.Injection(target)
	if target.ContentOfA() != "ABC" {
		t.Error("error: test not pass")
	}
}
