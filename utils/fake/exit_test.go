// Copyright 2019, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package fake_test

import (
	"testing"

	"github.com/qmcloud/admin-common/utils/assert"
	"github.com/qmcloud/admin-common/utils/fake"
)

func TestWithFakeExit(t *testing.T) {
	var er fake.ExitResult
	er = fake.WithFakeOsExit(func() {
		fake.Os_Exit(1)
	})
	assert.Equal(t, true, er.HasExit)
	assert.Equal(t, 1, er.ExitCode)

	er = fake.WithFakeOsExit(func() {
	})
	assert.Equal(t, false, er.HasExit)

	er = fake.WithFakeOsExit(func() {
		fake.Os_Exit(2)
	})
	assert.Equal(t, true, er.HasExit)
	assert.Equal(t, 2, er.ExitCode)
}
