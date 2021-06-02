// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

type FloatOpt struct {
	name         string
	defaultValue float64
}

func NewFloatOpt(name string) *FloatOpt {
	return &FloatOpt{
		name: name,
	}
}

func (o *FloatOpt) WithDefault(value float64) *FloatOpt {
	o.defaultValue = value
	return o
}

func (o *FloatOpt) Name() string {
	return o.name
}

func (o *FloatOpt) Default() float64 {
	return o.defaultValue
}
