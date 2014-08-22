package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"code.google.com/p/go.text/encoding/charmap"
	"code.google.com/p/go.text/transform"
	"github.com/temoto/heroshi/heroshi"
	"github.com/temoto/phonetic-semantics/phonosem"
)

const (
	DefaultConnectTimeout    = 10 * time.Second
	DefaultFetchTimeout      = 20 * time.Second
	DefaultIOTimeout         = 10 * time.Second
	DefaultKeepaliveTimeout  = 60 * time.Second
	DefaultNameServerAddress = "127.0.1.1:53"
	DefaultReadLimit         = 1 << 20
	DefaultUserAgent         = "PhonoSemantics/1"
)

var (
	transport *heroshi.Transport
)

func Request(url *url.URL, method string, headers map[string]string, body io.Reader) (result *heroshi.FetchResult) {
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return heroshi.ErrorResult(url, err.Error())
	}
	req.Header.Set("User-Agent", DefaultUserAgent)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	options := &heroshi.RequestOptions{
		ConnectTimeout:   DefaultConnectTimeout,
		ReadTimeout:      DefaultIOTimeout,
		WriteTimeout:     DefaultIOTimeout,
		ReadLimit:        DefaultReadLimit,
		KeepaliveTimeout: DefaultKeepaliveTimeout,
		Stat:             new(heroshi.RequestStat),
	}
	result = heroshi.Fetch(transport, req, options, DefaultFetchTimeout)
	result.Stat = options.Stat
	transport.CloseIdleConnections(false)

	return result
}

func Dial(netw, addr string, options *heroshi.RequestOptions) (net.Conn, error) {
	addrString, addrPort := heroshi.SplitPort(addr)
	var ip net.IP
	addrs, resolveTime, err := heroshi.ResolveName(addrString, DefaultNameServerAddress)
	if err != nil {
		log.Println("Dial: resolve error:", err.Error(), "in", resolveTime.String())
		return nil, err
	}
	if len(addrs) > 0 {
		ip = addrs[0]
		// log.Printf("Dial: resolved: %s -> %s in %s", addr, ip.String(), resolveTime.String())
		addr = ip.String()
	}
	if options != nil && options.Stat != nil {
		var tcpAddr net.TCPAddr
		tcpAddr.IP = ip
		options.Stat.RemoteAddr = &tcpAddr
		options.Stat.ResolveTime = resolveTime
	}

	// Return port back
	addr += ":" + addrPort

	var conn net.Conn
	if options != nil && options.ConnectTimeout != 0 {
		conn, err = net.DialTimeout(netw, addr, options.ConnectTimeout)
	} else {
		conn, err = net.Dial(netw, addr)
	}
	if err != nil {
		return conn, err
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return conn, errors.New("Dial: conn->TCPConn type assertion failed.")
	}
	tcpConn.SetKeepAlive(true)
	tcpConn.SetLinger(0)
	tcpConn.SetNoDelay(true)
	return tcpConn, err
}

func encode1251(s string) string {
	to1251 := charmap.Windows1251.NewEncoder()
	result, _, err := transform.String(to1251, s)
	check(err)
	return result
}

func decode1251(b []byte) []byte {
	from1251 := charmap.Windows1251.NewDecoder()
	result, _, err := transform.Bytes(from1251, b)
	check(err)
	log.Printf("decode1251 %v %s", b, string(result))
	return result
}

func CrawlWord(s string) (*phonosem.Item, error) {
	log.Printf("CrawlWord begin %s", s)
	urlString := "http://psi-technology.net/servisfonosemantika.php"
	// urlString := "http://localhost:9092"
	uri, err := url.Parse(urlString)
	check(err)

	headers := map[string]string{
	// "Content-Type": "application/x-www-form-urlencoded",
	}
	bodyString := fmt.Sprintf("slovo=%s&sub=", url.QueryEscape(encode1251(s)))
	body := strings.NewReader(bodyString)
	fetchResult := Request(uri, "POST", headers, body)
	log.Printf("CrawlWord %s %v", s, fetchResult)

	if !fetchResult.Success || fetchResult.StatusCode != 200 {
		log.Printf("CrawlWord request error word=%s", s)
		return nil, errors.New("CrawlWord HTTP error")
	}
	features, err := ParseFeatures(decode1251(fetchResult.Body))
	if err != nil {
		return nil, err
	}

	result := &phonosem.Item{
		Key:      s,
		Features: features,
	}
	return result, nil
}

func init() {
	transport = &heroshi.Transport{
		Dial:                Dial,
		MaxIdleConnsPerHost: 2,
	}
}
