package utils

import (
	"encoding/hex"
	"errors"
	"net"
	"strings"
)

func GeneratePacket(mac string) ([]byte, error, int) {
	mac = strings.ReplaceAll(mac, ":", "")
	macBytes, err := hex.DecodeString(mac)
	if err != nil {
		return nil, err, 500
	}
	if len(macBytes) != 6 {
		return nil, errors.New("invalid MAC address"), 400
	}
	packet := make([]byte, 102)
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}

	for i := 0; i < 16; i++ {
		copy(packet[6+i*6:], macBytes)
	}

	return packet, nil, 200

}

func SendPacket(packet []byte) error {
	conn, err := net.Dial("udp", "255.255.255.255:9")
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet)
	return err
}
