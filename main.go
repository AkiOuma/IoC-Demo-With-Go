package main

import (
	"fmt"
	"ioc-demo/ioc"
	"ioc-demo/sample"
)

func main() {
	container := ioc.NewContainer()
	target := sample.NewTarget()
	container.BindInstance("a", sample.NewA("ABC"))
	container.Injection(target)
	fmt.Println(target.ContentOfA())
}
