package ioc

import (
	"fmt"
	"reflect"
	"sync"
)

type Container struct {
	sync.Mutex
	instances map[string]interface{}
}

// 构造函数类型
// type constructor func(param ...interface{}) interface{}

// 实例化容器
func NewContainer() *Container {
	return &Container{instances: make(map[string]interface{})}
}

// 绑定对象到容器
func (c *Container) BindInstance(name string, instance interface{}) {
	c.instances[name] = instance
}

// 获取一个实例对象
func (c *Container) GetInstance(name string) interface{} {
	return c.instances[name]
}

// 容器中的对象注入实例,
// target需要传入实例对象的指针
func (c *Container) Injection(target interface{}) (err error) {
	elementType := reflect.TypeOf(target).Elem()
	elementValue := reflect.ValueOf(target).Elem()
	for i := 0; i < elementType.NumField(); i++ {
		instanceName := elementType.Field(i).Tag.Get("ioc")
		instance := c.GetInstance(instanceName)
		if instance == nil {
			err = instnaceNotFoundErr(instanceName)
			return
		}
		elementValue.Field(i).Set(reflect.ValueOf(instance))
	}
	return
}

func instnaceNotFoundErr(name string) error {
	return fmt.Errorf("error: instance[%s] was not found in container", name)
}
