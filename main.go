package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//func main() {
//	result := weather.CurrentWeather("101010100")
//	fmt.Println(result)
//}

type hr struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

func main() {
	fmt.Println(123, os.Args)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("收到请求：", r.URL.Path)
		fmt.Fprintf(w, "欢迎使用人力资源管理系统")
	})

	var db []hr
	//响应/insert，从传入的参数新增人员信息
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				fmt.Fprintln(w, "错误的请求")
			} else {
				name := r.FormValue("name")
				// 将string转为int
				age, _ := strconv.Atoi(r.FormValue("age"))
				gender, _ := strconv.Atoi(r.FormValue("gender"))
				db = append(db, hr{Name: name, Age: age, Gender: gender})
				fmt.Fprintln(w, "添加了"+name)
			}
		}
	})

	//给定format，可按json格式输出，默认格式为字符串
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "错误的请求")
		} else {
			format := r.FormValue("format")
			if format == "json" {
				data, err := json.Marshal(db)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Fprintln(w, string(data))
				}
			} else {
				for i := 0; i < len(db); i++ {
					fmt.Fprintln(w, db[i].Name, db[i].Age, db[i].Gender)
				}
			}
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("启动服务失败，错误信息：", err)
	} else {
		fmt.Println("启动服务成功！")
	}
}
