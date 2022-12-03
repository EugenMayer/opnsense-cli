package unbound

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (opn *UnboundApi) ServiceRestart() error {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "service", "restart")

	// create our Request
	request, reqCreationErr := http.NewRequest("POST", endpoint, nil)
	if reqCreationErr != nil {
		return reqCreationErr
	}

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return reqErr
	}

	if response.StatusCode == 200 {
		// else
		return nil
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("%s:%s", container.Message, reqErr))
	}
}

func (opn *UnboundApi) ServiceReconfigure() error {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "service", "reconfigure")

	// create our Request
	request, reqCreationErr := http.NewRequest("POST", endpoint, nil)
	if reqCreationErr != nil {
		return reqCreationErr
	}

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return reqErr
	}

	if response.StatusCode == 200 {
		// else
		return nil
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("%s:%s", container.Message, reqErr))
	}
}

func (opn *UnboundApi) ServiceStatus() (string, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "service", "status")

	// create our Request
	request, reqCreationErr := http.NewRequest("GET", endpoint, nil)
	if reqCreationErr != nil {
		return "", reqCreationErr
	}

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		var container struct {
			Status string `json:"status"`
		}
		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return "", err
		}
		// else
		return container.Status, nil
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
}
