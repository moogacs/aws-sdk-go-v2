// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Unsubscribes from receiving notifications when a dataset is modified by another
// device.This API can only be called with temporary credentials provided by
// Cognito Identity. You cannot call this API with developer credentials.
// UnsubscribeFromDataset The following examples have been edited for readability.
// POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZ-REQUESTSUPERTRACE: true
// X-AMZN-REQUESTID: 676896d6-14ca-45b1-8029-6d36b10a077e X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.UnsubscribeFromDataset
// HOST: cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T195446Z
// X-AMZ-SECURITY-TOKEN: AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#UnsubscribeFromDataset", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "IdentityId": "IDENTITY_ID", "DatasetName":
// "Rufus", "DeviceId": "5cd28fbe-dd83-47ab-9f83-19093a5fb014" } } 1.1 200 OK
// x-amzn-requestid: 676896d6-14ca-45b1-8029-6d36b10a077e date: Sat, 04 Oct 2014
// 19:54:46 GMT content-type: application/json content-length: 103 { "Output": {
// "__type": "com.amazonaws.cognito.sync.model#UnsubscribeFromDatasetResponse" },
// "Version": "1.0" }
func (c *Client) UnsubscribeFromDataset(ctx context.Context, params *UnsubscribeFromDatasetInput, optFns ...func(*Options)) (*UnsubscribeFromDatasetOutput, error) {
	if params == nil {
		params = &UnsubscribeFromDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnsubscribeFromDataset", params, optFns, c.addOperationUnsubscribeFromDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnsubscribeFromDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to UnsubscribeFromDataset.
type UnsubscribeFromDatasetInput struct {

	// The name of the dataset from which to unsubcribe.
	//
	// This member is required.
	DatasetName *string

	// The unique ID generated for this device by Cognito.
	//
	// This member is required.
	DeviceId *string

	// Unique ID for this identity.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. The ID of the pool to which this identity belongs.
	//
	// This member is required.
	IdentityPoolId *string
}

// Response to an UnsubscribeFromDataset request.
type UnsubscribeFromDatasetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUnsubscribeFromDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUnsubscribeFromDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUnsubscribeFromDataset{}, middleware.After)
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
	if err = addOpUnsubscribeFromDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnsubscribeFromDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnsubscribeFromDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "UnsubscribeFromDataset",
	}
}
