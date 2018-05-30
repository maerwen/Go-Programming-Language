package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func main() {
	testType()
	// testValue()
	// fmt.Println(formatAtom(reflect.ValueOf(map[string]string{"a": "a"})))
}
func test() {
	fmt.Println(reflect.TypeOf(nil))  //<nil>
	fmt.Println(reflect.ValueOf(nil)) //<invalid reflect.Value>
}
func testType() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String()) //int
	fmt.Println(t)          //int

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) //*os.File

	fmt.Printf("%T\n", 3) //int	%T后台进行了typeOf操作
	fmt.Printf("%v\n", 3)
}
func testValue() {
	v := reflect.ValueOf(3)
	fmt.Println(v)          //3
	fmt.Printf("%v\n", v)   //3
	fmt.Println(v.String()) //<int Value>
	t := v.Type()
	fmt.Println(t.String()) //int
	// valueOf逆操作
	v = reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i) //3
}
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: //	reflect.Array,reflect.Struct,reflect.Interface
		return v.Type().String() + "value"
	}
}
