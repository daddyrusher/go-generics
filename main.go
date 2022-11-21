package main

import "fmt"

// MyOwnInt is derived type
type MyOwnInt int

// Results is a slice of generic type Subtractable
type Results[T Subtractable] []T

// Subtractable is a constraint generic type of ints and floats
type Subtractable interface {
	~int | int32 | int64 | float32 | float64
}

type Movable[S Subtractable] interface {
	Move(S)
}

func Move[V Movable[S], S Subtractable](v V, distance, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}

type Person[S Subtractable] struct {
	Name string
}

func (p Person[S]) Move(meters S) {
	fmt.Printf("%s moved %d meters\n", p.Name, meters)
}

type Car[S Subtractable] struct {
	Name string
}

func (c Car[S]) Move(meters S) {
	fmt.Printf("%s moved %d meters\n", c.Name, meters)
}

func main() {
	var a MyOwnInt = 10
	var b MyOwnInt = 20

	var c float32 = 40.1
	var d float32 = 50.6

	var e int32 = 35
	var f int32 = 71

	//subtract ints
	intResult := Subtract(b, a)
	//subtract floats
	floatResult := Subtract(d, c)
	//subtract int32 values
	int32Result := Subtract(f, e)

	// generic slice of ints, because Subtractable is used for type params only
	var resultStorage Results[MyOwnInt]

	resultStorage = append(resultStorage, intResult)

	fmt.Println("int result:", intResult)
	fmt.Println("float result:", floatResult)
	fmt.Println("int32 result:", int32Result)

	var p = Person[int]{Name: "John"}
	var car = Car[float32]{Name: "BMW"}

	Move(p, 1000, 100)
	Move[Car[float32], float32](car, 10000.500, 100)
}

func Subtract[V Subtractable](a, b V) V {
	return a - b
}
