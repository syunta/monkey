package evaluator

import "github.com/syunta/monkey/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			return NULL
		},
	},
}
