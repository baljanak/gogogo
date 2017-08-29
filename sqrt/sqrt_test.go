package sqrt_test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/baljanak/gogogo/sqrt"
)

func ExamplePositiveRealSqrt() {
	fmt.Println(sqrt.RealSqrt(2))
	// Output:
	// 1.4142135623730951
}

func TestNegativeRealSqrt(t *testing.T) {

	f := func() {
		sqrt.RealSqrt(-2) // Code should panic
	}

	assertPanic(t, f)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("fn %v did not panic.", GetFunctionName(f))
		} else {
			t.Logf("code panic as expected: %v", r)
		}
	}()
	f()
}
