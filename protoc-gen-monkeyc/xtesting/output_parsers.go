package xtesting

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func ParseByte(t *testing.T, input string) byte {
	v, err := strconv.ParseUint(input, 10, 8)
	require.NoError(t, err)
	return byte(v)
}

func ParseInt[T constraints.Integer](t *testing.T, input string) T {
	v, err := strconv.ParseInt(input, 10, 64)
	require.NoError(t, err)
	return T(v)
}

func ParseFloat(t *testing.T, input string) float32 {
	v, err := strconv.ParseFloat(input, 32)
	require.NoError(t, err)
	return float32(v)
}

func ParseBool(t *testing.T, input string) bool {
	v, err := strconv.ParseBool(input)
	require.NoError(t, err)
	return v
}

func ParseArray[T any](t *testing.T, input string, f func(*testing.T, string) T) []T {
	var result []T
	input = strings.TrimSpace(strings.Trim(input, "[]"))
	if input == "" {
		return result
	}
	parts := strings.Split(input, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		result = append(result, f(t, p))
	}
	return result
}
