package dto

import (
	"OceanLearn/model"
	pb "OceanLearn/proto"
	"encoding/json"
	"fmt"
)

var (
	dict model.DictResponse
)

func ToDictDto(resp *pb.DictResponse) model.DictResponse {
	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("err：", err)

	}
	err = json.Unmarshal(jsonData, &dict)
	if err != nil {
		fmt.Println("err：", err)

	}
	return dict
}
