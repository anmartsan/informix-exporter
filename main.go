package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	conf "github.com/anmartsan/informix-exporter/config"
	exporter "github.com/anmartsan/informix-exporter/exporter"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	configfile = flag.String("configfile", "config.yaml", "Configuration File")
	serverfile = flag.String("configserver", "servers.yaml", "Servers File")
	puerto     = flag.String("port", "8087", "Listen Port")
	ConfigYaml *conf.ConfigYaml
	db         sql.DB
	Instances  *conf.InstanceList
)

func main() {

	ConfigYaml, err := conf.LoadConfig(configfile)
	if err != nil {
		log.Fatal("Error en  fichero Yaml:", err)

	}

	Instances, err = conf.LoadConfig2(serverfile)
	if err != nil {
		log.Fatal("Error en  fichero Yaml:", err)

	}
	fmt.Println(Instances)

	e := exporter.NewExporter(ConfigYaml, Instances)
	prometheus.MustRegister(e)

	log.Println("Arrancando servidor V0.3 en puerto:", *puerto)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+*puerto, nil))
	os.Exit(0)

}
