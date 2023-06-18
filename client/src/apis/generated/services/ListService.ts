/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { UsersList } from '../models/UsersList';
import type { UsersOfWordsList } from '../models/UsersOfWordsList';
import type { WordsList } from '../models/WordsList';
import type { WordsOfUsersList } from '../models/WordsOfUsersList';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class ListService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * アクセスしているuserのwordたち
     * アクセスしているuserの登録しているwordの取得
     * @returns WordsList Successful retrieval
     * @throws ApiError
     */
    public getListUserMe(): CancelablePromise<WordsList> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/list/user/me',
        });
    }

    /**
     * あるuserのwordたち
     * userの登録しているwordの取得
     * @param userId ID of the user
     * @returns WordsList Successful retrieval
     * @throws ApiError
     */
    public getListUser(
        userId: string,
    ): CancelablePromise<WordsList> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/list/user/{userId}',
            path: {
                'userId': userId,
            },
            errors: {
                404: `User not found`,
            },
        });
    }

    /**
     * ある単語を見ているuserたち
     * ある単語を登録しているuserの取得
     * @param word The word to search for
     * @returns UsersList Successful retrieval
     * @throws ApiError
     */
    public getListWord(
        word: string,
    ): CancelablePromise<UsersList> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/list/word/{word}',
            path: {
                'word': word,
            },
            errors: {
                404: `Word not found`,
            },
        });
    }

    /**
     * あるuserのwordたちを登録しているuserたち
     * あるuserが登録しているすべてのwordのそれぞれを登録しているusersの取得
     * @param userId ID of the user
     * @returns UsersOfWordsList Successful retrieval
     * @throws ApiError
     */
    public getListUserUsers(
        userId: string,
    ): CancelablePromise<UsersOfWordsList> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/list/user/{userId}/users',
            path: {
                'userId': userId,
            },
            errors: {
                404: `User not found`,
            },
        });
    }

    /**
     * あるwordのuserたちが登録しているwordたち
     * あるwordを登録しているすべてのuserのそれぞれが登録しているwordsの取得
     * @param word The word to search for
     * @returns WordsOfUsersList Successful retrieval
     * @throws ApiError
     */
    public getListWordWords(
        word: string,
    ): CancelablePromise<WordsOfUsersList> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/list/word/{word}/words',
            path: {
                'word': word,
            },
            errors: {
                404: `Word not found`,
            },
        });
    }

}
