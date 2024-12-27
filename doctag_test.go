package doctag

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

type NestedStruct struct {
	FieldA int `doc:"another field A"`
}

type EmbeddedStruct struct {
	FieldD string `example:"d"`
}

type SomeStruct struct {
	FieldA         string `doc:"This is field A"`
	FieldB         int
	FieldC         string       `example:"c"`
	Other          NestedStruct `doc:"other"`
	EmbeddedStruct `doc:"embed"`
}

func Test_Write(t *testing.T) {
	buf := &bytes.Buffer{}
	require.NoError(t, Write(buf, SomeStruct{}))
	require.NotEmpty(t, buf.String())
}
