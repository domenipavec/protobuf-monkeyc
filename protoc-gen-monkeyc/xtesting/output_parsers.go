package xtesting

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func ParseByteArray(t *testing.T, input string) []byte {
	var result []byte
	parts := strings.Split(strings.Trim(input, "[]"), ",")
	for _, p := range parts {
		v, err := strconv.ParseUint(strings.TrimSpace(p), 10, 8)
		require.NoError(t, err)
		result = append(result, byte(v))
	}
	return result
}
