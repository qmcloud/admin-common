// Copyright 2024, Chef.  All rights reserved.
// https://github.com/qmcloud/admin-common
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: QMY (407193275@qq.com)

//go:build windows
// +build windows

package nazalog

func (l *logger) writeLevelStringIfNeeded(level Level) {
	if l.core.option.LevelFlag {
		// windows系统不用写带颜色的日志级别字段
		l.core.buf.WriteString(levelToString[level])
	}
}
