package testish

import (
	"encoding/json"
	"github.com/omidfth/testish/internal/types/serviceNames"
)

type Option struct {
	ServiceName serviceNames.ServiceName
	ExposePort  int
	DumpPath    string
}

func castInterfaceToOption(i interface{}) *Option {
	j, _ := json.Marshal(i)
	var option Option
	json.Unmarshal(j, &option)
	return &option
}

func NewOption(serviceName serviceNames.ServiceName, exposePort int, dumpPath string) *Option {
	return &Option{
		ServiceName: serviceName,
		ExposePort:  exposePort,
		DumpPath:    dumpPath,
	}
}
