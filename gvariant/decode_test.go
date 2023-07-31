package gvariant

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeInt8(t *testing.T) {
	data := []byte{0x10}
	var expect int8 = 16
	var result int8 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeUint8(t *testing.T) {
	data := []byte{0x10}
	var expect uint8 = 16
	var result uint8 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeInt16(t *testing.T) {
	data := []byte{0x10, 0x10}
	var expect int16 = 4112
	var result int16 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeUint16(t *testing.T) {
	data := []byte{0x10, 0x10}
	var expect uint16 = 4112
	var result uint16 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeInt32(t *testing.T) {
	data := []byte{0x10, 0x10, 0x10, 0x10}
	var expect int32 = 269488144
	var result int32 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeUint32(t *testing.T) {
	data := []byte{0x10, 0x10, 0x10, 0x10}
	var expect uint32 = 269488144
	var result uint32 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeInt64(t *testing.T) {
	data := []byte{0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10}
	var expect int64 = 1157442765409226768
	var result int64 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeUint64(t *testing.T) {
	data := []byte{0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10}
	var expect uint64 = 1157442765409226768
	var result uint64 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeFloat32(t *testing.T) {
	data := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	var expect float32 = 3.1415927
	var result float32 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeFloat64(t *testing.T) {
	data := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	var expect float64 = 3.141592653589793
	var result float64 = 0
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expect, result)
}

func TestDecodeBool(t *testing.T) {
	data := []byte{0x01}
	var result bool = false
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.True(t, result)

	data = []byte{0x00}
	err = Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.False(t, result)
}

func TestDecodeString(t *testing.T) {
	data := []byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x00}
	expected := "ABCDEFG"
	result := ""
	err := Unmarshal(data, &result)

	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expected, result)
}

func TestArrayInt32(t *testing.T) {
	data := []byte{
		0x88, 0x02, 0x00, 0x00,
		0x89, 0x02, 0x00, 0x00,
		0x8A, 0x02, 0x00, 0x00,
		0x8B, 0x02, 0x00, 0x00,
	}
	expected := []int32{648, 649, 650, 651}
	result := []int32{}
	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.ElementsMatch(t, expected, result)
}

func TestArrayStrings(t *testing.T) {
	data := []byte{
		0x69, 0x00, 0x63, 0x61, 0x6E, 0x00, 0x68, 0x61, 0x73,
		0x00, 0x73, 0x74, 0x72, 0x69, 0x6E, 0x67, 0x73, 0x3F, 0x00,
		0x02, 0x06, 0x0a, 0x13,
	}

	expected := []string{"i", "can", "has", "strings?"}
	result := []string{}
	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.ElementsMatch(t, expected, result)
}

func TestStruct(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int32
	}

	data := []byte{0x66, 0x6F, 0x6F, 0x00, 0xff, 0xff, 0xff, 0xff, 0x04}
	expected := testStruct{"foo", -1}
	result := testStruct{}
	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expected, result)
}

func TestPaddedStruct(t *testing.T) {
	type testStruct struct {
		Field1 uint8
		Field2 int32
	}

	data := []byte{0x70, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00}
	expected := testStruct{0x70, 96}
	result := testStruct{}
	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expected, result)
}

func TestArrayFixedLenStuctsWithPadding(t *testing.T) {
	type testStruct struct {
		Field1 int32
		Field2 uint8
	}

	data := []byte{
		0x60, 0x00, 0x00, 0x00, 0x70, 0x00, 0x00, 0x00,
		0x88, 0x02, 0x00, 0x00, 0xf7, 0x00, 0x00, 0x00,
	}

	expected := []testStruct{
		{96, 0x70},
		{648, 0xf7},
	}
	result := []testStruct{}

	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expected, result)
}

func TestArrayVariableLenStuctsWithPadding(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int32
	}

	data := []byte{
		0x68, 0x69, 0x00, 0x00, 0xfe, 0xff, 0xff, 0xff,
		0x03, 0x00, 0x00, 0x00, 0x62, 0x79, 0x65, 0x00,
		0xff, 0xff, 0xff, 0xff, 0x04, 0x09,
		// This byte added to be consistent with the
		// specification of an array containing variable
		// width types.  However it is not included in
		// the example
		0x15,
	}

	expected := []testStruct{
		{"hi", -2},
		{"bye", -1},
	}
	result := []testStruct{}

	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expected, result)
}

func TestDictionary(t *testing.T) {
	data := []byte{
		0x61, 0x20, 0x6b, 0x65, 0x79, 0x00, 0x00, 0x00,
		0x02, 0x02, 0x00, 0x00, 0x06,
	}
	expected := map[string]int32{
		"a key": 514,
	}
	result := map[string]int32{}

	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expected, result)
}

func TestVariant(t *testing.T) {
	data := []byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x00, 0x00, 0x73}
	expected := Variant{
		Data:   []byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x00},
		Format: "s",
	}
	result := Variant{}

	err := Unmarshal(data, &result)
	assert.Nil(t, err, "Unmarshal error")
	assert.Equal(t, expected, result)
}

// TestOSTreeCommit tests that a serialised OSTree commit file contents
// deserialises.
// The commit type format is `(a{sv}aya(say)sstayay)`
//
//	a{sv} - Metadata
//	ay - parent checksum (empty string for initial)
//	a(say) - Related objects
//	s - subject
//	s - body
//	t - Timestamp in seconds since the epoch (UTC, big-endian)
//	ay - Root tree contents
//	ay - Root tree metadata
func TestOSTreeCommit(t *testing.T) {
	data, err := os.ReadFile("testdata/commit.dat")
	if err != nil {
		t.Fatalf("could not read test commit data: %s", err)
	}

	type commit struct {
		Metadata       []map[string]Variant
		ParentCheckSum []uint8
		RelatedObjects []struct {
			Name string
			Ref  []uint8
		}
		Subject          string
		Body             string
		Timestamp        uint64
		RootTreeContents []uint8
		RootTreeMetadata []uint8
	}

	mayPanic := func() {
		result := commit{}
		err = UnmarshalBigEndian(data, &result)
		assert.Nil(t, err, "Unmarshal error")
	}

	assert.NotPanics(t, mayPanic)
}

// TestOSTreeDirMeta tests that a serialised OSTree dirmeta file contents
// deserialises.
// The meta type format is `(a{sv}aya(say)sstayay)`
//
// u - uid (big-endian)
// u - gid (big-endian)
// u - mode (big-endian)
// a(ayay) - xattrs
func TestOSTreeDirMeta(t *testing.T) {
	data, err := os.ReadFile("testdata/dirmeta.dat")
	if err != nil {
		t.Fatalf("could not read test commit data: %s", err)
	}

	type dirmeta struct {
		UID    uint32
		GID    uint32
		Mode   uint32
		XAttrs []struct {
			Bytes1 []uint8
			Bytes2 []uint8
		}
	}

	mayPanic := func() {
		result := dirmeta{}
		err = UnmarshalBigEndian(data, &result)
		assert.Nil(t, err, "Unmarshal error")
	}

	assert.NotPanics(t, mayPanic)
}

// TestOSTreeDirTree tests that a serialised OSTree dirmeta file contents
// deserialises.
// The tree type form * - a(say) - array of (filename, checksum) for files
//
// - a(say) - array of (filename, checksum) for files
// - a(sayay) - array of (dirname, tree_checksum, meta_checksum) for directories
func TestOSTreeDirTree(t *testing.T) {
	data, err:= os.ReadFile("testdata/dirtree.dat")
	if err != nil {
		t.Fatalf("could not read test commit data: %s", err)
	}

	type dirtree struct {
		Files []struct {
			Filename string
			Checksum []uint8
		}
		Directories []struct {
			Dirname      string
			TreeChecksum []uint8
			MetaChecksum []uint8
		}
	}

	mayPanic := func() {
		result := dirtree{}
		err = UnmarshalBigEndian(data, &result)
		assert.Nil(t, err, "Unmarshal error")
	}

	assert.NotPanics(t, mayPanic)
}
