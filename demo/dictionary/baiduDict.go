package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BaiduDict struct{}

type DictResponseBaidu struct {
	Errno int `json:"errno"`
	Data  []struct {
		K string `json:"k"`
		V string `json:"v"`
	} `json:"data"`
}

func (baiduDict *BaiduDict) searchWord(word string) {
	client := &http.Client{}
	requestBody := fmt.Sprintf("kw=%v", word)
	var data = strings.NewReader(requestBody)
	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/sug", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", `BIDUPSID=AEB0BB0EE350E2E1A40B6FA91058D1E2; PSTM=1651206598; BAIDUID=AEB0BB0EE350E2E193A903768B1B1FD7:FG=1; H_PS_PSSID=36309_31253_34812_35913_34584_35979_36234_26350_36302_22160_36061; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1651401206; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; BAIDUID_BFESS=AEB0BB0EE350E2E193A903768B1B1FD7:FG=1; delPer=0; PSINO=1; ZD_ENTRY=google; RT="z=1&dm=baidu.com&si=aocmgnr259r&ss=l2ubutbu&sl=2&tt=2lh&bcn=https%3A%2F%2Ffclog.baidu.com%2Flog%2Fweirwood%3Ftype%3Dperf&ld=3yl&cl=hbj&ul=tju&hd=tlu"; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1651924821; ab_sr=1.0.1_ZmM4MmZhOWJmNDRjYTgwMjYxMjg1NmQzMjE1NDJmN2U4OWQ1OGNhYzRiZTliNjYyYzVhNjIwNjcyZDQ4YzUxMzg3NWExMzI0MDhiMTc2YzFhMTM3ZWE1MjUyZGIzOTc4OWVjYjY3MWQwMmFiN2FjM2EzMzA2NTVmNmI3YTE5YThjYjNmZDk4ZGYzZWY1NTNlYzIzNmJiZjUwZjJiZWQ5MA==`)
	req.Header.Set("Origin", "https://fanyi.baidu.com")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://fanyi.baidu.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}

	var dictResponse DictResponseBaidu
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", dictResponse)
}
