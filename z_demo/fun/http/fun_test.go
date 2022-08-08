package http

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"testing"
)

var header = []string{
	"Authorization",
	"RequestStack",
}

var getDataHeader map[string]string
var setDataHeader map[string]string

type Resp struct {
	State int      `json:"state"`
	Data  RespData `json:"data"`
	Msg   string   `json:"msg"`
}

type RespData struct {
	Items []ListItem `json:"items"`
}

type ListItem struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func TestHttp(t *testing.T) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	//获取数据url
	getDataUrl := `http://dp-5e16c665c2e35.qcloud-prod.oneitfarm.com/main.php/json/proxy/call?url=/main.php/json/task/all`
	//保存数据url
	saveDataUrl := `http://dp-5e620de0033d3.qcloud-prod.oneitfarm.com/main.php/json/proxy/call?url=/main.php/json/task/add`
	respData := getUmlData(req, resp, getDataUrl)
	clearHeader(req)
	setUmlData(req, resp, respData, saveDataUrl)
	t.Log("success")
}

func getUmlData(req *fasthttp.Request, resp *fasthttp.Response, getDataUrl string) *Resp {
	setReqHeader(req, getDataHeader)
	req.PostArgs().Set("filter[type]", "task_demand_usecase")
	req.PostArgs().Set("order", "order by id desc")
	req.PostArgs().Set("function_id", "432F0fdaC8164717")
	req.SetRequestURI(getDataUrl)

	if err := fasthttp.Do(req, resp); err != nil {
		log.Fatal("请求失败:", err.Error())
		return nil
	}
	data := &Resp{}
	json.Unmarshal(resp.Body(), data)
	return data
}

func setUmlData(req *fasthttp.Request, resp *fasthttp.Response, data *Resp, saveDataUrl string) {
	req.SetRequestURI(saveDataUrl)
	setReqHeader(req, setDataHeader)
	for _, v := range data.Data.Items {
		req.PostArgs().Set("name", v.Name)
		req.PostArgs().Set("type", "task_demand_usecase")
		req.PostArgs().Set("category_id", "992")
		req.PostArgs().Set("content", v.Content)
		req.PostArgs().Set("function_id", "617d6a002334c0e0")
		if err := fasthttp.Do(req, resp); err != nil {
			log.Fatal("请求失败:", err.Error())
			continue
		}
		log.Println(string(resp.Body()))
	}
}

func setReqHeader(req *fasthttp.Request, headerVal map[string]string) {
	for k, v := range headerVal {
		req.Header.Add(k, v)
	}
}

func clearHeader(req *fasthttp.Request) {
	for _, h := range header {
		req.Header.Del(h)
	}
}

func initHeaderData() {
	getDataHeader = make(map[string]string)
	getDataHeader["Authorization"] = "C5sFtPgkbrUJmAIa.u2vNPXRcUOiaeY1B.e2e287bff2a661c6130156b7baccec31"
	getDataHeader["RequestStack"] = `[{"appid":"fnue42okwlghuw5tohvmbsxqiz1amgcz","appkey":"2532ebecd93d470b9d9d8275fb941188","channel":"2"},{"appid":"uh49y8vwmxp5rsiqthfjm62ynxbajofc","appkey":"3b960c7ddfda45a3b1a815e0335ae08c","channel":"159"},{"appid":"opygwmeutz6kt15umavlwyrcjqqok0ni","channelAlias":"default"}]`
	setDataHeader = make(map[string]string)
	setDataHeader["Authorization"] = "sY952B0lSQ4gf8NA.XvHfT7LZlrAqoR89.80fcd1a985e84a0db0088e6d8ebdbd17"
	setDataHeader["RequestStack"] = `[{"appid":"fnue42okwlghuw5tohvmbsxqiz1amgcz","appkey":"97d22b840069424ead42668532163d30","channel":"23"},{"appid":"uh49y8vwmxp5rsiqthfjm62ynxbajofc","appkey":"edc8af2cfdaf438e9e1dc301234e13b9","channel":"335"},{"appid":"opygwmeutz6kt15umavlwyrcjqqok0ni","channelAlias":"default"}]`
}
