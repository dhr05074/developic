/* tslint:disable */
/* eslint-disable */
/**
 * Developic API
 * Developic 서비스에서 사용하는 REST API 입니다.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface GetMe200Response
 */
export interface GetMe200Response {
    /**
     * Developic에서 사용되는 유저의 이름입니다.
     * @type {string}
     * @memberof GetMe200Response
     */
    'nickname': string;
    /**
     * Developic에서 유저의 실력을 가늠하는 ELO 점수입니다.
     * @type {number}
     * @memberof GetMe200Response
     */
    'elo_score': number;
}
/**
 * 
 * @export
 * @interface GetRecords200Response
 */
export interface GetRecords200Response {
    /**
     * 
     * @type {Array<Record>}
     * @memberof GetRecords200Response
     */
    'records': Array<Record>;
}
/**
 * 
 * @export
 * @interface ModelError
 */
export interface ModelError {
    /**
     * 오류 구분 코드입니다.
     * @type {string}
     * @memberof ModelError
     */
    'code': string;
    /**
     * 상세한 오류 정보를 알려주는 메시지입니다.
     * @type {string}
     * @memberof ModelError
     */
    'message': string;
}
/**
 * 
 * @export
 * @interface Problem
 */
export interface Problem {
    /**
     * Developic에서 출제한 문제의 고유 ID입니다.
     * @type {string}
     * @memberof Problem
     */
    'id': string;
    /**
     * Developic에서 출제한 문제의 타이틀입니다.
     * @type {string}
     * @memberof Problem
     */
    'title': string;
    /**
     * 문제의 자세한 설명입니다.
     * @type {string}
     * @memberof Problem
     */
    'description': string;
    /**
     * Developic에서 사용되는 코드 데이터입니다. 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다. 
     * @type {string}
     * @memberof Problem
     */
    'code': string;
}
/**
 * Developic에서 사용할 프로그래밍 언어입니다.
 * @export
 * @enum {string}
 */

export const ProgrammingLanguage = {
    Go: 'Go',
    Javascript: 'Javascript',
    Cpp: 'Cpp'
} as const;

export type ProgrammingLanguage = typeof ProgrammingLanguage[keyof typeof ProgrammingLanguage];


/**
 * 
 * @export
 * @interface Record
 */
export interface Record {
    /**
     * Developic에서 생성된 결과 보고서의 고유 ID입니다.
     * @type {string}
     * @memberof Record
     */
    'id': string;
    /**
     * Developic에서 출제한 문제의 고유 ID입니다.
     * @type {string}
     * @memberof Record
     */
    'problem_id': string;
    /**
     * Developic에서 출제한 문제의 타이틀입니다.
     * @type {string}
     * @memberof Record
     */
    'problem_title': string;
    /**
     * 결과 보고서에서 사용자가 취득한 총점입니다.
     * @type {number}
     * @memberof Record
     */
    'efficiency': number;
    /**
     * 결과 보고서에서 사용자가 취득한 총점입니다.
     * @type {number}
     * @memberof Record
     */
    'readability': number;
    /**
     * 결과 보고서에서 사용자가 취득한 총점입니다.
     * @type {number}
     * @memberof Record
     */
    'robustness': number;
    /**
     * Developic에서 사용되는 코드 데이터입니다. 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다. 
     * @type {string}
     * @memberof Record
     */
    'code': string;
}
/**
 * 
 * @export
 * @interface RequestProblem202Response
 */
export interface RequestProblem202Response {
    /**
     * Developic에서 출제한 문제의 고유 ID입니다.
     * @type {string}
     * @memberof RequestProblem202Response
     */
    'problem_id': string;
}
/**
 * 
 * @export
 * @interface RequestProblemRequest
 */
export interface RequestProblemRequest {
    /**
     * 
     * @type {ProgrammingLanguage}
     * @memberof RequestProblemRequest
     */
    'language': ProgrammingLanguage;
    /**
     * Developic에서 유저의 실력을 가늠하는 ELO 점수입니다.
     * @type {number}
     * @memberof RequestProblemRequest
     */
    'elo_score'?: number;
}


/**
 * 
 * @export
 * @interface SubmitSolution202Response
 */
export interface SubmitSolution202Response {
    /**
     * Developic에서 생성된 결과 보고서의 고유 ID입니다.
     * @type {string}
     * @memberof SubmitSolution202Response
     */
    'record_id': string;
}
/**
 * 
 * @export
 * @interface SubmitSolutionRequest
 */
export interface SubmitSolutionRequest {
    /**
     * Developic에서 출제한 문제의 고유 ID입니다.
     * @type {string}
     * @memberof SubmitSolutionRequest
     */
    'problem_id': string;
    /**
     * Developic에서 사용되는 코드 데이터입니다. 코드 데이터의 Escape를 방지하기 위해 Base64로 인코딩되어 전송, 보관됩니다. 
     * @type {string}
     * @memberof SubmitSolutionRequest
     */
    'code': string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getMe: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/me`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {string} id 문제 생성 시 발급받은 ID입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getProblem: async (id: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getProblem', 'id', id)
            const localVarPath = `/problems/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 입력한 ID를 가진 채점 보고서를 조회합니다.
         * @param {string} id 조회하고자 하는 채점 보고서의 ID입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRecord: async (id: string, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getRecord', 'id', id)
            const localVarPath = `/records/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {number} [page] 조회할 페이지입니다.
         * @param {number} [limit] 한 페이지당 조회할 아이템의 수입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRecords: async (page?: number, limit?: number, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/records`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)

            if (page !== undefined) {
                localVarQueryParameter['page'] = page;
            }

            if (limit !== undefined) {
                localVarQueryParameter['limit'] = limit;
            }


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {RequestProblemRequest} requestProblemRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestProblem: async (requestProblemRequest: RequestProblemRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'requestProblemRequest' is not null or undefined
            assertParamExists('requestProblem', 'requestProblemRequest', requestProblemRequest)
            const localVarPath = `/problems`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(requestProblemRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @param {SubmitSolutionRequest} [submitSolutionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        submitSolution: async (submitSolutionRequest?: SubmitSolutionRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/submit`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication apiKey required
            await setApiKeyToObject(localVarHeaderParameter, "Authorization", configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(submitSolutionRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getMe(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetMe200Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getMe(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {string} id 문제 생성 시 발급받은 ID입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getProblem(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Problem>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getProblem(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 입력한 ID를 가진 채점 보고서를 조회합니다.
         * @param {string} id 조회하고자 하는 채점 보고서의 ID입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getRecord(id: string, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Record>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getRecord(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {number} [page] 조회할 페이지입니다.
         * @param {number} [limit] 한 페이지당 조회할 아이템의 수입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getRecords(page?: number, limit?: number, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetRecords200Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getRecords(page, limit, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {RequestProblemRequest} requestProblemRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async requestProblem(requestProblemRequest: RequestProblemRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RequestProblem202Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.requestProblem(requestProblemRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @param {SubmitSolutionRequest} [submitSolutionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async submitSolution(submitSolutionRequest?: SubmitSolutionRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SubmitSolution202Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.submitSolution(submitSolutionRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getMe(options?: any): AxiosPromise<GetMe200Response> {
            return localVarFp.getMe(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {string} id 문제 생성 시 발급받은 ID입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getProblem(id: string, options?: any): AxiosPromise<Problem> {
            return localVarFp.getProblem(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 입력한 ID를 가진 채점 보고서를 조회합니다.
         * @param {string} id 조회하고자 하는 채점 보고서의 ID입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRecord(id: string, options?: any): AxiosPromise<Record> {
            return localVarFp.getRecord(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {number} [page] 조회할 페이지입니다.
         * @param {number} [limit] 한 페이지당 조회할 아이템의 수입니다.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getRecords(page?: number, limit?: number, options?: any): AxiosPromise<GetRecords200Response> {
            return localVarFp.getRecords(page, limit, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {RequestProblemRequest} requestProblemRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestProblem(requestProblemRequest: RequestProblemRequest, options?: any): AxiosPromise<RequestProblem202Response> {
            return localVarFp.requestProblem(requestProblemRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @param {SubmitSolutionRequest} [submitSolutionRequest] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        submitSolution(submitSolutionRequest?: SubmitSolutionRequest, options?: any): AxiosPromise<SubmitSolution202Response> {
            return localVarFp.submitSolution(submitSolutionRequest, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getMe(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getMe(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {string} id 문제 생성 시 발급받은 ID입니다.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getProblem(id: string, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getProblem(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 입력한 ID를 가진 채점 보고서를 조회합니다.
     * @param {string} id 조회하고자 하는 채점 보고서의 ID입니다.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getRecord(id: string, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getRecord(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {number} [page] 조회할 페이지입니다.
     * @param {number} [limit] 한 페이지당 조회할 아이템의 수입니다.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getRecords(page?: number, limit?: number, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getRecords(page, limit, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {RequestProblemRequest} requestProblemRequest 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public requestProblem(requestProblemRequest: RequestProblemRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).requestProblem(requestProblemRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @param {SubmitSolutionRequest} [submitSolutionRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public submitSolution(submitSolutionRequest?: SubmitSolutionRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).submitSolution(submitSolutionRequest, options).then((request) => request(this.axios, this.basePath));
    }
}

