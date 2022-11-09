package helpers

import (
	"fmt"
)

type MetaParam struct {
	Limit  int64
	Page   int64
	Offset int64
}

func GetMetaParam(cpage int, climit int) (MetaParam, error) {
	var int_limit int64 = 10
	var int_page int64 = 1
	var int_offset int64

	if int64(climit) < 1 {
		int_page = 1
	} else {
		int_page = int64(cpage)
	}
	if int64(climit) < 1 {
		return MetaParam{}, fmt.Errorf("Limit have to more than 0")
	}else{
		int_limit = int64(climit)
	}

	if int_page == 1 {
		int_offset = 0
	} else {
		int_offset = (int64(int_page) - 1) * int64(int_limit)
	}

	var meta_param MetaParam
	meta_param.Page = int_page
	meta_param.Limit = int_limit
	meta_param.Offset = int_offset
	return meta_param, nil
}
