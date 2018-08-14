package main

import (
	"log"

	bs "microService/proto/business"
	gw "microService/proto/gateway"
	query "microService/proto/query"

	"context"
	"fmt"
	"microService/common"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"

	nats "github.com/nats-io/go-nats"
)

type Gateway struct{}

var (
	err      error
	queryRsp chan *query.Response
	nc       *nats.Conn
	bsClient bs.BusinessService
)

// Gateway.Search is a method which will be served by http request /gateway/search
// In the event we see /[service]/[method] the [service] is used as part of the method
// E.g /gateway/search goes to go.micro.api.search Gateway.Search
func (gateway *Gateway) Search(ctx context.Context, req *gw.QueryRequest, rsp *gw.QueryResponse) error {
	log.Print("Received Gateway.Search request")

	if len(req.Req.Q) == 0 {
		return errors.BadRequest("gateway", "please enter mandatory fields")
	}

	start := time.Now()

	b := &bs.BusinessRequest{}
	b.Req = req.Req
	res, err := bsClient.Search(context.TODO(), b)
	if nil != err {
		common.PrintError(err)
	}

	fmt.Println("got bs.search reply ", res, " and time taken to get this is ", time.Since(start))
	rsp.Rsp = res.Rsp
	return nil
}

func main() {
	//	create client for business service
	nc, err = common.ConnectNatsServer()
	if err != nil {
		return
	}
	defer nc.Close()
	service := micro.NewService(
		micro.Name("go.micro.api.gateway"),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init()

	// Subscribe to gateway message from nats server
	nc.Subscribe("gateway", func(m *nats.Msg) {
		intf, err := common.DecQueryRsp(m.Data)
		rsp := intf.(query.Response)
		if nil != err {
			common.PrintError(err)
		}
		fmt.Println("Received curl async message: ", rsp.Response)
	})

	bsClient = bs.NewBusinessService("business", service.Client())

	// register gateway handler
	gw.RegisterGatewayHandler(service.Server(), &Gateway{})
	if err = service.Run(); err != nil {
		common.PrintError(err)
	}
}
