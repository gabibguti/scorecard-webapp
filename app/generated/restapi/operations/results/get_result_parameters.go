// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetResultParams creates a new GetResultParams object
//
// There are no default values defined in the spec.
func NewGetResultParams() GetResultParams {

	return GetResultParams{}
}

// GetResultParams contains all the bound params for the get result operation
// typically these are obtained from a http.Request
//
// swagger:parameters getResult
type GetResultParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*SHA1 commit hash expressed in hexadecimal format
	  Pattern: ^[0-9a-fA-F]{40}$
	  In: query
	*/
	Commit *string
	/*Name of the owner/organization of the repository
	  Required: true
	  In: path
	*/
	Org string
	/*VCS platform. eg. github.com
	  Required: true
	  In: path
	*/
	Platform string
	/*Name of the repository
	  Required: true
	  In: path
	*/
	Repo string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetResultParams() beforehand.
func (o *GetResultParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCommit, qhkCommit, _ := qs.GetOK("commit")
	if err := o.bindCommit(qCommit, qhkCommit, route.Formats); err != nil {
		res = append(res, err)
	}

	rOrg, rhkOrg, _ := route.Params.GetOK("org")
	if err := o.bindOrg(rOrg, rhkOrg, route.Formats); err != nil {
		res = append(res, err)
	}

	rPlatform, rhkPlatform, _ := route.Params.GetOK("platform")
	if err := o.bindPlatform(rPlatform, rhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	rRepo, rhkRepo, _ := route.Params.GetOK("repo")
	if err := o.bindRepo(rRepo, rhkRepo, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCommit binds and validates parameter Commit from query.
func (o *GetResultParams) bindCommit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Commit = &raw

	if err := o.validateCommit(formats); err != nil {
		return err
	}

	return nil
}

// validateCommit carries on validations for parameter Commit
func (o *GetResultParams) validateCommit(formats strfmt.Registry) error {

	if err := validate.Pattern("commit", "query", *o.Commit, `^[0-9a-fA-F]{40}$`); err != nil {
		return err
	}

	return nil
}

// bindOrg binds and validates parameter Org from path.
func (o *GetResultParams) bindOrg(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Org = raw

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetResultParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Platform = raw

	return nil
}

// bindRepo binds and validates parameter Repo from path.
func (o *GetResultParams) bindRepo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Repo = raw

	return nil
}
