// Copyright 2022, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package nazasync

import (
	"sync"
	"sync/atomic"
)

type StdOnce struct {
	core sync.Once
}

func (o *StdOnce) Do(f func()) {
	o.core.Do(f)
}

type NonblockingOnce struct {
	done uint32
}

func (o *NonblockingOnce) Do(f func()) {
	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		f()
	}
}
