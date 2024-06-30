/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { WordDelete } from '../models/WordDelete';
import type { WordRequest } from '../models/WordRequest';
import type { WordsAllList } from '../models/WordsAllList';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class WordsService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * wordの登録
     * wordの登録
     * @param requestBody
     * @returns any Successful registration
     * @throws ApiError
     */
    public postWords(
        requestBody: WordRequest,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/words',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid input`,
            },
        });
    }
    /**
     * wordの削除
     * wordの削除
     * @param requestBody
     * @returns any Successful deletion
     * @throws ApiError
     */
    public deleteWords(
        requestBody: WordDelete,
    ): CancelablePromise<any> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/words',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Invalid input`,
            },
        });
    }
    /**
     * 全データの取得
     * 全データの取得
     * @returns WordsAllList Successful retrieval
     * @throws ApiError
     */
    public getWords(): CancelablePromise<WordsAllList> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/words',
        });
    }
}
