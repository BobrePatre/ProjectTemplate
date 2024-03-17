import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { ApilalaExampleRequest } from '../models/ApilalaExampleRequest';
import { ApilalaExampleResponse } from '../models/ApilalaExampleResponse';
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
    public exampleServiceExampleWithHttpInfo(body: ApilalaExampleRequest, _options?: Configuration): Promise<HttpInfo<ApilalaExampleResponse>> {
        const result = this.api.exampleServiceExampleWithHttpInfo(body, _options);
        return result.toPromise();
    }

    /**
     * Create an example
     * @param body 
     */
    public exampleServiceExample(body: ApilalaExampleRequest, _options?: Configuration): Promise<ApilalaExampleResponse> {
        const result = this.api.exampleServiceExample(body, _options);
        return result.toPromise();
    }


}



