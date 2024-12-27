package doctag

import (
	"errors"
	"io"
	"reflect"
)

const (
	defaultDoc     = "description"
	defaultExample = "example"
)

var notImplemented = errors.New("not implemented")

type Document struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name        string
	Description string
	Example     string
	Fields      []*Field
}

func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}

func ToDocument(i any) ([]Document, error) {
	return nil, notImplemented
}

func WithJSONName(w io.Writer, i any) error {
	return notImplemented
}

func WithStructFieldName(w io.Writer, i any) error {
	return notImplemented
}

func Write(w io.Writer, i any) error {
	return notImplemented
}
