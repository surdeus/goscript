package packages

import (
	"errors"
	"reflect"

	"github.com/mojosa-software/goscript/env"
)

func init() {
	env.Packages["errors"] = map[string]reflect.Value{
		"New": reflect.ValueOf(errors.New),
	}
}
