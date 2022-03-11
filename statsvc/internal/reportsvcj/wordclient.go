package reportsvcj

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	_querystring "github.com/google/go-querystring/query"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/registry"
)

type WordClient struct {
	provider registry.IServiceProvider
	client   *resty.Client
	rootPath string
}

func (receiver *WordClient) SetRootPath(rootPath string) {
	receiver.rootPath = rootPath
}

func (receiver *WordClient) SetProvider(provider registry.IServiceProvider) {
	receiver.provider = provider
}

func (receiver *WordClient) SetClient(client *resty.Client) {
	receiver.client = client
}
func (receiver *WordClient) PostWord(ctx context.Context, _headers map[string]string,
	bodyJSON RequestPayload) (ret ResultString, _resp *resty.Response, err error) {
	var _err error

	_req := receiver.client.R()
	_req.SetContext(ctx)
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_req.SetBody(bodyJSON)

	_resp, _err = _req.Post("/word")
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	if _err = json.Unmarshal(_resp.Body(), &ret); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	return
}

// GetWordQueryTabColumnSize 查询表格列数量
func (receiver *WordClient) GetWordQueryTabColumnSize(ctx context.Context, _headers map[string]string,
	queryParams struct {
		// required
		TemplateUrl string `json:"templateUrl,omitempty" url:"templateUrl"`
		// required
		TabFieldName string `json:"tabFieldName,omitempty" url:"tabFieldName"`
	}) (ret ResultInteger, _resp *resty.Response, err error) {
	var _err error

	_req := receiver.client.R()
	_req.SetContext(ctx)
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_queryParams, _ := _querystring.Values(queryParams)
	_req.SetQueryParamsFromValues(_queryParams)

	_resp, _err = _req.Get("/word/queryTabColumnSize")
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	if _err = json.Unmarshal(_resp.Body(), &ret); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	return
}

// PostWordPreview 预览图表
func (receiver *WordClient) PostWordPreview(ctx context.Context, _headers map[string]string,
	queryParams struct {
		// required
		TemplateUrl string `json:"templateUrl,omitempty" url:"templateUrl"`
	},
	bodyJSON ParagraphWrapper) (ret ResultString, _resp *resty.Response, err error) {
	var _err error

	_req := receiver.client.R()
	_req.SetContext(ctx)
	if len(_headers) > 0 {
		_req.SetHeaders(_headers)
	}
	_queryParams, _ := _querystring.Values(queryParams)
	_req.SetQueryParamsFromValues(_queryParams)
	_req.SetBody(bodyJSON)

	_resp, _err = _req.Post("/word/preview")
	if _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	if _resp.IsError() {
		err = errors.New(_resp.String())
		return
	}
	if _err = json.Unmarshal(_resp.Body(), &ret); _err != nil {
		err = errors.Wrap(_err, "")
		return
	}
	return
}

func NewWord(opts ...ddhttp.DdClientOption) *WordClient {
	defaultProvider := ddhttp.NewServiceProvider("REPORT_SVC_J")
	defaultClient := ddhttp.NewClient()

	svcClient := &WordClient{
		provider: defaultProvider,
		client:   defaultClient,
	}

	for _, opt := range opts {
		opt(svcClient)
	}

	svcClient.client.OnBeforeRequest(func(_ *resty.Client, request *resty.Request) error {
		request.URL = svcClient.provider.SelectServer() + svcClient.rootPath + request.URL
		return nil
	})

	svcClient.client.SetPreRequestHook(func(_ *resty.Client, request *http.Request) error {
		traceReq, _ := nethttp.TraceRequest(opentracing.GlobalTracer(), request,
			nethttp.OperationName(fmt.Sprintf("HTTP %s: %s", request.Method, request.URL.Path)))
		*request = *traceReq
		return nil
	})

	svcClient.client.OnAfterResponse(func(_ *resty.Client, response *resty.Response) error {
		nethttp.TracerFromRequest(response.Request.RawRequest).Finish()
		return nil
	})

	return svcClient
}
