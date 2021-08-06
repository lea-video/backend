/*
 * LEAV
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package leav

type Track struct {
	Id string `json:"id,omitempty"`

	Filename string `json:"filename,omitempty"`

	Filesize int32 `json:"filesize,omitempty"`

	Language string `json:"language,omitempty"`

	AudioBitrate int32 `json:"audioBitrate,omitempty"`

	AudioEncoding string `json:"audioEncoding,omitempty"`

	VideoBitrate int32 `json:"videoBitrate,omitempty"`

	VideoEncoding string `json:"videoEncoding,omitempty"`
}
