package api

import (
	"fmt"
	"net/url"
	"net/http"
	"crypto/tls"
)

type OPNsense struct {
	BaseUrl   url.URL
	ApiKey    string
	ApiSecret string
	NoSslVerify bool
}

func (opn *OPNsense) Send(request *http.Request) (*http.Response, error) {
	var client = &http.Client{}

	if opn.NoSslVerify {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	request.SetBasicAuth(opn.ApiKey, opn.ApiSecret)
	return client.Do(request)
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