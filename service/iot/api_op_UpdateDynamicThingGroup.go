// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a dynamic thing group.
func (c *Client) UpdateDynamicThingGroup(ctx context.Context, params *UpdateDynamicThingGroupInput, optFns ...func(*Options)) (*UpdateDynamicThingGroupOutput, error) {
	if params == nil {
		params = &UpdateDynamicThingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDynamicThingGroup", params, optFns, c.addOperationUpdateDynamicThingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDynamicThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDynamicThingGroupInput struct {

	// The name of the dynamic thing group to update.
	//
	// This member is required.
	ThingGroupName *string

	// The dynamic thing group properties to update.
	//
	// This member is required.
	ThingGroupProperties *types.ThingGroupProperties

	// The expected version of the dynamic thing group to update.
	ExpectedVersion *int64

	// The dynamic thing group index to update. Currently one index is supported:
	// 'AWS_Things'.
	IndexName *string

	// The dynamic thing group search query string to update.
	QueryString *string

	// The dynamic thing group query version to update. Currently one query version is
	// supported: "2017-09-30". If not specified, the query version defaults to this
	// value.
	QueryVersion *string
}

type UpdateDynamicThingGroupOutput struct {

	// The dynamic thing group version.
	Version int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateDynamicThingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDynamicThingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDynamicThingGroup{}, middleware.After)
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
	if err = addOpUpdateDynamicThingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDynamicThingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDynamicThingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateDynamicThingGroup",
	}
}
