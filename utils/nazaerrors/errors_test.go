// Copyright 2021, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package nazaerrors

import (
	"errors"
	"io"
	"testing"

	"github.com/qmcloud/admin-common/utils/assert"
)

func TestCombineErrors(t *testing.T) {
	golden := []error{
		nil,
		io.EOF,
		errors.New("1"),
	}

	var err error

	err = CombineErrors(nil)
	assert.Equal(t, nil, err)

	err = CombineErrors(golden[1])
	assert.Equal(t, golden[1], err)

	err = CombineErrors(golden[2])
	assert.Equal(t, golden[2], err)

	err = CombineErrors(nil, golden[1])
	assert.Equal(t, golden[1], err)

	err = CombineErrors(nil, golden[2])
	assert.Equal(t, golden[2], err)

	err = CombineErrors(golden[1], nil)
	assert.Equal(t, golden[1], err)

	err = CombineErrors(nil, golden[1], golden[2])
	assert.Equal(t, golden[1], err)

	err = CombineErrors(golden[1], nil, golden[2])
	assert.Equal(t, golden[1], err)

	err = CombineErrors(golden[1], golden[2], nil)
	assert.Equal(t, golden[1], err)
}
