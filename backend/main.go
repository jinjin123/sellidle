package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
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

func InitialScan(hostname string) []PortStatus {
	var results []PortStatus
	for i := 17001; i <= 17005; i++ {
		results = append(results, ScanPort("tcp", i, hostname))
	}
	return results
}

func CheckPort(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	portlist := InitialScan("localhost")
	var data = &PortList{
		Status: 200,
		Port:   portlist,
	}
	var ret, _ = json.Marshal(&data)
	w.Header().Set("Content-Type", "application-json")
	w.Write(ret)

}
func ProxyPort(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	fmt.Printf(" got / request. body:\n%s\n", body)
	//var data = &PortList{
	//	Status: 200,
	//
	//}
	//var ret, _ = json.Marshal(&data)
	w.Header().Set("Content-Type", "application-json")
	w.Write(body)
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
