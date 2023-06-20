// +build go1.9

package packages

import (
	"reflect"
	"sync"

	"github.com/surdeus/goscript/env"
)

func syncGo19() {
	env.PackageTypes["sync"]["Map"] = reflect.TypeOf(sync.Map{})
}
