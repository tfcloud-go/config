# TFCloud Go Config

Go package for config. Inspired By OpenStack [oslo.config](https://opendev.org/openstack/oslo.config) library.

## Usage

Install package

```shell
go get github.com/tfcloud-go/oslo-config
```

Register configurations

```go
// conf/server.go

package conf

// ...

var CONF = config.CONF

func init() {
	group := config.NewOptGroup("server")
	CONF.RegisterGroup(group)

	host := config.NewStrOpt("host").WithDefault("127.0.0.1")
	port := config.NewIntOpt("port").WithDefault(8080)
	CONF.RegisterOpts(group, host, port)
}
```

Parse configurations

```go
// main.go

package main

// ...

var CONF = config.CONF

func main() {
	// ...

	CONF.ParseFile("your_config_file_name.conf")

	// ...
}

```

Get configuration value

```go
// foo/server.go

package foo

// ...

var CONF = config.CONF

func startServer() {
	host := CONF.GetString("server", "host")
	port, _ := CONF.GetInt("server", "port")
	addr := net.JoinHostPort(host, port)

	// ...
}
```

Use configuration file

```shell
foo-server --config your_config_file_name.conf
```

## Copyright

The TFCloud Go Team. [Apache License 2.0](./LICENSE)
