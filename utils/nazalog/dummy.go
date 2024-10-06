// Copyright 2021, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package nazalog

var DummyLogger Logger

func init() {
	DummyLogger, _ = New(func(option *Option) {
		option.Level = LevelLogNothing
	})
}
