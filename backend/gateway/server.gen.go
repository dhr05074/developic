// Package gateway provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
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

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	ApiKeyScopes = "apiKey.Scopes"
)

// Defines values for ProgrammingLanguage.
const (
	Cpp        ProgrammingLanguage = "Cpp"
	Go         ProgrammingLanguage = "Go"
	Javascript ProgrammingLanguage = "Javascript"
)

// Code Developic에서 사용되는 코드 데이터입니다.
// 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다.
type Code = string

// ELOScore Developic에서 유저의 실력을 가늠하는 ELO 점수입니다.
type ELOScore = int32

// Error defines model for Error.
type Error struct {
	// Code 오류 구분 코드입니다.
	Code string `json:"code"`

	// Message 상세한 오류 정보를 알려주는 메시지입니다.
	Message string `json:"message"`
}

// Problem defines model for Problem.
type Problem struct {
	// Code Developic에서 사용되는 코드 데이터입니다.
	// 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다.
	Code Code `json:"code"`

	// Description 문제의 자세한 설명입니다.
	Description string `json:"description"`

	// Id Developic에서 출제한 문제의 고유 ID입니다.
	Id ProblemID `json:"id"`

	// Title Developic에서 출제한 문제의 타이틀입니다.
	Title ProblemTitle `json:"title"`
}

// ProblemID Developic에서 출제한 문제의 고유 ID입니다.
type ProblemID = string

// ProblemTitle Developic에서 출제한 문제의 타이틀입니다.
type ProblemTitle = string

// ProgrammingLanguage Developic에서 사용할 프로그래밍 언어입니다.
type ProgrammingLanguage string

// Record defines model for Record.
type Record struct {
	// Code Developic에서 사용되는 코드 데이터입니다.
	// 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다.
	Code Code `json:"code"`

	// Efficiency 결과 보고서에서 사용자가 취득한 총점입니다.
	Efficiency Score `json:"efficiency"`

	// Id Developic에서 생성된 결과 보고서의 고유 ID입니다.
	Id RecordID `json:"id"`

	// ProblemId Developic에서 출제한 문제의 고유 ID입니다.
	ProblemId ProblemID `json:"problem_id"`

	// ProblemTitle Developic에서 출제한 문제의 타이틀입니다.
	ProblemTitle ProblemTitle `json:"problem_title"`

	// Readability 결과 보고서에서 사용자가 취득한 총점입니다.
	Readability Score `json:"readability"`

	// Robustness 결과 보고서에서 사용자가 취득한 총점입니다.
	Robustness Score `json:"robustness"`
}

// RecordID Developic에서 생성된 결과 보고서의 고유 ID입니다.
type RecordID = string

// Score 결과 보고서에서 사용자가 취득한 총점입니다.
type Score = int32

// N200GetProblem defines model for 200GetProblem.
type N200GetProblem = Problem

// N200GetRecord defines model for 200GetRecord.
type N200GetRecord = Record

// N200GetRecords defines model for 200GetRecords.
type N200GetRecords struct {
	Records []Record `json:"records"`
}

// N202Submit defines model for 202Submit.
type N202Submit struct {
	// RecordId Developic에서 생성된 결과 보고서의 고유 ID입니다.
	RecordId RecordID `json:"record_id"`
}

// InternalServerError defines model for InternalServerError.
type InternalServerError = Error

// SubmitCode defines model for SubmitCode.
type SubmitCode struct {
	// Code Developic에서 사용되는 코드 데이터입니다.
	// 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다.
	Code Code `json:"code"`

	// ProblemId Developic에서 출제한 문제의 고유 ID입니다.
	ProblemId ProblemID `json:"problem_id"`
}

// RequestProblemJSONBody defines parameters for RequestProblem.
type RequestProblemJSONBody struct {
	// EloScore Developic에서 유저의 실력을 가늠하는 ELO 점수입니다.
	EloScore *ELOScore `json:"elo_score,omitempty"`

	// Language Developic에서 사용할 프로그래밍 언어입니다.
	Language ProgrammingLanguage `json:"language"`
}

// GetRecordsParams defines parameters for GetRecords.
type GetRecordsParams struct {
	// Page 조회할 페이지입니다.
	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	// Limit 한 페이지당 조회할 아이템의 수입니다.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

// SubmitSolutionJSONBody defines parameters for SubmitSolution.
type SubmitSolutionJSONBody struct {
	// Code Developic에서 사용되는 코드 데이터입니다.
	// 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다.
	Code Code `json:"code"`

	// ProblemId Developic에서 출제한 문제의 고유 ID입니다.
	ProblemId ProblemID `json:"problem_id"`
}

// RequestProblemJSONRequestBody defines body for RequestProblem for application/json ContentType.
type RequestProblemJSONRequestBody RequestProblemJSONBody

// SubmitSolutionJSONRequestBody defines body for SubmitSolution for application/json ContentType.
type SubmitSolutionJSONRequestBody SubmitSolutionJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /me)
	GetMe(ctx echo.Context) error

	// (POST /problems)
	RequestProblem(ctx echo.Context) error

	// (GET /problems/{id})
	GetProblem(ctx echo.Context, id ProblemID) error

	// (GET /records)
	GetRecords(ctx echo.Context, params GetRecordsParams) error

	// (GET /records/{id})
	GetRecord(ctx echo.Context, id RecordID) error

	// (POST /submit)
	SubmitSolution(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetMe(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMe(ctx)
	return err
}

// RequestProblem converts echo context to params.
func (w *ServerInterfaceWrapper) RequestProblem(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RequestProblem(ctx)
	return err
}

// GetProblem converts echo context to params.
func (w *ServerInterfaceWrapper) GetProblem(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id ProblemID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(ApiKeyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProblem(ctx, id)
	return err
}

// GetRecords converts echo context to params.
func (w *ServerInterfaceWrapper) GetRecords(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRecordsParams
	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRecords(ctx, params)
	return err
}

// GetRecord converts echo context to params.
func (w *ServerInterfaceWrapper) GetRecord(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id RecordID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(ApiKeyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRecord(ctx, id)
	return err
}

// SubmitSolution converts echo context to params.
func (w *ServerInterfaceWrapper) SubmitSolution(ctx echo.Context) error {
	var err error

	ctx.Set(ApiKeyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SubmitSolution(ctx)
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

	router.GET(baseURL+"/me", wrapper.GetMe)
	router.POST(baseURL+"/problems", wrapper.RequestProblem)
	router.GET(baseURL+"/problems/:id", wrapper.GetProblem)
	router.GET(baseURL+"/records", wrapper.GetRecords)
	router.GET(baseURL+"/records/:id", wrapper.GetRecord)
	router.POST(baseURL+"/submit", wrapper.SubmitSolution)

}

type N200GetProblemJSONResponse Problem

type N200GetRecordJSONResponse Record

type N200GetRecordsJSONResponse struct {
	Records []Record `json:"records"`
}

type N202SubmitJSONResponse struct {
	// RecordId Developic에서 생성된 결과 보고서의 고유 ID입니다.
	RecordId RecordID `json:"record_id"`
}

type InternalServerErrorJSONResponse Error

type GetMeRequestObject struct {
}

type GetMeResponseObject interface {
	VisitGetMeResponse(w http.ResponseWriter) error
}

type GetMe200JSONResponse struct {
	// EloScore Developic에서 유저의 실력을 가늠하는 ELO 점수입니다.
	EloScore ELOScore `json:"elo_score"`
}

func (response GetMe200JSONResponse) VisitGetMeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type RequestProblemRequestObject struct {
	Body *RequestProblemJSONRequestBody
}

type RequestProblemResponseObject interface {
	VisitRequestProblemResponse(w http.ResponseWriter) error
}

type RequestProblem202JSONResponse struct {
	// ProblemId Developic에서 출제한 문제의 고유 ID입니다.
	ProblemId ProblemID `json:"problem_id"`
}

func (response RequestProblem202JSONResponse) VisitRequestProblemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)

	return json.NewEncoder(w).Encode(response)
}

type RequestProblemdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response RequestProblemdefaultJSONResponse) VisitRequestProblemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetProblemRequestObject struct {
	Id ProblemID `json:"id"`
}

type GetProblemResponseObject interface {
	VisitGetProblemResponse(w http.ResponseWriter) error
}

type GetProblem200JSONResponse struct{ N200GetProblemJSONResponse }

func (response GetProblem200JSONResponse) VisitGetProblemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetProblem404Response struct {
}

func (response GetProblem404Response) VisitGetProblemResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetProblem409Response struct {
}

func (response GetProblem409Response) VisitGetProblemResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type GetProblemdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetProblemdefaultJSONResponse) VisitGetProblemResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetRecordsRequestObject struct {
	Params GetRecordsParams
}

type GetRecordsResponseObject interface {
	VisitGetRecordsResponse(w http.ResponseWriter) error
}

type GetRecords200JSONResponse struct{ N200GetRecordsJSONResponse }

func (response GetRecords200JSONResponse) VisitGetRecordsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetRecordsdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetRecordsdefaultJSONResponse) VisitGetRecordsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type GetRecordRequestObject struct {
	Id RecordID `json:"id"`
}

type GetRecordResponseObject interface {
	VisitGetRecordResponse(w http.ResponseWriter) error
}

type GetRecord200JSONResponse struct{ N200GetRecordJSONResponse }

func (response GetRecord200JSONResponse) VisitGetRecordResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetRecord404Response struct {
}

func (response GetRecord404Response) VisitGetRecordResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetRecord409Response struct {
}

func (response GetRecord409Response) VisitGetRecordResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type GetRecorddefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetRecorddefaultJSONResponse) VisitGetRecordResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type SubmitSolutionRequestObject struct {
	Body *SubmitSolutionJSONRequestBody
}

type SubmitSolutionResponseObject interface {
	VisitSubmitSolutionResponse(w http.ResponseWriter) error
}

type SubmitSolution202JSONResponse struct{ N202SubmitJSONResponse }

func (response SubmitSolution202JSONResponse) VisitSubmitSolutionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)

	return json.NewEncoder(w).Encode(response)
}

type SubmitSolutiondefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response SubmitSolutiondefaultJSONResponse) VisitSubmitSolutionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /me)
	GetMe(ctx context.Context, request GetMeRequestObject) (GetMeResponseObject, error)

	// (POST /problems)
	RequestProblem(ctx context.Context, request RequestProblemRequestObject) (RequestProblemResponseObject, error)

	// (GET /problems/{id})
	GetProblem(ctx context.Context, request GetProblemRequestObject) (GetProblemResponseObject, error)

	// (GET /records)
	GetRecords(ctx context.Context, request GetRecordsRequestObject) (GetRecordsResponseObject, error)

	// (GET /records/{id})
	GetRecord(ctx context.Context, request GetRecordRequestObject) (GetRecordResponseObject, error)

	// (POST /submit)
	SubmitSolution(ctx context.Context, request SubmitSolutionRequestObject) (SubmitSolutionResponseObject, error)
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

// GetMe operation middleware
func (sh *strictHandler) GetMe(ctx echo.Context) error {
	var request GetMeRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetMe(ctx.Request().Context(), request.(GetMeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMe")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetMeResponseObject); ok {
		return validResponse.VisitGetMeResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// RequestProblem operation middleware
func (sh *strictHandler) RequestProblem(ctx echo.Context) error {
	var request RequestProblemRequestObject

	var body RequestProblemJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.RequestProblem(ctx.Request().Context(), request.(RequestProblemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "RequestProblem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(RequestProblemResponseObject); ok {
		return validResponse.VisitRequestProblemResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// GetProblem operation middleware
func (sh *strictHandler) GetProblem(ctx echo.Context, id ProblemID) error {
	var request GetProblemRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetProblem(ctx.Request().Context(), request.(GetProblemRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetProblem")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetProblemResponseObject); ok {
		return validResponse.VisitGetProblemResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// GetRecords operation middleware
func (sh *strictHandler) GetRecords(ctx echo.Context, params GetRecordsParams) error {
	var request GetRecordsRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetRecords(ctx.Request().Context(), request.(GetRecordsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetRecords")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetRecordsResponseObject); ok {
		return validResponse.VisitGetRecordsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// GetRecord operation middleware
func (sh *strictHandler) GetRecord(ctx echo.Context, id RecordID) error {
	var request GetRecordRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetRecord(ctx.Request().Context(), request.(GetRecordRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetRecord")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetRecordResponseObject); ok {
		return validResponse.VisitGetRecordResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("Unexpected response type: %T", response)
	}
	return nil
}

// SubmitSolution operation middleware
func (sh *strictHandler) SubmitSolution(ctx echo.Context) error {
	var request SubmitSolutionRequestObject

	var body SubmitSolutionJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.SubmitSolution(ctx.Request().Context(), request.(SubmitSolutionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SubmitSolution")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(SubmitSolutionResponseObject); ok {
		return validResponse.VisitSubmitSolutionResponse(ctx.Response())
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
