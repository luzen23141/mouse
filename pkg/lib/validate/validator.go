package validate

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	if v.kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *DefaultValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("validate")

		// add any custom validations etc. here
	})
}

func (v *DefaultValidator) kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
