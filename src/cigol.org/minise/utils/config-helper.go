package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Config map[string]string

var config Config = Config{}

func GetConfig(name string) Config {
	if len(config) > 0 {
		return config
	}
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	b := make([]byte, 2048)
	var s string
	for {
		n, err := file.Read(b)
		if n > 0 {
			s = s + string(b[:n])
		}
		if strings.Contains(s, "\n") {
			str := strings.Split(s, "\n")
			for i, c := range str {
				// skip the last one, coz we do not know when it meets the end
				if i == len(str)-1 {
					break
				}
				c = strings.TrimSpace(c)
				if !strings.HasPrefix(c, "#") {
					if strings.Contains(c, "=") {
						i := strings.Index(c, "=")
						config[c[:i]] = c[i+1:]
					}
				}
			}
			s = str[len(str)-1]
		}

		if err == io.EOF {
			s = strings.TrimSpace(s)
			if !strings.HasPrefix(s, "#") {
				if strings.Contains(s, "=") {
					i := strings.Index(s, "=")
					config[s[:i]] = s[i+1:]
				}
			}
			break
		}

		if err != nil {
			break
		}
	}
	return config
}
