package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CurrentWeather(cityCode string) string {
	//使用net包发起Get请求
	resp, err := http.Get("https://devapi.qweather.com/v7/weather/now?location=" + cityCode + "&key=[您自己申请的AppKey]")
	if err != nil {
		fmt.Println("HTTP请求失败：", err)
		panic(err)
	}
	//使用断言关闭网络请求
	defer resp.Body.Close()
	//使用ioutil工具包获取服务端响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网络响应失败：", err)
		panic(err)
	}
	return string(body)
}
