package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {

	t.Run("Adds two operands", func(t *testing.T) {
		got := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, got, expected)
	})

	t.Run("Subtracts second operand from the first", func(t *testing.T) {
		got := Subtracts(4, 2)
		expected := 2
		assertCorrectMessage(t, got, expected)
	})

	t.Run("Multiplies both operands", func(t *testing.T) {
		got := Multiplies(3, 3)
		expected := 9
		assertCorrectMessage(t, got, expected)
	})

	t.Run("Divides the numerator by the denominator.", func(t *testing.T) {
		got, err := Divides(10, 5)
		expected := 2
		if err != nil {
			t.Error(err)
		}
		assertCorrectMessage(t, got, expected)
	})

	t.Run("Divide by zero test", func(t *testing.T) {
		got, _ := Divides(10, 0)
		expected := 0
		assertCorrectMessage(t, got, expected)
	})

	t.Run("The square root of a number", func(t *testing.T) {
		got, err := Sqrt(64.0)
		expected := 8.0

		if err != nil {
			t.Error(err)
		}
		if got != expected {
			t.Errorf("got %.2f expected %.2f", got, expected)
		}

	})
}

func assertCorrectMessage(t testing.TB, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
