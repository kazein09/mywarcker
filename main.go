package main

import (
    // Standard library packages
    "fmt"
    "strconv"
    "log"
    "net"
    "net/http"

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "github.com/skratchdot/open-golang/open"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}
