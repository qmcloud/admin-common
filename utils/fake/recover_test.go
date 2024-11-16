// Copyright 2020, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package fake_test

import (
	"testing"

	"github.com/qmcloud/admin-common/utils/fake"
)

func TestWithRecover(t *testing.T) {
	fake.WithRecover(func() {
		// noop
	})
	fake.WithRecover(func() {
		panic(0)
	})
}
