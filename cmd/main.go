package main

import (
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
	"github.com/eno-intel/helloworld-device/driver"
)

const (
	version     string = "1.0.0"
	serviceName string = "helloworld-device"
)

func main() {
	helloWordDriver := driver.HelloWorldDriver{}
	startup.Bootstrap(serviceName, version, &helloWordDriver)
}
