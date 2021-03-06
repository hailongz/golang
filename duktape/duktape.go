package duktape

/*
#cgo !windows CFLAGS: -std=c99 -O3 -Wall -fomit-frame-pointer -fstrict-aliasing
#cgo windows CFLAGS: -O3 -Wall -fomit-frame-pointer -fstrict-aliasing
#cgo linux LDFLAGS: -lm
#cgo freebsd LDFLAGS: -lm
#cgo openbsd LDFLAGS: -lm

#include "duk_config.h"
#include "duktape.h"
#include "kk.h"

extern duk_ret_t goFunctionCall(struct duk_hthread *ctx);
extern duk_ret_t goFinalizeCall(struct duk_hthread *ctx);
*/
import "C"

import (
	"fmt"
	"strconv"
	"sync"
	"unsafe"
)

type scope struct {
	id      int64
	autoId  int
	objects map[int]interface{}
}

var scopeAutoId int64 = 0
var scopeSet map[int64]*scope = map[int64]*scope{}
var scopeLock sync.RWMutex

func newScope() *scope {
	v := scope{}
	v.autoId = 0
	v.objects = map[int]interface{}{}

	scopeLock.Lock()
	defer scopeLock.Unlock()

	v.id = scopeAutoId + 1
	scopeAutoId = v.id
	scopeSet[v.id] = &v

	return &v
}

func deleteScope(id int64) {
	scopeLock.Lock()
	defer scopeLock.Unlock()
	delete(scopeSet, id)
}

func (s *scope) Add(object interface{}) int {
	id := s.autoId + 1
	s.autoId = id
	s.objects[id] = object
	return id
}

func (s *scope) Remove(id int) {
	delete(s.objects, id)
}

func (s *scope) Call(id int) int {
	v, ok := s.objects[id]
	if ok {
		fn, ok := v.(func() int)
		if ok {
			return fn()
		}
	}
	return 0
}

func (s *scope) Get(id int) interface{} {
	v, ok := s.objects[id]
	if ok {
		return v
	}
	return 0
}

type Context struct {
	s           *scope
	duk_context *C.struct_duk_hthread
}

func New() *Context {
	v := Context{
		s:           newScope(),
		duk_context: C.kk_duk_create_heap(),
	}
	return &v
}

func (d *Context) Recycle() {
	deleteScope(d.s.id)
	C.duk_destroy_heap(d.duk_context)
}

func (d *Context) ToError(idx int) error {

	err := &Error{}

	for _, key := range []string{"name", "message", "fileName", "lineNumber", "stack"} {

		d.GetPropString(idx, key)

		switch key {
		case "name":
			err.Type = d.SafeToString(-1)
		case "message":
			err.Message = d.SafeToString(-1)
		case "fileName":
			err.FileName = d.SafeToString(-1)
		case "lineNumber":
			if d.IsNumber(-1) {
				err.LineNumber = d.GetInt(-1)
			}
		case "stack":
			err.Stack = d.SafeToString(-1)
		}

		d.Pop()
	}

	return err
}

func (d *Context) ToBytes(idx int) []byte {

	if d.IsBuffer(idx) {
		v, n := d.ToBuffer(idx)
		return C.GoBytes(v, C.int(n))
	} else if d.IsBufferData(idx) {
		v, n := d.GetBufferData(idx)
		return C.GoBytes(v, C.int(n))
	}

	return nil
}

func (d *Context) PushBytes(data []byte) {

	if data == nil {
		d.PushUndefined()
		return
	}

	n := len(data)

	src := C.CBytes(data)

	dest := d.PushBuffer(n, false)

	C.memcpy(dest, src, C.size_t(n))

}

func (d *Context) PushGlobalGoFunction(key string, fn func() int) {
	d.PushGlobalObject()
	d.PushGoFunction(fn)
	d.PutPropString(-2, key)
	d.Pop()
}

func (d *Context) PushGoFunction(fn func() int) {
	s := d.s
	id := s.Add(fn)

	C.duk_push_c_function(d.duk_context, (*[0]byte)(C.goFunctionCall), C.DUK_VARARGS)

	setScope(d.duk_context, -1, s)
	setFunctionId(d.duk_context, -1, id)

	C.duk_push_c_function(d.duk_context, (*[0]byte)(C.goFinalizeCall), C.duk_idx_t(1))
	C.duk_set_finalizer(d.duk_context, C.duk_idx_t(-2))

}

func (d *Context) PushGoObject(object interface{}) {
	s := d.s
	id := s.Add(object)

	C.duk_push_object(d.duk_context)

	setScope(d.duk_context, -1, s)
	setFunctionId(d.duk_context, -1, id)

	C.duk_push_c_function(d.duk_context, (*[0]byte)(C.goFinalizeCall), C.duk_idx_t(1))
	C.duk_set_finalizer(d.duk_context, C.duk_idx_t(-2))

}

func (d *Context) ToGoObject(idx int) interface{} {
	if d.IsObject(idx) {
		s := getScope(d.duk_context, idx)
		if s != nil {
			id := getFunctionId(d.duk_context, idx)
			if id != 0 {
				return s.Get(id)
			}
		}
	}
	return nil
}

func setScope(ctx *C.struct_duk_hthread, idx int, s *scope) {

	key := C.CString("__scope")
	id := C.CString(fmt.Sprintf("%d", s.id))

	C.duk_push_string(ctx, key)
	C.duk_push_string(ctx, id)

	C.duk_def_prop(ctx, C.duk_idx_t(idx-2), C.DUK_DEFPROP_HAVE_VALUE|C.DUK_DEFPROP_HAVE_CONFIGURABLE|C.DUK_DEFPROP_CLEAR_ENUMERABLE|C.DUK_DEFPROP_CLEAR_WRITABLE)

	C.free(unsafe.Pointer(key))
	C.free(unsafe.Pointer(id))
}

func setFunctionId(ctx *C.struct_duk_hthread, idx int, id int) {

	key := C.CString("__id")

	C.duk_push_string(ctx, key)
	C.duk_push_int(ctx, C.duk_int_t(id))
	C.duk_def_prop(ctx, C.duk_idx_t(idx-2), C.DUK_DEFPROP_HAVE_VALUE|C.DUK_DEFPROP_HAVE_CONFIGURABLE|C.DUK_DEFPROP_HAVE_ENUMERABLE)

	C.free(unsafe.Pointer(key))
}

func getScope(ctx *C.struct_duk_hthread, idx int) *scope {

	var s *scope = nil

	key := C.CString("__scope")

	C.duk_get_prop_string(ctx, C.duk_idx_t(idx), key)

	if C.duk_is_string(ctx, C.duk_idx_t(-1)) != C.duk_bool_t(0) {

		p := C.duk_to_string(ctx, C.duk_idx_t(-1))

		if p != nil {
			id := C.GoString(p)
			scopeLock.RLock()
			defer scopeLock.RUnlock()
			iid, _ := strconv.ParseInt(id, 10, 64)
			s = scopeSet[iid]
		}
	}

	C.duk_pop(ctx)

	C.free(unsafe.Pointer(key))

	return s
}

func getFunctionId(ctx *C.struct_duk_hthread, idx int) int {

	var id int = 0

	key := C.CString("__id")

	C.duk_get_prop_string(ctx, C.duk_idx_t(idx), key)

	if C.duk_is_number(ctx, C.duk_idx_t(-1)) != C.duk_bool_t(0) {
		id = int(C.duk_to_int(ctx, C.duk_idx_t(-1)))
	}

	C.duk_pop(ctx)

	C.free(unsafe.Pointer(key))

	return id
}

//export goFunctionCall
func goFunctionCall(ctx *C.struct_duk_hthread) C.duk_ret_t {

	C.duk_push_current_function(ctx)

	s := getScope(ctx, -1)
	id := getFunctionId(ctx, -1)

	C.duk_pop(ctx)

	if id != 0 && s != nil {
		return C.duk_ret_t(s.Call(id))
	}

	return 0
}

//export goFinalizeCall
func goFinalizeCall(ctx *C.struct_duk_hthread) C.duk_ret_t {

	s := getScope(ctx, -1)
	id := getFunctionId(ctx, -1)

	if id != 0 && s != nil {
		s.Remove(id)
	}

	return 0
}
