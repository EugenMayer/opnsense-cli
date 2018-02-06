package api

import (
	"fmt"
	"net/url"
	"net/http"
	"crypto/tls"
	"github.com/joho/godotenv"
	"os"
	"errors"
	"crypto/x509"
)

type OPNsense struct {
	BaseUrl     url.URL
	ApiKey      string
	ApiSecret   string
	NoSslVerify bool
}

func (opn *OPNsense) Send(request *http.Request) (*http.Response, error) {
	var client = &http.Client{}

	certPool, _ := x509.SystemCertPool()
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: opn.NoSslVerify,
			RootCAs: certPool,
		},
	}

	request.SetBasicAuth(opn.ApiKey, opn.ApiSecret)
	return client.Do(request)
}

type NotFoundError struct {
	Name string
	Err error
}

func (f *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", f.Name )
}

func ConfigureFromEnv() (*OPNsense, error) {
	godotenv.Load()

	if _, isset := os.LookupEnv("OPN_URL"); !isset {
		return nil, errors.New(fmt.Sprintf("Please set the OPN_URL to your opnsense opnUrl like https://myopnsense:10443"))
	}

	if _, isset := os.LookupEnv("OPN_APIKEY"); !isset {
		return nil, errors.New(fmt.Sprintf("Please set OPN_APIKEY to your opnsense api apiKey"))
	}

	if _, isset := os.LookupEnv("OPN_APISECRET"); !isset {
		return nil, errors.New(fmt.Sprintf("Please set OPN_APISECRET to your opnsense api apiSecret"))
	}

	var parsedUrl, err = url.Parse(os.Getenv("OPN_URL"))
	if err != nil {
		return nil, err
	}

	return &OPNsense{
		BaseUrl:     *parsedUrl,
		ApiKey:      os.Getenv("OPN_APIKEY"),
		ApiSecret:   os.Getenv("OPN_APISECRET"),
		NoSslVerify: os.Getenv("OPN_NOSSLVERIFY") == "1",
	}, nil
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
