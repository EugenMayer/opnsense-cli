package api

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"net/url"
	"os"
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
			RootCAs:            certPool,
		},
	}

	request.SetBasicAuth(opn.ApiKey, opn.ApiSecret)
	return client.Do(request)
}

type NotFoundError struct {
	Name string
	Err  error
}

func (f *NotFoundError) Error() string {
	return fmt.Sprintf("not found: %s", f.Name)
}

func ConfigureFromEnv() (*OPNsense, error) {
	_ = godotenv.Load()

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

// EndpointForModule so basically api/<plugin>
func (opn *OPNsense) EndpointForModule(module string) string {
	return fmt.Sprintf("%s/api/%s", opn.BaseUrl.String(), module)
}

// EndpointForModuleController so basically api/<plugin>/<controller>
func (opn *OPNsense) EndpointForModuleController(module string, controller string) string {
	return fmt.Sprintf("%s/%s", opn.EndpointForModule(module), controller)
}

// EndpointForPluginControllerMethod so basically api/<plugin>/<controller>/<method>
func (opn *OPNsense) EndpointForPluginControllerMethod(module string, controller string, method string) string {
	return fmt.Sprintf("%s/%s", opn.EndpointForModuleController(module, controller), method)
}
