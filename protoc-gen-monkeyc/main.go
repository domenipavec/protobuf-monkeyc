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

const tab = "  "

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

	if len(proto.Dependency) > 0 {
		return nil, errors.New("import is not supported yet")
	}

	for _, msg := range proto.MessageType {
		err := generateMessage(buf, msg)
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
	}

	for _, field := range msg.Field {
		var typ string
		switch *field.Type {
		case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
			typ = "Double"
		case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
			typ = "Float"
		case descriptorpb.FieldDescriptorProto_TYPE_INT32, descriptorpb.FieldDescriptorProto_TYPE_UINT32, descriptorpb.FieldDescriptorProto_TYPE_SINT32, descriptorpb.FieldDescriptorProto_TYPE_FIXED32, descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
			typ = "Number"
		case descriptorpb.FieldDescriptorProto_TYPE_INT64, descriptorpb.FieldDescriptorProto_TYPE_UINT64, descriptorpb.FieldDescriptorProto_TYPE_SINT64, descriptorpb.FieldDescriptorProto_TYPE_FIXED64, descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
			typ = "Long"
		case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
			typ = "Boolean"
		case descriptorpb.FieldDescriptorProto_TYPE_STRING, descriptorpb.FieldDescriptorProto_TYPE_BYTES:
			typ = "String"
		case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, descriptorpb.FieldDescriptorProto_TYPE_ENUM:
			parts := strings.Split(*field.TypeName, ".")
			typ = parts[len(parts)-1]
		default:
			return fmt.Errorf("unknown type: %s", *field.Type)
		}
		if *field.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
			typ = fmt.Sprintf("Array<%s>", typ)
		}
		fmt.Fprintf(ind, "public var %s as %s;\n", strcase.LowerCamelCase(*field.Name), typ)
	}
	fmt.Fprintln(ind)

	generateMessageEncode(ind, msg)

	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	return nil
}

func generateMessageEncode(w io.Writer, msg *descriptorpb.DescriptorProto) error {
	fmt.Fprintln(w, "public function Encode() as ByteArray {")
	ind := indent.New(w, tab)

	fmt.Fprintln(ind, "var result = []b;")

	for _, field := range msg.Field {
		if field.Options != nil && field.Options.Packed != nil {
			log.Println(*field.Options.Packed)
		}
		fmt.Fprintf(ind, "result.addAll(Protobuf.encodeFieldVarint(%d, %s));\n", *field.Number, strcase.LowerCamelCase(*field.Name))
	}

	fmt.Fprintln(ind, "return result;")
	fmt.Fprintln(w, "}")

	return nil
}

func str(s string) *string {
	return &s
}
