/* 
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package logicmonitor

type RestPropertyPagination struct {

	Total int32 `json:"total,omitempty"`

	SearchId string `json:"searchId,omitempty"`

	Items []RestProperty `json:"items,omitempty"`
}