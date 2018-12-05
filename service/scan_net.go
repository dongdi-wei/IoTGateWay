package service

import (
	"IoTGateWay/consts"
	"errors"
	"fmt"
	"github.com/j-keck/arping"
	"net"
	"regexp"
	"strings"
	"sync"
	"time"
)

type NetScanner struct {
}

type Device struct {
	Ip   string
	Mac  string
	Type int
}
type InterFace struct {
	Name string
	Ip   string
	Mac  string
	Type int
}

func NewDevice(ip, mac string, t int) *Device {
	device := new(Device)
	device.Ip = ip
	device.Mac = mac
	device.Type = t
	return device
}

func NewInterFace(name, ip, mac string, t int) *InterFace {
	inter := new(InterFace)
	inter.Name = name
	inter.Ip = ip
	inter.Mac = mac
	inter.Type = t
	return inter
}

var (
	scannerOnce sync.Once
	netScan     *NetScanner
)

func GetNetScanner() *NetScanner {
	scannerOnce.Do(func() {
		netScan = new(NetScanner)
	})
	return netScan
}

func (s *NetScanner) Detect(ip string) ([]*Device , error){
	//校验是不是ip
	if m, err := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",
		ip); !m||err != nil {
			return nil,errors.New("input ip is not standard")
	}
	retList := []*Device{}
	//轮询ip
	ipFormat := ip[:strings.LastIndex(ip, ".")+1] + "%d"
	for i := 1; i <= 255; i++ {
		nextIp := fmt.Sprintf(ipFormat, i)
		if nextIp != ip {
			hwAddr, duration, err := s.Mac(nextIp)
			if err == arping.ErrTimeout {
				Logger.Error("IP %s is offline.\n", nextIp)
			} else if err != nil {
				Logger.Error("IP %s :%s\n", nextIp, err.Error())
			} else {
				Logger.Info("%s (%s) %d use", nextIp, hwAddr, duration/1000)
				if err != nil {
					Logger.Error("scanner Detect err:%v", err)
					continue
				}
				retList = append(retList, NewDevice(nextIp, hwAddr.String(), consts.TYPE_OTHER_DEVICE))
			}
		}
	}
	return retList,nil

}

func (s *NetScanner)Mac(ip string) (net.HardwareAddr, time.Duration, error) {
	dstIP := net.ParseIP(ip)
	return arping.Ping(dstIP)
}

func (s *NetScanner)ExternalIP() (string, string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", "", err
	}
	for _, iface := range ifaces {
		fmt.Println(iface)
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), iface.HardwareAddr.String(), nil
		}
	}
	return "", "", errors.New("not connected to the network")
}

func (s *NetScanner)InterFaces() ([]*InterFace, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	retList := []*InterFace{}
	for _, iface := range ifaces {
		fmt.Println(iface)
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			retList = append(retList, &InterFace{iface.Name, ip.String(), iface.HardwareAddr.String(), consts.TYPE_OWN_DEVICE})
		}
	}
	if retList != nil {
		return retList, nil
	}
	return nil, errors.New("not connected to the network")
}
