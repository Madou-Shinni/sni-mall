package ini

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var Config *ini.File
var iniErr error

func init() {
	// ini
	Config, iniErr = ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

}
