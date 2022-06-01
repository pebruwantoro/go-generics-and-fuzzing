package main

import "fmt"

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

		// optional to call generics function
		// Using type argument ==> SumIntsOrFloats[string, int64](ints)
   		// or without type argument ==> SumIntsOrFloats(floats)
		/* we are able to call the function without type argument, 
		but it is not always possible. So we should to declare a type constraint
		
		*/

	// calling the function use type argument
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// calling function without type argument
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))

	// calling generic function with constraint
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	Print([]int{1, 2, 3, 4})
	Print([]string{"one", "two", "three", "four"})

	
	intValues := []int64{2, 3, 15, 27, 11, 13}
	fmt.Println("Int values max: ", Max(intValues))
	floatValues := []float64{2.1, 3.2, 15.4, 27.9, 11.1, 13.2}
	fmt.Println("Float values max: ", Max(floatValues))

	m := &mechine[int64, float64]{Data: 100}
	fmt.Println(m.Upgrade(9))

}

//type declare
type Number interface {
    int64 | float64
}

// ====================Common function without golang generics=======================================
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// ============================== End ====================================


// ============================== Function with golang generics ===========================================

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// generic function using type declare
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// function is using any as parameters
func Print[T any] (s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Max[T Number](array []T) T {
	var result T = 0
	for key, _ := range array {
		if array[key] > result {
			result = array[key]
		}
	}
	return result
}

// method use generics
type Upgrader[X int64, Y float64] interface {
	Upgrade(X) Y
}

type mechine[X int64, Y float64] struct {
	Data X
}

func (m mechine[X, Y]) Upgrade(in X) Y {
	return 100.00
}