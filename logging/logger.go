package logging

import (
	"github.com/8tomat8/GoRepost/task"
	"os"
	"encoding/json"
	"errors"
	"net/http"
	"bytes"
	"github.com/golang/glog"
	"net/url"
)

const (
	resultsPath = "./resultsPath/"
	fileFormat  = ".json"
)

func WriteLog(t *task.Task) (err error) {
	jsonData, err := json.Marshal(t)
	if err != nil {
		return
	}

	fInf, err := os.Stat(resultsPath)
	if err == nil {
		if !fInf.IsDir() {
			return errors.New("Can't create " + resultsPath + " folder!")
		}
	} else if os.IsNotExist(err) {
		err = os.Mkdir(resultsPath, os.FileMode(0777))
		if err != nil {
			return
		}
	} else {
		return
	}

	var f *os.File
	f, err = os.Create(resultsPath + t.Id + fileFormat)
	if err != nil {
		return
	}
	defer f.Close()

	var n int
	n, err = f.Write(jsonData)
	if err != nil || len(jsonData) != n {
		return
	}

	if _, err = url.Parse(t.CallBackUrl); err == nil {
		_, err := http.Post(t.CallBackUrl, "application/json", bytes.NewReader(jsonData))
		if err != nil {
			glog.Error(err)
		}
	}

	return
}

func GetLog(id *string) (*os.File, error) {
	file, err := os.Open(resultsPath + *id + fileFormat)
	if err != nil {
		return nil, err
	}
	return file, nil
}
