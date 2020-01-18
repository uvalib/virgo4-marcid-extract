package main

import (
	"flag"
	"log"
)

// ServiceConfig defines all of the service configuration parameters
type ServiceConfig struct {
	InFileName    string
	OutFileName   string
	Id            string
}

// LoadConfiguration will load the service configuration from env/cmdline
// and return a pointer to it. Any failures are fatal.
func LoadConfiguration() *ServiceConfig {

	var cfg ServiceConfig

	flag.StringVar(&cfg.InFileName, "infile", "", "Input file name")
	flag.StringVar(&cfg.OutFileName, "outfile", "", "Output file name")
	flag.StringVar(&cfg.Id, "id", "", "id to extract")

	flag.Parse()

	if len(cfg.InFileName) == 0 {
		log.Fatalf("InFileName cannot be blank")
	}

	if len(cfg.OutFileName) == 0 {
		log.Fatalf("OutFileName cannot be blank")
	}

	if len(cfg.Id) == 0 {
		log.Fatalf("Id cannot be blank")
	}
	return &cfg
}

//
// end of file
//
