package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jae2274/goutils/llog"
)

type apiLogger struct {
	postUrl string
}

func (al *apiLogger) Log(llog *llog.LLog) error {
	return postApi(al.postUrl, llog)
}

func postApi(postUrl string, llog *llog.LLog) error {
	b, err := json.Marshal(llog)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(b)

	resp, err := http.Post(postUrl, "application/json", buff)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("status code is not 201, but %d", resp.StatusCode)
	}
	return nil
}
