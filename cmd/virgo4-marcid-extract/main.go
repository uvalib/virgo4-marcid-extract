package main

import (
	"io"
	"log"
	"os"
)

//
// main entry point
//
func main() {

	cfg := LoadConfiguration()
	localName := cfg.InFileName

	var err error

	loader, err := NewRecordLoader(cfg.InFileName, localName)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := loader.First(true)
	for {
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		id, err := rec.Id()
		if err != nil {
			log.Fatal(err)
		}

		if id == cfg.Id {

			err = writeRecord( cfg, rec )
			if err != nil {
				log.Fatal(err)
			}

			loader.Done()

			log.Printf("Written %s, terminating normally", cfg.OutFileName )

			// terminate normally
			os.Exit( 0 )		}

		rec, err = loader.Next(true)
		if err == io.EOF {
			break
		}
	}

	loader.Done()

	log.Printf("WARNING: Failed to locate ID %s", cfg.Id )

	// ID not located
	os.Exit( 1 )
}

func writeRecord( config *ServiceConfig, record Record ) error {

	outputFile, err := os.Create( config.OutFileName )
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// write the payload
	_, err = outputFile.Write( record.Raw())
	if err != nil {
		return err
	}

	return nil
}

//
// end of file
//
