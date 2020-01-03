package http

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	xhttp "net/http"
	xurl "net/url"
	"strings"
	"time"

	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/json"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const OptionTypeUrlencode = "application/x-www-form-urlencoded"
const OptionTypeJson = "application/json"
const OptionTypeText = "text/plain"
const OptionTypeXml = "text/xml"
const OptionTypeMultipart = "multipart/form-data"

const OptionResponseTypeText = "text"
const OptionResponseTypeJson = "json"
const OptionResponseTypeByte = "byte"
const OptionResponseTypeResponse = "response"
const OptionResponseTypeAuto = "auto"

type Options struct {
	Url             string
	Method          string
	Type            string
	ResponseType    string
	Data            interface{}
	Headers         map[string]string
	RedirectCount   int
	Timeout         time.Duration
	ResponseCharset string
}

type Error struct {
	StatusCode int               `json:"statusCode"`
	Status     string            `json:"status"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func (E *Error) Error() string {
	return fmt.Sprintf("[HTTP] [ERROR] %d %s", E.StatusCode, E.Status)
}

var ca *x509.CertPool
var client *xhttp.Client

func init() {
	ca = x509.NewCertPool()
	ca.AppendCertsFromPEM(pemCerts)
	client = &xhttp.Client{
		Transport: &xhttp.Transport{
			TLSClientConfig:   &tls.Config{RootCAs: ca},
			DisableKeepAlives: false,
		},
	}
}

func CA() *x509.CertPool {
	return ca
}

func NewClient() *xhttp.Client {
	return &xhttp.Client{
		Transport: &xhttp.Transport{
			TLSClientConfig:   &tls.Config{RootCAs: ca},
			DisableKeepAlives: false,
			IdleConnTimeout:   6 * time.Second,
		},
	}
}

func Send(options *Options) (interface{}, error) {
	return SendWithClient(client, options)
}

func getContent(resp *xhttp.Response, charset string) ([]byte, error) {

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-Type")

	if contentType == "" {
		contentType = resp.Header.Get("content-type")
	}

	contentType = strings.ToLower(contentType)

	if strings.Contains(contentType, "charset=gbk") || charset == "gbk" {
		rd := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())
		b, err = ioutil.ReadAll(rd)
		if err != nil {
			return nil, err
		}
	} else if strings.Contains(contentType, "charset=gb2312") || charset == "gb2312" {
		rd := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GB18030.NewDecoder())
		b, err = ioutil.ReadAll(rd)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func SendWithClient(client *xhttp.Client, options *Options) (interface{}, error) {

	var url = options.Url
	var resp *xhttp.Response
	var req *xhttp.Request
	var err error

	if options.Method == "POST" {

		contentType := options.Type

		var body []byte = nil

		if strings.Contains(options.Type, "json") {

			contentType = options.Type + "; charset=utf-8"

			body, err = json.Marshal(options.Data)

			if err != nil {
				return nil, err
			}

		} else if strings.Contains(options.Type, "text") {

			contentType = options.Type + "; charset=utf-8"

			body = []byte(dynamic.StringValue(options.Data, ""))

		} else if strings.Contains(options.Type, "multipart") {
			b := bytes.NewBuffer(nil)
			w := multipart.NewWriter(b)
			var fw io.Writer = nil

			dynamic.Each(options.Data, func(key interface{}, item interface{}) bool {

				skey := dynamic.StringValue(key, "")

				log.Println("[HTTP] [multipart]", skey)

				if dynamic.Get(item, "name") == nil {
					w.WriteField(skey, dynamic.StringValue(item, ""))
				} else {

					name := dynamic.StringValue(dynamic.Get(item, "name"), "")
					content := dynamic.Get(item, "content")
					fw, err = w.CreateFormFile(skey, name)

					if err != nil {
						return false
					}

					for content != nil {

						{
							s, ok := content.(string)
							if ok {
								fw.Write([]byte(s))
								break
							}
						}

						{
							s, ok := content.([]byte)
							if ok {
								fw.Write(s)
								break
							}
						}
						break
					}

				}

				return true
			})

			w.Close()

			if err != nil {
				return nil, err
			}

			contentType = w.FormDataContentType()

			body = b.Bytes()

		} else {

			contentType = options.Type + "; charset=utf-8"

			idx := 0
			b := bytes.NewBuffer(nil)

			dynamic.Each(options.Data, func(key interface{}, value interface{}) bool {

				if idx != 0 {
					b.WriteString("&")
				}

				b.WriteString(dynamic.StringValue(key, ""))
				b.WriteString("=")
				b.WriteString(xurl.QueryEscape(dynamic.StringValue(value, "")))

				idx = idx + 1

				return true
			})

			body = b.Bytes()
		}

		req, err = xhttp.NewRequest("POST", url, bytes.NewReader(body))

		if err == nil {

			if options.Timeout > 0 {
				ctx, cancel := context.WithTimeout(context.TODO(), options.Timeout)
				defer cancel()
				req = req.WithContext(ctx)
			}

			req.Header.Set("Content-Type", contentType)
			req.Header.Set("Connection", "keepalive")

			if options.Headers != nil {
				for key, value := range options.Headers {
					req.Header.Set(key, value)
				}
			}

			resp, err = client.Do(req)

		}

	} else {

		idx := 0

		b := bytes.NewBuffer(nil)

		dynamic.Each(options.Data, func(key interface{}, value interface{}) bool {

			if idx != 0 {
				b.WriteString("&")
			}

			b.WriteString(dynamic.StringValue(key, ""))
			b.WriteString("=")
			b.WriteString(xurl.QueryEscape(dynamic.StringValue(value, "")))

			idx = idx + 1

			return true
		})

		idx = strings.Index(url, "?")

		if idx >= 0 {
			if idx+1 == len(url) {
				url = url + b.String()
			} else {
				url = url + "&" + b.String()
			}
		} else {
			url = url + "?" + b.String()
		}

		req, err = xhttp.NewRequest("GET", url, nil)

		if err == nil {

			if options.Timeout > 0 {
				ctx, cancel := context.WithTimeout(context.TODO(), options.Timeout)
				defer cancel()
				req = req.WithContext(ctx)
			}

			if options.Headers != nil {
				for key, value := range options.Headers {
					req.Header.Add(key, value)
				}
			}
			resp, err = client.Do(req)
		}

	}

	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 200 {

		b, err := getContent(resp, options.ResponseCharset)

		if err != nil && err != io.EOF {
			return nil, err
		}

		if options.ResponseType == OptionResponseTypeAuto {
			contentType := resp.Header.Get("Content-Type")
			if contentType == "" {
				contentType = resp.Header.Get("content-type")
			}
			if strings.Contains(contentType, "json") {
				options.ResponseType = OptionResponseTypeJson
			} else if strings.Contains(contentType, "text") {
				options.ResponseType = OptionResponseTypeText
			} else {
				options.ResponseType = OptionResponseTypeByte
			}
		}

		if options.ResponseType == "json" {
			var data interface{} = nil
			err := json.Unmarshal(b, &data)
			if err != nil {
				return nil, err
			}
			return data, nil
		} else if options.ResponseType == "byte" {
			return b, nil
		} else {
			return string(b), nil
		}

	} else {

		b, err := getContent(resp, options.ResponseCharset)

		if err != nil {
			return nil, err
		}

		if options.ResponseType == "response" {
			return &Error{StatusCode: resp.StatusCode, Status: resp.Status, Headers: map[string]string{}, Body: string(b)}, nil
		}

		if resp.StatusCode == 302 && options.RedirectCount > 0 {
			options.Url = resp.Header.Get("Location")
			options.RedirectCount = options.RedirectCount - 1
			fmt.Println("[KK] Redirect", options.Url)
			return Send(options)
		}

		log.Println("[HTTP] [ERROR]", resp.StatusCode, string(b))

		e := Error{StatusCode: resp.StatusCode, Status: resp.Status, Headers: map[string]string{}, Body: string(b)}

		for key, h := range resp.Header {
			e.Headers[key] = h[0]
		}

		return nil, &e
	}
}