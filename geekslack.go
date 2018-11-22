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
		postImage(kannaImage())
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

	mes = "なに？"
	return
}

func kannaImage() string {
	urls := []string{
		"https://pbs.twimg.com/profile_images/1009220670826295297/qfHNpNPM_400x400.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/b/b9/Hashimoto_Kanna_at_Opening_Ceremony_of_the_Tokyo_International_Film_Festival_2017_%2840169184582%29.jpg/220px-Hashimoto_Kanna_at_Opening_Ceremony_of_the_Tokyo_International_Film_Festival_2017_%2840169184582%29.jpg",
		"https://dot.asahi.com/S2000/upload/2018101500008_1.jpg",
		"https://news.walkerplus.com/article/166398/970981_615.jpg",
		"https://img.sirabee.com/wp/wp-content/uploads/2018/03/sirabee20180318hasimotokanna1-600x400.jpg",
		"https://iwiz-chie.c.yimg.jp/im_siggYFrYURY5uMNYh993Ack.nw---x320-y320-exp5m-n1/d/iwiz-chie/que-10188509055",
		"http://sabziblog.com/wp-content/uploads/2018/03/568867_615.jpg",
		"https://pbs.twimg.com/media/DBl9HtdVoAAzt1A.jpg",
		"https://realsound.jp/wp-content/uploads/2018/03/20180305-hashimoto2-950x534.jpg",
		"https://realsound.jp/wp-content/uploads/2017/10/20171020_kanna.jpg",
		"https://news-sokuhou.site/wp-content/uploads/2017/10/%E6%A9%8B%E6%9C%AC%E7%92%B0%E5%A5%88%E3%81%AE%E3%83%9E%E3%83%8D%E3%83%BC%E3%82%B8%E3%83%A3%E3%83%BC%E3%81%AF%E8%AA%B0%EF%BC%9F%E5%90%8C%E5%B1%85%E3%82%82%E3%81%97%E3%81%A6%E3%81%84%E3%81%9F%E3%80%8C%E7%BE%A8%E3%81%BE%E3%81%97%E3%81%99%E3%81%8E%E3%82%8B%EF%BC%81%E3%80%8D.jpg",
		"https://i2.wp.com/unicorntomo.com/wp-content/uploads/2018/10/hashimoto10.jpg?fit=700%2C875&ssl=1",
		"https://img.cinematoday.jp/a/N0091045/_size_640x/_v_1492503321/main.jpg",
		"https://bikuchan.com/wp-content/uploads/2017/07/87b742db495de5becade67ea7621bbb6.jpg",
		"https://cdn-ak.f.st-hatena.com/images/fotolife/g/gaigar2444/20180228/20180228202908.jpg",
		"http://スマホ壁紙無料.com/wp-content/uploads/2016/11/hasimotokanna-95-a.jpg",
		"https://i.ytimg.com/vi/08daZMHVlFE/maxresdefault.jpg",
		"http://image.news.livedoor.com/newsimage/1/b/1b2af63a0fb0dc78318cf4f3a49c67c7.jpg",
		"https://grapee.jp/wp-content/uploads/45100_main.jpg",
		"https://mykneecap.com/wp-content/uploads/2018/08/hashimotokanna_pile0022.jpg",
		"https://news.walkerplus.com/article/163509/951502_615.jpg",
		"https://instagram.com/p/Be933O_l52j/media/?size=l",
		"https://o.aolcdn.com/images/dims3/GLOB/crop/2714x1359+0+707/resize/630x315!/format/jpg/quality/85/http%3A%2F%2Fo.aolcdn.com%2Fhss%2Fstorage%2Fmidas%2F57d13f080185c48549e8ef48f2520721%2F206029083%2Fjpp025325840.jpg",
		"https://amd.c.yimg.jp/im_siggyIWiLTdp8lkTrpsgrTGcsg---x400-y267-q90-exp3h-pril/amd/20180830-00000018-flix-000-1-view.jpg",
		"https://tak16.com/wp-content/uploads/2017/06/2017-06-07_18h21_14.png",
		"https://img.dmenumedia.jp/ent/wp-content/uploads/2018/01/285px-Kanna_hashimoto2017.jpg",
		"https://芸能人の実家住所まとめ.com/wp-content/uploads/2018/01/Screenshot-2018-01-13_16-32-21.png",
	}
	return urls[rand.Intn(len(urls))]
}

func kannaWordWithImage() string {
	replys := []string{
		"私の画像ねー？　はい！",
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
			"fallback":  "環奈の画像だよ！",
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
