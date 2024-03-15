import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { V1ExampleRequest } from '../models/V1ExampleRequest';
import { V1ExampleResponse } from '../models/V1ExampleResponse';

import { ObservableExampleServiceApi } from "./ObservableAPI";
import { ExampleServiceApiRequestFactory, ExampleServiceApiResponseProcessor} from "../apis/ExampleServiceApi";

export interface ExampleServiceApiExampleServiceExampleRequest {
    /**
     * 
     * @type V1ExampleRequest
     * @memberof ExampleServiceApiexampleServiceExample
     */
    body: V1ExampleRequest
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
    public exampleServiceExampleWithHttpInfo(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<HttpInfo<V1ExampleResponse>> {
        return this.api.exampleServiceExampleWithHttpInfo(param.body,  options).toPromise();
    }

    /**
     * Create an example
     * @param param the request object
     */
    public exampleServiceExample(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<V1ExampleResponse> {
        return this.api.exampleServiceExample(param.body,  options).toPromise();
    }

}
