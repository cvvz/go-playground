package main

import (
	"fmt"
	"net"
	"os"

	"sigs.k8s.io/cloud-provider-azure/pkg/nodeipam/ipam/cidrset"
)

func main() {
	vnetAddressSpace := "10.224.0.0/12"

	existingSubnetAddressSpaces := []string{
		"10.224.0.0/16",
		"10.225.0.0/16",
		"10.226.0.0/16",
		"10.227.0.0/16",
		"10.227.0.0/16",
		"10.228.0.0/16",
		"10.229.0.0/16",
		"10.230.0.0/16",
		"10.231.0.0/16",
		"10.232.0.0/16",
		"10.233.0.0/16",
		"10.234.0.0/16",
		"10.235.0.0/16",
		"10.236.0.0/16",
		"10.237.0.0/16",
		"10.238.0.0/16",
		// "10.239.0.0/15",
	}
	cidr, err := getAvailableSubnetAddressRange(vnetAddressSpace, existingSubnetAddressSpaces, 16)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(cidr)

}
func getAvailableSubnetAddressRange(vnetAddressSpace string, existingSubnetAddressSpaces []string, maskSize int) (string, error) {
	_, vnetCIDR, _ := net.ParseCIDR(vnetAddressSpace)
	cidrSet, err := cidrset.NewCIDRSet(vnetCIDR, maskSize)
	if err != nil {
		return "", err
	}

	var cidr *net.IPNet
	for cidr, err = cidrSet.AllocateNext(); err == nil; cidr, err = cidrSet.AllocateNext() {
		// 判断cidr和subnetAddressSpaces是否overlap
		for _, subnetAddressSpace := range existingSubnetAddressSpaces {
			_, subnetCIDR, _ := net.ParseCIDR(subnetAddressSpace)
			if cidrIntersect(cidr, subnetCIDR) {
				goto CONTINUE
			}
		}
		return cidr.String(), nil
	CONTINUE:
	}

	return "", err
}

func cidrIntersect(cidr1, cidr2 *net.IPNet) bool {
	if cidr1.Contains(cidr2.IP) || cidr2.Contains(cidr1.IP) {
		fmt.Printf("cidr1: %s, cidr2: %s\n", cidr1, cidr2)
		return true
	}
	return false
}
