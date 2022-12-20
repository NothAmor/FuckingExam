package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.design/x/clipboard"
)

type AnswerResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

func main() {
	var (
		question          string
		answerResponse    AnswerResponse
		clipboardQuestion string
	)

	err := clipboard.Init()
	if err != nil {
		log.Println("clipboard init error")
	}

	for {
		clipboardQuestion = string(clipboard.Read(clipboard.FmtText))

		if clipboardQuestion == question {
			continue
		}

		question = clipboardQuestion

		log.Println("题目: ", clipboardQuestion)

		answerUrl := "http://cx.icodef.com/wyn-nb?v=4"

		log.Println("开始查询答案...")

		resp, err := http.PostForm(answerUrl, url.Values{
			"question": {clipboardQuestion},
		})
		if err != nil {
			log.Println("查询出错")
			log.Println(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		json.Unmarshal(body, &answerResponse)

		fmt.Println("答案内容: ", answerResponse.Data)

		log.Println("发送到QQ...")

		qqSend := "http://124.221.152.192:8086/send_private_msg"

		resp, err = http.PostForm(qqSend, url.Values{
			"user_id": {"1565481748"},
			"message": {fmt.Sprintf("考试题目:%s\n题库的答案是: %s", clipboardQuestion, answerResponse.Data)},
		})
		if err != nil {
			log.Println(err)
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(body))

		time.Sleep(time.Second * 1)
	}
}
