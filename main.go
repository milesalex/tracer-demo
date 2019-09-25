package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/milesalex/tracer-demo/util"
	"github.com/opentracing/opentracing-go"
)

func main() {
	// init tracer
	closer, err := util.InitTracer()
	if err != nil {
		log.Fatalf("failed to new tracer, err: %v", err)
		return
	}
	defer closer.Close()

	http.HandleFunc("/", HelloServer)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	tracer := opentracing.GlobalTracer()
	clientSpan := tracer.StartSpan("client")

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	fmt.Printf("Responded: Hello, %s!\n", r.URL.Path[1:])
	clientSpan.Finish()
}
