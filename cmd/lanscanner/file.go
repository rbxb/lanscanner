package main

import (
	"os"
	"io/ioutil"
	"bytes"
	"net/http"
	"net/http/httputil"
	"time"
)

func saveResponse(req * http.Request, resp * http.Response, name string) error {
	if err := saveResponseMeta(req,resp,name); err != nil {
		return err
	}
	return saveResponseBody(resp,name)
}

func saveResponseMeta(req * http.Request, resp * http.Response, name string) error {
	var buf bytes.Buffer
	buf.Write([]byte("Time: " + time.Now().String() + "\r\n\r\n"))
	b, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		return err
	}
	buf.Write(b)
	b, err = httputil.DumpResponse(resp, false)
	if err != nil {
		return err
	}
	buf.Write(b)
	return ioutil.WriteFile(name + ".txt", buf.Bytes(), os.ModePerm)
}

func saveResponseBody(resp * http.Response, name string) error {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return ioutil.WriteFile(name + ".html", b, os.ModePerm)
}