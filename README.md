# Go IoC Demo
A simple demo for IoC(Inversion of control) using Go

## Usage Sample

For example, there is a struct name A as follow:
```go
type A struct {
	content string
}

func NewA(content string) *A {
	return &A{content}
}

```

And there is a struct name Target, which contains A, and we define the instace name in container with tag "ioc"
```go
type Target struct {
	A *A `ioc:"a"`
}

func NewTarget() *Target {
	return &Target{}
}

func (t *Target) ContentOfA() string {
	return t.A.content
}

```

Then we can initialize A in container, then inject A to target

```go
func main() {
	container := ioc.NewContainer()
	target := sample.NewTarget()
	container.BindInstance("a", sample.NewA("ABC"))
	container.Injection(target)
	fmt.Println(target.ContentOfA())  // will print ABC
}
```