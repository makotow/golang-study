package main

import (
	"os"
	"log"
	"os/exec"

	"path/filepath"
	"fmt"
)

func main() {

	// thanks http://hackgeo.com/foss/sqlite-how-to-import-csv
	// sqlite3 -separator , HackGeo.db ".import Counties.csv counties"
	pwd := filepath.Join(os.Getenv("HOME"), "tmp", "db", "HackGeo")
	command := "./load.sh"

	cmd := exec.Command(command)
	cmd.Dir = pwd

	fmt.Println(cmd.Dir)
	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to execute external command. %s", err)
		os.Exit(-1)
	}
}
