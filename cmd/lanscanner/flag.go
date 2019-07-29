package main

import (
	"net"
	"strconv"
	"encoding/binary"
	"errors"
)

type ipFlags []uint32

func(a * ipFlags) String() string {
	return "ip list"
}

func(a * ipFlags) Set(v string) error {
	ipn, err := strtoipn(v)
	*a = append(*a, ipn)
	return err
}

type portFlags []uint16

func(a * portFlags) String() string {
	return "port list"
}

func(a * portFlags) Set(v string) error {
	n, err := strconv.Atoi(v)
	*a = append(*a, uint16(n))
	return err
}

func strtoipn(str string) (uint32, error) {
	ip := net.ParseIP(str)
	if ip == nil {
		return 0, errors.New("Couldn't parse IP string '" + str + "'.")
	}
	switch len(ip) {
	case 16:
		return binary.BigEndian.Uint32(ip[12:]), nil
	default:
		return binary.BigEndian.Uint32(ip), nil
	}
}