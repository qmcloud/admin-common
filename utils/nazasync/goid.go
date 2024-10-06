// Copyright 2021, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package nazasync

import (
	"bytes"
	"errors"
	"runtime"
	"strconv"
)

// NOTICE copy from https://github.com/golang/net/blob/master/http2/gotrack.go

var ErrObtainGoroutineId = errors.New("nazasync: obtain current goroutine id failed")

func CurGoroutineId() (int64, error) {
	var goroutineSpace = []byte("goroutine ")

	b := make([]byte, 128)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		return -1, ErrObtainGoroutineId
	}
	return strconv.ParseInt(string(b[:i]), 10, 64)
}
