package indodax

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
