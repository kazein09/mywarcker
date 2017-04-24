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

ifaces, err := net.Interfaces()
// handle err
for _, i := range ifaces {
    addrs, err := i.Addrs()
    // handle err
    for _, addr := range addrs {
        var ip net.IP
        switch v := addr.(type) {
        case *net.IPNet:
                ip = v.IP
        case *net.IPAddr:
                ip = v.IP
        }
        // process IP address
    }
}
