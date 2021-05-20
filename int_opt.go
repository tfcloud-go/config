// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

type IntOpt struct {
	name         string
	defaultValue int
}

func NewIntOpt(name string) *IntOpt {
	return &IntOpt{
		name: name,
	}
}

func (o *IntOpt) WithDefault(value int) *IntOpt {
	o.defaultValue = value
	return o
}

func (o *IntOpt) Name() string {
	return o.name
}

func (o *IntOpt) Default() int {
	return o.defaultValue
}
