package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f := fib()

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello Word!!!\n")

		for i := 1; i <= 500; i++ {
			io.WriteString(w, strconv.Itoa(f())+"\n")
		}

		log.Info("Hello world called")

	})
	http.ListenAndServe(":80", nil)
}

func currentTimeInMillis() int64 {
	tv := new(syscall.Timeval)
	syscall.Gettimeofday(tv)
	return (int64(tv.Sec)*1e3 + int64(tv.Usec)/1e3)
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
