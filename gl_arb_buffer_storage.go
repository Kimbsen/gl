// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	BUFFER_IMMUTABLE_STORAGE         = C.GL_BUFFER_IMMUTABLE_STORAGE
	BUFFER_STORAGE_FLAGS             = C.GL_BUFFER_STORAGE_FLAGS
	CLIENT_STORAGE_BIT               = C.GL_CLIENT_STORAGE_BIT
	CLIENT_MAPPED_BUFFER_BARRIER_BIT = C.GL_CLIENT_MAPPED_BUFFER_BARRIER_BIT
	DYNAMIC_STORAGE_BIT              = C.GL_DYNAMIC_STORAGE_BIT
	MAP_COHERENT_BIT                 = C.GL_MAP_COHERENT_BIT
	MAP_PERSISTENT_BIT               = C.GL_MAP_PERSISTENT_BIT

//	MAP_READ_BIT                     = C.GL_MAP_READ_BIT
//	MAP_WRITE_BIT                    = C.GL_MAP_WRITE_BIT
)

func BufferStorage(target GLenum, size uint32, data unsafe.Pointer, flags uint32) {
	C.glBufferStorage(C.GLenum(target), C.GLsizeiptr(size), data, C.GLbitfield(flags))
}

func NamedBufferStorage(buffer Buffer, size uint32, data unsafe.Pointer, flags uint32) {
	C.glNamedBufferStorageEXT(C.GLuint(buffer), C.GLsizeiptr(size), data, C.GLbitfield(flags))
}
