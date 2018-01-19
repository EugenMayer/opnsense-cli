package opnsense

import (
	"net"
	"net/http"
	"fmt"
)


func (opn *OPNsense) CcdCreate(commonName string, ip net.IP, ip net.IPMask) {

}

func (opn *OPNsense) CcdExists(commonName string) (bool, error){
	var endpoint = opn.EndpointForPluginControllerMedthod("openvpn","ccd","getCcd")

	var request = fmt.Sprintf("%s/%s", endpoint, commonName)
	var resp, err = http.Get(request)

	if err != nil {
		return true, err
	}

	if resp.StatusCode == 200 {
		return true, nil
	}
	// else
	return false, nil
}

