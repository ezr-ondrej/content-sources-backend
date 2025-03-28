/* tslint:disable */
/* eslint-disable */
/**
 * ContentSourcesBackend
 * The API for the repositories of the content sources that you can use to create and manage repositories between third-party applications and the [Red Hat Hybrid Cloud Console](https://console.redhat.com). With these repositories, you can build and deploy images using Image Builder for Cloud, on-Premise, and Edge. You can handle tasks, search for required RPMs, fetch a GPGKey from the URL, and list the features within applications. 
 *
 * The version of the OpenAPI document: v1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { ApiPublicRepositoryResponse } from './ApiPublicRepositoryResponse';
import {
    ApiPublicRepositoryResponseFromJSON,
    ApiPublicRepositoryResponseFromJSONTyped,
    ApiPublicRepositoryResponseToJSON,
    ApiPublicRepositoryResponseToJSONTyped,
} from './ApiPublicRepositoryResponse';
import type { ApiResponseMetadata } from './ApiResponseMetadata';
import {
    ApiResponseMetadataFromJSON,
    ApiResponseMetadataFromJSONTyped,
    ApiResponseMetadataToJSON,
    ApiResponseMetadataToJSONTyped,
} from './ApiResponseMetadata';
import type { ApiLinks } from './ApiLinks';
import {
    ApiLinksFromJSON,
    ApiLinksFromJSONTyped,
    ApiLinksToJSON,
    ApiLinksToJSONTyped,
} from './ApiLinks';

/**
 * 
 * @export
 * @interface ApiPublicRepositoryCollectionResponse
 */
export interface ApiPublicRepositoryCollectionResponse {
    /**
     * 
     * @type {Array<ApiPublicRepositoryResponse>}
     * @memberof ApiPublicRepositoryCollectionResponse
     */
    data?: Array<ApiPublicRepositoryResponse>;
    /**
     * 
     * @type {ApiLinks}
     * @memberof ApiPublicRepositoryCollectionResponse
     */
    links?: ApiLinks;
    /**
     * 
     * @type {ApiResponseMetadata}
     * @memberof ApiPublicRepositoryCollectionResponse
     */
    meta?: ApiResponseMetadata;
}

/**
 * Check if a given object implements the ApiPublicRepositoryCollectionResponse interface.
 */
export function instanceOfApiPublicRepositoryCollectionResponse(value: object): value is ApiPublicRepositoryCollectionResponse {
    return true;
}

export function ApiPublicRepositoryCollectionResponseFromJSON(json: any): ApiPublicRepositoryCollectionResponse {
    return ApiPublicRepositoryCollectionResponseFromJSONTyped(json, false);
}

export function ApiPublicRepositoryCollectionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiPublicRepositoryCollectionResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'data': json['data'] == null ? undefined : ((json['data'] as Array<any>).map(ApiPublicRepositoryResponseFromJSON)),
        'links': json['links'] == null ? undefined : ApiLinksFromJSON(json['links']),
        'meta': json['meta'] == null ? undefined : ApiResponseMetadataFromJSON(json['meta']),
    };
}

export function ApiPublicRepositoryCollectionResponseToJSON(json: any): ApiPublicRepositoryCollectionResponse {
    return ApiPublicRepositoryCollectionResponseToJSONTyped(json, false);
}

export function ApiPublicRepositoryCollectionResponseToJSONTyped(value?: ApiPublicRepositoryCollectionResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'data': value['data'] == null ? undefined : ((value['data'] as Array<any>).map(ApiPublicRepositoryResponseToJSON)),
        'links': ApiLinksToJSON(value['links']),
        'meta': ApiResponseMetadataToJSON(value['meta']),
    };
}

