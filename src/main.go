package main

import (
	"involutiongo/src/cache"
	"involutiongo/src/db"
	"involutiongo/src/injector"
)

func main() {
	beanFactory := injector.NewBeanFactory()
	beanFactory.Set(cache.NewRedisOps())
	beanFactory.Set(db.NewDbOps())
}
