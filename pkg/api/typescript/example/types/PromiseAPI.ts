import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { V1ExampleRequest } from '../models/V1ExampleRequest';
import { V1ExampleResponse } from '../models/V1ExampleResponse';
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
    public exampleServiceExampleWithHttpInfo(body: V1ExampleRequest, _options?: Configuration): Promise<HttpInfo<V1ExampleResponse>> {
        const result = this.api.exampleServiceExampleWithHttpInfo(body, _options);
        return result.toPromise();
    }

    /**
     * Create an example
     * @param body 
     */
    public exampleServiceExample(body: V1ExampleRequest, _options?: Configuration): Promise<V1ExampleResponse> {
        const result = this.api.exampleServiceExample(body, _options);
        return result.toPromise();
    }


}



