package main

import (
	//"database/sql"
	_ "expvar"
	//"fmt"
	"net/http"
	"os"

	"aph-go-service-master/transport"

	"github.com/go-kit/kit/log"

	_ "github.com/lib/pq"
)

func main() {

	transport.GetInitDb()
	

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}
