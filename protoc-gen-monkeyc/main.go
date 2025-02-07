package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pborman/indent"
	"github.com/stoewer/go-strcase"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const tab = "    "

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if len(os.Args) > 1 {
		return fmt.Errorf("unknown argument %q (this program should be run by protoc, not directly)", os.Args[1])
	}
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		return err
	}

	resp := &pluginpb.CodeGeneratorResponse{}
	for _, fn := range req.FileToGenerate {
		proto := findProto(req.ProtoFile, fn)
		result, err := generateFile(fn, proto)
		if err != nil {
			resp.Error = str(err.Error())
			break
		}
		resp.File = append(resp.File, result)
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, bytes.NewReader(data))
	if err != nil {
		return err
	}

	return nil
}

func findProto(protos []*descriptorpb.FileDescriptorProto, fn string) *descriptorpb.FileDescriptorProto {
	for _, proto := range protos {
		if proto.GetName() == fn {
			return proto
		}
	}
	return nil
}

func generateFile(fn string, proto *descriptorpb.FileDescriptorProto) (*pluginpb.CodeGeneratorResponse_File, error) {
	buf := &bytes.Buffer{}

	fmt.Fprintln(buf, "// Code generated by protoc-gen-monkeyc. DO NOT EDIT.")
	fmt.Fprintln(buf)

	fmt.Fprintln(buf, "import Toybox.Lang;")

	if len(proto.Dependency) > 0 {
		return nil, errors.New("import is not supported yet")
	}

	for _, msg := range proto.MessageType {
		fmt.Fprintln(buf)
		err := generateMessage(buf, msg)
		if err != nil {
			return nil, err
		}
	}

	for _, e := range proto.EnumType {
		fmt.Fprintln(buf)
		err := generateEnum(buf, e)
		if err != nil {
			return nil, err
		}
	}

	return &pluginpb.CodeGeneratorResponse_File{
		Name:    str(strings.TrimSuffix(fn, ".proto") + ".pb.mc"),
		Content: str(buf.String()),
	}, nil
}

func generateMessage(w io.Writer, msg *descriptorpb.DescriptorProto) error {
	fmt.Fprintf(w, "class %s {\n", *msg.Name)
	ind := indent.New(w, tab)

	for _, nested := range msg.NestedType {
		err := generateMessage(ind, nested)
		if err != nil {
			return err
		}
		fmt.Fprintln(w)
	}
	for _, e := range msg.EnumType {
		err := generateEnum(ind, e)
		if err != nil {
			return err
		}
		fmt.Fprintln(w)
	}

	for _, field := range msg.Field {
		var typ string
		switch *field.Type {
		case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
			typ = "Float"
		case descriptorpb.FieldDescriptorProto_TYPE_INT32,
			descriptorpb.FieldDescriptorProto_TYPE_UINT32,
			descriptorpb.FieldDescriptorProto_TYPE_SINT32,
			descriptorpb.FieldDescriptorProto_TYPE_FIXED32,
			descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
			typ = "Number"
		case descriptorpb.FieldDescriptorProto_TYPE_INT64,
			descriptorpb.FieldDescriptorProto_TYPE_UINT64,
			descriptorpb.FieldDescriptorProto_TYPE_SINT64,
			descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
			descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
			typ = "Long"
		case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
			typ = "Boolean"
		case descriptorpb.FieldDescriptorProto_TYPE_STRING:
			typ = "String"
		case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
			typ = "ByteArray"
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			typ = typeNameToMonkey(*field.TypeName)
		default:
			return fmt.Errorf("unsupported type: %s", *field.Type)
		}
		if *field.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
			typ = fmt.Sprintf("Array<%s>", typ)
		}
		fmt.Fprintf(ind, "public var %s as %s;\n", strcase.LowerCamelCase(*field.Name), typ)
	}
	fmt.Fprintln(ind)

	err := generateMessageInitialize(ind, msg)
	if err != nil {
		return err
	}
	fmt.Fprintln(w)

	err = generateMessageEncode(ind, msg)
	if err != nil {
		return err
	}

	err = generateMessageDecode(ind, msg)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "}")

	return nil
}

func typeNameToMonkey(typeName string) string {
	parts := strings.Split(typeName, ".")
	return parts[len(parts)-1]
}

func generateMessageInitialize(w io.Writer, msg *descriptorpb.DescriptorProto) error {
	fmt.Fprintln(w, "public function initialize() {")
	ind := indent.New(w, tab)

	for _, field := range msg.Field {
		var zero string
		switch *field.Type {
		case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
			zero = "0.0"
		case descriptorpb.FieldDescriptorProto_TYPE_INT32,
			descriptorpb.FieldDescriptorProto_TYPE_UINT32,
			descriptorpb.FieldDescriptorProto_TYPE_SINT32,
			descriptorpb.FieldDescriptorProto_TYPE_FIXED32,
			descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
			zero = "0"
		case descriptorpb.FieldDescriptorProto_TYPE_INT64,
			descriptorpb.FieldDescriptorProto_TYPE_UINT64,
			descriptorpb.FieldDescriptorProto_TYPE_SINT64,
			descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
			descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
			zero = "0l"
		case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
			zero = "false"
		case descriptorpb.FieldDescriptorProto_TYPE_STRING:
			zero = `""`
		case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
			zero = `[]b`
		case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			zero = fmt.Sprintf("0 as %s", typeNameToMonkey(*field.TypeName))
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
			zero = fmt.Sprintf("new %s()", typeNameToMonkey(*field.TypeName))
		default:
			return fmt.Errorf("unsupported type: %s", *field.Type)
		}
		if *field.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
			zero = "[]"
		}
		fmt.Fprintf(ind, "%s = %s;\n", strcase.LowerCamelCase(*field.Name), zero)
	}

	fmt.Fprintln(w, "}")

	return nil
}

func generateMessageEncode(w io.Writer, msg *descriptorpb.DescriptorProto) error {
	fmt.Fprintln(w, "public function Encode() as ByteArray {")
	ind := indent.New(w, tab)

	fmt.Fprintln(ind, "var result = []b;")

	for _, field := range msg.Field {
		var encoder string
		switch *field.Type {
		case descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
			descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
			encoder = "Protobuf.encodeField64(%d, %s, %t)"
		case descriptorpb.FieldDescriptorProto_TYPE_FLOAT,
			descriptorpb.FieldDescriptorProto_TYPE_FIXED32,
			descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
			encoder = "Protobuf.encodeField32(%d, %s, %t)"
		case descriptorpb.FieldDescriptorProto_TYPE_INT32,
			descriptorpb.FieldDescriptorProto_TYPE_INT64,
			descriptorpb.FieldDescriptorProto_TYPE_UINT64,
			descriptorpb.FieldDescriptorProto_TYPE_UINT32,
			descriptorpb.FieldDescriptorProto_TYPE_BOOL,
			descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			encoder = "Protobuf.encodeFieldVarint(%d, %s, %t)"
		case descriptorpb.FieldDescriptorProto_TYPE_SINT32,
			descriptorpb.FieldDescriptorProto_TYPE_SINT64:
			encoder = "Protobuf.encodeFieldVarint(%d, Protobuf.toSignedInt(%s), %t)"
		case descriptorpb.FieldDescriptorProto_TYPE_STRING,
			descriptorpb.FieldDescriptorProto_TYPE_BYTES:
			encoder = "Protobuf.encodeFieldLen(%d, %s, %t)"
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
			encoder = "Protobuf.encodeFieldLen(%d, %s.Encode(), %t)"
		default:
			return fmt.Errorf("unsupported type: %s", *field.Type)
		}
		if *field.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
			if isPacked(field) {
				fmt.Fprintln(ind, "{")
				ind2 := indent.New(ind, tab)
				fmt.Fprintln(ind2, "var packed = []b;")
				fmt.Fprintf(ind2, "for (var i = 0; i < %s.size(); i++) {\n", strcase.LowerCamelCase(*field.Name))
				fmt.Fprintf(ind2, tab+"packed.addAll(%s);\n", fmt.Sprintf(encoder, 0, strcase.LowerCamelCase(*field.Name)+"[i]", true))
				fmt.Fprintln(ind2, "}")
				fmt.Fprintf(ind2, "result.addAll(Protobuf.encodeFieldLen(%d, packed, false));\n", *field.Number)
				fmt.Fprintln(ind, "}")
			} else {
				fmt.Fprintf(ind, "for (var i = 0; i < %s.size(); i++) {\n", strcase.LowerCamelCase(*field.Name))
				fmt.Fprintf(ind, tab+"result.addAll(%s);\n", fmt.Sprintf(encoder, *field.Number, strcase.LowerCamelCase(*field.Name)+"[i]", true))
				fmt.Fprintln(ind, "}")
			}
		} else {
			fmt.Fprintf(ind, "result.addAll(%s);\n", fmt.Sprintf(encoder, *field.Number, strcase.LowerCamelCase(*field.Name), false))
		}
	}

	fmt.Fprintln(ind, "return result;")
	fmt.Fprintln(w, "}")

	return nil
}

func generateMessageDecode(w io.Writer, msg *descriptorpb.DescriptorProto) error {
	fmt.Fprintln(w, "public function Decode(input as ByteArray) as Void {")
	ind := indent.New(w, tab)
	fmt.Fprintln(ind, "var d = new Protobuf.Decoder(input);")
	fmt.Fprintln(ind, "while (d.remaining() > 0) {")
	loop := indent.New(ind, tab)
	fmt.Fprintln(loop, "var tag = d.varint32();")
	fmt.Fprintln(loop, "switch (tag >> 3) {")
	sw := indent.New(loop, tab)

	for _, field := range msg.Field {
		fmt.Fprintf(sw, "case %d: {\n", *field.Number)
		cs := indent.New(sw, tab)

		var decoder, wireType string
		switch *field.Type {
		case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
			decoder = "d.varint32() != 0"
			wireType = "VARINT"
		case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			decoder = fmt.Sprintf("d.varint32() as %s", typeNameToMonkey(*field.TypeName))
			wireType = "VARINT"
		case descriptorpb.FieldDescriptorProto_TYPE_INT32,
			descriptorpb.FieldDescriptorProto_TYPE_UINT32:
			decoder = "d.varint32()"
			wireType = "VARINT"
		case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
			decoder = "Protobuf.fromSignedNumber(d.varint32())"
			wireType = "VARINT"
		case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
			decoder = "Protobuf.fromSignedLong(d.varint64())"
			wireType = "VARINT"
		case descriptorpb.FieldDescriptorProto_TYPE_INT64,
			descriptorpb.FieldDescriptorProto_TYPE_UINT64:
			decoder = "d.varint64()"
			wireType = "VARINT"
		case descriptorpb.FieldDescriptorProto_TYPE_FIXED32, descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
			decoder = "d.number()"
			wireType = "I32"
		case descriptorpb.FieldDescriptorProto_TYPE_FIXED64, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
			decoder = "d.long()"
			wireType = "I64"
		case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
			decoder = "d.float()"
			wireType = "I32"
		case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
			decoder = "d.data()"
			wireType = "LEN"
		case descriptorpb.FieldDescriptorProto_TYPE_STRING:
			decoder = "d.string()"
			wireType = "LEN"
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
			decoder = "%s.Decode(d.data())"
			wireType = "LEN"
		default:
			// TODO reenable the error
			// return fmt.Errorf("unsupported type: %s", *field.Type)
		}
		if decoder != "" { // TODO remove
			if *field.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
				switch wireType {
				case "VARINT", "I32", "I64":
					fmt.Fprintln(cs, "switch (tag & 7) {")
					wtsw := indent.New(cs, tab)
					wtcs := indent.New(wtsw, tab)
					fmt.Fprintf(wtsw, "case Protobuf.%s:\n", wireType)
					addRepeatedFieldValue(wtcs, field, decoder)
					fmt.Fprintln(wtcs, "break;")
					fmt.Fprintln(wtsw, "case Protobuf.LEN:")
					fmt.Fprintln(wtcs, "for (var endRemaining = d.remaining() - d.varint32(); d.remaining() > endRemaining;) {")
					wtfor := indent.New(wtcs, tab)
					addRepeatedFieldValue(wtfor, field, decoder)
					fmt.Fprintln(wtcs, "}")
					fmt.Fprintln(wtcs, "break;")
					fmt.Fprintln(wtsw, "default:")
					fmt.Fprintln(wtcs, `throw new Protobuf.Exception("invalid wire type");`)
					fmt.Fprintln(cs, "}")
				default:
					fmt.Fprintf(cs, "Protobuf.assertWireType(tag, Protobuf.%s);\n", wireType)
					addRepeatedFieldValue(cs, field, decoder)
				}
			} else {
				fmt.Fprintf(cs, "Protobuf.assertWireType(tag, Protobuf.%s);\n", wireType)
				if *field.Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
					fmt.Fprintf(cs, decoder+";\n", strcase.LowerCamelCase(*field.Name))
				} else {
					fmt.Fprintf(cs, "%s = %s;\n", strcase.LowerCamelCase(*field.Name), decoder)
				}
			}
		}
		fmt.Fprintln(cs, "break;")
		fmt.Fprintln(sw, "}")
	}
	fmt.Fprintln(loop, "}")
	fmt.Fprintln(ind, "}")
	fmt.Fprintln(w, "}")

	return nil
}

func addRepeatedFieldValue(w io.Writer, field *descriptorpb.FieldDescriptorProto, decoder string) {
	if *field.Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		fmt.Fprintf(w, "var msg = new %s();\n", typeNameToMonkey(*field.TypeName))
		fmt.Fprintf(w, decoder+";\n", "msg")
		fmt.Fprintf(w, "%s.add(msg);\n", strcase.LowerCamelCase(*field.Name))
	} else {
		fmt.Fprintf(w, "%s.add(%s);\n", strcase.LowerCamelCase(*field.Name), decoder)
	}
}

func isPacked(field *descriptorpb.FieldDescriptorProto) bool {
	if field.Options != nil && field.Options.Packed != nil {
		return *field.Options.Packed
	}
	switch *field.Type {
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64,
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT,
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32,
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32,
		descriptorpb.FieldDescriptorProto_TYPE_INT32,
		descriptorpb.FieldDescriptorProto_TYPE_INT64,
		descriptorpb.FieldDescriptorProto_TYPE_UINT64,
		descriptorpb.FieldDescriptorProto_TYPE_UINT32,
		descriptorpb.FieldDescriptorProto_TYPE_BOOL,
		descriptorpb.FieldDescriptorProto_TYPE_ENUM,
		descriptorpb.FieldDescriptorProto_TYPE_SINT32,
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return true
	default:
		return false
	}
}

func generateEnum(w io.Writer, e *descriptorpb.EnumDescriptorProto) error {
	fmt.Fprintf(w, "enum %s {\n", *e.Name)
	ind := indent.New(w, tab)
	for _, v := range e.Value {
		fmt.Fprintf(ind, "%s = %v,\n", *v.Name, *v.Number)
	}
	fmt.Fprintln(w, "}")
	return nil
}

func str(s string) *string {
	return &s
}
