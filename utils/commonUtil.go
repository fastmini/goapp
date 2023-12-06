// Package utils
// @Description:
// @Author AN 2023-12-06 23:17:04
package utils

import (
	"fiber/global"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func CTimeCost(start time.Time) {
	tc := time.Since(start)
	global.BLog.Infof("执行完成, 耗时：%v", tc)
}

func CIn(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

func CInMap(target string, result map[interface{}]interface{}) bool {
	var output []string
	for i, _ := range result {
		if str, ok := i.(string); ok {
			output = append(output, str)
		}
		if ii, ok := i.(int64); ok {
			output = append(output, strconv.FormatInt(ii, 10))
		}
		if ii, ok := i.(int32); ok {
			output = append(output, strconv.Itoa(int(ii)))
		}
		if ii, ok := i.(int); ok {
			output = append(output, strconv.Itoa(ii))
		}
	}
	if len(output) > 0 {
		return CIn(target, output)
	}
	return false
}

func CMakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func CTransportUrl(url string) string {
	if !strings.HasPrefix(url, "http") {
		url = strings.Join([]string{"http://", url}, "")
	}
	return url
}

func CString2Int(strArr []string) []int {
	res := make([]int, len(strArr))
	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}
	return res
}

func CString2Int64(strArr []string) []int64 {
	res := make([]int64, len(strArr))
	for index, val := range strArr {
		res[index], _ = strconv.ParseInt(val, 10, 64)
	}
	return res
}

func CStructSliceToMap(source interface{}, filedName string) map[interface{}][]interface{} {
	filedIndex := 0
	v := reflect.ValueOf(source) // 判断，interface转为[]interface{}
	if v.Kind() != reflect.Slice {
		panic("ERROR: Unknown type, slice expected.")
	}
	l := v.Len()
	retList := make([]interface{}, l)
	for i := 0; i < l; i++ {
		retList[i] = v.Index(i).Interface()
	}
	if len(retList) > 0 {
		firstObj := retList[0]
		objT := reflect.TypeOf(firstObj)
		for i := 0; i < objT.NumField(); i++ {
			if objT.Field(i).Name == filedName {
				filedIndex = i
			}
		}
	}

	resMap := make(map[interface{}][]interface{})
	for _, elem := range retList {
		key := reflect.ValueOf(elem).Field(filedIndex).Interface()
		value := make([]interface{}, 0)
		resMap[key] = value
	}

	for _, elem := range retList {
		key := reflect.ValueOf(elem).Field(filedIndex).Interface()
		resMap[key] = append(resMap[key], elem)
	}
	return resMap
}
