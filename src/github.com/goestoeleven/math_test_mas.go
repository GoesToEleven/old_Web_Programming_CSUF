package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

/*
It's a good idea to test many different combinations of possibilities

This is a very common way to setup tests (abundant examples can be found
in the source code for the packages included with Go). We create a struct
to represent the inputs and outputs for the function. Then we create a list
of these structs (pairs). Then we loop through each one and run the function.

*/
