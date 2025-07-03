package greet

import "testing"

func TestGreet(t *testing.T){
	tests:= []struct {
        input    string
        expected string
    }{
        {"Alice", "Hello, Alice!"},
        {"", "Hello, world!"},
        {"Bob", "Hello, Bob!"},
    }

    for _, test := range tests {
        result := Greet(test.input)
        if result != test.expected {
            t.Errorf("Greet(%s) = %s, want %s", test.input, result, test.expected)
        }
    }
}
