package storage

import "unsafe"

/*
GENERATED FILE. DO NOT EDIT
*/

/* bool */

func (h *Header) Bools() []bool      { return *(*[]bool)(unsafe.Pointer(h)) }
func (h *Header) SetB(i int, x bool) { h.Bools()[i] = x }
func (h *Header) GetB(i int) bool    { return h.Bools()[i] }

/* int */

func (h *Header) Ints() []int       { return *(*[]int)(unsafe.Pointer(h)) }
func (h *Header) SetI(i int, x int) { h.Ints()[i] = x }
func (h *Header) GetI(i int) int    { return h.Ints()[i] }

/* int8 */

func (h *Header) Int8s() []int8       { return *(*[]int8)(unsafe.Pointer(h)) }
func (h *Header) SetI8(i int, x int8) { h.Int8s()[i] = x }
func (h *Header) GetI8(i int) int8    { return h.Int8s()[i] }

/* int16 */

func (h *Header) Int16s() []int16       { return *(*[]int16)(unsafe.Pointer(h)) }
func (h *Header) SetI16(i int, x int16) { h.Int16s()[i] = x }
func (h *Header) GetI16(i int) int16    { return h.Int16s()[i] }

/* int32 */

func (h *Header) Int32s() []int32       { return *(*[]int32)(unsafe.Pointer(h)) }
func (h *Header) SetI32(i int, x int32) { h.Int32s()[i] = x }
func (h *Header) GetI32(i int) int32    { return h.Int32s()[i] }

/* int64 */

func (h *Header) Int64s() []int64       { return *(*[]int64)(unsafe.Pointer(h)) }
func (h *Header) SetI64(i int, x int64) { h.Int64s()[i] = x }
func (h *Header) GetI64(i int) int64    { return h.Int64s()[i] }

/* uint */

func (h *Header) Uints() []uint      { return *(*[]uint)(unsafe.Pointer(h)) }
func (h *Header) SetU(i int, x uint) { h.Uints()[i] = x }
func (h *Header) GetU(i int) uint    { return h.Uints()[i] }

/* uint8 */

func (h *Header) Uint8s() []uint8      { return *(*[]uint8)(unsafe.Pointer(h)) }
func (h *Header) SetU8(i int, x uint8) { h.Uint8s()[i] = x }
func (h *Header) GetU8(i int) uint8    { return h.Uint8s()[i] }

/* uint16 */

func (h *Header) Uint16s() []uint16      { return *(*[]uint16)(unsafe.Pointer(h)) }
func (h *Header) SetU16(i int, x uint16) { h.Uint16s()[i] = x }
func (h *Header) GetU16(i int) uint16    { return h.Uint16s()[i] }

/* uint32 */

func (h *Header) Uint32s() []uint32      { return *(*[]uint32)(unsafe.Pointer(h)) }
func (h *Header) SetU32(i int, x uint32) { h.Uint32s()[i] = x }
func (h *Header) GetU32(i int) uint32    { return h.Uint32s()[i] }

/* uint64 */

func (h *Header) Uint64s() []uint64      { return *(*[]uint64)(unsafe.Pointer(h)) }
func (h *Header) SetU64(i int, x uint64) { h.Uint64s()[i] = x }
func (h *Header) GetU64(i int) uint64    { return h.Uint64s()[i] }

/* uintptr */

func (h *Header) Uintptrs() []uintptr         { return *(*[]uintptr)(unsafe.Pointer(h)) }
func (h *Header) SetUintptr(i int, x uintptr) { h.Uintptrs()[i] = x }
func (h *Header) GetUintptr(i int) uintptr    { return h.Uintptrs()[i] }

/* float32 */

func (h *Header) Float32s() []float32     { return *(*[]float32)(unsafe.Pointer(h)) }
func (h *Header) SetF32(i int, x float32) { h.Float32s()[i] = x }
func (h *Header) GetF32(i int) float32    { return h.Float32s()[i] }

/* float64 */

func (h *Header) Float64s() []float64     { return *(*[]float64)(unsafe.Pointer(h)) }
func (h *Header) SetF64(i int, x float64) { h.Float64s()[i] = x }
func (h *Header) GetF64(i int) float64    { return h.Float64s()[i] }

/* complex64 */

func (h *Header) Complex64s() []complex64   { return *(*[]complex64)(unsafe.Pointer(h)) }
func (h *Header) SetC64(i int, x complex64) { h.Complex64s()[i] = x }
func (h *Header) GetC64(i int) complex64    { return h.Complex64s()[i] }

/* complex128 */

func (h *Header) Complex128s() []complex128   { return *(*[]complex128)(unsafe.Pointer(h)) }
func (h *Header) SetC128(i int, x complex128) { h.Complex128s()[i] = x }
func (h *Header) GetC128(i int) complex128    { return h.Complex128s()[i] }

/* string */

func (h *Header) Strings() []string      { return *(*[]string)(unsafe.Pointer(h)) }
func (h *Header) SetStr(i int, x string) { h.Strings()[i] = x }
func (h *Header) GetStr(i int) string    { return h.Strings()[i] }

/* unsafe.Pointer */

func (h *Header) UnsafePointers() []unsafe.Pointer         { return *(*[]unsafe.Pointer)(unsafe.Pointer(h)) }
func (h *Header) SetUnsafePointer(i int, x unsafe.Pointer) { h.UnsafePointers()[i] = x }
func (h *Header) GetUnsafePointer(i int) unsafe.Pointer    { return h.UnsafePointers()[i] }
