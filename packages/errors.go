package packages

import (
	"errors"
	"reflect"

	"github.com/surdeus/goscript/env"
)

func init() {
	env.Packages["errors"] = map[string]reflect.Value{
		"New": reflect.ValueOf(errors.New),
	}
}
