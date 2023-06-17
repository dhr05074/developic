// Package gateway provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package gateway

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"portfolio/schema"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Defines values for ReportStatus.
const (
	Done       ReportStatus = "done"
	Failed     ReportStatus = "failed"
	Pending    ReportStatus = "pending"
	Processing ReportStatus = "processing"
)

// Feedback 사용자의 포트폴리오 분석 보고서의 피드백 구성 요소입니다.
type Feedback = schema.Feedback

// Project 사용자가 포트폴리오에 작성한 프로젝트입니다.
type Project = schema.Project

// Recommendation 사용자의 포트폴리오 분석 보고서의 추천 구성 요소입니다.
type Recommendation = schema.Recommendation

// ReportStatus 포트폴리오 분석 진행 상태입니다.
type ReportStatus string

// RequestId 모든 요청에 사용되는 요청 ID입니다. 서버에서 생성되는 값입니다.
type RequestId = string

// SubmitPortfolioJSONBody defines parameters for SubmitPortfolio.
type SubmitPortfolioJSONBody struct {
	// CareerYears 사용자의 경력입니다. 신입은 0으로 입력합니다.
	CareerYears int `json:"career_years"`

	// Job 사용자가 원하는 직군입니다.
	Job string `json:"job"`

	// PreferredCompany 사용자가 원하는 회사의 특징을 기술합니다.
	PreferredCompany *string `json:"preferred_company,omitempty"`

	// Projects 사용자가 작성한 프로젝트입니다.
	Projects []Project `json:"projects"`

	// TechStacks 사용자가 사용할 수 있는 기술 스택입니다.
	TechStacks []string `json:"tech_stacks"`
}

// SubmitPortfolioJSONRequestBody defines body for SubmitPortfolio for application/json ContentType.
type SubmitPortfolioJSONRequestBody SubmitPortfolioJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 포트폴리오 분석 리포트 조회
	// (GET /report/{request_id})
	GetReportByRequestID(ctx echo.Context, requestId RequestId) error
	// 포트폴리오 제출
	// (POST /submit)
	SubmitPortfolio(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetReportByRequestID converts echo context to params.
func (w *ServerInterfaceWrapper) GetReportByRequestID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "request_id" -------------
	var requestId RequestId

	err = runtime.BindStyledParameterWithLocation("simple", false, "request_id", runtime.ParamLocationPath, ctx.Param("request_id"), &requestId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter request_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReportByRequestID(ctx, requestId)
	return err
}

// SubmitPortfolio converts echo context to params.
func (w *ServerInterfaceWrapper) SubmitPortfolio(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SubmitPortfolio(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/report/:request_id", wrapper.GetReportByRequestID)
	router.POST(baseURL+"/submit", wrapper.SubmitPortfolio)

}

type GetReportByRequestIDRequestObject struct {
	RequestId RequestId `json:"request_id"`
}

type GetReportByRequestIDResponseObject interface {
	VisitGetReportByRequestIDResponse(w http.ResponseWriter) error
}

type GetReportByRequestID200JSONResponse struct {
	// ProjectFeedbacks 사용자의 프로젝트에 대한 피드백입니다.
	ProjectFeedbacks []Feedback `json:"project_feedbacks"`

	// ProjectRecommendations 사용자의 프로젝트에 대한 추천입니다.
	ProjectRecommendations []Recommendation `json:"project_recommendations"`

	// Status 포트폴리오 분석 진행 상태입니다.
	Status ReportStatus `json:"status"`

	// TechStackFeedbacks 사용자의 테크 스택에 대한 피드백입니다.
	TechStackFeedbacks []Feedback `json:"tech_stack_feedbacks"`

	// TechStackRecommendations 사용자의 테크 스택에 대한 추천입니다.
	TechStackRecommendations []Recommendation `json:"tech_stack_recommendations"`
}

func (response GetReportByRequestID200JSONResponse) VisitGetReportByRequestIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetReportByRequestID404JSONResponse struct {
	// Message 요청 ID에 해당하는 리포트가 존재하지 않거나, 아직 생성되지 않았습니다.
	Message string `json:"message"`
}

func (response GetReportByRequestID404JSONResponse) VisitGetReportByRequestIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type SubmitPortfolioRequestObject struct {
	Body *SubmitPortfolioJSONRequestBody
}

type SubmitPortfolioResponseObject interface {
	VisitSubmitPortfolioResponse(w http.ResponseWriter) error
}

type SubmitPortfolio202JSONResponse struct {
	// RequestId 모든 요청에 사용되는 요청 ID입니다. 서버에서 생성되는 값입니다.
	RequestId RequestId `json:"request_id"`
}

func (response SubmitPortfolio202JSONResponse) VisitSubmitPortfolioResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// 포트폴리오 분석 리포트 조회
	// (GET /report/{request_id})
	GetReportByRequestID(ctx context.Context, request GetReportByRequestIDRequestObject) (GetReportByRequestIDResponseObject, error)
	// 포트폴리오 제출
	// (POST /submit)
	SubmitPortfolio(ctx context.Context, request SubmitPortfolioRequestObject) (SubmitPortfolioResponseObject, error)
}

type StrictHandlerFunc func(ctx echo.Context, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetReportByRequestID operation middleware
func (sh *strictHandler) GetReportByRequestID(ctx echo.Context, requestId RequestId) error {
	var request GetReportByRequestIDRequestObject

	request.RequestId = requestId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetReportByRequestID(ctx.Request().Context(), request.(GetReportByRequestIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetReportByRequestID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetReportByRequestIDResponseObject); ok {
		return validResponse.VisitGetReportByRequestIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// SubmitPortfolio operation middleware
func (sh *strictHandler) SubmitPortfolio(ctx echo.Context) error {
	var request SubmitPortfolioRequestObject

	var body SubmitPortfolioJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.SubmitPortfolio(ctx.Request().Context(), request.(SubmitPortfolioRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SubmitPortfolio")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(SubmitPortfolioResponseObject); ok {
		return validResponse.VisitSubmitPortfolioResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}