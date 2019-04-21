package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
        "os"
)

type ServerVersion struct {
	Version string
	IP string
	Port  int
}


var ser ServerVersion
func (this *ServerVersion)sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is new version : %s\n", this.Version)
	log.Printf("This is new version : %s", ser.Version)
}

func init(){
        fmt.Println("TEMP file : ", os.TempDir())
	flag.StringVar(&ser.Version, "version", "1.0.0","set service version.")
	flag.StringVar(&ser.IP, "ip", "127.0.0.1","ip.")
	flag.IntVar(&ser.Port, "port", 9090,"port.")
	flag.Parse()
	log.Println("This is NEW version :", ser.Version)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ser.sayhello)
	server := &http.Server{
		Addr:fmt.Sprintf("%s:%d", ser.IP, ser.Port),
		Handler: mux,
	}
	server.ListenAndServe()
	log.Printf("Start listenAndServe...\n IP:%s\n PORT: %d \n Version: %s", ser.IP, ser.Port, ser.Version)
}

