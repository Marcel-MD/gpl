package evaluator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Marcel-MD/gpl/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments to `len`. got=%d, want=1",
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
				return newError("wrong number of arguments to `pop`. got=%d, want=1 | 2",
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
				return newError("wrong number of arguments to `push`. got=%d, want=2 | 3",
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
	"rand": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments to `rand`. got=%d, want=1",
					len(args))
			}

			seed := rand.NewSource(time.Now().UnixNano())
			r := rand.New(seed)

			switch arg := args[0].(type) {
			case *object.Integer:
				if arg.Value < 1 {
					return newError("argument to `rand` should be bigger equal or greater than 1")
				}
				return &object.Integer{Value: r.Int63n(arg.Value)}
			case *object.Float:
				if arg.Value < 1 {
					return newError("argument to `rand` should be bigger equal or greater than 1")
				}
				return &object.Float{Value: r.Float64() * arg.Value}
			default:
				return newError("argument to `rand` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments to `int`. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				i, err := strconv.ParseInt(arg.Value, 0, 64)
				if err != nil {
					return newError("bad argument to `int`, got %s", arg.Value)
				}
				return &object.Integer{Value: i}
			case *object.Float:
				return &object.Integer{Value: int64(arg.Value)}
			case *object.Boolean:
				i := 0
				if arg.Value {
					i = 1
				}
				return &object.Integer{Value: int64(i)}
			default:
				return newError("argument to `int` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"float": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments to `float`. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				f, err := strconv.ParseFloat(arg.Value, 64)
				if err != nil {
					return newError("bad argument to `float`, got %s", arg.Value)
				}
				return &object.Float{Value: f}
			case *object.Integer:
				return &object.Float{Value: float64(arg.Value)}
			default:
				return newError("argument to `float` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"str": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments to `str`. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Integer:
				return &object.String{Value: fmt.Sprintf("%d", arg.Value)}
			case *object.Float:
				return &object.String{Value: fmt.Sprintf("%f", arg.Value)}
			case *object.Boolean:
				s := "false"
				if arg.Value {
					s = "true"
				}
				return &object.String{Value: s}
			default:
				return newError("argument to `str` not supported, got %s",
					args[0].Type())
			}
		},
	},
}
