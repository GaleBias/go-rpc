package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"go-rpc/server"
)

type Input struct{ Name string }
type Output struct{ Msg string }
type Hello struct {
	FuncFiled func(in *Input) (*Output, error)
}

func (h *Hello) GetServiceName() string {
	return "hello"
}

func SetFuncFeild(val Service) {
	v := reflect.ValueOf(val).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		feild := t.Field(i)
		feildvalue := v.Field(i)
		if feildvalue.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {

				in := args[0].Interface() // in = &Input{Name:"kuangfeng"}
				out := reflect.New(feild.Type.Out(0).Elem()).Interface()

				inData, err := json.Marshal(in) // inData = []byte(`{"name":"kuangfeng"}`)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				serviceName := val.GetServiceName()
				//cfg,_ := icp.GetServiceConfig(serviceName)
				//cfg,_ := ycp.GetServiceConfig(serviceName)
				cfg, _ := App.CfgProvider.GetServiceConfig(serviceName)

				client := http.Client{}
				//resp, err := client.Post(cfg.Endpoint, "application/json", bytes.NewReader(inData))
				req, err = http.NewRequest("POST", cfg.Endpoint, bytes.NewReader(inData))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("sparrow-service", serviceName)
				req.Header.Set("sparrow-service-method", feild.Name)
				resp, err = client.Do(req)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				data, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				err = json.Unmarshal(data, out)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				return []reflect.Value{reflect.ValueOf(out), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			feildvalue.Set(reflect.MakeFunc(feild.Type, fn))
		}
	}
}
func main() {
	icp := NewInMemoryConfigProvider()        // 从内存中初始化配置信息
	_ = InitApplication(WithCfgProvider(icp)) // 将初始化的配置信息,传入给App

	h := &Hello{}
	SetFuncFeild(h)
	data, _ := h.FuncFiled(&Input{
		Name: "nulang",
	})
	fmt.Println(data)
}
