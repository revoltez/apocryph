package proto

import (
	"errors"
	"io"
	"os"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var UnmarshalFormats = map[string]func(b []byte, m protoreflect.ProtoMessage) error{
	"json": protojson.Unmarshal,
	"pb":   proto.Unmarshal,
	"pbtext": prototext.Unmarshal,
}

var UnmarshalFormatNames []string

func init() {
	UnmarshalFormatNames := make([]string, 0, len(UnmarshalFormats))
	for name := range UnmarshalFormats {
		UnmarshalFormatNames = append(UnmarshalFormatNames, name)
	}
}

func Unmarshal(format string, bytes []byte, m protoreflect.ProtoMessage) error {
	Unmarshal := UnmarshalFormats[format]
	if Unmarshal == nil {
		return errors.New("Unknown format: " + format)
	}
	return Unmarshal(bytes, m)
}

func DetectFormat(path string) string {
	for name := range UnmarshalFormats {
		if strings.HasSuffix(path, "." + name) {
			return name
		}
	}
	return path
}

func UnmarshalFile(path string, format string, m protoreflect.ProtoMessage) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if format == "" {
		format = DetectFormat(path)
	}

	return Unmarshal(format, bytes, m)
}
