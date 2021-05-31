// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package config

import (
	"net/url"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

func newCfg() *ini.File {
	return ini.Empty()
}

type ConfigOpts struct {
	cfg *ini.File
}

func newConfigOpts() *ConfigOpts {
	return &ConfigOpts{
		cfg: newCfg(),
	}
}

func (c *ConfigOpts) ParseFile(file string) error {
	if c.cfg == nil {
		c.cfg = newCfg()
	}

	if err := c.cfg.Append(file); err != nil {
		return err
	}

	return nil
}

func (c *ConfigOpts) RegisterGroup(group *OptGroup) {
	if group == nil {
		return
	}

	if c.cfg == nil {
		c.cfg = newCfg()
	}

	n := group.Name()
	if _, err := c.cfg.GetSection(n); err != nil {
		_, _ = c.cfg.NewSection(n)
	}
}

func (c *ConfigOpts) RegisterOpt(group *OptGroup, opt interface{}) {
	if group == nil {
		return
	}

	n := group.Name()
	section, err := c.cfg.GetSection(n)
	if err != nil {
		section, _ = c.cfg.NewSection(n)
	}

	switch v := opt.(type) {
	case *StrOpt:
		c.registerStrOpt(section, v)
	case *BoolOpt:
		c.registerBoolOpt(section, v)
	case *IntOpt:
		c.registerIntOpt(section, v)
	case *ListOpt:
		c.registerListOpt(section, v)
	case *URIOpt:
		c.registerURIOpt(section, v)
	}
}

func (c *ConfigOpts) RegisterOpts(group *OptGroup, opts ...interface{}) {
	for _, opt := range opts {
		c.RegisterOpt(group, opt)
	}
}

func (c *ConfigOpts) Get(group, option string) string {
	return c.cfg.Section(group).Key(option).Value()
}

func (c *ConfigOpts) registerStrOpt(section *ini.Section, opt *StrOpt) {
	if section == nil || opt == nil {
		return
	}

	if section.HasKey(opt.Name()) {
		return
	}

	value := opt.Default()
	_, _ = section.NewKey(opt.Name(), value)
}

func (c *ConfigOpts) GetString(group, option string) string {
	return c.cfg.Section(group).Key(option).String()
}

func (c *ConfigOpts) registerBoolOpt(section *ini.Section, opt *BoolOpt) {
	if section == nil || opt == nil {
		return
	}

	if section.HasKey(opt.Name()) {
		return
	}

	value := strconv.FormatBool(opt.Default())
	_, _ = section.NewKey(opt.Name(), value)
}

func (c *ConfigOpts) GetBool(group, option string) (bool, error) {
	return c.cfg.Section(group).Key(option).Bool()
}

func (c *ConfigOpts) registerIntOpt(section *ini.Section, opt *IntOpt) {
	if section == nil || opt == nil {
		return
	}

	if section.HasKey(opt.Name()) {
		return
	}

	value := strconv.Itoa(opt.Default())
	_, _ = section.NewKey(opt.Name(), value)
}

func (c *ConfigOpts) GetInt(group, option string) (int, error) {
	return c.cfg.Section(group).Key(option).Int()
}

func (c *ConfigOpts) registerListOpt(section *ini.Section, opt *ListOpt) {
	if section == nil || opt == nil {
		return
	}

	if section.HasKey(opt.Name()) {
		return
	}

	value := strings.Join(opt.Default(), ",")
	_, _ = section.NewKey(opt.Name(), value)
}

func (c *ConfigOpts) GetList(group, option string) []string {
	return c.cfg.Section(group).Key(option).Strings(",")
}

func (c *ConfigOpts) registerURIOpt(section *ini.Section, opt *URIOpt) {
	if section == nil || opt == nil {
		return
	}

	if section.HasKey(opt.Name()) {
		return
	}

	value := opt.Default()
	if value == nil {
		return
	}
	_, _ = section.NewKey(opt.Name(), value.String())
}

func (c *ConfigOpts) GetURI(group, option string) (url.URL, error) {
	value := c.cfg.Section(group).Key(option).Value()

	u, err := url.Parse(value)
	return *u, err
}
