/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"os"
)

// Linger please
var (
	_ _context.Context
)

// PetApiService PetApi service
type PetApiService service

type apiAddPetRequest struct {
	ctx _context.Context
	apiService *PetApiService
	body *Pet
}


func (r apiAddPetRequest) Body(body Pet) apiAddPetRequest {
	r.body = &body
	return r
}

/*
AddPet Add a new pet to the store
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiAddPetRequest
*/
func (a *PetApiService) AddPet(ctx _context.Context) apiAddPetRequest {
	return apiAddPetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request

*/
func (r apiAddPetRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.AddPet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiDeletePetRequest struct {
	ctx _context.Context
	apiService *PetApiService
	petId int64
	apiKey *string
}


func (r apiDeletePetRequest) ApiKey(apiKey string) apiDeletePetRequest {
	r.apiKey = &apiKey
	return r
}

/*
DeletePet Deletes a pet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId Pet id to delete
@return apiDeletePetRequest
*/
func (a *PetApiService) DeletePet(ctx _context.Context, petId int64) apiDeletePetRequest {
	return apiDeletePetRequest{
		apiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
Execute executes the request

*/
func (r apiDeletePetRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.DeletePet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.apiKey != nil {
		localVarHeaderParams["api_key"] = parameterToString(*r.apiKey, "")
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiFindPetsByStatusRequest struct {
	ctx _context.Context
	apiService *PetApiService
	status *[]string
}


func (r apiFindPetsByStatusRequest) Status(status []string) apiFindPetsByStatusRequest {
	r.status = &status
	return r
}

/*
FindPetsByStatus Finds Pets by status
Multiple status values can be provided with comma separated strings
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFindPetsByStatusRequest
*/
func (a *PetApiService) FindPetsByStatus(ctx _context.Context) apiFindPetsByStatusRequest {
	return apiFindPetsByStatusRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return []Pet
*/
func (r apiFindPetsByStatusRequest) Execute() ([]Pet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Pet
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.FindPetsByStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/findByStatus"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.status == nil {
		return localVarReturnValue, nil, reportError("status is required and must be specified")
	}

	localVarQueryParams.Add("status", parameterToString(*r.status, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiFindPetsByTagsRequest struct {
	ctx _context.Context
	apiService *PetApiService
	tags *[]string
}


func (r apiFindPetsByTagsRequest) Tags(tags []string) apiFindPetsByTagsRequest {
	r.tags = &tags
	return r
}

/*
FindPetsByTags Finds Pets by tags
Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFindPetsByTagsRequest
*/
func (a *PetApiService) FindPetsByTags(ctx _context.Context) apiFindPetsByTagsRequest {
	return apiFindPetsByTagsRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request
 @return []Pet
*/
func (r apiFindPetsByTagsRequest) Execute() ([]Pet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Pet
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.FindPetsByTags")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/findByTags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.tags == nil {
		return localVarReturnValue, nil, reportError("tags is required and must be specified")
	}

	localVarQueryParams.Add("tags", parameterToString(*r.tags, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiGetPetByIdRequest struct {
	ctx _context.Context
	apiService *PetApiService
	petId int64
}


/*
GetPetById Find pet by ID
Returns a single pet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet to return
@return apiGetPetByIdRequest
*/
func (a *PetApiService) GetPetById(ctx _context.Context, petId int64) apiGetPetByIdRequest {
	return apiGetPetByIdRequest{
		apiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
Execute executes the request
 @return Pet
*/
func (r apiGetPetByIdRequest) Execute() (Pet, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Pet
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.GetPetById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["api_key"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["api_key"] = key
			}
		}
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiUpdatePetRequest struct {
	ctx _context.Context
	apiService *PetApiService
	body *Pet
}


func (r apiUpdatePetRequest) Body(body Pet) apiUpdatePetRequest {
	r.body = &body
	return r
}

/*
UpdatePet Update an existing pet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiUpdatePetRequest
*/
func (a *PetApiService) UpdatePet(ctx _context.Context) apiUpdatePetRequest {
	return apiUpdatePetRequest{
		apiService: a,
		ctx: ctx,
	}
}

/*
Execute executes the request

*/
func (r apiUpdatePetRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UpdatePet")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiUpdatePetWithFormRequest struct {
	ctx _context.Context
	apiService *PetApiService
	petId int64
	name *string
	status *string
}


func (r apiUpdatePetWithFormRequest) Name(name string) apiUpdatePetWithFormRequest {
	r.name = &name
	return r
}

func (r apiUpdatePetWithFormRequest) Status(status string) apiUpdatePetWithFormRequest {
	r.status = &status
	return r
}

/*
UpdatePetWithForm Updates a pet in the store with form data
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet that needs to be updated
@return apiUpdatePetWithFormRequest
*/
func (a *PetApiService) UpdatePetWithForm(ctx _context.Context, petId int64) apiUpdatePetWithFormRequest {
	return apiUpdatePetWithFormRequest{
		apiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
Execute executes the request

*/
func (r apiUpdatePetWithFormRequest) Execute() (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UpdatePetWithForm")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
		
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.name != nil {
		localVarFormParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.status != nil {
		localVarFormParams.Add("status", parameterToString(*r.status, ""))
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type apiUploadFileRequest struct {
	ctx _context.Context
	apiService *PetApiService
	petId int64
	additionalMetadata *string
	file **os.File
}


func (r apiUploadFileRequest) AdditionalMetadata(additionalMetadata string) apiUploadFileRequest {
	r.additionalMetadata = &additionalMetadata
	return r
}

func (r apiUploadFileRequest) File(file *os.File) apiUploadFileRequest {
	r.file = &file
	return r
}

/*
UploadFile uploads an image
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet to update
@return apiUploadFileRequest
*/
func (a *PetApiService) UploadFile(ctx _context.Context, petId int64) apiUploadFileRequest {
	return apiUploadFileRequest{
		apiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
Execute executes the request
 @return ApiResponse
*/
func (r apiUploadFileRequest) Execute() (ApiResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pet/{petId}/uploadImage"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
		
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.additionalMetadata != nil {
		localVarFormParams.Add("additionalMetadata", parameterToString(*r.additionalMetadata, ""))
	}
	localVarFormFileName = "file"
	var localVarFile *os.File
	if r.file != nil {
		localVarFile = *r.file
	}
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiUploadFileWithRequiredFileRequest struct {
	ctx _context.Context
	apiService *PetApiService
	petId int64
	requiredFile **os.File
	additionalMetadata *string
}


func (r apiUploadFileWithRequiredFileRequest) RequiredFile(requiredFile *os.File) apiUploadFileWithRequiredFileRequest {
	r.requiredFile = &requiredFile
	return r
}

func (r apiUploadFileWithRequiredFileRequest) AdditionalMetadata(additionalMetadata string) apiUploadFileWithRequiredFileRequest {
	r.additionalMetadata = &additionalMetadata
	return r
}

/*
UploadFileWithRequiredFile uploads an image (required)
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param petId ID of pet to update
@return apiUploadFileWithRequiredFileRequest
*/
func (a *PetApiService) UploadFileWithRequiredFile(ctx _context.Context, petId int64) apiUploadFileWithRequiredFileRequest {
	return apiUploadFileWithRequiredFileRequest{
		apiService: a,
		ctx: ctx,
		petId: petId,
	}
}

/*
Execute executes the request
 @return ApiResponse
*/
func (r apiUploadFileWithRequiredFileRequest) Execute() (ApiResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiResponse
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "PetApiService.UploadFileWithRequiredFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fake/{petId}/uploadImageWithRequiredFile"
	localVarPath = strings.Replace(localVarPath, "{"+"petId"+"}", _neturl.PathEscape(parameterToString(r.petId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	
	
	if r.requiredFile == nil {
		return localVarReturnValue, nil, reportError("requiredFile is required and must be specified")
	}
	
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.additionalMetadata != nil {
		localVarFormParams.Add("additionalMetadata", parameterToString(*r.additionalMetadata, ""))
	}
	localVarFormFileName = "requiredFile"
	localVarFile := *r.requiredFile
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
