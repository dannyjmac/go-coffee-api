// Package classification Product API.
//
// Documentation for Product API
//
//     Schemes: http, https
//     BasePath: /
//     Version: 1.0.0
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
// swagger:meta
package handlers

import "github.com/dannyjmac/go-micro-3/data"

// Note - Types defined here are purely for documentation purposes

// Generic error message returned as a string
// swagger: response errorResponse
type ErrorResponseWrapper struct {
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger: response errorValidation
type ErrorValidationWrapper struct {
	Body ValidationError
}

// A list of products returned in the response
// swagger:response productsResponse
type ProductsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// Returned when successfully deleting/updating a product
// swagger:response noContent
type ProductsNoContent struct{}

// swagger:response productResponse
type ProductResponseWrapper struct {
	// Newly added product
	// in: body
	Body data.Product
}

// swagger:parameters deleteProduct
type ProductIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}
