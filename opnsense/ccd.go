package opnsense

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"errors"
	"bytes"
)

type Ccd struct {
	CommonName string `json:"common_name"`
	TunnelNetwork string `json:"tunnel_network"`
	TunnelNetwork6 string `json:"tunnel_network6"`
	RemoteNetwork string `json:"remote_network"`
	RemoteNetwork6 string `json:"remote_network6"`
	LocalNetwork string `json:"local_network"`
	LocalNetwork6 string `json:"local_network6"`
	Block string `json:"block"`
	PushReset string `json:"push_reset"`
}

func (opn *OPNsense) CcdCreate(ccd Ccd, update bool) (string, error) {
	// endpoint
	var endpoint string

	if update {
		endpoint = opn.EndpointForPluginControllerMedthod("openvpn","ccd","setCcdByName")
	} else {
		endpoint = opn.EndpointForPluginControllerMedthod("openvpn","ccd","setCcd")
	}

	var container struct {
		Ccd Ccd `json:"ccd"`
	}

	container.Ccd = ccd
	// create our Request
	jsonBody := new(bytes.Buffer)
	encodingErr := json.NewEncoder(jsonBody).Encode(container)
	if encodingErr != nil {
		return "", encodingErr
	}

	request, reqCreationErr := http.NewRequest("POST", endpoint, jsonBody)

	if reqCreationErr != nil {
		return "", reqCreationErr
	}

	var response, reqErr = opn.send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		var resultContainer struct {
			Result string `json:"result"`
			Uuid string `json:"modified_uuid"`
		}
		jsonError := json.NewDecoder(response.Body).Decode(&resultContainer)

		if jsonError != nil {
			return "", jsonError
		}
		// else
		return resultContainer.Uuid, nil
	} else {
		var resultContainer struct {
			Result string `json:"result"`
		}
		jsonError := json.NewDecoder(response.Body).Decode(&resultContainer)
		if jsonError != nil {
			return "", jsonError
		}
		return "", errors.New(resultContainer.Result)
	}
	// else
	return "", nil
}


func (opn *OPNsense) CcdRemove(commonName string) (string, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("openvpn","ccd","delCcdByName")
	var reqUrl = fmt.Sprintf("%s/%s", endpoint, commonName)

	// create our Request
	jsonBody := new(bytes.Buffer)
	json.NewEncoder(jsonBody).Encode(make([]string, 0))
	request, reqCreationErr := http.NewRequest("POST", reqUrl, jsonBody)

	if reqCreationErr != nil {
		return "", reqCreationErr
	}

	var response, reqErr = opn.send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		var container struct {
			Uuid string `json:"removed_uuid"`
		}
		jsonError := json.NewDecoder(response.Body).Decode(&container)

		if jsonError != nil {
			return "", jsonError
		}
		// else
		return container.Uuid, nil
	} else {
		return "", errors.New("error in response")
	}
	// else
	return "", nil
}

func (opn *OPNsense) CcdGet(commonName string) (Ccd, error){
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("openvpn","ccd","getCcdByName")

	// final request URL
	var reqUrl = fmt.Sprintf("%s/%s", endpoint, commonName)
	// create our Request
	var request, reqCreationErr = http.NewRequest("GET", reqUrl, nil)

	if reqCreationErr != nil {
		return Ccd{}, reqCreationErr
	}

	// send it to the server
	var response, reqErr = opn.send(request)
	if reqErr != nil {
		return Ccd{}, reqErr
	}

	if response.StatusCode == 200 {
		var container struct {
			Ccd Ccd `json:"ccd"`
		}
		jsonError := json.NewDecoder(response.Body).Decode(&container)

		if jsonError != nil {
			return Ccd{}, jsonError
		}
		// else
		return container.Ccd, nil
	} else {
		return Ccd{}, errors.New("error in response")
	}
	// else
	return Ccd{}, nil
}

func (opn *OPNsense) CcdExists(commonName string) (bool, error){
	var ccd, err = opn.CcdGet(commonName)

	if err != nil {
		return true, err
	}

	if ccd.CommonName != "" {
		return true, nil
	}
	// else
	return false, nil
}

func handleBasicResponse(response http.Response, container interface{}) (interface{}, error){
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Ccd{}, err
	}

	if response.StatusCode == 200 {
		jsonError := json.Unmarshal(body, &container)

		if jsonError != nil {
			return nil, jsonError
		}
		// else
		return container, nil
	} else {
		return nil, errors.New("error in response")
	}
}