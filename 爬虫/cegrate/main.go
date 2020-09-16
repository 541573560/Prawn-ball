package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.joymy.net/")

	if err != nil {
	}
	defer resp.Body.Close()                // 函数结束时关闭Body
	body, err := ioutil.ReadAll(resp.Body) // 读取Body
	fmt.Print(string(body))
}
