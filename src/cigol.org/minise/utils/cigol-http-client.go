package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	// hosts are like 'http://10.126.3.160:9200'
	hosts []string
}

func NewClient(hosts []string) Client {
	return Client{hosts: hosts}
}

func (c *Client) Get(path, data string) (string, error) {
	if len(c.hosts) == 0 {
		return "", nil
	}
	var e error
	for _, host := range c.hosts {
		var resp *http.Response
		var err error
		if data == "" {
			resp, err = http.Get(host + "/" + path)
		} else {
			reader := strings.NewReader(data)
			req, err := http.NewRequest("GET", host+"/"+path, reader)
			if err != nil {
				e = err
				continue
			}
			resp, err = http.DefaultClient.Do(req)
		}
		if err != nil {
			e = err
			continue
		}
		defer resp.Body.Close()

		content := ""
		b := make([]byte, 8*1024)
		for {
			n, err := resp.Body.Read(b)
			if n > 0 {
				content = content + string(b[:n])
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				return content, err
			}
		}

		return content, nil
	}
	return "", e
}

// put method, puts a document to the given url
func (c *Client) Put(index, t, id, data string) error {
	if len(c.hosts) == 0 {
		return nil
	}
	for _, host := range c.hosts {
		reader := strings.NewReader(data)
		req, err := http.NewRequest("PUT", host+"/"+index+"/"+t+"/"+id, reader)
		if err != nil {
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		content := ""
		b := make([]byte, 8*1024)
		for {
			n, err := resp.Body.Read(b)
			if n > 0 {
				content = content + string(b[:n])
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
		}
		fmt.Println(content)

		return nil
	}
	return nil
}
