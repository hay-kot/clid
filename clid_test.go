package clid

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

type Nested struct {
	FlagString string `cli:"nested-string"`
}

type testStruct struct {
	Nested      Nested
	NestedPtr   *Nested
	FlagString  string  `cli:"string"`
	FlagInt     int     `cli:"int"`
	FlagInt8    int8    `cli:"int8"`
	FlagInt16   int16   `cli:"int16"`
	FlagInt32   int32   `cli:"int32"`
	FlagInt64   int64   `cli:"int64"`
	FlagFloat32 float32 `cli:"float32"`
	FlagFloat64 float64 `cli:"float"`

	FlagUint   uint   `cli:"uint"`
	FlagUint8  uint8  `cli:"uint8"`
	FlagUint16 uint16 `cli:"uint16"`
	FlagUint32 uint32 `cli:"uint32"`
	FlagUint64 uint64 `cli:"uint64"`
}

func TestString(t *testing.T) {
	ts := testStruct{
		NestedPtr: &Nested{}, // Initialize Pointer
	}

	set := flag.NewFlagSet("test", 0)

	set.String("string", "value1", "test flag1")

	set.Float64("float", 10.0, "test float")
	set.Float64("float32", 5.0, "test float32")

	set.Int("int", 1, "test int")
	set.Int64("int64", 1, "test int64")
	set.Int64("int32", 1, "test int32")
	set.Int64("int16", 1, "test int16")
	set.Int64("int8", 1, "test int8")

	set.Uint("uint", 1, "test uint")
	set.Uint64("uint64", 2, "test uint64")
	set.Uint64("uint32", 3, "test uint32")
	set.Uint64("uint16", 4, "test uint16")
	set.Uint64("uint8", 5, "test uint8")

	set.String("nested-string", "value2", "test nested string")

	mock := cli.NewContext(nil, set, nil)

	err := Decode(mock, &ts)

	if err != nil {
		t.Error(err)
	}

	// Strings
	assert.Equal(t, "value1", ts.FlagString)

	// Floats
	assert.InDelta(t, 10.0, ts.FlagFloat64, 0.0001)
	assert.InDelta(t, 5.0, ts.FlagFloat32, 0.0001)

	// Ints
	assert.Equal(t, 1, ts.FlagInt)
	assert.Equal(t, int8(1), ts.FlagInt8)
	assert.Equal(t, int16(1), ts.FlagInt16)
	assert.Equal(t, int32(1), ts.FlagInt32)
	assert.Equal(t, int64(1), ts.FlagInt64)

	// Nested Struct
	assert.Equal(t, "value2", ts.Nested.FlagString)
	assert.Equal(t, "value2", ts.NestedPtr.FlagString)

	// Units
	assert.Equal(t, uint(1), ts.FlagUint)
	assert.Equal(t, uint8(5), ts.FlagUint8)
	assert.Equal(t, uint16(4), ts.FlagUint16)
	assert.Equal(t, uint32(3), ts.FlagUint32)
	assert.Equal(t, uint64(2), ts.FlagUint64)
}
