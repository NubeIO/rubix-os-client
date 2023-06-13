package rubixoscli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestClient_ProxyPOST(t *testing.T) {
	plan, _ := ioutil.ReadFile("proxy-example.json")
	var data interface{}
	err := json.Unmarshal(plan, &data)
	fmt.Println(err)

	post, err := client.ProxyPOST("rc", "/api/system/scanner", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}
