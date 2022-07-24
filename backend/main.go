package main

import (
	"encoding/json"
	"errors"
	"os"

	//"errors"
	"fmt"
	"os/exec"

	//"io/ioutil"
	"net"
	"net/http"
	//"os"
	"strconv"
	"time"

)

type PortStatus struct {
	Port  int
	State string
}
type PortList struct {
	Status int          `json:"status"`
	Port   []PortStatus `json:"data"`
}
type BindPort struct {
	Outside string    `json:"outside"`
	Inside  string    `json:"inside"`
	Ip      string `json:"ip"`
}

type InvokeScan interface {
	InitialScan(hostname string)  []PortStatus
	InitialScanSigle(hostname string, port int)  PortStatus
}
type Scan struct{
}

func (s Scan) InitialScan(hostname string) []PortStatus {
	var results []PortStatus
	for i := 17001; i <= 17999; i++ {
		results = append(results, ScanPort("tcp", i, hostname))
	}
	return results
}

func (s Scan) InitialScanSigle(hostname string, port int) PortStatus {
	var results PortStatus
	results =  ScanPort("tcp", port, hostname)
	return results
}

func ScanPort(protocol string, port int, hostname string) PortStatus {
	result := PortStatus{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 5*time.Second)
	if err != nil {
		result.State = "close"
		return result
	}
	defer conn.Close()
	result.State = "open"
	return result
}

func CheckPort(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	var invoke InvokeScan
	s := new(Scan)
	invoke = s
	portlist := invoke.InitialScan("localhost")
	var data = &PortList{
		Status: 200,
		Port:   portlist,
	}
	var ret, _ = json.Marshal(&data)
	w.Header().Set("Content-Type", "application-json")
	w.Write(ret)

}
func ProxyPort(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t  BindPort
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	fmt.Printf("data: %s, req:%s-%s\n",t,r.Header.Get("X-Real-Ip"),r.Header.Get("X-Forwarded-For"))
	tmp := make(map[string]interface{})
	if t.Ip == "192.168.1.110"{
		tmp["status"] = 401
		var ret ,_ = json.Marshal(&tmp)
		w.Write(ret)
		return
	}
	var invoke InvokeScan
	s := new(Scan)
	invoke = s
	intVar,err := strconv.Atoi(t.Outside)
	//var p PortStatus
	p := invoke.InitialScanSigle("localhost",intVar)
	fmt.Println(p)
	// port idle check
	if  p.Port < 17001 ||  p.Port >=17999  {
		tmp["data"] = 401
		var ret ,_ = json.Marshal(&tmp)
		w.Write(ret)
		return
	}
	if p.State != "close"  {
			tmp["data"] = 401
			var ret ,_ = json.Marshal(&tmp)
			w.Write(ret)
			return
	}
	go func(t BindPort){
		cmd := exec.Command("/bin/bash", "/root/vps/tpl.sh",t.Ip, t.Outside,t.Inside)
		_,err := cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}(t)

	w.Header().Set("Content-Type", "application-json")
	tmp["status"] = 200
	tmp["data"] = t
	var ret ,_ = json.Marshal(&tmp)
	w.Write(ret)
}

func main() {
	http.HandleFunc("/check", CheckPort)
	http.HandleFunc("/update", ProxyPort)

	err := http.ListenAndServe(":85", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
