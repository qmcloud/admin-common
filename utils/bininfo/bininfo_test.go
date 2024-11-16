// Copyright 2019, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package bininfo

import (
	"fmt"
	"testing"
)

func TestStringifySingleLine(t *testing.T) {
	fmt.Println(StringifySingleLine())
}

func TestStringifyMultiLine(t *testing.T) {
	fmt.Println(StringifyMultiLine())
}

func TestCorner(t *testing.T) {
	GitStatus = ""
	beauty()
}
