// Code generated by go generate; DO NOT EDIT.
// This file was generated by go-doudou at
// 2021-05-29 14:23:58.622958 +0800 CST m=+0.003111083
package vo

import (
	"encoding/json"
	"github.com/unionj-cloud/go-doudou/name/strategies"
)


func (object PageFilter) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap[strategies.LowerCaseConvert("Name")] = object.Name
	objectMap[strategies.LowerCaseConvert("Dept")] = object.Dept
	return json.Marshal(objectMap)
}

func (object Order) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap[strategies.LowerCaseConvert("Col")] = object.Col
	objectMap[strategies.LowerCaseConvert("Sort")] = object.Sort
	return json.Marshal(objectMap)
}

func (object Page) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap[strategies.LowerCaseConvert("Orders")] = object.Orders
	objectMap[strategies.LowerCaseConvert("PageNo")] = object.PageNo
	objectMap[strategies.LowerCaseConvert("Size")] = object.Size
	return json.Marshal(objectMap)
}

func (object PageQuery) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap[strategies.LowerCaseConvert("Filter")] = object.Filter
	objectMap[strategies.LowerCaseConvert("Page")] = object.Page
	return json.Marshal(objectMap)
}

func (object PageRet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap[strategies.LowerCaseConvert("Items")] = object.Items
	objectMap[strategies.LowerCaseConvert("PageNo")] = object.PageNo
	objectMap[strategies.LowerCaseConvert("PageSize")] = object.PageSize
	objectMap[strategies.LowerCaseConvert("Total")] = object.Total
	objectMap[strategies.LowerCaseConvert("HasNext")] = object.HasNext
	return json.Marshal(objectMap)
}

func (object UserVo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap[strategies.LowerCaseConvert("Id")] = object.Id
	objectMap[strategies.LowerCaseConvert("Name")] = object.Name
	objectMap[strategies.LowerCaseConvert("Phone")] = object.Phone
	objectMap[strategies.LowerCaseConvert("Dept")] = object.Dept
	return json.Marshal(objectMap)
}

