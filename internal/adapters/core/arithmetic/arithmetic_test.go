package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()
	a, b := 1, 1
	answer, err := arith.Addition(a, b)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.NoError(t, err)
	require.Equal(t, 2, answer)
}
func TestSubtraction(t *testing.T) {
	arith := NewAdapter()
	a, b := 1, 1
	answer, err := arith.Subtraction(a, b)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.NoError(t, err)
	require.Equal(t, 0, answer)
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()
	a, b := 1, 1
	answer, err := arith.Multiplication(a, b)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.NoError(t, err)
	require.Equal(t, 1, answer)
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()
	a, b := 1, 1
	answer, err := arith.Division(a, b)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.NoError(t, err)
	require.Equal(t, 1, answer)
}
