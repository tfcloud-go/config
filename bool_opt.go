// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

type BoolOpt struct {
	name         string
	defaultValue bool
}

func NewBoolOpt(name string) *BoolOpt {
	return &BoolOpt{
		name: name,
	}
}

func (o *BoolOpt) WithDefailt(value bool) *BoolOpt {
	o.defaultValue = value
	return o
}

func (o *BoolOpt) Name() string {
	return o.name
}

func (o *BoolOpt) Default() bool {
	return o.defaultValue
}
