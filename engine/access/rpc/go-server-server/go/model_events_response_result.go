/*
 * flow/access/access.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type EventsResponseResult struct {

	BlockId string `json:"blockId,omitempty"`

	BlockHeight string `json:"blockHeight,omitempty"`

	Events []EntitiesEvent `json:"events,omitempty"`

	BlockTimestamp time.Time `json:"blockTimestamp,omitempty"`
}
