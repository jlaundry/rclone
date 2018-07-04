package media

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ContentKeyPoliciesClient is the client for the ContentKeyPolicies methods of the Media service.
type ContentKeyPoliciesClient struct {
	BaseClient
}

// NewContentKeyPoliciesClient creates an instance of the ContentKeyPoliciesClient client.
func NewContentKeyPoliciesClient(subscriptionID string) ContentKeyPoliciesClient {
	return NewContentKeyPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContentKeyPoliciesClientWithBaseURI creates an instance of the ContentKeyPoliciesClient client.
func NewContentKeyPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ContentKeyPoliciesClient {
	return ContentKeyPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Content Key Policy in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// contentKeyPolicyName - the Content Key Policy name.
// parameters - the request parameters
func (client ContentKeyPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy) (result ContentKeyPolicy, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ContentKeyPolicyProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ContentKeyPolicyProperties.Options", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("media.ContentKeyPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, accountName, contentKeyPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ContentKeyPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"contentKeyPolicyName": autorest.Encode("path", contentKeyPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ContentKeyPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ContentKeyPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ContentKeyPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Content Key Policy in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// contentKeyPolicyName - the Content Key Policy name.
func (client ContentKeyPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, contentKeyPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ContentKeyPoliciesClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"contentKeyPolicyName": autorest.Encode("path", contentKeyPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ContentKeyPoliciesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ContentKeyPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the details of a Content Key Policy in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// contentKeyPolicyName - the Content Key Policy name.
func (client ContentKeyPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (result ContentKeyPolicy, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, contentKeyPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ContentKeyPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"contentKeyPolicyName": autorest.Encode("path", contentKeyPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ContentKeyPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ContentKeyPoliciesClient) GetResponder(resp *http.Response) (result ContentKeyPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPolicyPropertiesWithSecrets get a Content Key Policy including secret values
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// contentKeyPolicyName - the Content Key Policy name.
func (client ContentKeyPoliciesClient) GetPolicyPropertiesWithSecrets(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (result ContentKeyPolicyProperties, err error) {
	req, err := client.GetPolicyPropertiesWithSecretsPreparer(ctx, resourceGroupName, accountName, contentKeyPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "GetPolicyPropertiesWithSecrets", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPolicyPropertiesWithSecretsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "GetPolicyPropertiesWithSecrets", resp, "Failure sending request")
		return
	}

	result, err = client.GetPolicyPropertiesWithSecretsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "GetPolicyPropertiesWithSecrets", resp, "Failure responding to request")
	}

	return
}

// GetPolicyPropertiesWithSecretsPreparer prepares the GetPolicyPropertiesWithSecrets request.
func (client ContentKeyPoliciesClient) GetPolicyPropertiesWithSecretsPreparer(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"contentKeyPolicyName": autorest.Encode("path", contentKeyPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}/getPolicyPropertiesWithSecrets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPolicyPropertiesWithSecretsSender sends the GetPolicyPropertiesWithSecrets request. The method will close the
// http.Response Body if it receives an error.
func (client ContentKeyPoliciesClient) GetPolicyPropertiesWithSecretsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetPolicyPropertiesWithSecretsResponder handles the response to the GetPolicyPropertiesWithSecrets request. The method always
// closes the http.Response Body.
func (client ContentKeyPoliciesClient) GetPolicyPropertiesWithSecretsResponder(resp *http.Response) (result ContentKeyPolicyProperties, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the Content Key Policies in the account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// filter - restricts the set of items returned.
// top - specifies a non-negative integer n that limits the number of items returned from a collection. The
// service returns the number of available items up to but not greater than the specified value n.
// orderby - specifies the the key by which the result collection should be ordered.
func (client ContentKeyPoliciesClient) List(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result ContentKeyPolicyCollectionPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ckpc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "List", resp, "Failure sending request")
		return
	}

	result.ckpc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ContentKeyPoliciesClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ContentKeyPoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ContentKeyPoliciesClient) ListResponder(resp *http.Response) (result ContentKeyPolicyCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ContentKeyPoliciesClient) listNextResults(lastResults ContentKeyPolicyCollection) (result ContentKeyPolicyCollection, err error) {
	req, err := lastResults.contentKeyPolicyCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContentKeyPoliciesClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, orderby string) (result ContentKeyPolicyCollectionIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, accountName, filter, top, orderby)
	return
}

// Update updates an existing Content Key Policy in the Media Services account
// Parameters:
// resourceGroupName - the name of the resource group within the Azure subscription.
// accountName - the Media Services account name.
// contentKeyPolicyName - the Content Key Policy name.
// parameters - the request parameters
func (client ContentKeyPoliciesClient) Update(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy) (result ContentKeyPolicy, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, contentKeyPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.ContentKeyPoliciesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ContentKeyPoliciesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":          autorest.Encode("path", accountName),
		"contentKeyPolicyName": autorest.Encode("path", contentKeyPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ContentKeyPoliciesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ContentKeyPoliciesClient) UpdateResponder(resp *http.Response) (result ContentKeyPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
