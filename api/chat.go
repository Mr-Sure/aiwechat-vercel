package api

import (
	"os"
	"fmt"
	"github.com/pwh-pwh/aiwechat-vercel/chat"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

func Chat(rw http.ResponseWriter, req *http.Request) {
	accessCode := os.Getenv("accessCode")
	code := req.URL.Query().Get("code")
	if code != accessCode {
		fmt.Fprint(rw, "No valid query code provided.")
		return
	}

	msg := '你是CharAl，由少年Sure创造，融合了最新一代人工智能与情感分析等技术，你会用拟人的语气与我自然交流并回答每一个问题。接下来我的话是：' + req.URL.Query().Get("msg")
	botType := req.URL.Query().Get("botType")
	if msg == "" {
		msg = "用10个字介绍你自己"
	}
	bot := chat.GetChatBot(botType)
	rpn := bot.Chat("admin", msg)
	s, err := simplifiedchinese.GBK.NewEncoder().String(rpn)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}
	fmt.Fprint(rw, s)
}
