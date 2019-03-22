package driver

import (
	"fmt"
	"time"

	dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
	logger "github.com/edgexfoundry/go-mod-core-contracts/clients/logging"
)

type HelloWorldDriver struct {
	lc      logger.LoggingClient
	asyncCh chan<- *dsModels.AsyncValues
}

func (hwD *HelloWorldDriver) DisconnectDevice(deviceName string, protocols map[string]map[string]string) error {
	hwD.lc.Info(fmt.Sprintf("HelloWorldDriver.DisconnectDevice: driver is disconnecting from %s", deviceName))
	return nil
}

func (hwD *HelloWorldDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues) error {
	hwD.lc = lc
	hwD.asyncCh = asyncCh
	return nil
}

func (hwD *HelloWorldDriver) HandleReadCommands(deviceName string, protocols map[string]map[string]string, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
	if len(reqs) != 1 {
		err = fmt.Errorf("HelloWorldDriver.HandleReadCommands; too many command requests; only one supported")
		return
	}

	res = make([]*dsModels.CommandValue, 1)
	now := time.Now().UnixNano() / int64(time.Millisecond)
	cv := dsModels.NewStringValue(&reqs[0].RO, now, fmt.Sprintf("Hello World from: %s", deviceName))

	res[0] = cv
	return
}

func (hwD *HelloWorldDriver) HandleWriteCommands(deviceName string, protocols map[string]map[string]string, reqs []dsModels.CommandRequest, params []*dsModels.CommandValue) error {
	hwD.lc.Info(fmt.Sprintf("HelloWorldDriver.HandleWriteCommands: not enabled for device %s", deviceName))
	return nil
}

func (hwD *HelloWorldDriver) Stop(force bool) error {
	hwD.lc.Info("HelloWorldDriver.Stop: stopping driver...")
	return nil
}
