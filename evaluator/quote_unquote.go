package evaluator

import  (
	"github.com/syunta/monkey/ast"
	"github.com/syunta/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
