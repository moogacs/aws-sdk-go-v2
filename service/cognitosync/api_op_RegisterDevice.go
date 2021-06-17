// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers a device to receive push sync notifications.This API can only be
// called with temporary credentials provided by Cognito Identity. You cannot call
// this API with developer credentials. RegisterDevice The following examples have
// been edited for readability. POST / HTTP/1.1 CONTENT-TYPE: application/json
// X-AMZN-REQUESTID: 368f9200-3eca-449e-93b3-7b9c08d8e185 X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.RegisterDevice HOST:
// cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T194643Z
// X-AMZ-SECURITY-TOKEN: AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#RegisterDevice",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "IdentityId": "IDENTITY_ID", "Platform": "GCM",
// "Token": "PUSH_TOKEN" } } 1.1 200 OK x-amzn-requestid:
// 368f9200-3eca-449e-93b3-7b9c08d8e185 date: Sat, 04 Oct 2014 19:46:44 GMT
// content-type: application/json content-length: 145 { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#RegisterDeviceResponse", "DeviceId":
// "5cd28fbe-dd83-47ab-9f83-19093a5fb014" }, "Version": "1.0" }
func (c *Client) RegisterDevice(ctx context.Context, params *RegisterDeviceInput, optFns ...func(*Options)) (*RegisterDeviceOutput, error) {
	if params == nil {
		params = &RegisterDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterDevice", params, optFns, c.addOperationRegisterDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to RegisterDevice.
type RegisterDeviceInput struct {

	// The unique ID for this identity.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. Here, the ID of the pool that the identity belongs
	// to.
	//
	// This member is required.
	IdentityPoolId *string

	// The SNS platform type (e.g. GCM, SDM, APNS, APNS_SANDBOX).
	//
	// This member is required.
	Platform types.Platform

	// The push token.
	//
	// This member is required.
	Token *string
}

// Response to a RegisterDevice request.
type RegisterDeviceOutput struct {

	// The unique ID generated for this device by Cognito.
	DeviceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationRegisterDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRegisterDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterDevice(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRegisterDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "RegisterDevice",
	}
}
