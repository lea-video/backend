/*
 * LEAV
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Season struct {

	Title *Title `json:"title,omitempty"`

	Series *Series `json:"series,omitempty"`

	Episode []Media `json:"episode,omitempty"`
}