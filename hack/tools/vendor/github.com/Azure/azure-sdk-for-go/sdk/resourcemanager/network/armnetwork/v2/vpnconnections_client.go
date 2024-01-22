//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VPNConnectionsClient contains the methods for the VPNConnections group.
// Don't use this type directly, use NewVPNConnectionsClient() instead.
type VPNConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVPNConnectionsClient creates a new instance of VPNConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVPNConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".VPNConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VPNConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing
// connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the connection.
//   - vpnConnectionParameters - Parameters supplied to create or Update a VPN Connection.
//   - options - VPNConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *VPNConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VPNConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VPNConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, vpnConnectionParameters)
}

// BeginDelete - Deletes a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the connection.
//   - options - VPNConnectionsClientBeginDeleteOptions contains the optional parameters for the VPNConnectionsClient.BeginDelete
//     method.
func (client *VPNConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*runtime.Poller[VPNConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, connectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VPNConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the vpn connection.
//   - options - VPNConnectionsClientGetOptions contains the optional parameters for the VPNConnectionsClient.Get method.
func (client *VPNConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientGetOptions) (VPNConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VPNConnectionsClient) getHandleResponse(resp *http.Response) (VPNConnectionsClientGetResponse, error) {
	result := VPNConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNConnection); err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVPNGatewayPager - Retrieves all vpn connections for a particular virtual wan vpn gateway.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - options - VPNConnectionsClientListByVPNGatewayOptions contains the optional parameters for the VPNConnectionsClient.NewListByVPNGatewayPager
//     method.
func (client *VPNConnectionsClient) NewListByVPNGatewayPager(resourceGroupName string, gatewayName string, options *VPNConnectionsClientListByVPNGatewayOptions) *runtime.Pager[VPNConnectionsClientListByVPNGatewayResponse] {
	return runtime.NewPager(runtime.PagingHandler[VPNConnectionsClientListByVPNGatewayResponse]{
		More: func(page VPNConnectionsClientListByVPNGatewayResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNConnectionsClientListByVPNGatewayResponse) (VPNConnectionsClientListByVPNGatewayResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByVPNGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VPNConnectionsClientListByVPNGatewayResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VPNConnectionsClientListByVPNGatewayResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VPNConnectionsClientListByVPNGatewayResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVPNGatewayHandleResponse(resp)
		},
	})
}

// listByVPNGatewayCreateRequest creates the ListByVPNGateway request.
func (client *VPNConnectionsClient) listByVPNGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNConnectionsClientListByVPNGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVPNGatewayHandleResponse handles the ListByVPNGateway response.
func (client *VPNConnectionsClient) listByVPNGatewayHandleResponse(resp *http.Response) (VPNConnectionsClientListByVPNGatewayResponse, error) {
	result := VPNConnectionsClientListByVPNGatewayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNConnectionsResult); err != nil {
		return VPNConnectionsClientListByVPNGatewayResponse{}, err
	}
	return result, nil
}

// BeginStartPacketCapture - Starts packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - gatewayName - The name of the gateway.
//   - vpnConnectionName - The name of the vpn connection.
//   - options - VPNConnectionsClientBeginStartPacketCaptureOptions contains the optional parameters for the VPNConnectionsClient.BeginStartPacketCapture
//     method.
func (client *VPNConnectionsClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStartPacketCaptureOptions) (*runtime.Poller[VPNConnectionsClientStartPacketCaptureResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.startPacketCapture(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientStartPacketCaptureResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNConnectionsClientStartPacketCaptureResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// StartPacketCapture - Starts packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) startPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStartPacketCaptureOptions) (*http.Response, error) {
	req, err := client.startPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *VPNConnectionsClient) startPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStartPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/startpacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if vpnConnectionName == "" {
		return nil, errors.New("parameter vpnConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnConnectionName}", url.PathEscape(vpnConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// BeginStopPacketCapture - Stops packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - gatewayName - The name of the gateway.
//   - vpnConnectionName - The name of the vpn connection.
//   - options - VPNConnectionsClientBeginStopPacketCaptureOptions contains the optional parameters for the VPNConnectionsClient.BeginStopPacketCapture
//     method.
func (client *VPNConnectionsClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStopPacketCaptureOptions) (*runtime.Poller[VPNConnectionsClientStopPacketCaptureResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stopPacketCapture(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientStopPacketCaptureResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VPNConnectionsClientStopPacketCaptureResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// StopPacketCapture - Stops packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) stopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStopPacketCaptureOptions) (*http.Response, error) {
	req, err := client.stopPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *VPNConnectionsClient) stopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStopPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/stoppacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if vpnConnectionName == "" {
		return nil, errors.New("parameter vpnConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnConnectionName}", url.PathEscape(vpnConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}
