package uuid

import (
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/sony/sonyflake"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{
	MachineID: func() (uint16, error) { return podIP(), nil },
})

// SnowFlakeIDUint64 return sequence uuid
// 39 bits for time in units of 10 msec
// 8 bits for a sequence number
// 16 bits for a machine id
func SnowFlakeIDUint64() uint64 {
	id, _ := sf.NextID()
	return id
}

// SnowFlakeID is string format SnowFlakeIDUint64
func SnowFlakeID() string {
	return strconv.FormatUint(SnowFlakeIDUint64(), 10)
}

func podIP() uint16 {
	podIP := os.Getenv("POD_IP")
	if podIP == "" {
		podIP = RandomIpV4Address()
	}
	ip := net.ParseIP(podIP)
	return uint16(ip[8])<<7 + uint16(ip[9])<<6 +
		uint16(ip[10])<<5 + uint16(ip[11])<<4 +
		uint16(ip[12])<<3 + uint16(ip[13])<<2 +
		uint16(ip[14])<<1 + uint16(ip[15])
}

// RandomIpV4Address returns a valid IPv4 address as string
func RandomIpV4Address() string {
	var blocks []string
	for i := 0; i < 4; i++ {
		number := rand.Intn(255)
		blocks = append(blocks, strconv.Itoa(number))
	}
	return strings.Join(blocks, ".")
}
