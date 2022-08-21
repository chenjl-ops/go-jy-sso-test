package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"net/http"
	"strings"
)

// Request 统一请求json格式数据
/*
usage:

type Data struct {
	Name string `json:"name"`
    Age  int    `json:"age"`
}

type D struct {
	Code     string `json:"code"`
    Message  string `json:"message"`
    Data     []Data `json:"data"`
}

var data D

url := fmt.Sprintf("http://xxxx.xxx.xxx")
err := Request(url, &data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(data)

*/
func Request(url string, data interface{}) error {
	resp, err := http.Get(url)

	//logrus.Println("请求地址: ", url)
	if err != nil {
		logrus.Println("请求失败: ", err)
		return err
	}
	defer resp.Body.Close()

	err1 := json.NewDecoder(resp.Body).Decode(data)
	if err1 != nil {
		logrus.Println("解析失败: ", err1)
		return err1
	}
	return nil
}

// RequestMethod 统一request请求
/*
Usage:


data := map[string]interface{}{}
result := new(interface{})
headers := map[string]string{"Content-Type": "application/json"}
url := fmt.Sprintf("http://xxxx.xxx.xxx")

GET:
err := RequestMethod(url, "GET", headers, data, &result)
if err != nil {
	fmt.Println(err)
}
fmt.Println(result)

POST|PUT|DELETE:
err := RequestMethod(url, "POST|PUT|DELETE", headers, data, &result)
if err != nil {
	fmt.Println(err)
}
fmt.Println(result)
*/
func RequestMethod(url string, method string, headers map[string]string, requestData interface{}, responseData interface{}) error {
	methods := []string{"GET", "PUT", "POST", "DELETE"}
	if funk.Contains(methods, strings.ToUpper(method)) {
		bytesData, err := json.Marshal(requestData)
		if err != nil {
			logrus.Println("json error", err)
			return err
		} else {
			request, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewReader(bytesData))
			if err != nil {
				logrus.Println("json error", err)
				return err
			}
			for k, v := range headers {
				request.Header.Set(k, v)
			}
			// request.Header.Set("authorization", b.GetAuthorization())
			client := http.Client{}
			resp, err := client.Do(request)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			err1 := json.NewDecoder(resp.Body).Decode(responseData)
			if err1 != nil {
				return err1
			}
			defer resp.Body.Close()
			return nil
		}
	} else {
		return errors.New(fmt.Sprintf("%s not allow", method))
	}
}
