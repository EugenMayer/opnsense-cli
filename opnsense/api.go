package opnsense

import (
	"fmt"
	"net/url"
)

type OPNsense struct {
	BaseUrl   url.URL
	ApiKey    string
	ApiSecret string
}

// so basically api/<plugin>
func (opn *OPNsense) EndpointForPlugin(plugin string) string {
	return fmt.Sprintf("%s/api/%s", opn.BaseUrl.String(), plugin)
}

// so basically api/<plugin>/<controller>
func (opn *OPNsense) EndpointForPluginController(plugin string, controller string) string {
	return fmt.Sprintf("%s/%s", opn.EndpointForPlugin(plugin), controller)
}

// so basically api/<plugin>/<controller>/<method>
func (opn *OPNsense) EndpointForPluginControllerMedthod(plugin string, controller string, method string) string {
	return fmt.Sprintf("%s/%s", opn.EndpointForPluginController(plugin, controller), method)
}