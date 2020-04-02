package util

import (
	"context"
	"io"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	genericrest "k8s.io/apiserver/pkg/registry/generic/rest"
	"k8s.io/apiserver/pkg/registry/rest"
	"net/http"
	"net/url"
)


// LocationStreamer is a resource that streams the contents of a particular
// location URL.
type LocationStreamer struct {
	Ip              string
	Request         FileNodeRequest
	Location        *url.URL
	Transport       http.RoundTripper
	ContentType     string
	Flush           bool
	ResponseChecker genericrest.HttpResponseChecker
	RedirectChecker func(req *http.Request, via []*http.Request) error
}

// a LocationStreamer must implement a rest.ResourceStreamer
var _ rest.ResourceStreamer = &LocationStreamer{}

func (obj *LocationStreamer) GetObjectKind() schema.ObjectKind {
	return schema.EmptyObjectKind
}
func (obj *LocationStreamer) DeepCopyObject() runtime.Object {
	panic("rest.LocationStreamer does not implement DeepCopyObject")
}

// InputStream returns a stream with the contents of the URL location. If no location is provided,
// a null stream is returned.
func (s *LocationStreamer) InputStream(ctx context.Context, apiVersion, acceptHeader string) (stream io.ReadCloser, flush bool, contentType string, err error) {
	//if s.Location == nil {
	//	// If no location was provided, return a null stream
	//	return nil, false, "", nil
	//}
	//transport := s.Transport
	//if transport == nil {
	//	transport = http.DefaultTransport
	//}
	//
	//client := &http.Client{
	//	Transport:     transport,
	//	CheckRedirect: s.RedirectChecker,
	//}
	//req, err := http.NewRequest("GET", s.Location.String(), nil)
	//if err != nil {
	//	return nil, false, "", fmt.Errorf("failed to construct request for %s, got %v", s.Location.String(), err)
	//}
	//// Pass the parent context down to the request to ensure that the resources
	//// will be release properly.
	//req = req.WithContext(ctx)
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	return nil, false, "", err
	//}
	//
	//if s.ResponseChecker != nil {
	//	if err = s.ResponseChecker.Check(resp); err != nil {
	//		return nil, false, "", err
	//	}
	//}
	//
	//contentType = s.ContentType
	//if len(contentType) == 0 {
	//	contentType = resp.Header.Get("Content-Type")
	//	if len(contentType) > 0 {
	//		contentType = strings.TrimSpace(strings.SplitN(contentType, ";", 2)[0])
	//	}
	//}
	//flush = s.Flush
	//stream = resp.Body
	stream = GetPodReader(s.Request,s.Ip)
	flush = false
	contentType = s.ContentType
	return
}
