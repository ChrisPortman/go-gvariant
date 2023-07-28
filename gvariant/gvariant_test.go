package gvariant

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFixedWidth(t *testing.T) {
	testUint8 := uint8(1)
	testInt8 := int8(1)
	testUint16 := uint16(1)
	testInt16 := int16(1)
	testUint32 := uint32(1)
	testInt32 := int32(1)
	testUint64 := uint64(1)
	testInt64 := int64(1)
	testBool := true
	testFloat32 := float32(1.1)
	testFloat64 := float64(1.1)
	testString := "test"

	assert.True(t, isFixedWidth(reflect.ValueOf(&testUint8).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testInt8).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testUint16).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testInt16).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testUint32).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testInt32).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testUint64).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testInt64).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testBool).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testFloat32).Elem()))
	assert.True(t, isFixedWidth(reflect.ValueOf(&testFloat64).Elem()))

	assert.False(t, isFixedWidth(reflect.ValueOf(&testString)))
}

func TestTypeWidth(t *testing.T) {
	testUint8 := uint8(1)
	testInt8 := int8(1)
	testUint16 := uint16(1)
	testInt16 := int16(1)
	testUint32 := uint32(1)
	testInt32 := int32(1)
	testUint64 := uint64(1)
	testInt64 := int64(1)
	testBool := true
	testFloat32 := float32(1.1)
	testFloat64 := float64(1.1)
	testString := "test"

	testFixedStruct := struct {
		Field1 uint8
		Field2 uint32
	}{0x70, 96}

	assert.Equal(t, 1, typeWidth(reflect.ValueOf(&testUint8).Elem()))
	assert.Equal(t, 1, typeWidth(reflect.ValueOf(&testInt8).Elem()))
	assert.Equal(t, 2, typeWidth(reflect.ValueOf(&testUint16).Elem()))
	assert.Equal(t, 2, typeWidth(reflect.ValueOf(&testInt16).Elem()))
	assert.Equal(t, 4, typeWidth(reflect.ValueOf(&testUint32).Elem()))
	assert.Equal(t, 4, typeWidth(reflect.ValueOf(&testInt32).Elem()))
	assert.Equal(t, 8, typeWidth(reflect.ValueOf(&testUint64).Elem()))
	assert.Equal(t, 8, typeWidth(reflect.ValueOf(&testInt64).Elem()))
	assert.Equal(t, 8, typeWidth(reflect.ValueOf(&testFloat32).Elem()))
	assert.Equal(t, 8, typeWidth(reflect.ValueOf(&testFloat32).Elem()))
	assert.Equal(t, 8, typeWidth(reflect.ValueOf(&testFloat64).Elem()))
	assert.Equal(t, 8, typeWidth(reflect.ValueOf(&testFloat64).Elem()))
	assert.Equal(t, 1, typeWidth(reflect.ValueOf(&testBool).Elem()))
	assert.Equal(t, 8, typeWidth(reflect.ValueOf(&testFixedStruct).Elem()))

	assert.Equal(t, 0, typeWidth(reflect.ValueOf(&testString).Elem()))
}

func TestUnshift(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5}
	firstTwo := unshift(&mySlice, 2)

	assert.Equal(t, mySlice, []int{3, 4, 5})
	assert.Equal(t, firstTwo, []int{1, 2})
}
