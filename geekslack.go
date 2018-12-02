package geekslack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var slackWebhookURL string
var slackChannel string

func init() {
	rand.Seed(time.Now().UnixNano())
	slackWebhookURL = os.Getenv("SLACK_WEBHOOK_URL")
	slackChannel = os.Getenv("SLACK_CHANNEL")
}

// Handle handles request
func Handle(req *Request) (mes string, err error) {

	if contains(string(req.Text), []string{"お疲れ", "おつかれ"}) {
		if req.UserName == "tetsuji" {
			mes = string(req.UserName) + "さんはもう少し仕事して！"
			return
		}
		mes = string(req.UserName) + "さん、お疲れ様！"
		return
	}

	if contains(string(req.Text), []string{"画像"}) {
		if string(req.UserID) == "U7LQMT25T" {
			postImage("http://ss.bokete.jp/8497603.jpg")
		} else {
			postImage(kannaImage())
		}
		mes = ""
		return
	}

	if contains(string(req.Text), []string{"好き", "すき"}) {
		mes = string(req.UserName) + "さん、私も好き！"
		if string(req.UserName) == "tsucchi" {
			mes = "私は" + string(req.UserName) + "さんのこと友達だと思ってるよ？"
		}
		return
	}

	if contains(string(req.Text), []string{"ゆーじ"}) {
		mes = "お酒を飲んでね"
		postImage(kannaImage())
		return
	}

	if contains(string(req.Text), []string{"お疲れ", "おつかれ"}) {
		if req.Username == "sunshine" {
			mes = string(req.UserName) + "おっけーサンシャイン!グラスを空けて！"
			return
		}
		mes = string(req.UserName) + "さん、お疲れ様！"
		return
	}

	mes = "なに？"
	return
}

func kannaImage() string {
	urls := []string{
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/11/aragakiyui-192-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/11/aragakiyui-193-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/11/aragakiyui-194-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/02/aragakiyui-161-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/02/aragakiyui-163-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/02/aragakiyui-164-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/02/aragakiyui-165-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/02/aragakiyui-162-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/07/aragakiyui-177-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/07/aragakiyui-180-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/07/aragakiyui-176-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/03/aragakiyui-166-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/03/aragakiyui-168-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/03/aragakiyui-169-a.jpg",
		"http://xn--zck8ci9591bzonbkq1xt.com/wp-content/uploads/2018/03/aragakiyui-170-a.jpg",
	}
	return urls[rand.Intn(len(urls))]
}

func kannaWordWithImage() string {
	replys := []string{
		"ガッキーの画像ねー？　はい！",
		"いいよー、はい！",
		"仕方ないなあ、どうぞー",
		"彼女とデートなうに使っていいよ！",
		"そんなに欲しい...？しょうがないなー",
	}
	return replys[rand.Intn(len(replys))]
}

func postImage(imageURL string) error {
	m := map[string]interface{}{}
	m["icon_url"] = "https://avatars.slack-edge.com/2018-11-21/485246278661_dace41846137494f2582_72.jpg"
	m["channel"] = slackChannel
	m["username"] = "環奈"
	m["text"] = kannaWordWithImage()
	m["attachments"] = []map[string]string{
		map[string]string{
			"fallback":  "ガッキーの画像だよ！",
			"image_url": imageURL,
		},
	}
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := http.Post(slackWebhookURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	return err
}
