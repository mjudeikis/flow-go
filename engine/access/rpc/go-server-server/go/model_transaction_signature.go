/*
 * flow/access/access.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TransactionSignature struct {

	Address string `json:"address,omitempty"`

	KeyId int64 `json:"keyId,omitempty"`

	Signature string `json:"signature,omitempty"`
}
