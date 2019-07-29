package main

import (
	"flag"
	"net/http"
	"time"
	"github.com/rbxb/lanscanner"
	"strconv"
	"path/filepath"
	"os"
)

var (
	ipns ipFlags
	ri bool
	ports portFlags
	rp bool
	savepath string
	delay uint64
	timeout uint64
)

func init() {
	flag.Var(&ipns, "ip", "List/range of IPs. (*required*)")
	flag.Var(&ports, "port", "List/range of ports. (*required*)")
	flag.BoolVar(&ri, "rip", true, "Range IPs?")
	flag.BoolVar(&rp, "rport", false, "Range ports?")
	flag.StringVar(&savepath, "save", "./responses", "Responses save path.")
	flag.Uint64Var(&delay, "delay", 100, "Request delay duration in milliseconds.")
	flag.Uint64Var(&timeout, "timeout", 4, "Request timeout duration in seconds.")
}

func responseHandler() lanscanner.HandleResponseFunc {
	uniqueName := make(chan int, 1)
	uniqueName <- 0
	return func(req * http.Request, resp * http.Response, err error){
		if err != nil {
			return
		}
		n := <- uniqueName
		uniqueName <- n + 1
		name := filepath.Join(savepath,strconv.Itoa(n))
		if err = saveResponse(req,resp,name); err != nil {
			panic(err)
		}
	}
}

func main() {
	flag.Parse()
	savepath = filepath.Join(savepath, strconv.Itoa(int(time.Now().Unix())))
	if err := os.MkdirAll(savepath, os.ModePerm); err != nil {
		panic(err)
	}
	scn := lanscanner.Scanner{
		Client: http.Client{
			Timeout: time.Second * time.Duration(timeout),
		},
		HandleFunc: responseHandler(),
		SleepDuration: time.Millisecond * time.Duration(delay),
	}
	if err := scn.Scan(ipns,ri,ports,rp); err != nil {
		panic(err)
	}
}