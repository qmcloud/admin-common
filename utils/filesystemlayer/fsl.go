// Copyright 2021, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

package filesystemlayer

import "os"

var _ IFileSystemLayer = &FslDisk{}
var _ IFileSystemLayer = &FslMemory{}

var _ IFile = &os.File{}
var _ IFile = &file{}
