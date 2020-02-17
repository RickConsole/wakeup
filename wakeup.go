package wakeup

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"runtime"
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
	//Converting response into string result
	body, _ := ioutil.ReadAll(response.Body)
	vendor := string(body)
	return vendor
}

func CheckUsingMAC() bool {
	//Do basic check if vendor matches VMware or virtualbox
	//Returns false if machine is a VM
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

func CheckUsingSysinfo() bool {
	//Uses system info commands to detect VM probability
	//Returns false if machine is a VM
	opsys := runtime.GOOS

	if opsys == "windows" {
		cmd := exec.Command("SYSTEMINFO")
		out, err := cmd.CombinedOutput()

		if err != nil {
			panic(err)
		} else if strings.Contains(string(out), "VMware") {
			return false
		} else if strings.Contains(string(out), "VirtualBox") {
			return false
		} else if strings.Contains(string(out), "hypervisor") {
			return false
		}

	} else if opsys == "linux" {
		cmd := exec.Command("cat", "/sys/class/dmi/id/product_name")
		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		} else if strings.Contains(string(out), "VMware") {
			return false
		} else if strings.Contains(string(out), "Virtual") {
			return false
		}

	}
	return true
}