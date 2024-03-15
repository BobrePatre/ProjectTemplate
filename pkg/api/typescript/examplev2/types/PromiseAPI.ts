import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'

import { Examplev2ExampleRequest } from '../models/Examplev2ExampleRequest';
import { Examplev2ExampleResponse } from '../models/Examplev2ExampleResponse';
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
     * @param body 
     */
    public exampleServiceExampleWithHttpInfo(body: Examplev2ExampleRequest, _options?: Configuration): Promise<HttpInfo<Examplev2ExampleResponse>> {
        const result = this.api.exampleServiceExampleWithHttpInfo(body, _options);
        return result.toPromise();
    }

    /**
     * @param body 
     */
    public exampleServiceExample(body: Examplev2ExampleRequest, _options?: Configuration): Promise<Examplev2ExampleResponse> {
        const result = this.api.exampleServiceExample(body, _options);
        return result.toPromise();
    }


}



