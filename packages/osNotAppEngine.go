// +build !appengine

package packages

import (
	"os"
	"reflect"

	"github.com/mojosa-software/goscript/env"
)

func osNotAppEngine() {
	env.Packages["os"]["Getppid"] = reflect.ValueOf(os.Getppid)
}
