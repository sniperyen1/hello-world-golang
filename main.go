package main

import (
	"io"
	"net/http"
	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello Word!!!")

		log.Info("Hello world called")

	})
	http.ListenAndServe(":80", nil)
}

func currentTimeInMillis() int64 {
	tv := new(syscall.Timeval)
	syscall.Gettimeofday(tv)
	return (int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3)
}
