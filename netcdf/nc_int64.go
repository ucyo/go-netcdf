// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.go
// DO NOT EDIT (except nc_double.go).

package netcdf

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteInt64s writes data as the entire data for variable v.
func (v Var) WriteInt64s(data []int64) error {
	if err := okData(v, INT64, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_longlong(C.int(v.ds), C.int(v.id), (*C.longlong)(unsafe.Pointer(&data[0]))))
}

// ReadInt64s reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadInt64s(data []int64) error {
	if err := okData(v, INT64, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_longlong(C.int(v.ds), C.int(v.id), (*C.longlong)(unsafe.Pointer(&data[0]))))
}

// WriteInt64s sets the value of attribute a to val.
func (a Attr) WriteInt64s(val []int64) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_longlong(C.int(a.v.ds), C.int(a.v.id), cname,
		C.nc_type(INT64), C.size_t(len(val)), (*C.longlong)(unsafe.Pointer(&val[0]))))
}

// ReadInt64s reads the entire attribute value into val.
func (a Attr) ReadInt64s(val []int64) (err error) {
	if err := okData(a, INT64, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_longlong(C.int(a.v.ds), C.int(a.v.id), cname,
		(*C.longlong)(unsafe.Pointer(&val[0]))))
	return
}

// ReadInt64At returns a value via index position
func (v Var) ReadInt64At(idx []uint64) (val int64, err error) {
	err = newError(C.nc_get_var1_longlong(C.int(v.ds), C.int(v.id),
		(*C.size_t)(unsafe.Pointer(&idx[0])), (*C.longlong)(unsafe.Pointer(&val))))
	return
}

// WriteInt64At sets a value via its index position
func (v Var) WriteInt64At(idx []uint64, val int64) (err error) {
	err = newError(C.nc_put_var1_longlong(C.int(v.ds), C.int(v.id),
		(*C.size_t)(unsafe.Pointer(&idx[0])), (*C.longlong)(unsafe.Pointer(&val))))
	return
}

// Int64sReader is a interface that allows reading a sequence of values of fixed length.
type Int64sReader interface {
	Len() (n uint64, err error)
	ReadInt64s(val []int64) (err error)
}

// GetInt64s reads the entire data in r and returns it.
func GetInt64s(r Int64sReader) (data []int64, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]int64, n)
	err = r.ReadInt64s(data)
	return
}

// testReadInt64s writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteInt64s(v Var, n uint64) error {
	data := make([]int64, n)
	for i := 0; i < int(n); i++ {
		data[i] = int64(i + 10)
	}
	return v.WriteInt64s(data)
}

// testReadInt64s reads data from v and checks that it's the same as what
// was written by testWriteDouble. N is v.Len().
// This function is only used for testing.
func testReadInt64s(v Var, n uint64) error {
	data := make([]int64, n)
	if err := v.ReadInt64s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := int64(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %v; expected %v", i, data[i], val)
		}
	}
	return nil
}

func testReadInt64At(v Var, n uint64) error {
	data := make([]int64, n)
	if err := v.ReadInt64s(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		shape, _ := v.LenDims()
		coords, _ := UnravelIndex(uint64(i), shape)
		expected := int64(i + 10)
		val, _ := v.ReadInt64At(coords)
		if val != data[i] {
			return fmt.Errorf("data at position %v is %v; expected %v", i, val, expected)
		}
	}
	return nil
}

func testWriteInt64At(v Var, n uint64) error {
	shape, _ := v.LenDims()
	ndim := len(shape)
	coord := make([]uint64, ndim)
	for i := 0; i < ndim; i++ {
		for k := 0; k < ndim; k++ {
			coord[k] = uint64(i)
		}
		v.WriteInt64At(coord, int64(i))
	}
	for i := 0; i < ndim; i++ {
		for k := 0; k < ndim; k++ {
			coord[k] = uint64(i)
		}
		val, _ := v.ReadInt64At(coord)
		if val != int64(i) {
			return fmt.Errorf("data at position %v is %v; expected %v", coord, val, int(i))
		}
	}
	return nil
}
