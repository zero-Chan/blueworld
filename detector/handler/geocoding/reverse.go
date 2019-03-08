package geocoding

import (
	"net/http"
	"io/ioutil"
	"log"

	"zero-Chan/blueworld/detector/handler/base"
	"k8s.io/apimachinery/pkg/util/json"
	"fmt"
)

type ReverseRequest struct {
	// 经度
	Longitude float64

	// 纬度
	Latitude float64
}

type ReverseResponse struct {
	// 国家
	Country string

	// 省
	Province string

	// 市
	City string
}

type reverseCore struct {
	base.BaseHttpCore

	// 高德地图url
	AmapApi string

	// 高德地图key
	AmapApiKey string
}

func NewReverseCore() *reverseCore {
	core := &reverseCore{
		AmapApi:    "https://restapi.amap.com/v3/geocode/regeo",
		AmapApiKey: "236a43ab49661fa2eb504bbb72a7b987",
	}

	return core
}

func (this *reverseCore) ServeHTTP(httpRespw http.ResponseWriter, httpReq *http.Request) {
	// read http body
	bodyStream, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		log.Printf("read http body fail: %s", err)
		return
	}

	request := ReverseRequest{}
	err = json.Unmarshal(bodyStream, &request)
	if err != nil {
		log.Printf("json unmarshal http body fail: %s", err)
		return
	}

	err = this.CallAmap(&request)
	if err != nil {

	}
}

func (this *reverseCore) NewRequest() interface{} {
	return &ReverseRequest{}
}

func (this *reverseCore) CallAmap(req *ReverseRequest) (resp ReverseResponse, err error) {
	// 广州
	// curl 'https://restapi.amap.com/v3/geocode/regeo?key=236a43ab49661fa2eb504bbb72a7b987&location=113.27,23.13'
	// response:
	// {"status":"1","regeocode":{"addressComponent":{"city":"广州市","province":"广东省","adcode":"440104","district":"越秀区","towncode":"440104003000","streetNumber":{"number":"3号","location":"113.270037,23.1302619","direction":"北","distance":"29.3722","street":"都府街"},"country":"中国","township":"北京街道","businessAreas":[{"location":"113.2833179485815,23.13343970212761","name":"建设","id":"440104"},{"location":"113.27805496082468,23.13224423161515","name":"东风","id":"440104"},{"location":"113.29835622352945,23.131624339792353","name":"东风东","id":"440104"}],"building":{"name":[],"type":[]},"neighborhood":{"name":"都府居住小区","type":"商务住宅;住宅区;住宅小区"},"citycode":"020"},"formatted_address":"广东省广州市越秀区北京街道都府居住小区"},"info":"OK","infocode":"10000"}

	httpResp, err := http.Get(fmt.Sprintf("%s?key=%s&location=%f,%f", this.AmapApi, this.AmapApiKey, req.Longitude, req.Latitude))
	if err != nil {
		return
	}

	return
}

// 外部包
func GetModule() *reverseCore {
	core := NewReverseCore()
	return core
}
