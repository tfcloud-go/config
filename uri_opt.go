// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

import "net/url"

type URIOpt struct {
	name         string
	defaultValue *url.URL
}

func NewURIOpt(name string) *URIOpt {
	return &URIOpt{
		name: name,
	}
}

func (o *URIOpt) WithDefault(value url.URL) *URIOpt {
	o.defaultValue = &value
	return o
}

func (o *URIOpt) Name() string {
	return o.name
}

func (o *URIOpt) Default() *url.URL {
	return o.defaultValue
}
