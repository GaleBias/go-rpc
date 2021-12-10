package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

type Service interface {
	GetServiceName() string
}
type Input struct{ Name string }
type Output struct{ Msg string }
type Hello struct {
	FuncFiled func(in *Input) (*Output, error)
}

func (h *Hello) GetServiceName() string{
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
				cfg,_ := ycp.GetServiceConfig(serviceName)

				client := http.Client{}
				resp, err := client.Post(cfg.Endpoint, "application/json", bytes.NewReader(inData))
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				data, err := ioutil.ReadAll(resp.Body)
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
func main(){
	h := &Hello{}
	SetFuncFeild(h)
	data,_ := h.FuncFiled(&Input{
		Name: "nulang",
	})
	fmt.Println(data)
}