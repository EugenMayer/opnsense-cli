package openvpn

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"errors"
	"bytes"
	coreapi "github.com/eugenmayer/opnsense-cli/opnsense/api"
)

type OpenVpnApi struct {
	*coreapi.OPNsense
}

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

func (opn *OpenVpnApi) CcdCreate(ccd Ccd, update bool) (string, error) {
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
			Uuid string `json:"uuid"`
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
		return resultContainer.Data.Uuid, nil
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


func (opn *OpenVpnApi) CcdRemove(commonName string) (string, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("openvpn","ccd","delCcdByName")
	var reqUrl = fmt.Sprintf("%s/%s", endpoint, commonName)

	// create our Request
	jsonBody := new(bytes.Buffer)
	json.NewEncoder(jsonBody).Encode(make([]string, 0))
	request, reqCreationErr := http.NewRequest("POST", reqUrl, jsonBody)
	request.Header.Set("Content-Type", "application/json")

	if reqCreationErr != nil {
		return "", reqCreationErr
	}

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		var container struct {
			Uuid string `json:"uuid"`
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

func (opn *OpenVpnApi) CcdGet(commonName string) (Ccd, error){
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
	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return Ccd{}, errors.New(fmt.Sprintf("%s:%s",bodyString, reqErr))
	}

	if response.StatusCode == 200 {
		var container struct {
			Status string `json:"status"`
			Ccd Ccd `json:"data"`
		}
		jsonError := json.NewDecoder(response.Body).Decode(&container)

		if jsonError != nil {
			return Ccd{}, jsonError
		}
		// else
		return container.Ccd, nil
	} else {
		var container struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}

		jsonError := json.NewDecoder(response.Body).Decode(&container)
		if jsonError != nil {
			return Ccd{}, jsonError
		}
		return Ccd{}, errors.New(fmt.Sprintf("%s",container.Message))
	}
	// else
	return Ccd{}, nil
}


func (opn *OpenVpnApi) CcdList() ([]Ccd, error){
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMedthod("openvpn","ccd","getCcd")

	// create our Request
	var request, reqCreationErr = http.NewRequest("GET", endpoint, nil)

	if reqCreationErr != nil {
		return []Ccd{}, reqCreationErr
	}

	// send it to the server
	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return []Ccd{}, errors.New(fmt.Sprintf("%s:%s",bodyString, reqErr))
	}

	if response.StatusCode == 200 {
		var container struct {
			Status string `json:"status"`
			Ccds []Ccd `json:"data"`
		}

		jsonError := json.NewDecoder(response.Body).Decode(&container)

		if jsonError != nil {
			return []Ccd{}, jsonError
		}
		// else
		return container.Ccds, nil
	} else {
		var container struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}

		jsonError := json.NewDecoder(response.Body).Decode(&container)
		if jsonError != nil {
			return []Ccd{}, jsonError
		}
		return []Ccd{}, errors.New(fmt.Sprintf("%s:%s",container.Message, reqErr))
	}
	// else
	return []Ccd{}, nil
}

func (opn *OpenVpnApi) CcdExists(commonName string) (bool, error){
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