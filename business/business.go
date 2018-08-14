package main

import (
	"encoding/json"
	"fmt"
	"log"
	bs "microService/proto/business"
	ds "microService/proto/datastore"
	query "microService/proto/query"
	"net/http"

	"microService/common"

	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	nats "github.com/nats-io/go-nats"
)

type Business struct{}

var (
	sub      *nats.Subscription
	nc       *nats.Conn
	err      error
	req      query.Request
	dsClient ds.DatastoreService
	apiReq   common.ApiReq
	msg      chan string
)

// Business.Search is a method which will be served by http request /business/search
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /business/search goes to Business.Search
func (business *Business) Search(ctx context.Context, req *bs.BusinessRequest, rsp *bs.BusinessResponse) error {
	log.Print("Received Business.Search request")
	if len(req.Req.Q) == 0 {
		return errors.BadRequest("business", "please enter mandatory fields")
	}

	d := &ds.DatastoreRequest{}
	d.Req = req.Req
	res, err := dsClient.Search(context.TODO(), d)
	if nil != err {
		common.PrintError(err)
	}

	qRsp := &query.Response{}

	if res.Rsp == "" {
		//		post message to trigger external API request
		reqBytes, err := common.EncByteArray(req.Req)
		if err != nil {
			common.PrintError(err)
		}
		common.Publish("sendExternalApiReq", reqBytes)
		//		wait for response from external api
		qRsp.Response = <-msg
		//		same response in DB
		saveReq := &ds.SaveRequest{}
		saveReq.Req = req.Req
		saveReq.Recipe = qRsp.Response
		dsClient.Save(context.TODO(), saveReq)
	} else {
		fmt.Println("got ds.search reply ", res)
		qRsp.Response = res.Rsp
	}
	rsp.Rsp = qRsp
	return nil
}

func getExternalApi(url string) (strResult string) {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "THERE IS SOME PROBLEM WITH URL"
	}
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "HTTP GET FAILED"
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record interface{}

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Printf("external API result %v\n", record)
	if nil != record {
		strResult = record.(string)
	} else {
		strResult = "HTTP GET FAILED"
	}
	msg <- strResult
	fmt.Printf("strResult ADDRESS IS %p\n", &strResult)
	return strResult
}

func main() {
	nc, err = common.ConnectNatsServer()
	if err != nil {
		return
	}
	defer nc.Close()

	//	get API Request params decoded from file and saved in global variable
	apiReq, err = common.GetApiReq("app")
	if nil != err {
		common.PrintError(err)
		return
	}

	service := micro.NewService(
		micro.Name("business"),
	)

	service.Init()

	//make string channel
	msg = make(chan string)
	// Subscribe to business message from nats server
	nc.Subscribe("businessReq", func(m *nats.Msg) {
		//		store request to process further
		reqIntf, err := common.DecQueryReq(m.Data)
		if nil != err {
			common.PrintError(err)
		}
		req = reqIntf.(query.Request)
		common.Publish("datastore", m.Data)
	})

	nc.Subscribe("sendExternalApiReq", func(m *nats.Msg) {
		intf, err := common.DecQueryReq(m.Data)
		if nil != err {
			common.PrintError(err)
		}
		req := intf.(query.Request)
		//			call external API
		//		https://api.edamam.com/search?q=chicken&app_id=${YOUR_APP_ID}&app_key=${YOUR_APP_KEY}

		//			url := fmt.Sprintf("https://api.edamam.com/search?q=%s%c%s%s}%c%s%s}", req.Q, 38, "app_id=${", req.AppId, 38, "app_key=${", req.AppKey)
		url := fmt.Sprintf("https://api.edamam.com/search?q=%s&app_id=${%s}&app_key=${%s}", req.Q, apiReq.AppId, apiReq.AppKey)

		rsp := getExternalApi(url)
		fmt.Printf("rsp ADDRESS IS %p\n", &rsp)
		fmt.Println("rsp is ", rsp)
		fmt.Println("url is ", url)
		//		msg <- rsp
	})

	dsClient = ds.NewDatastoreService("datastore", service.Client())
	// register business handler
	bs.RegisterBusinessHandler(service.Server(), &Business{})
	if err = service.Run(); err != nil {
		common.PrintError(err)
	}
}
