/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { RecommendedWords } from '../models/RecommendedWords';
import type { SimilarUsers } from '../models/SimilarUsers';
import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';
export class SimilarService {
    constructor(public readonly httpRequest: BaseHttpRequest) {}
    /**
     * 似たような者を探す
     * 特定のuserと同じような単語を登録しているuserの取得
     * @param userId ID of the user
     * @returns SimilarUsers OK
     * @throws ApiError
     */
    public getUsersWithSimilarWords(
        userId: string,
    ): CancelablePromise<SimilarUsers> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/similar/{userId}',
            path: {
                'userId': userId,
            },
        });
    }
    /**
     * おすすめの単語を出す
     * 特定のuserと同じような単語を登録しているuserが登録している単語の取得
     * @param userId ID of the user
     * @returns RecommendedWords OK
     * @throws ApiError
     */
    public getRecommendedWordsForUser(
        userId: string,
    ): CancelablePromise<RecommendedWords> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/similar/{userId}/recommend',
            path: {
                'userId': userId,
            },
        });
    }
}
