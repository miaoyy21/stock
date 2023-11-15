package main

import (
	"encoding/json"
	"fmt"
	"github.com/liushuochen/gotable"
	"log"
	"net/http"
)

type ResponseData struct {
	Date          string `json:"d"`   // 交易时间
	Open          string `json:"o"`   // 开盘价（元）
	High          string `json:"h"`   // 最高价（元）
	Low           string `json:"l"`   // 最低价（元）
	Close         string `json:"c"`   // 收盘价（元）
	Volume        string `json:"v"`   // 成交量（手）
	Turnover      string `json:"e"`   // 成交额（元）
	AmplitudeRate string `json:"zf"`  // 振幅（%）
	TurnoverRate  string `json:"hs"`  // 换手率（%）
	ChangeRate    string `json:"zd"`  // 涨跌幅（%）
	RangeRate     string `json:"zde"` // 涨跌额（元）
}

func main() {
	url := "http://api.biyingapi.com/hszbl/fsjy/600130/dn/1376636348bc8c31a3"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("http.NewRequest ERR : %s \n", err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("http.DefaultClient ERR : %s \n", err.Error())
	}

	var ds []ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&ds); err != nil {
		log.Fatalf("json.Decode ERR : %s \n", err.Error())
	}

	table, err := gotable.Create("", "交易时间", "开盘价（元）", "收盘价（元）", "最高价（元）", "最低价（元）", "成交量（手）", "成交额（元）", "振幅（%）", "换手率（%）", "涨跌幅（%）", "涨跌额（元）")
	if err != nil {
		log.Fatalf("gotable.Create ERR : %s \n", err.Error())
	}

	table.Align("成交量（手）", gotable.Right)
	table.Align("成交额（元）", gotable.Right)

	for i, d := range ds[len(ds)-200:] {
		v := make([]string, 0, 12)

		v = append(v, fmt.Sprintf("%d", i+1))
		v = append(v, d.Date)
		v = append(v, d.Open)
		v = append(v, d.Close)
		v = append(v, d.High)
		v = append(v, d.Low)
		v = append(v, d.Volume)
		v = append(v, d.Turnover)
		v = append(v, d.AmplitudeRate)
		v = append(v, d.TurnoverRate)
		v = append(v, d.ChangeRate)
		v = append(v, d.RangeRate)

		if err := table.AddRow(v); err != nil {
			log.Fatalf("gotable.Create ERR : %s \n", err.Error())
		}
	}

	fmt.Println(table)
}
