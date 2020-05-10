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
	err      *Error
}

func New(url string, options ...Option) *DingTalk {
	dt := &DingTalk{
		url:     url,
		timeout: 5 * time.Second,
		err:     &Error{},
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

func (dt *DingTalk) Request(req Requester) error {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	if err := dt.checkURL(); err != nil {
		return dt.newError("checkURL", "", err)
	}
	body, err := dt.request(req)
	if err != nil {
		return dt.newError("request", body, err)
	}
	if err := dt.checkResponse(req); err != nil {
		return dt.newError("checkResponse", body, err)
	}
	return nil
}

func (dt *DingTalk) checkURL() error {
	_, err := url.Parse(dt.url)
	if err != nil {
		return err
	}
	return nil
}

func (dt *DingTalk) request(req Requester) (string, error) {
	method := req.GetMethod()
	header := req.GetHeader()
	body, err := req.GetBody()
	if err != nil {
		return "", err
	}
	log.Debugf("url: %s, timeout: %s, method: %s, header: %v, body: %s",
		dt.url, dt.timeout, method, header, body)

	dt.response, err = dt.client.Request(method, header, body)
	if err != nil {
		return string(body), err
	}
	return string(body), nil
}

func (dt *DingTalk) checkResponse(req Requester) error {
	defer dt.response.Body.Close()
	data, err := ioutil.ReadAll(dt.response.Body)
	if err != nil {
		return err
	}
	dt.response.Body = ioutil.NopCloser(bytes.NewReader(data))

	if dt.response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid http status %d, body: %s", dt.response.StatusCode, data)
	}

	respMsg := ResponseMsg{}
	if err := json.Unmarshal(data, &respMsg); err != nil {
		return fmt.Errorf("body: %s, %w", data, err)
	}
	respMsg.ApplicationHost = dt.response.Header.Get("Application-Host")
	respMsg.ServiceHost = dt.response.Header.Get("Location-Host")
	if respMsg.ErrCode != req.GetSuccessCode() {
		return fmt.Errorf("%s", respMsg)
	}
	return nil
}

func (dt *DingTalk) newError(op, body string, err error) error {
	dt.err = newError(op, dt.url, body, err)
	return dt.err
}

func (dt *DingTalk) GetResponse() (*http.Response, error) {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	if dt.err.Err != nil {
		return nil, dt.newError("GetResponse", "", dt.err)
	}
	return dt.response, nil
}

func GetLogLevel() glog.Level {
	return log.Level()
}

func SetLogLevel(level glog.Level) {
	log.SetLevel(level)
}
