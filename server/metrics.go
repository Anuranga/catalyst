package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/kosatnkn/catalyst/app/config"
	"github.com/kosatnkn/catalyst/app/container"
	"github.com/kosatnkn/catalyst/app/metrics"
)

// Expose metrics as a seperate Prometheus metric server.
func exposeMetrics(cfg config.AppConfig, ctr *container.Container) {

	if !cfg.Metrics.Enabled {
		return
	}

	// register defined metrics
	metrics.Init()

	// set metric exposing port and endpoint
	address := cfg.Host + ":" + strconv.Itoa(cfg.Metrics.Port)
	http.Handle(cfg.Metrics.Route, promhttp.Handler())

	// run metric server in a goroutine so that it doesn't block
	go func() {

		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Println(err)
			panic("Metric server error...")
		}
	}()

	fmt.Println(fmt.Sprintf("Exposing metrics on %v ...", address))
}
