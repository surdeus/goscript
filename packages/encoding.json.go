package packages

import (
	"encoding/json"
	"reflect"

	"github.com/mojosa-software/goscript/env"
)

func init() {
	env.Packages["encoding/json"] = map[string]reflect.Value{
		"Marshal":   reflect.ValueOf(json.Marshal),
		"Unmarshal": reflect.ValueOf(json.Unmarshal),
	}
}
