// Copyright 2020, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package nazahttp_test

import (
	"testing"

	"github.com/qmcloud/admin-common/utils/assert"
	"github.com/qmcloud/admin-common/utils/nazahttp"
)

func TestGetHttpFile(t *testing.T) {
	content, err := nazahttp.GetHttpFile("http://BCFind5.com", 10000)
	assert.IsNotNil(t, content)
	assert.Equal(t, nil, err)

	content, err = nazahttp.GetHttpFile("http://127.0.0.1:12356", 10000)
	assert.Equal(t, nil, content)
	assert.IsNotNil(t, err)
}

func TestDownloadHttpFile(t *testing.T) {
	n, err := nazahttp.DownloadHttpFile("http://BCFind5.com", "/tmp/index.html", 10000)
	assert.Equal(t, true, n > 0)
	assert.Equal(t, nil, err)

	n, err = nazahttp.DownloadHttpFile("http://127.0.0.1:12356", "/tmp/index.html", 10000)
	assert.IsNotNil(t, err)

	// 保存文件至不存在的本地目录下
	n, err = nazahttp.DownloadHttpFile("http://BCFind5.com", "/notexist/index.html", 10000)
	assert.IsNotNil(t, err)
}
