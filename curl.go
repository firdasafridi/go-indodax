package indodax

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//
// callPublic call public API with specific method and parameters.
// On success it will return response body.
// On fail it will return an empty body with an error.
//
func (cl *Client) curlPublic(urlPath string) (body []byte, err error) {
	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"Content-Type": []string{
				"application/x-www-form-urlencoded",
			},
		},
	}

	req.URL, err = url.Parse(cl.env.BaseHostPublic + urlPath)
	if err != nil {
		return nil, fmt.Errorf("curlPublic: " + err.Error())
	}

	res, err := cl.conn.Do(req)
	if err != nil {
		return nil, fmt.Errorf("curlPublic: " + err.Error())
	}

	body, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("curlPublic: " + err.Error())
	}

	printDebug(string(body))

	return body, nil
}

//
// callPrivate call private API with specific method and parameters.
// On success it will return response body.
// On fail it will return an empty body with an error.
//
func (cl *Client) curlPrivate(method string, params url.Values) (
	body []byte, err error,
) {
	req, err := cl.newPrivateRequest(method, params)
	if err != nil {
		return nil, fmt.Errorf("curlPrivate: " + err.Error())
	}

	res, err := cl.conn.Do(req)
	if err != nil {
		return nil, fmt.Errorf("curlPrivate: " + err.Error())
	}

	body, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("curlPrivate: " + err.Error())
	}

	printDebug(string(body))

	return body, nil
}

//
// newPrivateRequest is method to generate authentication for private API.
// On success it will return http request.
// On fail it will return an error.
//
func (cl *Client) newPrivateRequest(apiMethod string, params url.Values) (
	req *http.Request, err error,
) {
	query := url.Values{
		"timestamp": []string{
			timestampAsString(),
		},
		"method": []string{
			apiMethod,
		},
	}

	virtualParams := map[string][]string(params)
	for k, v := range virtualParams {
		if len(v) > 0 {
			query.Set(k, v[0])
		}
	}

	reqBody := query.Encode()

	printDebug(fmt.Sprintf("newPrivateRequest >> request body:%s", reqBody))

	sign := cl.encodeToHmac512(reqBody)

	req = &http.Request{
		Method: http.MethodPost,
		Header: http.Header{
			"Content-Type": []string{
				"application/x-www-form-urlencoded",
			},
			"Key": []string{
				cl.env.apiKey,
			},
			"Sign": []string{
				sign,
			},
		},
		Body: ioutil.NopCloser(strings.NewReader(reqBody)),
	}

	req.URL, err = url.Parse(cl.env.BaseHostPrivate)
	if err != nil {
		err = fmt.Errorf("newPrivateRequest: " + err.Error())
		return nil, err
	}
	return req, nil
}

func (cl *Client) encodeToHmac512(param string) string {
	sign := hmac.New(sha512.New, []byte(cl.env.apiSecret))

	sign.Write([]byte(param))

	return hex.EncodeToString(sign.Sum(nil))
}
