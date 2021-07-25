package main

import "involutiongo/src/controller"

func main() {
	//beanFactory := injector.NewBeanFactory()
	//beanFactory.Set(cache.NewRedisOps())
	//beanFactory.Set(db.NewDbOps())

	controller.Start()
}
