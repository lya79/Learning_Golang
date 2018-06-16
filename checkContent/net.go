package checkContent

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

//ex: addr := "123.123.123.123:456"
func CheckIPv4Addr(addr string) error {
	if IsEmptyString(&addr) {
		return errors.New("addr is empty")
	}

	strArray := strings.Split(addr, ":")

	if len(strArray) != 2 {
		return errors.New("Invalid addr:" + addr)
	}

	if IsEmptyString(&strArray[0]) {
		return errors.New("Invalid addr:" + addr)
	}

	if IsEmptyString(&strArray[1]) {
		return errors.New("Invalid addr:" + addr)
	}

	var error error

	ipv4 := strArray[0]
	error = CheckIPv4(ipv4)
	if error != nil {
		return error
	}

	var port int64
	port, error = strconv.ParseInt(strArray[1], 10, 64)
	if error != nil {
		return error
	}

	error = CheckPort(port)
	if error != nil {
		return error
	}

	return nil
}

func CheckIPv4(ipv4 string) error {
	if IsEmptyString(&ipv4) {
		return errors.New("Invalid ip:" + ipv4)
	}

	ip := net.ParseIP(ipv4)
	if ip.To4() == nil {
		return errors.New(ipv4 + " is not an IPv4 address")
	}
	return nil
}

func CheckPort(port int64) error {
	if port < 0 || port > 65535 {
		str := strconv.FormatInt(port, 10)
		return errors.New("Invalid port:" + str)
	}
	return nil
}
