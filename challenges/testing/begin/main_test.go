// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		name string
		input string
		want int
	}{
		{"#00", "#00", 0},
		{"a2_32_^_&/)", "a2_32_^_&/)", 1},
		{"812_%6ab//", "812_%6ab//", 2},
	}

	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T)  {
			got := lc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = []struct {
		name string
		input string
		want int
	}{
		{"#00", "#00", 2},
		{"abc_1,?/", "abc_1,?/", 1},
		{"abc_12&8_^", "abc_12&8_^", 3},
	}

	nc := numberCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T)  {
			got := nc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCounter(t *testing.T) {
	var tests = []struct {
		name string
		input string
		want int
	}{
		{"#00", "#00", 1},
		{"abc_1,?/", "abc_1,?/", 4},
		{"abc_12&8_^", "abc_12&8_^", 4},
	}

	sc := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T)  {
			got := sc.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
