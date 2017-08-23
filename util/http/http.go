// httpHelper
package _http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//定义设置了超时时间的httpclient
var currClient *http.Client = &http.Client{
	Transport: &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*300)
			if err != nil {
				fmt.Println("dail timeout", err)
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 200,
	},
}

func HttpGet(url string) error {
	resp, err := currClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func HttpPost(url string, postbody string, bodyType string) error {
	postbytes := bytes.NewBuffer([]byte(postbody))

	if bodyType == "" {
		bodyType = "application/x-www-form-urlencoded"
	}
	resp, err := currClient.Post(url, bodyType, postbytes)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil

}
