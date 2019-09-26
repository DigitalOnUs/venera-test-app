package localnet

import (
	"fmt"
	"log"
	"net"
)

func GetLocalAddresses() ([]string, error) {
	IPs := make(map[string]struct{})

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range ifaces {
		// interface
		address, err := i.Addrs()
		if err != nil {
			log.Println(err)
			continue
		}
		// addresses to fetch
		for _, a := range address {
			switch v := a.(type) {
			case *net.IPAddr:
				maskSize, _ := v.IP.DefaultMask().Size()
				IPs[fmt.Sprintf("%s/%d", v.IP, maskSize)] = struct{}{}
			case *net.IPNet:
				maskSize, _ := v.Mask.Size()
				IPs[fmt.Sprintf("%s/%d", v.IP, maskSize)] = struct{}{}
			}
		}
	}

	List := make([]string, len(IPs))
	i := 0
	for value := range IPs {
		List[i] = value
		i++
	}

	return List, nil
}
