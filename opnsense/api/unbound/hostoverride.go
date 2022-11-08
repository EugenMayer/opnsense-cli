package unbound

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	coreapi "github.com/eugenmayer/opnsense-cli/opnsense/api"
	"io"
	"net/http"
)

//goland:noinspection GoNameStartsWithPackageName
type UnboundApi struct {
	*coreapi.OPNsense
}

type HostOverride struct {
	// 0 for disabled, 1 for enabled
	Enabled     string `json:"enabled"`
	Host        string `json:"hostname"`
	Domain      string `json:"domain"`
	Ip          string `json:"server"`
	Rr          string `json:"rr"`          // A, MX, CNAME...
	Mxprio      string `json:"mxprio"`      // 10, 20
	Mx          string `json:"mx"`          // mail.domain.tld ...
	Description string `json:"description"` // any arbitrary text
	Uuid        string `json:"uuid"`
}

func (opn *UnboundApi) HostOverrideCreateOrUpdate(hostOverride HostOverride) (string, error) {
	if hostOverride.Uuid == "" { // no uuid given use host / domain based fuzzy search
		var searchResult, _ = opn.HostEntryGetByFQDN(hostOverride.Host, hostOverride.Domain)

		if searchResult.Uuid != "" {
			fmt.Println(fmt.Sprintf("Found entry with same FQDN, doing update with uuid: %s", searchResult.Uuid))
			hostOverride.Uuid = searchResult.Uuid
		}
	}

	if hostOverride.Uuid == "" {
		return opn.HostOverrideCreate(hostOverride)
	} else {
		return opn.HostOverrideUpdate(hostOverride)
	}
}

func (opn *UnboundApi) HostOverrideUpdate(hostOverride HostOverride) (string, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "settings", "setHostOverride")
	var fullPath = fmt.Sprintf("%s/%s", endpoint, hostOverride.Uuid)

	var container struct {
		HostOverride HostOverride `json:"host"`
	}

	container.HostOverride = hostOverride
	// create our Request
	jsonBody := new(bytes.Buffer)
	if err := json.NewEncoder(jsonBody).Encode(container); err != nil {
		return "", err
	}

	request, reqCreationErr := http.NewRequest("POST", fullPath, jsonBody)
	if reqCreationErr != nil {
		return "", reqCreationErr
	}
	request.Header.Set("Content-Type", "application/json")

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return "", reqErr
	}

	if response.StatusCode == 200 {
		var resultContainer struct {
			ResultStatus string `json:"result"`
			Uuid         string `json:"uuid"`
		}

		if err := json.NewDecoder(response.Body).Decode(&resultContainer); err != nil {
			return "", err
		}
		// else
		return resultContainer.Uuid, nil
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
}

func (opn *UnboundApi) HostOverrideCreate(hostOverride HostOverride) (string, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "settings", "addHostOverride")

	var container struct {
		HostOverride HostOverride `json:"host"`
	}

	container.HostOverride = hostOverride
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
		var resultContainer struct {
			ResultStatus string `json:"result"`
			Uuid         string `json:"uuid"`
		}

		if err := json.NewDecoder(response.Body).Decode(&resultContainer); err != nil {
			return "", err
		}
		// else
		return resultContainer.Uuid, nil
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
}

func (opn *UnboundApi) HostEntryGetByFQDN(host string, domain string) (HostOverride, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "settings", "searchHostOverride")

	// we search for the host first and compare the domain later (searchPhrase does not support searching for FQDN)
	var reqUrl = fmt.Sprintf("%s?searchPhrase=%s", endpoint, host)
	// create our Request
	var request, reqCreationErr = http.NewRequest("GET", reqUrl, nil)

	if reqCreationErr != nil {
		return HostOverride{}, reqCreationErr
	}

	// send it to the server
	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		bodyBytes, _ := io.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return HostOverride{}, errors.New(fmt.Sprintf("%s:%s", bodyString, reqErr))
	}

	if response.StatusCode == 200 {
		var container struct {
			Overrides []HostOverride `json:"rows"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return HostOverride{}, err
		}
		// else

		var allWithMatchingDomain = Filter(container.Overrides, func(override HostOverride) bool {
			return override.Domain == domain
		})

		if len(allWithMatchingDomain) > 1 {
			return HostOverride{}, &coreapi.TooManyFoundError{
				Err:  nil,
				Name: "found more then one entry",
			}
		}

		if len(allWithMatchingDomain) == 0 {
			return HostOverride{}, &coreapi.NotFoundError{
				Err:  nil,
				Name: "hostentry",
			}
		}

		// else
		return allWithMatchingDomain[0], nil
	} else if response.StatusCode == 404 {
		return HostOverride{}, &coreapi.NotFoundError{
			Err:  nil,
			Name: "hostentry",
		}
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return HostOverride{}, err
		}
		return HostOverride{}, errors.New(fmt.Sprintf("%s", container.Message))
	}
}

func (opn *UnboundApi) HostOverrideList() ([]HostOverride, error) {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "settings", "searchHostOverride")

	// create our Request
	var request, reqCreationErr = http.NewRequest("GET", endpoint, nil)

	if reqCreationErr != nil {
		return []HostOverride{}, reqCreationErr
	}

	// send it to the server
	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		bodyBytes, _ := io.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return []HostOverride{}, errors.New(fmt.Sprintf("%s:%s", bodyString, reqErr))
	}

	if response.StatusCode == 200 {
		var container struct {
			Overrides []HostOverride `json:"rows"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return []HostOverride{}, err
		}

		// else
		return container.Overrides, nil
	} else {
		var container struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		if err := json.NewDecoder(response.Body).Decode(&container); err != nil {
			return []HostOverride{}, err
		}
		return []HostOverride{}, errors.New(fmt.Sprintf("%s:%s", container.Message, reqErr))
	}
}

func (opn *UnboundApi) HostEntryRemove(uuid string) error {
	// endpoint
	var endpoint = opn.EndpointForPluginControllerMethod("unbound", "settings", "delHostOverride")
	var fullPath = fmt.Sprintf("%s/%s", endpoint, uuid)

	// create our Request
	request, reqCreationErr := http.NewRequest("POST", fullPath, nil)
	if reqCreationErr != nil {
		return reqCreationErr
	}

	var response, reqErr = opn.Send(request)
	if reqErr != nil {
		return reqErr
	}

	if response.StatusCode == 200 {
		var resultContainer struct {
			Result string `json:"result"`
		}

		if err := json.NewDecoder(response.Body).Decode(&resultContainer); err != nil {
			return err
		}
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

func (opn *UnboundApi) HostEntryExists(host string, domain string) (bool, error) {
	if _, err := opn.HostEntryGetByFQDN(host, domain); err != nil {
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

func Filter[T any](vs []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
