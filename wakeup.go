package wakeup

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func GetMAC() string {
	var mac string
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				mac = i.HardwareAddr.String()
				break
			}
		}
	}
	return mac
}

func GetVendor() string {
	var mac string
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				mac = i.HardwareAddr.String()
				break
			}
		}
	}
	response, _ := http.Get("https://api.macaddress.io/v1?apiKey=at_Spa8oMJfznmFraFhS7gyTC6NO3qDU&output=vendor&search=" + mac)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	vendor := string(body)
	//Converting response into string result
	return vendor
}

func Check() bool {
	vendor := GetVendor()
	if strings.Contains(vendor, "VMware") {
		return false
	} else if strings.Contains(vendor, "Pcs Systemtechnik GmbH") {
		return false
	} else if strings.Contains(vendor, "Oracle Corp") {
		return false
	} else {
		return true
	}
}