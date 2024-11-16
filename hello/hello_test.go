package main

import "testing"

func TestOla(t *testing.T) {
	checkMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say Hello to people", func(t *testing.T) {
		result := Hello("Tiggas")
		expected := "Hello, Tiggas"

		checkMessage(t, result, expected)
	})

	t.Run("say 'Hello, World' when the function receives an empty string", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World"

		checkMessage(t, result, expected)
	})
}
