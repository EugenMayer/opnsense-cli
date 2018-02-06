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
	TunnelNetwork string `json:"tunnel_network"` // CIDR like 10.1.1.1/24
	TunnelNetwork6 string `json:"tunnel_network6"`
	RemoteNetwork string `json:"remote_network"`  // CIDR like 10.1.1.1/24
	RemoteNetwork6 string `json:"remote_network6"`
	LocalNetwork string `json:"local_network"`
	LocalNetwork6 string `json:"local_network6"` // CIDR like 10.1.1.1/24
	Block string `json:"block"` // 0 or 1 as string
	PushReset string `json:"push_reset"` // 0 or 1 as string
}

// factory to ensure sane defaults
func NewCCD(commonName string) Ccd{
	return Ccd{
		CommonName: commonName,
		Block: "0",
		PushReset: "0",
	}
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

	if err := json.NewEncoder(jsonBody).Encode(container); err != nil {
		return "", err
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

		if err := json.NewDecoder(response.Body).Decode(&resultContainer); err != nil {
			return "", err
		}
		// else
		return resultContainer.Data.Uuid, nil
	} else {
		var container struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return "", err
		}

		return "", errors.New(container.Message)
	}
}


func (opn *OpenVpnApi) CcdRemove(commonName string) (string, error) {
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

	request.Header.Set("Content-Type", "application/json")

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		var container struct {
			Uuid string `json:"uuid"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return "", err
		}
		// else
		return container.Uuid, nil
	} else {
		return "", errors.New("error in response")
	}
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
		bodyString := ""
		if response != nil {
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			bodyString = string(bodyBytes)
		}
		return Ccd{}, errors.New(fmt.Sprintf("%s:%s",reqErr, bodyString))
	}

	if response.StatusCode == 200 {
		var container struct {
			Status string `json:"status"`
			Ccd Ccd `json:"data"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return Ccd{}, err
		}
		// else
		return container.Ccd, nil

	} else if  response.StatusCode == 404 {
		return Ccd{}, &coreapi.NotFoundError{
			Err: nil,
			Name: "ccd",
		}
	} else {
		var container struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return Ccd{}, err
		}

		return Ccd{}, errors.New(fmt.Sprintf("%s",container.Message))
	}
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
		bodyString := ""
		if response != nil {
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			bodyString = string(bodyBytes)
		}
		return []Ccd{}, errors.New(fmt.Sprintf("%s:%s",reqErr, bodyString))
	}

	if response.StatusCode == 200 {
		var container struct {
			Status string `json:"status"`
			Ccds []Ccd `json:"data"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return []Ccd{}, err
		}
		// else
		return container.Ccds, nil
	} else {
		var container struct {
			Status string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return []Ccd{}, err
		}
		return []Ccd{}, errors.New(fmt.Sprintf("%s:%s",container.Message, reqErr))
	}
}

func (opn *OpenVpnApi) CcdExists(commonName string) (bool, error){
	if _, err := opn.CcdGet(commonName); err != nil {
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