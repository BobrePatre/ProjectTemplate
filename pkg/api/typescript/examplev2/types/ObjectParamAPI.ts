import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { Examplev2ExampleRequest } from '../models/Examplev2ExampleRequest';
import { Examplev2ExampleResponse } from '../models/Examplev2ExampleResponse';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';

import { ObservableExampleServiceApi } from "./ObservableAPI";
import { ExampleServiceApiRequestFactory, ExampleServiceApiResponseProcessor} from "../apis/ExampleServiceApi";

export interface ExampleServiceApiExampleServiceExampleRequest {
    /**
     * 
     * @type Examplev2ExampleRequest
     * @memberof ExampleServiceApiexampleServiceExample
     */
    body: Examplev2ExampleRequest
}

export class ObjectExampleServiceApi {
    private api: ObservableExampleServiceApi

    public constructor(configuration: Configuration, requestFactory?: ExampleServiceApiRequestFactory, responseProcessor?: ExampleServiceApiResponseProcessor) {
        this.api = new ObservableExampleServiceApi(configuration, requestFactory, responseProcessor);
    }

    /**
     * @param param the request object
     */
    public exampleServiceExampleWithHttpInfo(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<HttpInfo<Examplev2ExampleResponse>> {
        return this.api.exampleServiceExampleWithHttpInfo(param.body,  options).toPromise();
    }

    /**
     * @param param the request object
     */
    public exampleServiceExample(param: ExampleServiceApiExampleServiceExampleRequest, options?: Configuration): Promise<Examplev2ExampleResponse> {
        return this.api.exampleServiceExample(param.body,  options).toPromise();
    }

}
