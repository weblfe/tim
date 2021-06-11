package tim

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewCustomElem(t *testing.T) {
	var src = NewCustomElem("111111", "123123", "{}", "888")
	fmt.Println(toJson(src))
}

func TestNewImageElem(t *testing.T) {
	var src = NewImageElem(MsgImageContent{
		UUID: "1111",
		ImageFormat:1,
		ImageInfoArray:[]ImageInfoItem{
			{
				Type:0,
				Size:0,
				Width:0,
				Height:0,
				URL:"https://www.baidu.com",
			},
		},
	})
	fmt.Println(toJson(src))
}


func toJson(src interface{}) string {
	_json, err := json.Marshal(src)
	if err != nil {
		return ""
	}
	return string(_json)
}
