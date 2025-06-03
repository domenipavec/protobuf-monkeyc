package main_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/domenipavec/protoc-gen-monkeyc/example"
	"github.com/domenipavec/protoc-gen-monkeyc/xtesting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

const outputSeparator = "#"

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

func TestDecode(t *testing.T) {
	testCases := []struct {
		Name    string
		Message *example.ExampleMessage
	}{
		{
			Name:    "empty",
			Message: &example.ExampleMessage{},
		},
		{
			Name: "all",
			Message: &example.ExampleMessage{
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
				Rstr: []string{"", "abcdefghijklmnopqrstuvwxyz", "haha"},
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
				Oo: &example.ExampleMessage_Oostr{
					Oostr: "one of string",
				},
				Oi32: proto.Int32(667),
				Ostr: proto.String("optional string"),
				Ogm: &example.GlobalMessage{
					G1: 668,
				},
				Olm: &example.ExampleMessage_LocalMessage{
					L1: "optional lm",
				},
			},
		},
		{
			Name: "negative",
			Message: &example.ExampleMessage{
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
				Oo: &example.ExampleMessage_Oogm{
					Oogm: &example.GlobalMessage{
						G1: -666,
					},
				},
			},
		},
	}

	runner := xtesting.MonkeyRunner{}

	monkeyParts := []string{
		"pb.i32", "pb.i64", "pb.u32", "pb.u64", "pb.s32",
		"pb.s64", "pb.f32", "pb.f64", "pb.sf32", "pb.sf64",
		"pb.fl", "pb.str.toUtf8Array()", "pb.byt", "pb.b", "pb.ge",
		"pb.le", "pb.GetGm().g1", "pb.GetLm().l1.toUtf8Array()", "pb.ri64", "pb.rf32",
		"pb.rf64", "pb.rstr", "pb.rgm.size()", "pb.rgm.size() > 0 ? pb.rgm[0].g1 : 0", "pb.rgm.size() > 1 ? pb.rgm[1].g1 : 0",
		"pb.rgm.size() > 2 ? pb.rgm[2].g1 : 0", "pb.rpi64", "pb.rpf32", "pb.rpf64",
		"pb.GetOogm().g1", "pb.GetOostr()", "pb.GetOoi32()",
		"pb.GetOi32()", "pb.GetOstr()", "pb.GetOgm().g1", "pb.GetOlm().l1",
	}

	for _, tc := range testCases {
		data, err := proto.Marshal(tc.Message)
		require.NoError(t, err)

		runner.Run(xtesting.MonkeyRunnerCase{
			InitCode: fmt.Sprintf(`
	var pb = new ExampleMessage();
	pb.Decode(%s);
`, formatByteArray(data)),
			Function: "Lang.format",
			Args:     generateFormatArgs(monkeyParts...),
			Callback: func(output string) {
				t.Run(tc.Name, func(t *testing.T) {
					t.Log(output)

					parts := strings.Split(output, outputSeparator)
					require.Len(t, parts, len(monkeyParts))

					got := &example.ExampleMessage{}
					got.I32 = xtesting.ParseInt[int32](t, parts[0])
					got.I64 = xtesting.ParseInt[int64](t, parts[1])
					got.U32 = xtesting.ParseInt[uint32](t, parts[2])
					got.U64 = xtesting.ParseInt[uint64](t, parts[3])
					got.S32 = xtesting.ParseInt[int32](t, parts[4])
					got.S64 = xtesting.ParseInt[int64](t, parts[5])
					got.F32 = xtesting.ParseInt[uint32](t, parts[6])
					got.F64 = xtesting.ParseInt[uint64](t, parts[7])
					got.Sf32 = xtesting.ParseInt[int32](t, parts[8])
					got.Sf64 = xtesting.ParseInt[int64](t, parts[9])
					got.Fl = xtesting.ParseFloat(t, parts[10])
					got.Str = string(xtesting.ParseArray(t, parts[11], xtesting.ParseByte))
					got.Byt = xtesting.ParseArray(t, parts[12], xtesting.ParseByte)
					got.B = xtesting.ParseBool(t, parts[13])
					got.Ge = xtesting.ParseInt[example.GlobalEnum](t, parts[14])
					got.Le = xtesting.ParseInt[example.ExampleMessage_LocalEnum](t, parts[15])
					gmg1 := xtesting.ParseInt[int32](t, parts[16])
					if gmg1 != 0 {
						got.Gm = &example.GlobalMessage{
							G1: gmg1,
						}
					}
					lml1 := string(xtesting.ParseArray(t, parts[17], xtesting.ParseByte))
					if lml1 != "" {
						got.Lm = &example.ExampleMessage_LocalMessage{
							L1: lml1,
						}
					}
					got.Ri64 = xtesting.ParseArray(t, parts[18], xtesting.ParseInt[int64])
					got.Rf32 = xtesting.ParseArray(t, parts[19], xtesting.ParseInt[int32])
					got.Rf64 = xtesting.ParseArray(t, parts[20], xtesting.ParseInt[int64])
					got.Rstr = xtesting.ParseArray(t, parts[21], func(t *testing.T, s string) string { return s })
					rgmn := xtesting.ParseInt[int](t, parts[22])
					for i := 0; i < rgmn; i++ {
						got.Rgm = append(got.Rgm, &example.GlobalMessage{
							G1: xtesting.ParseInt[int32](t, parts[23+i]),
						})
					}
					got.Rpi64 = xtesting.ParseArray(t, parts[26], xtesting.ParseInt[int64])
					got.Rpf32 = xtesting.ParseArray(t, parts[27], xtesting.ParseInt[int32])
					got.Rpf64 = xtesting.ParseArray(t, parts[28], xtesting.ParseInt[int64])
					oogm := xtesting.ParseInt[int32](t, parts[29])
					if oogm != 0 {
						got.Oo = &example.ExampleMessage_Oogm{
							Oogm: &example.GlobalMessage{
								G1: oogm,
							},
						}
					}
					if parts[30] != "" {
						got.Oo = &example.ExampleMessage_Oostr{
							Oostr: parts[30],
						}
					}
					ooi32 := xtesting.ParseInt[int32](t, parts[31])
					if ooi32 != 0 {
						got.Oo = &example.ExampleMessage_Ooi32{
							Ooi32: ooi32,
						}
					}
					oi32 := xtesting.ParseInt[int32](t, parts[32])
					if oi32 != 0 {
						got.Oi32 = proto.Int32(oi32)
					}
					if parts[33] != "" {
						got.Ostr = proto.String(parts[33])
					}
					ogm := xtesting.ParseInt[int32](t, parts[34])
					if ogm != 0 {
						got.Ogm = &example.GlobalMessage{
							G1: ogm,
						}
					}
					if parts[35] != "" {
						got.Olm = &example.ExampleMessage_LocalMessage{
							L1: parts[35],
						}
					}

					assert.EqualExportedValues(t, tc.Message, got)
				})
			},
		})
	}

	runner.Execute(t)
}

func formatByteArray(input []byte) string {
	strs := make([]string, len(input))
	for i := range input {
		strs[i] = strconv.Itoa(int(input[i]))
	}
	return fmt.Sprintf("[%s]b", strings.Join(strs, ", "))
}

func generateFormatArgs(vars ...string) []any {
	first := `"`
	for i := range vars {
		if i != 0 {
			first += outputSeparator
		}
		first += fmt.Sprintf("$%d$", i+1)
	}
	first += `"`
	second := fmt.Sprintf("[%s]", strings.Join(vars, ", "))
	return []any{first, second}
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
	var gm = new GlobalMessage();
	gm.g1 = 34619;
	pb.gm = gm;
	var lm = new ExampleMessage.LocalMessage();
	lm.l1 = "čmrlj";
	pb.lm = lm;
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
	pb.oogm = new GlobalMessage();
	pb.oogm.g1 = 666;
	pb.oi32 = 667;
	pb.ostr = "optional str";
	pb.ogm = new GlobalMessage();
	pb.ogm.g1 = 668;
	pb.olm = new ExampleMessage.LocalMessage();
	pb.olm.l1 = "optional lm";
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
				Oo: &example.ExampleMessage_Oogm{
					Oogm: &example.GlobalMessage{
						G1: 666,
					},
				},
				Oi32: proto.Int32(667),
				Ostr: proto.String("optional str"),
				Ogm: &example.GlobalMessage{
					G1: 668,
				},
				Olm: &example.ExampleMessage_LocalMessage{
					L1: "optional lm",
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
	var gm = new GlobalMessage();
	gm.g1 = -34619;
	pb.gm = gm;
	pb.ri64 = [0l, -1l, -17244l, -116521165211652116l];
	pb.rf32 = [0, -2, -22222];
	pb.rf64 = [0l, -2l, -75295l, -7529575295l, -752957529575295752l];
	pb.rgm = [new GlobalMessage(), new GlobalMessage(), new GlobalMessage()];
	pb.rgm[1].g1 = -68537;
	pb.rgm[2].g1 = -76793;
	pb.rpi64 = [0l, -1l, -91403l, -914039140391403914l];
	pb.rpf32 = [0, -1, -56846];
	pb.rpf64 = [0l, -1l, -12175l, -1217512175l, -121751217512175121l];
	pb.ooi32 = -666;
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
				Oo: &example.ExampleMessage_Ooi32{
					Ooi32: -666,
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
					data := xtesting.ParseArray(t, output, xtesting.ParseByte)
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

					got, err := binary.ReadUvarint(bytes.NewReader(xtesting.ParseArray(t, output, xtesting.ParseByte)))
					require.NoError(t, err)

					assert.Equal(t, tc.Want, int64(got))
				})
			},
		})
	}

	runner.Execute(t)
}

func TestDecoderVarint32(t *testing.T) {
	testCases := []struct {
		Value int32
	}{
		{
			Value: 1,
		},
		{
			Value: 127,
		},
		{
			Value: 128,
		},
		{
			Value: math.MaxInt32,
		},
		{
			Value: -1,
		},
		{
			Value: -123,
		},
		{
			Value: math.MinInt32,
		},
		{
			Value: rand.Int32(),
		},
	}

	runner := xtesting.MonkeyRunner{}

	for _, tc := range testCases {
		data := binary.AppendUvarint([]byte{}, uint64(tc.Value))
		runner.Run(xtesting.MonkeyRunnerCase{
			InitCode: fmt.Sprintf("var d = new Protobuf.Decoder(%s);", formatByteArray(data)),
			Function: "d.varint32",
			Callback: func(output string) {
				t.Run(fmt.Sprint(tc.Value), func(t *testing.T) {
					t.Log(output)

					assert.Equal(t, fmt.Sprint(tc.Value), output)
				})
			},
		})
	}

	runner.Execute(t)
}

func TestDecoderVarint64(t *testing.T) {
	testCases := []struct {
		Value int64
	}{
		{
			Value: 1,
		},
		{
			Value: 127,
		},
		{
			Value: 128,
		},
		{
			Value: math.MaxInt32,
		},
		{
			Value: math.MaxInt64,
		},
		{
			Value: -1,
		},
		{
			Value: -123,
		},
		{
			Value: math.MinInt32,
		},
		{
			Value: math.MinInt64,
		},
		{
			Value: int64(rand.Int32()),
		},
		{
			Value: int64(rand.Int64()),
		},
	}

	runner := xtesting.MonkeyRunner{}

	for _, tc := range testCases {
		data := binary.AppendUvarint([]byte{}, uint64(tc.Value))
		runner.Run(xtesting.MonkeyRunnerCase{
			InitCode: fmt.Sprintf("var d = new Protobuf.Decoder(%s);", formatByteArray(data)),
			Function: "d.varint64",
			Callback: func(output string) {
				t.Run(fmt.Sprint(tc.Value), func(t *testing.T) {
					t.Log(output)

					assert.Equal(t, fmt.Sprint(tc.Value), output)
				})
			},
		})
	}

	runner.Execute(t)
}
