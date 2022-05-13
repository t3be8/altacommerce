package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	r "github.com/t3be8/altacommerce/delivery/views/order"
)

type ConfigMidtrans interface {
	CreatePayout(OrderID string, Total float64) map[string]interface{}
	PayOrder(order string) r.ResponsePayOrder
}

type Snap struct {
	snap snap.Client
}

func InitMidrans() *Snap {
	s := snap.Client{}
	s.New("SB-Mid-server-JRM7bi4-TPMjLFDEBmh0w-2h", midtrans.Sandbox)
	return &Snap{
		snap: s,
	}
}

func (s *Snap) CreatePayout(OrderID string, Total float64) map[string]interface{} {
	rb := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  OrderID,
			GrossAmt: int64(Total),
		},
		Callbacks: &snap.Callbacks{
			Finish: "http://18.136.202.111:9001/payorder/finish",
		},
	}
	data, _ := json.Marshal(rb)
	nb := bytes.NewBuffer(data)
	type ResponseWithMap map[string]interface{}
	Resp := ResponseWithMap{}
	err := s.snap.HttpClient.Call(http.MethodPost, "https://app.sandbox.midtrans.com/snap/v1/transactions", &s.snap.ServerKey, s.snap.Options, nb, &Resp)
	fmt.Println(err)
	return Resp
}

func (s *Snap) PayOrder(order string) r.ResponsePayOrder {
	url := fmt.Sprintf("https://api.sandbox.midtrans.com/v2/%s/status", order)
	method := "GET"

	payload := strings.NewReader("\n\n")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err, "err")
	}

	key := s.snap.ServerKey
	EncodeKey := base64.StdEncoding.EncodeToString([]byte(key))
	fmt.Println(EncodeKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", EncodeKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var response r.ResponsePayOrder
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &response)
	return response
}
