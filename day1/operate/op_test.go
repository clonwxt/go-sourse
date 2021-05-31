package operate

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMath(t *testing.T) {
	a := 10
	b := 10.0
	// 不同类型是不可以计算的
	// t.Log(a + b)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
}

func TestMath1(t *testing.T) {
	a, b := 10, 20
	t.Log(a + b)
	t.Log(a - b)
	t.Log(a * b)
	t.Log(a / b)
	t.Log(a % b)
}

func TestMath2(t *testing.T) {
	var a, b uint = 10, 20
	t.Log(a + b)
	t.Log(a - b)
	t.Log(a * b)
	t.Log(a / b)
	t.Log(a % b)
}

func TestMath3(t *testing.T) {
	var a, b float64 = 10, 20
	t.Log(a + b)
	t.Log(a - b)
	t.Log(a * b)
	t.Log(a / b)
}

func TestSpecOp(t *testing.T) {
	a := 10
	a++
	a++
	a++
	a--
	a--
	a--
	t.Log(a)
}

func TestBoolExpr(t *testing.T) {
	a, b := 10, 20
	t.Log(a > b)
	t.Log(a >= b)
	t.Log(a < b)
}

func TestBoolExpr1(t *testing.T) {
	a, b := 20, 10
	if a > b {
		a, b = b, a
	}
	t.Log(a, b)
}

func TestAdd(t *testing.T) {
	age := 90
	gender := "male"

	if age > 80 && gender == "male" {
		t.Log("old boy")
	} else {
		t.Log("too young,too simple")
	}
}

func TestBit(t *testing.T) {
	var (
		a uint = 60 // 0011 1100  2^5 + 2^4 + 2^3 + 2^2
		b uint = 85 // 0101 0101  2^6 + 2^4 + 2^2 + 2^0
	)
	t.Logf("a & b = %d", a&b)   // 0001 0100     2^4 + 2^2
	t.Logf("a | b = %d", a|b)   // 0111 1101     2^6 + 2^5 + 2^4 + 2^3 + 2^2 + 2^0
	t.Logf("a ^ b = %d", a^b)   // 0110 1001     2^6 + 2^5 + 2^3 + 2^0
	t.Logf("a >> 2 = %d", a>>2) // 0000 1111     15
	t.Logf("a << 2 = %d", a<<2) // 1111 0000     240
}

func TestF(t *testing.T) {
	a := 10
	a += 5
	t.Log(a)
}

func TestPointer(t *testing.T) {
	a := 10
	b := &a
	*b += 5
	t.Log(a)
}