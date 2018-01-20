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
	Host string `json:"host"`
	Domain string `json:"domain"`
	Ip string `json:"ip	"`
	Rr string `json:"rr"`
	Mxprio string `json:"mxprio"`
	Mx string `json:"mx"`
	Description string `json:"descr"`
}

func (opn *UnboundApi) HostEntryCreate(hostEntry HostEntry, update bool) (string, error) {
	// endpoint
	var endpoint string

	if update {
		endpoint = opn.EndpointForPluginControllerMedthod("unbound","hostentry","setHostEntryByFQDN")
	} else {
		endpoint = opn.EndpointForPluginControllerMedthod("unbound","hostentry","setHostEntry")
	}

	var container struct {
		HostEntry HostEntry `json:"hostentry"`
	}

	container.HostEntry = hostEntry
	// create our Request
	jsonBody := new(bytes.Buffer)
	encodingErr := json.NewEncoder(jsonBody).Encode(container)
	if encodingErr != nil {
		return "", encodingErr
	}

	request, reqCreationErr := http.NewRequest("POST", endpoint, jsonBody)
	request.Header.Set("Content-Type", "application/json")

	if reqCreationErr != nil {
		return "", reqCreationErr
	}

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
			Data data `json:"data"`
		}
		jsonError := json.NewDecoder(response.Body).Decode(&resultContainer)

		if jsonError != nil {
			return "", jsonError
		}
		// else
		return resultContainer.Data.FQDN, nil
	} else {
		var container struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}

		jsonError := json.NewDecoder(response.Body).Decode(&container)
		if jsonError != nil {
			return "", jsonError
		}

		return "", errors.New(container.Message)
	}
	// else
	return "", nil
}


func (opn *UnboundApi) HostEntryGet(host string, domain string) (HostEntry, error){
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("unbound","hostentry","getHostEntry")

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
		return HostEntry{}, errors.New(fmt.Sprintf("%s:%s",bodyString, reqErr))
	}

	if response.StatusCode == 200 {
		var container struct {
			Status string `json:"status"`
			HostEntry HostEntry `json:"data"`
		}
		jsonError := json.NewDecoder(response.Body).Decode(&container)

		if jsonError != nil {
			return HostEntry{}, jsonError
		}
		// else
		return container.HostEntry, nil
	} else {
		var container struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}

		jsonError := json.NewDecoder(response.Body).Decode(&container)
		if jsonError != nil {
			return HostEntry{}, jsonError
		}
		return HostEntry{}, errors.New(fmt.Sprintf("%s",container.Message))
	}
	// else
	return HostEntry{}, nil
}
