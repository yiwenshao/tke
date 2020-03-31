package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tkestack.io/tke/pkg/util/log"
)

const(
	LogagentPort = "8090"
)

type FileNodeRequest struct {
	PodName string `json:"podName"`
	Namespace string `json:"namespace"`
	Container string `json:"container"`
}




func GetPodFileTree(req FileNodeRequest, ip string) string {
	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	jsonStr, err := json.Marshal(req)
	if err != nil {
		log.Errorf("unable to marshal request to json %v", err)
		return ""
	}
	url := "http://" + ip + ":" + LogagentPort + "/v1/logfile/directory"
	log.Infof("url is %v", url)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	httpReq.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Errorf("unable to generate request %v", err)
		return ""
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Errorf("unable to connect to log-agent %v", err)
		return ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
