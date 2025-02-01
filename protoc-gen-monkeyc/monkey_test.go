package main_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"testing"

	"github.com/domenipavec/protoc-gen-monkeyc/xtesting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	err := xtesting.StartSimulator()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestEncode(t *testing.T) {

}

func TestEncodeVarint(t *testing.T) {
	rand64 := rand.Int64()
	rand32 := rand.Int32()

	testCases := []struct {
		Input any
		Want  int64
	}{
		{
			Input: "true",
			Want:  1,
		},
		{
			Input: "false",
			Want:  0,
		},
		{
			Input: 1,
			Want:  1,
		},
		{
			Input: 127,
			Want:  127,
		},
		{
			Input: 128,
			Want:  128,
		},
		{
			Input: math.MaxInt32,
			Want:  math.MaxInt32,
		},
		{
			Input: fmt.Sprintf("%dl", math.MaxInt64),
			Want:  math.MaxInt64,
		},
		{
			Input: -1,
			Want:  -1,
		},
		{
			Input: "-1l",
			Want:  -1,
		},
		{
			Input: math.MinInt32,
			Want:  math.MinInt32,
		},
		{
			Input: fmt.Sprintf("%dl", math.MinInt64),
			Want:  math.MinInt64,
		},
		{
			Input: fmt.Sprintf("%dl", rand64),
			Want:  rand64,
		},
		{
			Input: rand32,
			Want:  int64(rand32),
		},
	}

	runner := xtesting.MonkeyRunner{}

	for _, tc := range testCases {
		runner.Run(xtesting.MonkeyRunnerCase{
			Function: "Protobuf.encodeVarint",
			Args:     []any{tc.Input},
			Callback: func(output string) {
				t.Run(fmt.Sprint(tc.Input), func(t *testing.T) {
					t.Log(output)

					got, err := binary.ReadUvarint(bytes.NewReader(xtesting.ParseByteArray(t, output)))
					require.NoError(t, err)

					assert.Equal(t, tc.Want, int64(got))
				})
			},
		})
	}

	runner.Execute(t)
}
