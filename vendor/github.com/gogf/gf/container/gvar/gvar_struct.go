// Copyright 2020 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gvar

import "github.com/gogf/gf/util/gconv"

// Struct maps value of <v> to <pointer>.
// The parameter <pointer> should be a pointer to a struct instance.
// The parameter <mapping> is used to specify the key-to-attribute mapping rules.
func (v *Var) Struct(pointer interface{}, mapping ...map[string]string) error {
	return gconv.Struct(v.Val(), pointer, mapping...)
}

// Struct maps value of <v> to <pointer> recursively.
// The parameter <pointer> should be a pointer to a struct instance.
// The parameter <mapping> is used to specify the key-to-attribute mapping rules.
func (v *Var) StructDeep(pointer interface{}, mapping ...map[string]string) error {
	return gconv.StructDeep(v.Val(), pointer, mapping...)
}

// Structs converts and returns <v> as given struct slice.
func (v *Var) Structs(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.Structs(v.Val(), pointer, mapping...)
}

// StructsDeep converts and returns <v> as given struct slice recursively.
func (v *Var) StructsDeep(pointer interface{}, mapping ...map[string]string) (err error) {
	return gconv.StructsDeep(v.Val(), pointer, mapping...)
}
