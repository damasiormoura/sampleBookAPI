/*
 * Simple Inventory API
 *
 * Sample Books API
 *
 * API version: 1.0.0
 * Contact: damasiormoura@gmail.com.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Book struct {

	Isbn string `json:"isbn"`

	Title string `json:"title"`

	Author string `json:"author"`

	Edition string `json:"edition"`
}
