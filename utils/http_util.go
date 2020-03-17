package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type HTTPData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

type HTTPPageData struct {
	NumsPerPage int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Count       int         `json:"count"`
	TotalPages  int         `json:"totalPages"`
	Data        interface{} `json:"data"`
}

func ReturnHTTPSuccess(this *beego.Controller, val interface{}) {

	rtndata := HTTPData{
		ErrNo:  0,
		ErrMsg: "",
		Data:   val,
	}

	data, err := json.Marshal(rtndata)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
}

func GetHTTPRtnJsonData(errno int, errmsg string) interface{} {

	rtndata := HTTPData{
		ErrNo:  errno,
		ErrMsg: errmsg,
		Data:   nil,
	}
	data, _ := json.Marshal(rtndata)

	return json.RawMessage(string(data))

}

func GetPageData(rawData []orm.Params, page int, size int) HTTPPageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []orm.Params

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return HTTPPageData{
		NumsPerPage: size,
		CurrentPage: page,
		Count:       count,
		TotalPages:  totalpages,
		Data:        pagedata,
	}
}
