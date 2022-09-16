package db

import (
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// TODO https://linear.app/znki/issue/ZNK-18/improve-mapping-from-protobuf-to-database-model
func CreateDBQuery(schemas []proto.Message) string {
	if schemas == nil || len(schemas) == 0 {
		return ""
	}

	var dbQuery strings.Builder
	for _, message := range schemas {
		tableName := strings.ToLower(string(proto.MessageName(message).Name()))
		dbQuery.WriteString("CREATE TABLE IF NOT EXISTS " + tableName + " ( ")

		fields := message.ProtoReflect().Descriptor().Fields()
		for i := 0; i < fields.Len(); i++ {
			typ := ConvertTyp(fields.Get(i))
			dbQuery.WriteString(string(fields.Get(i).Name()) + " " + typ)
			if i == 0 {
				dbQuery.WriteString(" PRIMARY KEY")
			}

			if i != fields.Len()-1 {
				dbQuery.WriteString(", ")
			}
		}
		dbQuery.WriteString(");\n")
	}

	return dbQuery.String()
}

func ConvertTyp(field protoreflect.FieldDescriptor) string {
	switch field.Kind() {
	case protoreflect.Int32Kind:
		return "INTEGER"
	case protoreflect.StringKind:
		return "TEXT"
	case protoreflect.BoolKind:
		return "BOOL"
	case protoreflect.MessageKind:
		return "TEXT"
	default:
		return "TEXT"
	}
}
