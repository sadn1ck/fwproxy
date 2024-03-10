package http_proxy

import (
	"io"
	"log"
	"net/http"
)

var HOP_BY_HOP_HEADERS = []string{
	"Keep-Alive", "Transfer-Encoding", "TE", "Connection", "Trailer", "Upgrade", "Proxy-Authorization", "Proxy-Authenticate",
}

func removeHopByHopHeaders(header http.Header) http.Header {
	removeHeaderKeys := header.Values("Connection")
	removeHeaderKeys = append(removeHeaderKeys, HOP_BY_HOP_HEADERS...)
	for _, h := range removeHeaderKeys {
		header.Del(h)
	}
	return header
}

func proxyRequest(r *http.Request) (*http.Response, error) {
	proxyReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		return nil, err
	}
	proxyReq.Header = removeHopByHopHeaders(r.Header.Clone())
	proxyReq.Header.Add("Proxy-Connection", "keep-alive")
	proxyReq.Header.Add("X-Forwarded-For", r.RemoteAddr)

	res, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func cloneProxyResponseIntoWriter(res *http.Response, w http.ResponseWriter) {
	for header, values := range res.Header {
		for _, value := range values {
			w.Header().Add(header, value)
		}
	}
	w.WriteHeader(res.StatusCode)

	defer res.Body.Close()
	// Copy the body from the proxy response to the response writer
	io.Copy(w, res.Body)
}

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := proxyRequest(r)

		if err != nil {
			log.Printf("Error: %s", err)
			w.Write([]byte(err.Error()))
			return
		}

		cloneProxyResponseIntoWriter(res, w)

	})
	http.ListenAndServe(":4041", nil)
}
