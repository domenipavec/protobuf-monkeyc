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

	"github.com/domenipavec/protoc-gen-monkeyc/example"
	"github.com/domenipavec/protoc-gen-monkeyc/xtesting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestMain(m *testing.M) {
	err := xtesting.StartSimulator()
	if err != nil {
		log.Fatal(err)
	}
	err = xtesting.RebuildProto()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestEncode(t *testing.T) {
	testCases := []struct {
		Name    string
		Setters string
		Want    *example.ExampleMessage
	}{
		{
			Name: "empty",
			Want: &example.ExampleMessage{},
		},
		{
			Name: "all",
			Setters: `
	pb.i32 = 12345;
	pb.i64 = 1234567890123456789l;
	pb.u32 = 67890;
	pb.u64 = 987654321098765432l;
	pb.s32 = 92590;
	pb.s64 = 925909259092590925l;
	pb.f32 = 35604;
	pb.f64 = 356043560435604356l;
	pb.sf32 = 83987;
	pb.sf64 = 839878398783987839l;
	pb.fl = 123.456789;
	pb.str = "češnja";
	pb.byt = [123, 78, 93, 0, 1]b;
	pb.b = true;
	pb.ge = B;
	pb.le = ExampleMessage.LB;
	pb.gm.g1 = 34619;
	pb.lm.l1 = "čmrlj";
	pb.ri64 = [0l, 1l, 17244l, 116521165211652116l];
	pb.rf32 = [0, 2, 22222];
	pb.rf64 = [0l, 2l, 75295l, 7529575295l, 752957529575295752l];
	pb.rstr = ["", "abcdefghijklmnopqrstuvwxyz", "čšž", "", "haha"];
	pb.rgm = [new GlobalMessage(), new GlobalMessage(), new GlobalMessage()];
	pb.rgm[1].g1 = 68537;
	pb.rgm[2].g1 = 76793;
	pb.rpi64 = [0l, 1l, 91403l, 914039140391403914l];
	pb.rpf32 = [0, 1, 56846];
	pb.rpf64 = [0l, 1l, 12175l, 1217512175l, 121751217512175121l];
`,
			Want: &example.ExampleMessage{
				I32:  12345,
				I64:  1234567890123456789,
				U32:  67890,
				U64:  987654321098765432,
				S32:  92590,
				S64:  925909259092590925,
				F32:  35604,
				F64:  356043560435604356,
				Sf32: 83987,
				Sf64: 839878398783987839,
				Fl:   123.456789,
				Str:  "češnja",
				Byt:  []byte{123, 78, 93, 0, 1},
				B:    true,
				Ge:   example.GlobalEnum_B,
				Le:   example.ExampleMessage_LB,
				Gm: &example.GlobalMessage{
					G1: 34619,
				},
				Lm: &example.ExampleMessage_LocalMessage{
					L1: "čmrlj",
				},
				Ri64: []int64{
					0,
					1,
					17244,
					116521165211652116,
				},
				Rf32: []int32{
					0,
					2,
					22222,
				},
				Rf64: []int64{
					0,
					2,
					75295,
					7529575295,
					752957529575295752,
				},
				Rstr: []string{"", "abcdefghijklmnopqrstuvwxyz", "čšž", "", "haha"},
				Rgm: []*example.GlobalMessage{
					{},
					{
						G1: 68537,
					},
					{
						G1: 76793,
					},
				},
				Rpi64: []int64{
					0,
					1,
					91403,
					914039140391403914,
				},
				Rpf32: []int32{
					0,
					1,
					56846,
				},
				Rpf64: []int64{
					0,
					1,
					12175,
					1217512175,
					121751217512175121,
				},
			},
		},
		{
			Name: "negative",
			Setters: `
	pb.i32 = -12345;
	pb.i64 = -1234567890123456789l;
	pb.s32 = -92590;
	pb.s64 = -925909259092590925l;
	pb.sf32 = -83987;
	pb.sf64 = -839878398783987839l;
	pb.fl = -123.456789;
	pb.gm.g1 = -34619;
	pb.ri64 = [0l, -1l, -17244l, -116521165211652116l];
	pb.rf32 = [0, -2, -22222];
	pb.rf64 = [0l, -2l, -75295l, -7529575295l, -752957529575295752l];
	pb.rgm = [new GlobalMessage(), new GlobalMessage(), new GlobalMessage()];
	pb.rgm[1].g1 = -68537;
	pb.rgm[2].g1 = -76793;
	pb.rpi64 = [0l, -1l, -91403l, -914039140391403914l];
	pb.rpf32 = [0, -1, -56846];
	pb.rpf64 = [0l, -1l, -12175l, -1217512175l, -121751217512175121l];
`,
			Want: &example.ExampleMessage{
				I32:  -12345,
				I64:  -1234567890123456789,
				S32:  -92590,
				S64:  -925909259092590925,
				Sf32: -83987,
				Sf64: -839878398783987839,
				Fl:   -123.456789,
				Gm: &example.GlobalMessage{
					G1: -34619,
				},
				Ri64: []int64{
					0,
					-1,
					-17244,
					-116521165211652116,
				},
				Rf32: []int32{
					0,
					-2,
					-22222,
				},
				Rf64: []int64{
					0,
					-2,
					-75295,
					-7529575295,
					-752957529575295752,
				},
				Rgm: []*example.GlobalMessage{
					{},
					{
						G1: -68537,
					},
					{
						G1: -76793,
					},
				},
				Rpi64: []int64{
					0,
					-1,
					-91403,
					-914039140391403914,
				},
				Rpf32: []int32{
					0,
					-1,
					-56846,
				},
				Rpf64: []int64{
					0,
					-1,
					-12175,
					-1217512175,
					-121751217512175121,
				},
			},
		},
	}

	runner := xtesting.MonkeyRunner{}

	for _, tc := range testCases {
		runner.Run(xtesting.MonkeyRunnerCase{
			InitCode: "var pb = new ExampleMessage();\n" + tc.Setters,
			Function: "pb.Encode",
			Callback: func(output string) {
				t.Run(tc.Name, func(t *testing.T) {
					t.Log(output)

					got := &example.ExampleMessage{}
					data := xtesting.ParseByteArray(t, output)
					err := proto.Unmarshal(data, got)
					require.NoError(t, err)

					assert.EqualExportedValues(t, tc.Want, got)

					m, err := proto.Marshal(tc.Want)
					require.NoError(t, err)

					assert.Len(t, data, len(m))
				})
			},
		})
	}

	runner.Execute(t)
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
