package assert

import "testing"

// Error asserts that the given parameter is error.
func Error(t testing.TB, a error) {
	Equal(t, a != nil, true)
}

// NoError asserts that the given parameter is no error.
func NoError(t testing.TB, a error) {
	Equal(t, a == nil, true)
}
