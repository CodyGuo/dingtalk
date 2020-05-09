package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/CodyGuo/dingtalk/utils"
	"github.com/CodyGuo/glog"
)

var (
	log = glog.New(os.Stderr)
)

func init() {
	log.SetFlags(glog.LglogFlags)
}

type Requester interface {
	GetMethod() string
	GetHeader() map[string]string
	GetBody() ([]byte, error)
	GetSuccessCode() int64
}

type DingTalk struct {
	mu       sync.Mutex
	url      string
	secret   string
	timeout  time.Duration
	client   *utils.HttpClient
	response *http.Response
	err      *Err
}

func New(url string, options ...Option) *DingTalk {
	dt := &DingTalk{
		url:     url,
		timeout: 3 * time.Second,
		err:     newErr("new", nil),
	}
	for _, option := range options {
		option(dt)
	}
	dt.initClient()
	return dt
}

func (dt *DingTalk) GetSecret() string {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	return dt.secret
}

func (dt *DingTalk) SetSecret(secret string) {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	dt.secret = secret
}

func (dt *DingTalk) GetTimeout() time.Duration {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	return dt.client.GetTimeout()
}

func (dt *DingTalk) SetTimeout(timeout time.Duration) {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	dt.client.SetTimeout(timeout)
}

func (dt *DingTalk) Request(req Requester) error {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	if err := dt.check(); err != nil {
		dt.err = newErr("request check failed", err)
		return err
	}
	method := req.GetMethod()
	header := req.GetHeader()
	body, err := req.GetBody()
	if err != nil {
		dt.err = newErr("get body failed", err)
		return err
	}
	log.Debugf("url: %s, timeout: %s, method: %s, header: %v, body: %s",
		dt.url, dt.timeout, method, header, body)

	resp, err := dt.client.Request(method, header, body)
	if err != nil {
		dt.err = newErr("http request failed", err)
		return dt.err
	}
	defer resp.Body.Close()
	dt.response = resp

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		dt.err = newErr("read resp body failed", err)
		return dt.err
	}
	dt.response.Body = ioutil.NopCloser(bytes.NewReader(data))
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("invalid http status %d, detail body: %s", resp.StatusCode, data)
		dt.err = newErr("http request failed, "+string(data), err)
		return dt.err
	}
	jsonDecoder := json.NewDecoder(bytes.NewReader(data))
	if err := jsonDecoder.Decode(dt.err); err != nil {
		dt.err = newErr("json decode resp failed, "+string(data), err)
		return dt.err
	}
	dt.err.ApplicationHost = resp.Header.Get("Application-Host")
	dt.err.ServiceHost = resp.Header.Get("Location-Host")
	if dt.err.Code != req.GetSuccessCode() {
		dt.err = newErr("response status code failed", dt.err)
		return dt.err
	}
	return nil
}

func (dt *DingTalk) GetResponse() (*http.Response, error) {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	if dt.err.Detail != nil {
		return nil, dt.err
	}
	return dt.response, nil
}

func (dt *DingTalk) initClient() {
	// 拼接请求参数
	step := "?"
	if strings.Contains(dt.url, "?") {
		step = "&"
	}
	params := dt.genQueryParams()
	dt.url = strings.Join([]string{dt.url, params}, step)
	dt.client = utils.NewHttpClient(dt.url, dt.timeout)
}

func (dt *DingTalk) genQueryParams() string {
	params := url.Values{}
	if dt.secret != "" {
		timestamp := time.Now().UnixNano() / 1e6
		sign := utils.ComputeSignature(timestamp, dt.secret)
		params.Add("timestamp", strconv.FormatInt(timestamp, 10))
		params.Add("sign", sign)
	}
	return params.Encode()
}

func (dt *DingTalk) check() error {
	_, err := url.Parse(dt.url)
	if err != nil {
		return err
	}
	return nil
}

func GetLogLevel() glog.Level {
	return log.Level()
}

func SetLogLevel(level glog.Level) {
	log.SetLevel(level)
}
