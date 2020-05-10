package util

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// todo 可以使用配置文件来动态的加载
	ERROR          = "请稍后再试"
	IdIsFalse      = "id不正确"
	IdIsNil        = "id不为空"
	ParameterError = "请求参数不正确"
)

// todo 根据不同的language 来返回
var ErrorMap map[string]error

func ProcessMongoErr(s string) error {
	if v, ok := ErrorMap[s]; ok {
		return v
	}
	return errors.New(ERROR)
}

func InitErrMap() {
	ErrorMap = make(map[string]error)
	ErrorMap[mongo.ErrNoDocuments.Error()] = errors.New("表不存在")

}
