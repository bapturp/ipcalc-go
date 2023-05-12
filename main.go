package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func BitCountMaskToN(s string) uint32 {
	// converts a bit count mask (i.e. 24) to integer
	mask, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	if mask < 0 && mask > 32 {
		log.Fatalf("Invalid CIDR mask: %d", mask)
	}

	return math.MaxUint32 << (32 - mask)
}

func DotDecToN(s string) (ip uint32) {
	// converts a quad-dotted decimal string (i.e. "192.168.0.1") to a number
	var byteArr [4]byte

	for i, s := range strings.Split(s, ".") {
		b, err := strconv.ParseUint(s, 10, 8) // uint8
		if err != nil {
			log.Fatal(err)
		}
		byteArr[i] = byte(b)
	}

	ip += uint32(byteArr[0]) << 24
	ip += uint32(byteArr[1]) << 16
	ip += uint32(byteArr[2]) << 8
	ip += uint32(byteArr[3])

	return ip
}

func IntToDotDec(n uint32) string {
	// converts a number to a quad-dotted decimal string
	var bytes [4]byte

	bytes[0] = byte((n >> 24) & 0xff)
	bytes[1] = byte((n >> 16) & 0xff)
	bytes[2] = byte((n >> 8) & 0xff)
	bytes[3] = byte(n & 0xff)

	return fmt.Sprintf("%d.%d.%d.%d", bytes[0], bytes[1], bytes[2], bytes[3])

}

func NetworkAddress(ip, mask uint32) uint32 {
	return ip & mask
}

func BroadcastAddress(ip, mask uint32) uint32 {
	return ip | (^mask & math.MaxUint32)
}

func ValidateAdress(s string) {
	sarr := strings.Split(s, ".")
	if len(sarr) != 4 {
		log.Fatal("Invalid IP address")
	}

	for _, v := range sarr {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Invalid IP address")
		}
		if tmp < 0 || tmp > 255 {
			log.Fatal("Invalid IP address")
		}
	}
}

func ValidateMask(s string) {
	mask, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Invalid netmask")
	}
	if mask < 0 || mask > 32 {
		log.Fatal("Invalid netmask")
	}
}

func main() {
	args := os.Args[1:]

	argv := strings.Split(args[0], "/")

	ValidateAdress(argv[0])
	ValidateMask(argv[1])

	ip := DotDecToN(argv[0])
	mask := BitCountMaskToN(argv[1])
	networdAddr := NetworkAddress(ip, mask)
	broadcastAddr := BroadcastAddress(ip, mask)

	fmt.Printf("Address:\t%s\n", IntToDotDec(ip))
	fmt.Printf("Netmask:\t%s\n", IntToDotDec(mask))
	fmt.Println("=>")
	fmt.Printf("Network:\t%s\n", IntToDotDec(networdAddr))
	fmt.Printf("Broadcast:\t%s\n", IntToDotDec(broadcastAddr))
	fmt.Printf("HostMin:\t%s\n", IntToDotDec(networdAddr+1))
	fmt.Printf("HostMax:\t%s\n", IntToDotDec(broadcastAddr-1))
	fmt.Printf("Hosts/Net:\t%d\n", ^mask-1)
}
