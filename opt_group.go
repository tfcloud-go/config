// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

type OptGroup struct {
	name string
}

func NewOptGroup(name string) *OptGroup {
	return &OptGroup{
		name: name,
	}
}

func (g *OptGroup) String() string {
	return g.name
}

func (g *OptGroup) Name() string {
	return g.name
}
