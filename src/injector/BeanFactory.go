package injector

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper)}
}

func (this *BeanFactoryImpl) Set(vlist ...interface{}) {
	if vlist == nil || len(vlist) == 0 {
		return
	}
	for _, v := range vlist {
		this.beanMapper.add(v)
	}
}

func (this *BeanFactoryImpl) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	getV := this.beanMapper.get(v)
	if getV.IsValid() {
		return getV.Interface()
	}
	return nil
}
