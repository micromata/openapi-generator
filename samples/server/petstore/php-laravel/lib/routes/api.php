<?php

/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 * PHP version 7.1.3
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI-Generator
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 * Source files are located at:
 *
 * > https://github.com/OpenAPITools/openapi-generator/blob/master/modules/openapi-generator/src/main/resources/php-laravel/
 */

/**
 * patch call123TestSpecialTags
 * Summary: To test special tags
 * Notes: To test special tags and operation ID starting with number
 * Output-Formats: [application/json]
 */
Route::patch('/v2/another-fake/dummy', 'AnotherFakeController@call123TestSpecialTags');
/**
 * patch testClientModel
 * Summary: To test \&quot;client\&quot; model
 * Notes: To test \&quot;client\&quot; model
 * Output-Formats: [application/json]
 */
Route::patch('/v2/fake', 'FakeController@testClientModel');
/**
 * post testEndpointParameters
 * Summary: Fake endpoint for testing various parameters  假端點  偽のエンドポイント  가짜 엔드 포인트
 * Notes: Fake endpoint for testing various parameters  假端點  偽のエンドポイント  가짜 엔드 포인트

 */
Route::post('/v2/fake', 'FakeController@testEndpointParameters');
/**
 * get testEnumParameters
 * Summary: To test enum parameters
 * Notes: To test enum parameters

 */
Route::get('/v2/fake', 'FakeController@testEnumParameters');
/**
 * delete testGroupParameters
 * Summary: Fake endpoint to test group parameters (optional)
 * Notes: Fake endpoint to test group parameters (optional)

 */
Route::delete('/v2/fake', 'FakeController@testGroupParameters');
/**
 * put testBodyWithFileSchema
 * Summary: 
 * Notes: For this test, the body for this request much reference a schema named &#x60;File&#x60;.

 */
Route::put('/v2/fake/body-with-file-schema', 'FakeController@testBodyWithFileSchema');
/**
 * put testBodyWithQueryParams
 * Summary: 
 * Notes: 

 */
Route::put('/v2/fake/body-with-query-params', 'FakeController@testBodyWithQueryParams');
/**
 * post createXmlItem
 * Summary: creates an XmlItem
 * Notes: this route creates an XmlItem

 */
Route::post('/v2/fake/create_xml_item', 'FakeController@createXmlItem');
/**
 * post testInlineAdditionalProperties
 * Summary: test inline additionalProperties
 * Notes: 

 */
Route::post('/v2/fake/inline-additionalProperties', 'FakeController@testInlineAdditionalProperties');
/**
 * get testJsonFormData
 * Summary: test json serialization of form data
 * Notes: 

 */
Route::get('/v2/fake/jsonFormData', 'FakeController@testJsonFormData');
/**
 * post fakeOuterBooleanSerialize
 * Summary: 
 * Notes: Test serialization of outer boolean types
 * Output-Formats: [*_/_*]
 */
Route::post('/v2/fake/outer/boolean', 'FakeController@fakeOuterBooleanSerialize');
/**
 * post fakeOuterCompositeSerialize
 * Summary: 
 * Notes: Test serialization of object with outer number type
 * Output-Formats: [*_/_*]
 */
Route::post('/v2/fake/outer/composite', 'FakeController@fakeOuterCompositeSerialize');
/**
 * post fakeOuterNumberSerialize
 * Summary: 
 * Notes: Test serialization of outer number types
 * Output-Formats: [*_/_*]
 */
Route::post('/v2/fake/outer/number', 'FakeController@fakeOuterNumberSerialize');
/**
 * post fakeOuterStringSerialize
 * Summary: 
 * Notes: Test serialization of outer string types
 * Output-Formats: [*_/_*]
 */
Route::post('/v2/fake/outer/string', 'FakeController@fakeOuterStringSerialize');
/**
 * put testQueryParameterCollectionFormat
 * Summary: 
 * Notes: To test the collection format in query parameters

 */
Route::put('/v2/fake/test-query-paramters', 'FakeController@testQueryParameterCollectionFormat');
/**
 * patch testClassname
 * Summary: To test class name in snake case
 * Notes: To test class name in snake case
 * Output-Formats: [application/json]
 */
Route::patch('/v2/fake_classname_test', 'FakeClassnameTags123Controller@testClassname');
/**
 * post uploadFileWithRequiredFile
 * Summary: uploads an image (required)
 * Notes: 
 * Output-Formats: [application/json]
 */
Route::post('/v2/fake/{petId}/uploadImageWithRequiredFile', 'PetController@uploadFileWithRequiredFile');
/**
 * post addPet
 * Summary: Add a new pet to the store
 * Notes: 

 */
Route::post('/v2/pet', 'PetController@addPet');
/**
 * put updatePet
 * Summary: Update an existing pet
 * Notes: 

 */
Route::put('/v2/pet', 'PetController@updatePet');
/**
 * get findPetsByStatus
 * Summary: Finds Pets by status
 * Notes: Multiple status values can be provided with comma separated strings
 * Output-Formats: [application/xml, application/json]
 */
Route::get('/v2/pet/findByStatus', 'PetController@findPetsByStatus');
/**
 * get findPetsByTags
 * Summary: Finds Pets by tags
 * Notes: Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
 * Output-Formats: [application/xml, application/json]
 */
Route::get('/v2/pet/findByTags', 'PetController@findPetsByTags');
/**
 * delete deletePet
 * Summary: Deletes a pet
 * Notes: 

 */
Route::delete('/v2/pet/{petId}', 'PetController@deletePet');
/**
 * get getPetById
 * Summary: Find pet by ID
 * Notes: Returns a single pet
 * Output-Formats: [application/xml, application/json]
 */
Route::get('/v2/pet/{petId}', 'PetController@getPetById');
/**
 * post updatePetWithForm
 * Summary: Updates a pet in the store with form data
 * Notes: 

 */
Route::post('/v2/pet/{petId}', 'PetController@updatePetWithForm');
/**
 * post uploadFile
 * Summary: uploads an image
 * Notes: 
 * Output-Formats: [application/json]
 */
Route::post('/v2/pet/{petId}/uploadImage', 'PetController@uploadFile');
/**
 * get getInventory
 * Summary: Returns pet inventories by status
 * Notes: Returns a map of status codes to quantities
 * Output-Formats: [application/json]
 */
Route::get('/v2/store/inventory', 'StoreController@getInventory');
/**
 * post placeOrder
 * Summary: Place an order for a pet
 * Notes: 
 * Output-Formats: [application/xml, application/json]
 */
Route::post('/v2/store/order', 'StoreController@placeOrder');
/**
 * delete deleteOrder
 * Summary: Delete purchase order by ID
 * Notes: For valid response try integer IDs with value &lt; 1000. Anything above 1000 or nonintegers will generate API errors

 */
Route::delete('/v2/store/order/{order_id}', 'StoreController@deleteOrder');
/**
 * get getOrderById
 * Summary: Find purchase order by ID
 * Notes: For valid response try integer IDs with value &lt;&#x3D; 5 or &gt; 10. Other values will generated exceptions
 * Output-Formats: [application/xml, application/json]
 */
Route::get('/v2/store/order/{order_id}', 'StoreController@getOrderById');
/**
 * post createUser
 * Summary: Create user
 * Notes: This can only be done by the logged in user.

 */
Route::post('/v2/user', 'UserController@createUser');
/**
 * post createUsersWithArrayInput
 * Summary: Creates list of users with given input array
 * Notes: 

 */
Route::post('/v2/user/createWithArray', 'UserController@createUsersWithArrayInput');
/**
 * post createUsersWithListInput
 * Summary: Creates list of users with given input array
 * Notes: 

 */
Route::post('/v2/user/createWithList', 'UserController@createUsersWithListInput');
/**
 * get loginUser
 * Summary: Logs user into the system
 * Notes: 
 * Output-Formats: [application/xml, application/json]
 */
Route::get('/v2/user/login', 'UserController@loginUser');
/**
 * get logoutUser
 * Summary: Logs out current logged in user session
 * Notes: 

 */
Route::get('/v2/user/logout', 'UserController@logoutUser');
/**
 * delete deleteUser
 * Summary: Delete user
 * Notes: This can only be done by the logged in user.

 */
Route::delete('/v2/user/{username}', 'UserController@deleteUser');
/**
 * get getUserByName
 * Summary: Get user by user name
 * Notes: 
 * Output-Formats: [application/xml, application/json]
 */
Route::get('/v2/user/{username}', 'UserController@getUserByName');
/**
 * put updateUser
 * Summary: Updated user
 * Notes: This can only be done by the logged in user.

 */
Route::put('/v2/user/{username}', 'UserController@updateUser');

