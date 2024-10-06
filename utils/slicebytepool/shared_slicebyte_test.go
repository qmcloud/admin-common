// Copyright 2019, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package slicebytepool

import (
	"testing"

	"github.com/qmcloud/admin-common/utils/assert"
)

func TestNewSharedSliceByte(t *testing.T) {
	pool := NewSliceBytePool(StrategyMultiSlicePoolBucket)
	assert.IsNotNil(t, pool)
	ssb := NewSharedSliceByte(1000, WithPool(pool))
	assert.IsNotNil(t, ssb)
	ssb2 := ssb.Ref()
	assert.IsNotNil(t, ssb2)

	ssb.ReleaseIfNeeded()
	ssb2.ReleaseIfNeeded()
}

func TestWrapSharedSliceByte(t *testing.T) {
	b := make([]byte, 1000)
	pool := NewSliceBytePool(StrategyMultiSlicePoolBucket)
	assert.IsNotNil(t, pool)
	ssb := WrapSharedSliceByte(b, WithPool(pool))
	assert.IsNotNil(t, ssb)
}
