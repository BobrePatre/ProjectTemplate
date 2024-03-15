import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { Examplev1ExampleRequest } from '../models/Examplev1ExampleRequest';
import { Examplev1ExampleResponse } from '../models/Examplev1ExampleResponse';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';

import { ObservableExampleServiceApi } from "./ObservableAPI";
import { ExampleServiceApiRequestFactory, ExampleServiceApiResponseProcessor} from "../apis/ExampleServiceApi";

export interface ExampleServiceApiExampleServiceExampleRequest {
    /**
     * 
     * @type Examplev1ExampleRequest
     * @memberof ExampleServiceApiexampleServiceExample
     */
    body: Examplev1ExampleRequest
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
    public exampleServiceExampleWithHttpInfo(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<HttpInfo<Examplev1ExampleResponse>> {
        return this.api.exampleServiceExampleWithHttpInfo(param.body,  options).toPromise();
    }

    /**
     * Create an example
     * @param param the request object
     */
    public exampleServiceExample(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<Examplev1ExampleResponse> {
        return this.api.exampleServiceExample(param.body,  options).toPromise();
    }

}
