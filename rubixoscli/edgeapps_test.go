package rubixoscli

import (
	"fmt"
	"github.com/NubeIO/rubix-os/installer"
	"github.com/NubeIO/rubix-os/src/cli/constants"
	"github.com/go-resty/resty/v2"
	"testing"
)

var installerObj = installer.New(&installer.Installer{})
var client = New(&Client{
	Rest:          &resty.Client{},
	Installer:     installerObj,
	Ip:            "0.0.0.0",
	Port:          1662,
	HTTPS:         false,
	ExternalToken: composeToken("token"),
}, installerObj)

func TestApp_EdgeAppStatus(t *testing.T) {
	status, connectionErr, requestErr := client.EdgeAppStatus("rc", constants.RubixOs)
	fmt.Println("status", status)
	fmt.Println("connectionErr", connectionErr)
	fmt.Println("requestErr", requestErr)
}
