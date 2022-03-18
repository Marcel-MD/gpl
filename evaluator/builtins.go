package evaluator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Marcel-MD/gpl/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"pop": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 && len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1 | 2",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `pop` not supported, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array).Elements

			if len(args) == 1 {
				if len(arr) > 0 {
					return &object.Array{Elements: arr[:len(arr)-1]}
				} else {
					return newError("can't `pop` empty array")
				}
			}

			if args[1].Type() != object.INTEGER_OBJ {
				return newError("index to `pop` not supported, got %s",
					args[1].Type())
			}

			index := args[1].(*object.Integer).Value

			if index < 0 || int(index) >= len(arr) {
				return newError("cannot `pop` index out of range %d", index)
			}

			if index == 0 {
				return &object.Array{Elements: arr[1:]}
			}

			return &object.Array{Elements: append(arr[:index], arr[index+1:]...)}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 2 && len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=2 | 3",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array).Elements

			if len(args) == 2 {
				return &object.Array{Elements: append(arr, args[1])}
			}

			if args[2].Type() != object.INTEGER_OBJ {
				return newError("index to `push` must be INTEGER, got %s",
					args[2].Type())
			}
			index := args[2].(*object.Integer).Value

			if index < 0 || int(index) > len(arr) {
				return newError("cannot `push` index out of range %d", index)
			}

			if len(arr) == int(index) {
				return &object.Array{Elements: append(arr, args[1])}
			}

			newArr := append(arr[:index+1], arr[index:]...)
			newArr[index] = args[1]

			return &object.Array{Elements: newArr}
		},
	},
	"write": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect())
			}
			return NULL
		},
	},
	"writef": {
		Fn: func(args ...object.Object) object.Object {
			if args[0].Type() != object.STRING_OBJ {
				return newError("file path not string, got %s",
					args[0].Type())
			}

			str := ""
			for i := 1; i < len(args); i++ {
				str = str + fmt.Sprint(args[i].Inspect())
			}

			data := []byte(str)
			err := os.WriteFile(args[0].(*object.String).Value, data, 0644)

			if err != nil {
				return newError(err.Error())
			}

			return NULL
		},
	},
	"read": {
		Fn: func(args ...object.Object) object.Object {

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()

			return &object.String{Value: scanner.Text()}
		},
	},
	"readf": {
		Fn: func(args ...object.Object) object.Object {
			if args[0].Type() != object.STRING_OBJ {
				return newError("file path not string, got %s",
					args[0].Type())
			}

			content, err := ioutil.ReadFile(args[0].(*object.String).Value)

			if err != nil {
				return newError(err.Error())
			}

			return &object.String{Value: string(content)}
		},
	},
}
