import { ResponseContext, RequestContext, HttpFile, HttpInfo } from '../http/http';
import { Configuration} from '../configuration'
import { Observable, of, from } from '../rxjsStub';
import {mergeMap, map} from  '../rxjsStub';
import { ProtobufAny } from '../models/ProtobufAny';
import { RpcStatus } from '../models/RpcStatus';
import { V1ExampleRequest } from '../models/V1ExampleRequest';
import { V1ExampleResponse } from '../models/V1ExampleResponse';

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
     * Create an example
     * @param body 
     */
    public exampleServiceExampleWithHttpInfo(body: V1ExampleRequest, _options?: Configuration): Observable<HttpInfo<V1ExampleResponse>> {
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
     * Create an example
     * @param body 
     */
    public exampleServiceExample(body: V1ExampleRequest, _options?: Configuration): Observable<V1ExampleResponse> {
        return this.exampleServiceExampleWithHttpInfo(body, _options).pipe(map((apiResponse: HttpInfo<V1ExampleResponse>) => apiResponse.data));
    }

}
