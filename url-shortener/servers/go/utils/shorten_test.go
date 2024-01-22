package utils

import (
	"reflect"
	"testing"
)

func TestGenerateShortKey(t *testing.T) {
	shortKey := GenerateShortKey()

	if reflect.TypeOf(shortKey).Kind() != reflect.String || len(shortKey) != 6 {
		t.Error("invalid short key algorithm.")
	}
}
