package unbound

import (
	"encoding/json"
	"errors"
	"bytes"
	"net/http"
	coreapi "github.com/eugenmayer/opnsense-cli/opnsense/api"
	"io/ioutil"
	"fmt"
)

type UnboundApi struct {
	*coreapi.OPNsense
}

type HostEntry struct {
	Host        string `json:"host"`
	Domain      string `json:"domain"`
	Ip          string `json:"ip"`
	Rr          string `json:"rr"` // A, MX, CNAME...
	Mxprio      string `json:"mxprio"` // 10, 20
	Mx          string `json:"mx"` // mail.domain.tld ...
	Description string `json:"descr"` // any arbitrary text
}

func NewHostEntry(host, domain, ip string) HostEntry{
	return HostEntry{
		Host: host,
		Domain: domain,
		Ip: ip,
		Rr: "A",
	}
}

func (opn *UnboundApi) HostEntryCreateOrUpdate(hostEntry HostEntry) (string, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("unbound", "hostentry", "setHostEntry")

	var container struct {
		HostEntry HostEntry `json:"hostentry"`
	}

	container.HostEntry = hostEntry
	// create our Request
	jsonBody := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBody).Encode(container); err != nil {
		return "", err
	}

	request, reqCreationErr := http.NewRequest("POST", endpoint, jsonBody)
	if reqCreationErr != nil {
		return "", reqCreationErr
	}
	request.Header.Set("Content-Type", "application/json")

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		type data struct {
			FQDN string `json:"fqdn"`
		}
		var resultContainer struct {
			Result string `json:"status"`
			Data   data   `json:"data"`
		}

		if err := json.NewDecoder(response.Body).Decode(&resultContainer); err != nil {
			return "", err
		}
		// else
		return resultContainer.Data.FQDN, nil
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return "", err
		}

		return "", errors.New(container.Message)
	}
	// else
	return "", nil
}

func (opn *UnboundApi) HostEntryGet(host string, domain string) (HostEntry, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("unbound", "hostentry", "getHostEntry")

	var pseudoFQDN = fmt.Sprintf("%s|%s", host, domain)
	// final request URL
	var reqUrl = fmt.Sprintf("%s/%s", endpoint, pseudoFQDN)
	// create our Request
	var request, reqCreationErr = http.NewRequest("GET", reqUrl, nil)

	if reqCreationErr != nil {
		return HostEntry{}, reqCreationErr
	}

	// send it to the server
	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return HostEntry{}, errors.New(fmt.Sprintf("%s:%s", bodyString, reqErr))
	}

	if response.StatusCode == 200 {
		var container struct {
			Status    string    `json:"status"`
			HostEntry HostEntry `json:"data"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return HostEntry{}, err
		}
		// else
		return container.HostEntry, nil
	} else if response.StatusCode == 404 {
		return HostEntry{}, &coreapi.NotFoundError{
			Err: nil,
			Name: "hostentry",
		}
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return HostEntry{}, err
		}
		return HostEntry{}, errors.New(fmt.Sprintf("%s", container.Message))
	}
	// else
	return HostEntry{}, nil
}

func (opn *UnboundApi) HostEntryList() ([]HostEntry, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("unbound", "hostentry", "getHostEntry")

	// create our Request
	var request, reqCreationErr = http.NewRequest("GET", endpoint, nil)

	if reqCreationErr != nil {
		return []HostEntry{}, reqCreationErr
	}

	// send it to the server
	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return []HostEntry{}, errors.New(fmt.Sprintf("%s:%s", bodyString, reqErr))
	}

	if response.StatusCode == 200 {
		var container struct {
			Status string      `json:"status"`
			Ccds   []HostEntry `json:"data"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return []HostEntry{}, err
		}
		// else
		return container.Ccds, nil
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return []HostEntry{}, err
		}
		return []HostEntry{}, errors.New(fmt.Sprintf("%s:%s", container.Message, reqErr))
	}
	// else
	return []HostEntry{}, nil
}

func (opn *UnboundApi) HostEntryRemove(host string, domain string) (string, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("unbound", "hostentry", "delHostEntry")

	var container struct {
		HostEntry HostEntry `json:"hostentry"`
	}
	container.HostEntry = HostEntry{
		Host:   host,
		Domain: domain,
	}

	// create our Request
	jsonBody := new(bytes.Buffer)
	json.NewEncoder(jsonBody).Encode(container)

	request, reqCreationErr := http.NewRequest("POST", endpoint, jsonBody)
	if reqCreationErr != nil {
		return "", reqCreationErr
	}
	request.Header.Set("Content-Type", "application/json")

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		type data struct {
			FQDN string `json:"fqdn"`
		}
		var resultContainer struct {
			Result string `json:"status"`
			Data   data   `json:"data"`
		}

		if err := json.NewDecoder(response.Body).Decode(&resultContainer); err != nil {
			return "", err
		}
		// else
		return resultContainer.Data.FQDN, nil
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return "", err
		}
		return "", errors.New(fmt.Sprintf("%s:%s", container.Message, reqErr))
	}
	// else
	return "", nil
}

func (opn *UnboundApi) HostEntryExists(host string, domain string) (bool, error) {
	if _, err := opn.HostEntryGet(host, domain); err != nil {
		switch err.(type) {
		case *coreapi.NotFoundError:
			return false, nil
		default:
			return true, err
		}
	}
	// else found something
	return true, nil
}
