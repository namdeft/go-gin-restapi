package common

type SuccessRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SuccessResponse(data, paging, filter interface{}) *SuccessRes {
	return &SuccessRes{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccessResponse(data interface{}) *SuccessRes {
	return &SuccessRes{
		Data:   data,
		Paging: nil,
		Filter: nil,
	}
}
