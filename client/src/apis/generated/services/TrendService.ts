/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { TrendingWords } from '../models/TrendingWords';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class TrendService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * 今日のトレンド
     * 今日最も追加されたwordの取得
     * @param limit 返すwordの数
     * @returns TrendingWords OK
     * @throws ApiError
     */
    public getTodayTrendingWords(
        limit: number = 10,
    ): CancelablePromise<TrendingWords> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/trend/day/today',
            query: {
                'limit': limit,
            },
        });
    }

    /**
     * ある日のトレンド
     * 特定の日に最も追加されたwordの取得
     * @param day Specific day in the format "YYYY-MM-DD"
     * @param limit 返すwordの数
     * @returns TrendingWords OK
     * @throws ApiError
     */
    public getTrendingWordsForDay(
        day: string,
        limit: number = 10,
    ): CancelablePromise<TrendingWords> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/trend/day/{day}',
            path: {
                'day': day,
            },
            query: {
                'limit': limit,
            },
        });
    }

    /**
     * ある月のトレンド
     * 特定の月に最も追加されたwordの取得
     * @param month Specific month in the format "YYYY-MM"
     * @param limit 返すwordの数
     * @returns TrendingWords OK
     * @throws ApiError
     */
    public getTrendingWordsForMonth(
        month: string,
        limit: number = 10,
    ): CancelablePromise<TrendingWords> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/trend/month/{month}',
            path: {
                'month': month,
            },
            query: {
                'limit': limit,
            },
        });
    }

    /**
     * ある年のトレンド
     * 特定の年に最も追加されたwordの取得
     * @param year Specific year in the format "YYYY"
     * @param limit 返すwordの数
     * @returns TrendingWords OK
     * @throws ApiError
     */
    public getTrendingWordsForYear(
        year: string,
        limit: number = 10,
    ): CancelablePromise<TrendingWords> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/trend/year/{year}',
            path: {
                'year': year,
            },
            query: {
                'limit': limit,
            },
        });
    }

}
