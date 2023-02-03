// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClienta(accessKeyId *string, accessKeySecret *string) (_result *openapi.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("cdn.aliyuncs.com")
	_result = &openapi.Client{}
	_result, _err = openapi.NewClient(config)
	return _result, _err
}

/**
 * API 相关
 * @param path params
 * @return OpenApi.Params
 */
func CreateApiInfoa() (_result *openapi.Params) {
	params := &openapi.Params{
		// 接口名称
		Action: tea.String("RefreshObjectCaches"),
		// 接口版本
		Version: tea.String("2018-05-10"),
		// 接口协议
		Protocol: tea.String("HTTPS"),
		// 接口 HTTP 方法
		Method:   tea.String("POST"),
		AuthType: tea.String("AK"),
		Style:    tea.String("RPC"),
		// 接口 PATH
		Pathname: tea.String("/"),
		// 接口请求体内容格式
		ReqBodyType: tea.String("json"),
		// 接口响应体内容格式
		BodyType: tea.String("json"),
	}
	_result = params
	return _result
}

func _main(args []*string) (_err error) {
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClienta(tea.String(os.Getenv("ALIYUN_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIYUN_ACCESS_KEY_SECRET")))
	if _err != nil {
		return _err
	}

	params := CreateApiInfoa()
	// query params
	queries := map[string]interface{}{}
	queries["ObjectPath"] = tea.String("https://ali-cdn-poc.lego.cn/t.png")
	queries["ObjectType"] = tea.String("File")
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query: openapiutil.Query(queries),
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode
	_, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return _err
	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
