package resourcemover

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UnresolvedDependenciesClient is the a first party Azure service orchestrating the move of Azure resources from one
// Azure region to another or between zones within a region.
type UnresolvedDependenciesClient struct {
	BaseClient
}

// NewUnresolvedDependenciesClient creates an instance of the UnresolvedDependenciesClient client.
func NewUnresolvedDependenciesClient(subscriptionID string) UnresolvedDependenciesClient {
	return NewUnresolvedDependenciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUnresolvedDependenciesClientWithBaseURI creates an instance of the UnresolvedDependenciesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewUnresolvedDependenciesClientWithBaseURI(baseURI string, subscriptionID string) UnresolvedDependenciesClient {
	return UnresolvedDependenciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a list of unresolved dependencies.
// Parameters:
// resourceGroupName - the Resource Group Name.
// moveCollectionName - the Move Collection Name.
func (client UnresolvedDependenciesClient) Get(ctx context.Context, resourceGroupName string, moveCollectionName string) (result UnresolvedDependencyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UnresolvedDependenciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, moveCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcemover.UnresolvedDependenciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client UnresolvedDependenciesClient) GetPreparer(ctx context.Context, resourceGroupName string, moveCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"moveCollectionName": autorest.Encode("path", moveCollectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/unresolvedDependencies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UnresolvedDependenciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UnresolvedDependenciesClient) GetResponder(resp *http.Response) (result UnresolvedDependencyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
