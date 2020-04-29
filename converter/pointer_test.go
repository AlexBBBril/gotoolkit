package convert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stringPointer(t *testing.T) {
	maps := map[string]*string{
		"test":          StringPointer("test"),
		"adagds":        StringPointer("adagds"),
		"tes434y54t":    StringPointer("tes434y54t"),
		"teskytkghkykt": StringPointer("teskytkghkykt"),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestIntPointer(t *testing.T) {
	maps := map[int]*int{
		1:    IntPointer(1),
		12:   IntPointer(12),
		123:  IntPointer(123),
		1234: IntPointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestInt8Pointer(t *testing.T) {
	maps := map[int8]*int8{
		1:   Int8Pointer(1),
		12:  Int8Pointer(12),
		123: Int8Pointer(123),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestInt16Pointer(t *testing.T) {
	maps := map[int16]*int16{
		1:    Int16Pointer(1),
		12:   Int16Pointer(12),
		123:  Int16Pointer(123),
		1234: Int16Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestInt32Pointer(t *testing.T) {
	maps := map[int32]*int32{
		1:    Int32Pointer(1),
		12:   Int32Pointer(12),
		123:  Int32Pointer(123),
		1234: Int32Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestInt64Pointer(t *testing.T) {
	maps := map[int64]*int64{
		1:    Int64Pointer(1),
		12:   Int64Pointer(12),
		123:  Int64Pointer(123),
		1234: Int64Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestUintPointer(t *testing.T) {
	maps := map[uint]*uint{
		1:    UintPointer(1),
		12:   UintPointer(12),
		123:  UintPointer(123),
		1234: UintPointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestUint8Pointer(t *testing.T) {
	maps := map[uint8]*uint8{
		1:   Uint8Pointer(1),
		12:  Uint8Pointer(12),
		123: Uint8Pointer(123),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestUint16Pointer(t *testing.T) {
	maps := map[uint16]*uint16{
		1:    Uint16Pointer(1),
		12:   Uint16Pointer(12),
		123:  Uint16Pointer(123),
		1234: Uint16Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestUint32Pointer(t *testing.T) {
	maps := map[uint32]*uint32{
		1:    Uint32Pointer(1),
		12:   Uint32Pointer(12),
		123:  Uint32Pointer(123),
		1234: Uint32Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestUint64Pointer(t *testing.T) {
	maps := map[uint64]*uint64{
		1:    Uint64Pointer(1),
		12:   Uint64Pointer(12),
		123:  Uint64Pointer(123),
		1234: Uint64Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestUintptrPointer(t *testing.T) {
	maps := map[uintptr]*uintptr{
		1:    UintptrPointer(1),
		12:   UintptrPointer(12),
		123:  UintptrPointer(123),
		1234: UintptrPointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestBytePointer(t *testing.T) {
	maps := map[byte]*byte{
		1:   BytePointer(1),
		12:  BytePointer(12),
		123: BytePointer(123),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestRunePointer(t *testing.T) {
	maps := map[rune]*rune{
		1:    RunePointer(1),
		12:   RunePointer(12),
		123:  RunePointer(123),
		1234: RunePointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestFloat32Pointer(t *testing.T) {
	maps := map[float32]*float32{
		1.1:    Float32Pointer(1.1),
		12.1:   Float32Pointer(12.1),
		123.1:  Float32Pointer(123.1),
		1234.1: Float32Pointer(1234.1),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestFloat64Pointer(t *testing.T) {
	maps := map[float64]*float64{
		1.1:    Float64Pointer(1.1),
		12.1:   Float64Pointer(12.1),
		123.1:  Float64Pointer(123.1),
		1234.1: Float64Pointer(1234.1),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestComplex64Pointer(t *testing.T) {
	maps := map[complex64]*complex64{
		1:    Complex64Pointer(1),
		12:   Complex64Pointer(12),
		123:  Complex64Pointer(123),
		1234: Complex64Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestComplex128Pointer(t *testing.T) {
	maps := map[complex128]*complex128{
		1:    Complex128Pointer(1),
		12:   Complex128Pointer(12),
		123:  Complex128Pointer(123),
		1234: Complex128Pointer(1234),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}

func TestBooleanPointer(t *testing.T) {
	maps := map[bool]*bool{
		true:  BooleanPointer(true),
		false: BooleanPointer(false),
	}

	for k, v := range maps {
		assert.Equal(t, k, *v)
	}
}
