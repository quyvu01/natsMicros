package messageHelper

import (
	"reflect"
	"strings"
)

func GetMessageExchange[TMessage any]() string {
	var message TMessage
	t := reflect.TypeOf(&message).Elem()
	return strings.Replace(t.PkgPath()+"."+t.Name(), "/", "_", -1)
}
