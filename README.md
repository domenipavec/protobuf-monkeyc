# Protobuf Monkey C

Protobuf implementation for Garmin watches programmed with Monkey C.

Supports encoding and decoding of all basic types except `double`. GRPC is not supported.

## Usage

1. If you haven’t installed the protoc compiler, download the package and follow the instructions in the README.
2. Download `protoc-gen-monkeyc` from [releases](https://github.com/domenipavec/protobuf-monkeyc/releases) and install it in your PATH.
3. Run the compiler on your proto file:

```bash
protoc -I=$SRC_DIR --monkeyc_out=$DST_DIR $SRC_DIR/example.proto
```

4. This will generate `example.pb.mc` file, include it in your project.
5. Include `Protobuf` barrel in your `manifest.xml`.

### Encoding protobufs

```monkeyc
var src = new ExampleMessage();
// set any fields you want on src
var byteArray = src.Encode();
```

### Decoding protobufs

```monkeyc
var dst = new ExampleMessage();
dst.Decode(byteArray);
```

## For Developers

### Testing

Test are driven from go and require docker for compiling and simulating Monkey C.
Running `go test ./...` in `protoc-gen-monkeyc` should be enough to run all tests.
