package main

import (
	"fmt"
	"net"
)

// func main() {
// 	_, ipnet1, _ := net.ParseCIDR("192.168.0.0/16")
// 	_, ipnet2, _ := net.ParseCIDR("192.255.0.0/16")
// 	fmt.Println("Overlapping:", isOverlapping(ipnet1, ipnet2))

// 	ipmask := net.CIDRMask(12, 32)
// 	fmt.Println("IP Mask:", ipmask)
// }

// func isOverlapping(ipnet1 *net.IPNet, ipnet2 *net.IPNet) bool {
// 	for i := range ipnet1.IP {
// 		if (ipnet1.IP[i] & ipnet1.Mask[i]) != (ipnet2.IP[i] & ipnet1.Mask[i]) {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	// _, ipnet, _ := net.ParseCIDR("192.168.0.0/16")
	// for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
	// 	fmt.Println(ip)
	// }
	fmt.Printf("%b\n", 239)
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
