package main

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func ybApp()  []byte{
	r, _ := http.Get("http://192.168.31.26:18006/actuator/prometheus")
	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)
	return body
}


func main() {
	e := echo.New()
	e.GET("/metrics", func(context echo.Context) error {
		//fmt.Println(string(ybApp()))
		return context.String(
			http.StatusOK,
			string(ybApp()),
		)
	})
	e.Start(":8080")
}