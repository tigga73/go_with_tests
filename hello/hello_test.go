package main

import "testing"

func TestOla(t *testing.T) {
	result := Hello("Tiggas")
	expected := "Hello Tiggas"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
