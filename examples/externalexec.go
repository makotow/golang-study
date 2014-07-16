package examples

import "os/exec"

// cwd
// cmd
// return: return
func ExecuteCommand(cmdpath, cmd string) error {
	command := exec.Command(cmd)
	command.Dir = cmdpath
	err := command.Run()
	return err
}

//func main() {
//
//	// thanks http://hackgeo.com/foss/sqlite-how-to-import-csv
//	// sqlite3 -separator , HackGeo.db ".import Counties.csv counties"
//	pwd := filepath.Join(os.Getenv("HOME"), "tmp", "db", "HackGeo")
//	command := "./load.sh"
//
//	cmd := exec.Command(command)
//	cmd.Dir = pwd
//
//	err := cmd.Run()
//	if err != nil {
//		log.Fatal("failed to execute external command. %s", err)
//		os.Exit(-1)
//	}
//}
