package client

import (
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// NewEcsMgmtClient ...
func NewEcsMgmtClient(name, username, password string, nodes []string, reqTimeout, blockDur time.Duration) *EcsMgmtClient {
	return &EcsMgmtClient{
		Name:     name,
		username: username,
		password: password,
		ecs:      NewEcs(nodes),
		blockDur: blockDur,
		mutex:    &sync.Mutex{},
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: reqTimeout,
		},
	}
}

// EcsMgmtClient defines client for ECS mgmt
type EcsMgmtClient struct {
	Name     string
	username string
	password string
	ecs      *Ecs
	blockDur time.Duration
	token    *Token
	mutex    *sync.Mutex
	client   *http.Client
}

// PerformRequest sends request to ECS
func (e *EcsMgmtClient) PerformRequest(method, uri string, body io.Reader, bodyLength int64, headers http.Header, auth Authentication) (*http.Response, error) {
	// TODO, add retry, especially for token expiry
	host, hostIdx, err := e.ecs.NextAvailableNode()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, host+uri, body)
	if err != nil {
		return nil, err
	}

	// Set content length
	req.ContentLength = bodyLength

	// Set header
	for k, vs := range headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}
	log.Printf("[%s] performing request %s %s", e.Name, req.Method, req.URL)

	auth.SetAuth(req)

	resp, err := e.client.Do(req)
	if err != nil {
		log.Printf("[%s] error while performing request %s %s: %s", e.Name, req.Method, req.URL, err)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			e.ecs.BlockNode(hostIdx, e.blockDur)
		}
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 219 {
		return resp, nil
	}

	// Close the body on failure response
	resp.Body.Close()
	return nil, fmt.Errorf("[%s]: %s %s", e.Name, method, resp.Status)
}

// DELETE ...
func (e *EcsMgmtClient) DELETE(uri string, body io.Reader, bodyLength int64, headers http.Header, auth Authentication) (*http.Response, error) {
	return e.Query("DELETE", uri, body, bodyLength, headers)
}

// GET ...
func (e *EcsMgmtClient) GET(uri string, body io.Reader, bodyLength int64, headers http.Header, auth Authentication) (*http.Response, error) {
	return e.Query("GET", uri, body, bodyLength, headers)
}

// HEAD ...
func (e *EcsMgmtClient) HEAD(uri string, body io.Reader, bodyLength int64, headers http.Header, auth Authentication) (*http.Response, error) {
	return e.Query("HEAD", uri, body, bodyLength, headers)
}

// POST ...
func (e *EcsMgmtClient) POST(uri string, body io.Reader, bodyLength int64, headers http.Header, auth Authentication) (*http.Response, error) {
	return e.Query("POST", uri, body, bodyLength, headers)
}

// PUT ...
func (e *EcsMgmtClient) PUT(uri string, body io.Reader, bodyLength int64, headers http.Header, auth Authentication) (*http.Response, error) {
	return e.Query("PUT", uri, body, bodyLength, headers)
}

// MgmtLogin to login ECS mgmt interface
func (e *EcsMgmtClient) MgmtLogin() (err error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.token != nil && !e.token.Expired() {
		return nil
	}
	var resp *http.Response
	if resp, err = e.PerformRequest("GET", "/login", nil, 0, http.Header{}, &BasicAuth{e.username, e.password}); err != nil {
		return err
	}
	defer resp.Body.Close()
	token := resp.Header.Get("X-Sds-Auth-Token")
	if token == "" {
		return fmt.Errorf("No Token Granted")
	}
	// first time login
	if e.token == nil {
		e.token = NewToken(token)
	} else {
		prevToken := e.token.Refresh(token)
		// release previous token to avoid token leak
		go e.MgmtLogout(prevToken)
	}
	return nil
}

// Query does the general query to ECS
func (e *EcsMgmtClient) Query(method, uri string, body io.Reader, bodyLength int64, headers http.Header) (*http.Response, error) {
	if e.token == nil || e.token.Expired() {
		if err := e.MgmtLogin(); err != nil {
			return nil, err
		}
	}
	token, _ := e.token.Get()
	return e.PerformRequest(method, uri, body, bodyLength, headers, &TokenAuth{token})
}

// MgmtLogout logs out ECS mgmt interface
func (e *EcsMgmtClient) MgmtLogout(token string) error {
	resp, err := e.PerformRequest("GET", "/logout", nil, 0, http.Header{}, &TokenAuth{token})
	if err != nil {
		log.Printf("logout failed [%s]", err)
		return err
	}
	resp.Body.Close()
	return nil
}

const (
	// DecodeJSON ...
	DecodeJSON = iota
	// DecodeXML ...
	DecodeXML
)

// ErrUnsupportedDecodeType defines a new error type
var ErrUnsupportedDecodeType = fmt.Errorf("Unsupported Decode Type")

// GetMgmtQueryResultJSON gets query result in json format
func (e *EcsMgmtClient) GetMgmtQueryResultJSON(uri string, v interface{}) error {
	return e.getMgmtQueryResult("GET", uri, nil, 0, http.Header{}, v, DecodeJSON)
}

// GetMgmtQueryResultXML gets query result in xml format
func (e *EcsMgmtClient) GetMgmtQueryResultXML(uri string, v interface{}) error {
	return e.getMgmtQueryResult("GET", uri, nil, 0, http.Header{}, v, DecodeXML)
}

// GetMgmtQueryResultJSONByPost gets query result in json format
func (e *EcsMgmtClient) GetMgmtQueryResultJSONByPost(uri string, body io.Reader, bodyLength int64, headers http.Header, v interface{}) error {
	return e.getMgmtQueryResult("POST", uri, body, bodyLength, headers, v, DecodeJSON)
}

// GetMgmtQueryResultXMLByPost gets query result in xml format
func (e *EcsMgmtClient) GetMgmtQueryResultXMLByPost(uri string, body io.Reader, bodyLength int64, headers http.Header, v interface{}) error {
	return e.getMgmtQueryResult("POST", uri, body, bodyLength, headers, v, DecodeXML)
}

// GetMgmtQueryResult gets query result
func (e *EcsMgmtClient) getMgmtQueryResult(method, uri string, body io.Reader, bodyLength int64, headers http.Header, v interface{}, t int) error {
	resp, err := e.Query(method, uri, body, bodyLength, headers)
	if err != nil {
		log.Printf("Error while requesting %s: %v", uri, err)
		return err
	}

	defer resp.Body.Close()
	switch t {
	case DecodeJSON:
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			log.Printf("Error while decoding json from response: %v", err)
			return err
		}
	case DecodeXML:
		if err = xml.NewDecoder(resp.Body).Decode(v); err != nil {
			log.Printf("Error while decoding xml from response: %v", err)
			return err
		}
	default:
		return ErrUnsupportedDecodeType
	}

	return nil
}

// Close removes token and set pointer to nil
func (e *EcsMgmtClient) Close() {
	e.mutex.Lock()
	if e.token != nil {
		if prev, expired := e.token.Get(); !expired {
			e.MgmtLogout(prev)
		}
	}
	e.token = nil
	e.mutex.Unlock()
	e = nil
}
