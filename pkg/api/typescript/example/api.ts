/* tslint:disable */
/* eslint-disable */
/**
 * Example API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: v1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError, operationServerMap } from './base';

/**
 * 
 * @export
 * @interface ExampleExampleRequest
 */
export interface ExampleExampleRequest {
    /**
     * 
     * @type {string}
     * @memberof ExampleExampleRequest
     */
    'title'?: string;
    /**
     * 
     * @type {string}
     * @memberof ExampleExampleRequest
     */
    'description'?: string;
    /**
     * 
     * @type {string}
     * @memberof ExampleExampleRequest
     */
    'body'?: string;
}
/**
 * 
 * @export
 * @interface ExampleExampleResponse
 */
export interface ExampleExampleResponse {
    /**
     * 
     * @type {string}
     * @memberof ExampleExampleResponse
     */
    'title'?: string;
    /**
     * 
     * @type {string}
     * @memberof ExampleExampleResponse
     */
    'description'?: string;
    /**
     * 
     * @type {string}
     * @memberof ExampleExampleResponse
     */
    'body'?: string;
}
/**
 * 
 * @export
 * @interface ProtobufAny
 */
export interface ProtobufAny {
    [key: string]: object | any;

    /**
     * 
     * @type {string}
     * @memberof ProtobufAny
     */
    '@type'?: string;
}
/**
 * 
 * @export
 * @interface RpcStatus
 */
export interface RpcStatus {
    /**
     * 
     * @type {number}
     * @memberof RpcStatus
     */
    'code'?: number;
    /**
     * 
     * @type {string}
     * @memberof RpcStatus
     */
    'message'?: string;
    /**
     * 
     * @type {Array<ProtobufAny>}
     * @memberof RpcStatus
     */
    'details'?: Array<ProtobufAny>;
}

/**
 * ExampleServiceApi - axios parameter creator
 * @export
 */
export const ExampleServiceApiAxiosParamCreator = function (configuration?: Configuration) {
    // @ts-ignore
    return {
        /**
         * 
         * @summary Create an example
         * @param {ExampleExampleRequest} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        exampleServiceExample: async (body: ExampleExampleRequest, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            assertParamExists('exampleServiceExample', 'body', body)
            const localVarPath = `/grpc/v1/example/idi-na-xui`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication BearerAuth required
            // oauth required
            await setOAuthToObject(localVarHeaderParameter, "BearerAuth", [], configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(body, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ExampleServiceApi - functional programming interface
 * @export
 */
export const ExampleServiceApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = ExampleServiceApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Create an example
         * @param {ExampleExampleRequest} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async exampleServiceExample(body: ExampleExampleRequest, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ExampleExampleResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.exampleServiceExample(body, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['ExampleServiceApi.exampleServiceExample']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * ExampleServiceApi - factory interface
 * @export
 */
export const ExampleServiceApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = ExampleServiceApiFp(configuration)
    return {
        /**
         * 
         * @summary Create an example
         * @param {ExampleExampleRequest} body 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        exampleServiceExample(body: ExampleExampleRequest, options?: any): AxiosPromise<ExampleExampleResponse> {
            return localVarFp.exampleServiceExample(body, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * ExampleServiceApi - object-oriented interface
 * @export
 * @class ExampleServiceApi
 * @extends {BaseAPI}
 */
export class ExampleServiceApi extends BaseAPI {
    /**
     * 
     * @summary Create an example
     * @param {ExampleExampleRequest} body 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ExampleServiceApi
     */
    public exampleServiceExample(body: ExampleExampleRequest, options?: RawAxiosRequestConfig) {
        return ExampleServiceApiFp(this.configuration).exampleServiceExample(body, options).then((request) => request(this.axios, this.basePath));
    }
}



