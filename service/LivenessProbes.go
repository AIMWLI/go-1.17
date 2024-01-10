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

var resArray = mapset.NewSet("detection", "ai-nlp up", "pong", "hello world", "ok")

func StartLivenessProbesJob() {
	kv := setting.ProbesSetting.KV
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 5s", func() {
		client := util.GetHttpClient()
		for k, v := range kv {
			//fmt.Printf("Address of i=%d:\t%p\n", client, &client)
			resp, err := client.Get(v)
			if resp != nil {
				defer resp.Body.Close()
			}
			if err != nil {
				//fmt.Printf("%v %v", k, err)
				_ = fmt.Errorf("failed to fetch URL %s : %s", v, resp.Status)
			}
			if resp.StatusCode != http.StatusOK {
				SendMsgToBot("探针检测异常", k, v)
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

//receiveId["asr"] = "oc_aa573b380df9e8fde91a106f772dea35"
//receiveId["ocr"] = "oc_bbaa1949e2267b73c3cc32e2e36d448d"

func SendMsgToBot(msg string, k string, v string) {
	//receiveId["nlp"] = "oc_532f1e2b627dd390434eba76cbe2b79a"
	receiveId := map[string]string{"asr": "111", "ocr": "222",
		"nlp": "333", "wenda": "444",
		"stamp": "555",
	}
	botUrl := "http://1.2.3.4:8501/send_to_feishu/send_text"
	values := url.Values{}
	//values.Add("receive_id", "oc_ca3ca124143d8c10fd6cdc4c078aee10") // wenda
	//values.Add("receive_id", "oc_532f1e2b627dd390434eba76cbe2b79a") // nlp
	//values.Add("receive_id", "oc_bbaa1949e2267b73c3cc32e2e36d448d") // ocr
	//values.Add("receive_id", "oc_aa573b380df9e8fde91a106f772dea35") // asr
	values.Add("receive_id", receiveId[k])
	values.Add("text", "<at user_id=\\\\\\\"all\\\\\\\">ALL</at> "+"\\\\n "+k+" "+msg+",\\\\n url: "+v+", \\\\n please check")
	client := util.GetHttpClient()
	//client.Post(url+param, "application/json", bytes.NewReader([]byte(jsonString)))
	resp, err := client.PostForm(botUrl, values)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Printf("%v %v \n", k, err)
		//fmt.Errorf("failed to fetch URL %s : %s", v, resp.Status)
		return
	}
	status := resp.StatusCode
	if status != http.StatusOK {
		fmt.Printf("failed to fetch URL %s : %s", v, resp.Status)
		//fmt.Errorf("send msg to bot error %s : %s", v, resp.Status)
		//return
	}

	//body, err := io.ReadAll(resp.Body)
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s, send to msg to bot %s \n", time.Now().Format(tf), k+" "+msg+" "+v)

}
