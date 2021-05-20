// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

type StrOpt struct {
	name         string
	defaultValue string
}

func NewStrOpt(name string) *StrOpt {
	return &StrOpt{
		name: name,
	}
}

func (o *StrOpt) WithDefault(value string) *StrOpt {
	o.defaultValue = value
	return o
}

func (o *StrOpt) Name() string {
	return o.name
}

func (o *StrOpt) Default() string {
	return o.defaultValue
}
