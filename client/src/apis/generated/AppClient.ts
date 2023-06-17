/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BaseHttpRequest } from './core/BaseHttpRequest';
import type { OpenAPIConfig } from './core/OpenAPI';
import { FetchHttpRequest } from './core/FetchHttpRequest';

import { BotService } from './services/BotService';
import { ListService } from './services/ListService';
import { MeService } from './services/MeService';
import { SimilarService } from './services/SimilarService';
import { TrendService } from './services/TrendService';
import { WordsService } from './services/WordsService';

type HttpRequestConstructor = new (config: OpenAPIConfig) => BaseHttpRequest;

export class AppClient {

    public readonly bot: BotService;
    public readonly list: ListService;
    public readonly me: MeService;
    public readonly similar: SimilarService;
    public readonly trend: TrendService;
    public readonly words: WordsService;

    public readonly request: BaseHttpRequest;

    constructor(config?: Partial<OpenAPIConfig>, HttpRequest: HttpRequestConstructor = FetchHttpRequest) {
        this.request = new HttpRequest({
            BASE: config?.BASE ?? '/api',
            VERSION: config?.VERSION ?? '1.0.0',
            WITH_CREDENTIALS: config?.WITH_CREDENTIALS ?? false,
            CREDENTIALS: config?.CREDENTIALS ?? 'include',
            TOKEN: config?.TOKEN,
            USERNAME: config?.USERNAME,
            PASSWORD: config?.PASSWORD,
            HEADERS: config?.HEADERS,
            ENCODE_PATH: config?.ENCODE_PATH,
        });

        this.bot = new BotService(this.request);
        this.list = new ListService(this.request);
        this.me = new MeService(this.request);
        this.similar = new SimilarService(this.request);
        this.trend = new TrendService(this.request);
        this.words = new WordsService(this.request);
    }
}

