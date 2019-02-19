package first_class_func

import (
	"testing"
	"fmt"
)

func TestFFFunc(t *testing.T){
	takesAFunction(foo);
	var f Stringy = returnsAFunction();
	f();
	var baz Stringy = func()string{
		return "anonymous stringy\n"
	};
	fmt.Printf(baz());
}