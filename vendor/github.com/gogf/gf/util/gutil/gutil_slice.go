// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gutil

// CopySlice does a shallow copy of slice <data> for most commonly used slice type
// []interface{}.
func CopySlice(data []interface{}) []interface{} {
	newData := make([]interface{}, len(data))
	copy(newData, data)
	return newData
}
