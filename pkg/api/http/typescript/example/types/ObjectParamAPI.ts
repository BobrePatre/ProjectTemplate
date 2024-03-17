import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { ApiexampleExampleRequest } from '../models/ApiexampleExampleRequest';
import { ApiexampleExampleResponse } from '../models/ApiexampleExampleResponse';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';

import { ObservableExampleServiceApi } from "./ObservableAPI";
import { ExampleServiceApiRequestFactory, ExampleServiceApiResponseProcessor} from "../apis/ExampleServiceApi";

export interface ExampleServiceApiExampleServiceExampleRequest {
    /**
     * 
     * @type ApiexampleExampleRequest
     * @memberof ExampleServiceApiexampleServiceExample
     */
    body: ApiexampleExampleRequest
}

export class ObjectExampleServiceApi {
    private api: ObservableExampleServiceApi

    public constructor(configuration: Configuration, requestFactory?: ExampleServiceApiRequestFactory, responseProcessor?: ExampleServiceApiResponseProcessor) {
        this.api = new ObservableExampleServiceApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * Create an example
     * @param param the request object
     */
    public exampleServiceExampleWithHttpInfo(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<HttpInfo<ApiexampleExampleResponse>> {
        return this.api.exampleServiceExampleWithHttpInfo(param.body,  options).toPromise();
    }

    /**
     * Create an example
     * @param param the request object
     */
    public exampleServiceExample(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<ApiexampleExampleResponse> {
        return this.api.exampleServiceExample(param.body,  options).toPromise();
    }

}
