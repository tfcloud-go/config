// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

type ListOpt struct {
	name         string
	defaultValue []string
}

func NewListOpt(name string) *ListOpt {
	return &ListOpt{
		name: name,
	}
}

func (o *ListOpt) WithDefault(first string, others ...string) *ListOpt {
	if len(o.defaultValue) == 0 {
		o.defaultValue = []string{first}
	}

	for _, other := range others {
		o.defaultValue = append(o.defaultValue, other)
	}

	return o
}

func (o *ListOpt) Name() string {
	return o.name
}

func (o *ListOpt) Default() []string {
	return o.defaultValue
}
