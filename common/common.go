package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
	query "microService/proto/query"

	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"

	nats "github.com/nats-io/go-nats"
)

// define our own host type
type DbCfg struct {
	Address   string `json:"address"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	DbName    string `json:"dbName"`
	TableName string `json:"tableName"`
}

// define our own apiReq params
type ApiReq struct {
	AppKey string `json:"appKey"`
	AppId  string `json:"appId"`
}

var (
	//	sub *nats.Subscription
	nc  *nats.Conn
	err error
)

func GetApiReq(param string) (req ApiReq, _ error) {
	err = loadConfig()
	if nil != err {
		return req, err
	}
	if err = config.Get("config", param).Scan(&req); err != nil {
		return req, err
	}

	return req, nil
}

func loadConfig() error {
	//	read config file
	// load the config from a file source
	if err = config.Load(file.NewSource(
		file.WithPath("./config/config.json"),
	)); err != nil {
		return err
	}
	return nil
}

func GetDbConfig(param string) (dbCfg DbCfg, _ error) {
	err = loadConfig()
	if nil != err {
		return dbCfg, err
	}
	if err = config.Get("config", param).Scan(&dbCfg); err != nil {
		return dbCfg, err
	}

	return dbCfg, nil
}

func PrintError(e error) {
	fmt.Println("error in nats connect", e)
}

func ConnectNatsServer() (*nats.Conn, error) {
	//	connect to nats server
	nc, err = nats.Connect(nats.DefaultURL)
	if nil != err {
		PrintError(err)
		return nil, err
	}
	fmt.Println("connected to nats server")
	return nc, nil
}

func Publish(sub string, b []byte) {
	nc.Publish(sub, b)
	//	fmt.Println("nc is ", nc)
	//	m, err := sub.NextMsg(time.Microsecond * 100)
	//	if nil != err {
	//		fmt.Println("error in subsciber: ", err)
	//		rsp.Response = "error in search request"
	//	} else {
	//		intf, err := common.DecQueryReq(m.Data)
	//		y := intf.(query.QueryRequest)
	//		if nil != err {
	//			common.PrintError(err)
	//		} else {
	//			rsp.Response = "got your request " + y.Q
	//		}
	//	}
}

func EncByteArray(enc interface{}) ([]byte, error) {
	// gob encoding
	encBuf := &bytes.Buffer{}
	err := gob.NewEncoder(encBuf).Encode(enc)
	return encBuf.Bytes(), err
}

func DecQueryReq(b []byte) (interface{}, error) {
	// gob decoding
	decBuf := bytes.NewBuffer(b)
	req := query.Request{}
	err := gob.NewDecoder(decBuf).Decode(&req)
	return req, err
}

func DecQueryRsp(b []byte) (interface{}, error) {
	// gob decoding
	decBuf := bytes.NewBuffer(b)
	rsp := query.Response{}
	err := gob.NewDecoder(decBuf).Decode(&rsp)
	return rsp, err
}
