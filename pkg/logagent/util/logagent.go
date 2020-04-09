package util

import (
	"io"
)

const(
	LogagentPort = "8090"
)




//func GetPodFileTree(req FileNodeRequest, ip string) string {
//	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
//	jsonStr, err := json.Marshal(req)
//	if err != nil {
//		log.Errorf("unable to marshal request to json %v", err)
//		return ""
//	}
//	url := "http://" + ip + ":" + LogagentPort + "/v1/logfile/directory"
//	log.Infof("url is %v", url)
//	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//	httpReq.Header.Set("Content-Type", "application/json")
//	if err != nil {
//		log.Errorf("unable to generate request %v", err)
//		return ""
//	}
//
//	client := &http.Client{}
//	resp, err := client.Do(httpReq)
//	if err != nil {
//		log.Errorf("unable to connect to log-agent %v", err)
//		return ""
//	}
//	body, _ := ioutil.ReadAll(resp.Body)
//	return string(body)
//}

type ReaderCloserGetter interface {
	GetReaderCloser() io.ReadCloser
}


//func GetPodReader(req FileNodeRequest, ip string) io.ReadCloser {
//	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
//	jsonStr, err := json.Marshal(req)
//	if err != nil {
//		log.Errorf("unable to marshal request to json %v", err)
//		return nil
//	}
//	url := "http://" + ip + ":" + LogagentPort + "/v1/logfile/directory"
//	log.Infof("url is %v", url)
//	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
//	httpReq.Header.Set("Content-Type", "application/json")
//	if err != nil {
//		log.Errorf("unable to generate request %v", err)
//		return nil
//	}
//
//	client := &http.Client{}
//	resp, err := client.Do(httpReq)
//	if err != nil {
//		log.Errorf("unable to connect to log-agent %v", err)
//		return nil
//	}
//	return resp.Body
//}
