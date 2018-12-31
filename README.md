### go-httplogger

Golang apache common log format middleware. It sits between in the topmost level of your http mux stack. Logs to stdout.

### Example usage

	package main

	import (
	    "io"
	    "log"
	    "net/http"

	    httplogger "github.com/gleicon/go-httplogger"
	)

	func helloHandler(w http.ResponseWriter, r *http.Request) {
	    io.WriteString(w, "Hello world!")
	}

	func main() {
	    serveMux := http.NewServeMux()

	    serveMux.HandleFunc("/", helloHandler)
	    srv := http.Server{
		Addr:    ":8080",
		Handler: httplogger.HTTPLogger(serveMux),
	    }
	    log.Fatal(srv.ListenAndServe())
	}

 
### Example output

	2018/12/31 12:20:33 HTTP - 127.0.0.1 - - [31/Dec/2018:12:20:33 +0100] "GET / HTTP/1.1" 200 12 "" "Wget/1.17.1 (linux-gnu)" 100351us
