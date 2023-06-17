/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Bot } from '../models/Bot';
import type { WordBotSetting } from '../models/WordBotSetting';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class BotService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * bot投稿に対する通知の設定
     * 各wordに対してbotが投稿したものについても通知するか決める
     * @param requestBody
     * @returns any Successful deletion
     * @throws ApiError
     */
    public putWords(
        requestBody: WordBotSetting,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/words',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid input`,
            },
        });
    }

    /**
     * bot投稿に対する通知の一括設定
     * botが投稿したもの全てについて通知するか決める
     * @param requestBody
     * @returns any Successful deletion
     * @throws ApiError
     */
    public postWordsBot(
        requestBody: Bot,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/words/bot',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                404: `User not found`,
            },
        });
    }

}
