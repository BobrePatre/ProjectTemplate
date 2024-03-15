import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'
import { Observable, of, from } from '../rxjsStub';
import {mergeMap, map} from  '../rxjsStub';
import { Examplev2ExampleRequest } from '../models/Examplev2ExampleRequest';
import { Examplev2ExampleResponse } from '../models/Examplev2ExampleResponse';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';

import { ExampleServiceApiRequestFactory, ExampleServiceApiResponseProcessor} from "../apis/ExampleServiceApi";
export class ObservableExampleServiceApi {
    private requestFactory: ExampleServiceApiRequestFactory;
    private responseProcessor: ExampleServiceApiResponseProcessor;
    private configuration: Configuration;

    public constructor(
        configuration: Configuration,
        requestFactory?: ExampleServiceApiRequestFactory,
        responseProcessor?: ExampleServiceApiResponseProcessor
    ) {
        this.configuration = configuration;
        this.requestFactory = requestFactory || new ExampleServiceApiRequestFactory(configuration);
        this.responseProcessor = responseProcessor || new ExampleServiceApiResponseProcessor();
    }

    /**
     * @param body 
     */
    public exampleServiceExampleWithHttpInfo(body: Examplev2ExampleRequest, _options?: Configuration): Observable<HttpInfo<Examplev2ExampleResponse>> {
        const requestContextPromise = this.requestFactory.exampleServiceExample(body, _options);

        // build promise chain
        let middlewarePreObservable = from<RequestContext>(requestContextPromise);
        for (let middleware of this.configuration.middleware) {
            middlewarePreObservable = middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => middleware.pre(ctx)));
        }

        return middlewarePreObservable.pipe(mergeMap((ctx: RequestContext) => this.configuration.httpApi.send(ctx))).
            pipe(mergeMap((response: ResponseContext) => {
                let middlewarePostObservable = of(response);
                for (let middleware of this.configuration.middleware) {
                    middlewarePostObservable = middlewarePostObservable.pipe(mergeMap((rsp: ResponseContext) => middleware.post(rsp)));
                }
                return middlewarePostObservable.pipe(map((rsp: ResponseContext) => this.responseProcessor.exampleServiceExampleWithHttpInfo(rsp)));
            }));
    }

    /**
     * @param body 
     */
    public exampleServiceExample(body: Examplev2ExampleRequest, _options?: Configuration): Observable<Examplev2ExampleResponse> {
        return this.exampleServiceExampleWithHttpInfo(body, _options).pipe(map((apiResponse: HttpInfo<Examplev2ExampleResponse>) => apiResponse.data));
    }

}
