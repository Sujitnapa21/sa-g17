/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Patient
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntCustomerEdges,
    EntCustomerEdgesFromJSON,
    EntCustomerEdgesFromJSONTyped,
    EntCustomerEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCustomer
 */
export interface EntCustomer {
    /**
     * Email holds the value of the "Email" field.
     * @type {string}
     * @memberof EntCustomer
     */
    email?: string;
    /**
     * Username holds the value of the "Username" field.
     * @type {string}
     * @memberof EntCustomer
     */
    username?: string;
    /**
     * 
     * @type {EntCustomerEdges}
     * @memberof EntCustomer
     */
    edges?: EntCustomerEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntCustomer
     */
    id?: number;
}

export function EntCustomerFromJSON(json: any): EntCustomer {
    return EntCustomerFromJSONTyped(json, false);
}

export function EntCustomerFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCustomer {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': !exists(json, 'Email') ? undefined : json['Email'],
        'username': !exists(json, 'Username') ? undefined : json['Username'],
        'edges': !exists(json, 'edges') ? undefined : EntCustomerEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntCustomerToJSON(value?: EntCustomer | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Email': value.email,
        'Username': value.username,
        'edges': EntCustomerEdgesToJSON(value.edges),
        'id': value.id,
    };
}

