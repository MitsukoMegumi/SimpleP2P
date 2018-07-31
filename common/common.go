package common

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	upnp "github.com/NebulousLabs/go-upnp"
	"github.com/mitsukomegumi/GoP2P/fastping"
)

const (
	// NodeAvailableRep - global definition for value of node availability
	NodeAvailableRep = 10
)

/*
	BEGIN EXPORTED METHODS:
*/

// CheckAddress - check that specified IP address can be pinged, and is available on specified port
func CheckAddress(address string) error {
	p := fastping.NewPinger()                          // Create new pinger
	ipAddress, err := net.ResolveIPAddr("ip", address) // Resolve address
	p.AddIPAddr(ipAddress)                             // Add address to pinger

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) { // On correct address
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt) // Print address meta
		fmt.Printf("IP %s tested successfully \n", addr.String())        // Print address meta
	}
	p.OnIdle = func() { // On address timeout
		err = errors.New("Timed out with IP " + ipAddress.String() + "\n") // Assign meta to error
	}

	if err != nil { // Check for error
		return err // Return found error
	}

	err = p.Run()   // Assign to error
	if err != nil { // Check for errors
		if strings.Contains(err.Error(), "operation not permitted") { // Check for specific error
			return errors.New("operation requires root privileges") // Return custom error
		}

		return err // Return error
	}

	return nil // No error found, return nil
}

// GetExtIPAddrWithUpNP - retrieve the external IP address of the current machine via upnp
func GetExtIPAddrWithUpNP() (string, error) {
	// connect to router
	d, err := upnp.Discover()
	if err != nil { // Check for errors
		return "", err // return error
	}

	// discover external IP
	ip, err := d.ExternalIP()
	if err != nil { // Check for errors
		return "", err // return error
	}
	return ip, nil
}

// GetExtIPAddrWithoutUpNP - retrieve the external IP address of the current machine w/o upnp
func GetExtIPAddrWithoutUpNP() (string, error) {
	ip := make([]byte, 100) // Create IP buffer

	resp, err := http.Get("http://checkip.amazonaws.com/") // Attempt to check IP via aws
	if err != nil {                                        // Check for errors
		return "", err // Return error
	}

	defer resp.Body.Close()     // Close connection
	_, err = resp.Body.Read(ip) // Read IP

	if err != nil { // Check for errors
		return "", err // Return error
	}

	return string(ip[:len(ip)]), nil // Return ip
}

// GetCurrentTime - get current time in the UTC format
func GetCurrentTime() time.Time {
	return time.Now().UTC() // Returns current time in UTC
}

/*
	END EXPORTED METHODS
*/
