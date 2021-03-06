//  Copyright (c) 2013 Couchbase, Inc.

package builtin

import "fmt"

import "github.com/prataprc/monster/common"

var _ = fmt.Sprintf("dummy")

// Inc will increment a variable.
// args[0] - variable name
// args[1] - quantum of value to increment
// if variable name is present in local scope it will be used,
// otherwise variable name from global scope is used.
func Inc(scope common.Scope, args ...interface{}) interface{} {
	name, by := args[0].(string), int64(1)
	if len(args) > 1 {
		by = args[1].(int64)
	}
	value, g, ok := scope.Get(name)
	if ok {
		scope.Set(name, value.(int64)+by, g)
	}
	return ""
}
