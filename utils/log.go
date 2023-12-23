package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Out = getOutFilePath()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.DebugLevel)
}

func getOutFilePath() *os.File {
	dir, err := os.Getwd()
	if err != nil {
		return os.Stdout
	}

	dirpath := dir + "/logs/"
	err = os.MkdirAll(dirpath, 0777)
	if err != nil {
		fmt.Println(err)
		return os.Stdout
	}

	now := time.Now()
	filepath := now.Format("2006-01-02") + ".log"

	file, err := os.OpenFile(dirpath+filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return file
}
