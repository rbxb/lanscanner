package lanscanner

import (
	"errors"
	"net/http"
	"net"
	"encoding/binary"
	"strconv"
	"time"
	"sync"
)

type HandleResponseFunc func(req * http.Request, resp * http.Response, err error)

type Scanner struct {
	Client http.Client
	HandleFunc HandleResponseFunc
	SleepDuration time.Duration
}

func(scn * Scanner) Scan(ipns []uint32, ri bool, ports []uint16, rp bool) error {
	if ri && len(ipns) < 2 {
		return errors.New("IP range not given.")
	}
	if rp && len(ports) < 2 {
		return errors.New("Port range not given.")
	}
	wg := &sync.WaitGroup{}
	scn.iterateIps(ipns,ri,ports,rp,wg)
	wg.Wait()
	return nil
}

func(scn * Scanner) iterateIps(ipns []uint32, ri bool, ports []uint16, rp bool, wg * sync.WaitGroup) {
	if ri {
		for ipn := ipns[0]; ipn <= ipns[1]; ipn++ {
			scn.iteratePorts(ipn,ports,rp,wg)
		}
	} else {
		for _, ipn := range ipns {
			scn.iteratePorts(ipn,ports,rp,wg)
		}
	}
}

func(scn * Scanner) iteratePorts(ipn uint32, ports []uint16, rp bool, wg * sync.WaitGroup) {
	b := make(net.IP, 4)
	binary.BigEndian.PutUint32(b, ipn)
	ip := b.String()
	if rp {
		for port := ports[0]; port <= ports[1]; port++ {
			scn.get(ip,port,wg)
		}
	} else {
		for _, port := range ports {
			scn.get(ip,port,wg)
		}
	}
}

func(scn * Scanner) get(ip string, port uint16, wg * sync.WaitGroup) {
	addr := "http://" + ip + ":" + strconv.Itoa(int(port))
	req, _ := http.NewRequest("GET", addr, nil)
	wg.Add(1)
	go func() {
		resp, err := scn.Client.Do(req)
		if err == nil {
			scn.HandleFunc(req, resp, err)
		}
		wg.Done()
	}()
	time.Sleep(scn.SleepDuration)
}