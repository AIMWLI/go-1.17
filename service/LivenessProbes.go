package service

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/robfig/cron/v3"
	"go-gin/pkg/setting"
	"go-gin/util"
	"io"
	"net/http"
	"net/url"
	"time"
)

var tf = "2006-01-02 15:04:05"

var resArray = mapset.NewSet("detection", "ai-nlp up", "pong")

func StartLivenessProbesJob() {
	kv := setting.ProbesSetting.KV
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 20s", func() {
		client := util.GetHttpClient()
		for k, v := range kv {
			//fmt.Printf("Address of i=%d:\t%p\n", client, &client)
			resp, err := client.Get(v)
			if resp != nil {
				defer resp.Body.Close()
			}
			if err != nil || resp.StatusCode != http.StatusOK {
				fmt.Println(err)
				fmt.Errorf("failed to fetch URL %s : %s", v, resp.Status)
				SendMsgToBot("探针检测响应异常", k, v)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				SendMsgToBot("探针检测返回结果异常", k, v)
				continue
			}
			if resArray.Contains(string(body)) {
				fmt.Printf("%s, Service of %s ok\n", time.Now().Format(tf), k)
			}
		}
	})
	c.Start()
}
func SendMsgToBot(msg string, k string, v string) {
	botUrl := "http://xx.xx.xx.xx/send_to_feishu/send_text"
	values := url.Values{}
	values.Add("receive_id", "oc_xxxxxxxxx")
	values.Add("text", "<at user_id=\\\\\\\"all\\\\\\\">ALL</at> "+k+" exception: "+msg+", URL: "+v+" \\\\n Please Check")
	client := util.GetHttpClient()
	//client.Post(url+param, "application/json", bytes.NewReader([]byte(jsonString)))
	resp, err := client.PostForm(botUrl, values)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(err)
		fmt.Errorf("send msg to bot error %s : %s", v, resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s, send to msg to bot %s \n", time.Now().Format(tf), string(body))

}
