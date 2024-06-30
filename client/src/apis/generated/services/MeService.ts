/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Me } from '../models/Me';
import type { WordMeSetting } from '../models/WordMeSetting';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class MeService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * 自分の投稿に対する通知の設定
     * 各wordに対して自分が投稿したものについても通知するか決める
     * @param requestBody
     * @returns any Successful deletion
     * @throws ApiError
     */
    public putWordsMe(
        requestBody: WordMeSetting,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'PUT',
            url: '/words/me/',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid input`,
            },
        });
    }
    /**
     * 自分の投稿に対する通知の一括設定
     * 自分が投稿したもの全てについて通知するか決める
     * @param requestBody
     * @returns any Successful deletion
     * @throws ApiError
     */
    public postWordsMeAll(
        requestBody: Me,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/words/me/all',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                404: `User not found`,
            },
        });
    }
}
