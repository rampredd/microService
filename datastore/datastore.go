package main

import (
	"log"

	ds "microService/proto/datastore"
	query "microService/proto/query"

	"database/sql"
	"fmt"
	"microService/common"
	"microService/maria"

	"context"

	"github.com/micro/go-micro"
	nats "github.com/nats-io/go-nats"
)

type Datastore struct{}

var (
	sub *nats.Subscription
	nc  *nats.Conn
	err error
	db  *sql.DB
)

func (datastore *Datastore) Search(ctx context.Context, req *ds.DatastoreRequest, rsp *ds.DatastoreResponse) error {
	log.Print("Received Datastore.Search request")

	fmt.Printf("req is %v\n", req.Req)
	res, err := maria.GetRecipe(req.Req)
	if err != nil {
		common.PrintError(err)
	}
	if len(res) > 0 {
		fmt.Println("db returned ", res)
	} else {
		fmt.Println("query doesn't match any results")
	}

	rsp.Rsp = res
	return nil
}

func (datastore *Datastore) Save(ctx context.Context, req *ds.SaveRequest, rsp *ds.SaveResponse) error {
	log.Print("Received Datastore.Save request")

	fmt.Printf("req is %v\n", req.Req)
	res, err := maria.SaveRecipe(req)
	if err != nil {
		common.PrintError(err)
	} else {
		fmt.Println(res)
	}
	return nil
}

func search() string {
	return "RECIPE NOT FOUND"
}

func main() {
	//	connect to mysql DB server
	db, err = maria.Connect("db")
	if err != nil {
		common.PrintError(err)
		return
	}
	defer db.Close()

	//	connect to nats server
	nc, err = common.ConnectNatsServer()
	if err != nil {
		common.PrintError(err)
		return
	}
	defer nc.Close()

	service := micro.NewService(
		micro.Name("datastore"),
	)

	service.Init()

	// Subscribe to datastore message from nats server
	nc.Subscribe("datastore", func(m *nats.Msg) {
		reply := query.Response{Response: search()}
		rspBytes, err := common.EncByteArray(reply)
		if err != nil {
			common.PrintError(err)
		}
		common.Publish("businessRsp", rspBytes)
	})

	// register datastore handler
	ds.RegisterDatastoreHandler(service.Server(), &Datastore{})
	if err = service.Run(); err != nil {
		common.PrintError(err)
	}
}
