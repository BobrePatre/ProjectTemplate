import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { ApiexampleExampleRequest } from '../models/ApiexampleExampleRequest';
import { ApiexampleExampleResponse } from '../models/ApiexampleExampleResponse';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { ObservableExampleServiceApi } from './ObservableAPI';

import { ExampleServiceApiRequestFactory, ExampleServiceApiResponseProcessor} from "../apis/ExampleServiceApi";
export class PromiseExampleServiceApi {
    private api: ObservableExampleServiceApi

    public constructor(
        configuration: Configuration,
        requestFactory?: ExampleServiceApiRequestFactory,
        responseProcessor?: ExampleServiceApiResponseProcessor
    ) {
        this.api = new ObservableExampleServiceApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * Create an example
     * @param body 
     */
    public exampleServiceExampleWithHttpInfo(body: ApiexampleExampleRequest, _options?: Configuration): Promise<HttpInfo<ApiexampleExampleResponse>> {
        const result = this.api.exampleServiceExampleWithHttpInfo(body, _options);
        return result.toPromise();
    }

    /**
     * Create an example
     * @param body 
     */
    public exampleServiceExample(body: ApiexampleExampleRequest, _options?: Configuration): Promise<ApiexampleExampleResponse> {
        const result = this.api.exampleServiceExample(body, _options);
        return result.toPromise();
    }


}



