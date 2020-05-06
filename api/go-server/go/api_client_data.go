/*
 * worker-monitor management API
 *
 * The API manages worker-monitor client data.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A ClientDataApiController binds http requests to an api service and writes the service results to the http response
type ClientDataApiController struct {
	service ClientDataApiServicer
}

// NewClientDataApiController creates a default api controller
func NewClientDataApiController(s ClientDataApiServicer) Router {
	return &ClientDataApiController{ service: s }
}

// Routes returns all of the api route for the ClientDataApiController
func (c *ClientDataApiController) Routes() Routes {
	return Routes{ 
		{
			"ClientdataClientIDGet",
			strings.ToUpper("Get"),
			"/clientdata/{clientID}",
			c.ClientdataClientIDGet,
		},
		{
			"ClientdataGet",
			strings.ToUpper("Get"),
			"/clientdata",
			c.ClientdataGet,
		},
		{
			"ClientdataPost",
			strings.ToUpper("Post"),
			"/clientdata",
			c.ClientdataPost,
		},
	}
}

// ClientdataClientIDGet - Get target client data by client ID
func (c *ClientDataApiController) ClientdataClientIDGet(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	clientID, err := parseIntParameter(params["clientID"])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ClientdataClientIDGet(clientID)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ClientdataGet - Get target client data
func (c *ClientDataApiController) ClientdataGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.ClientdataGet()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// ClientdataPost - Register target client data
func (c *ClientDataApiController) ClientdataPost(w http.ResponseWriter, r *http.Request) { 
	clientData := &[]ClientData{}
	if err := json.NewDecoder(r.Body).Decode(&clientData); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.ClientdataPost(*clientData)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}