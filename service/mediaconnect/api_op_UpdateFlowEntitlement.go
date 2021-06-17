// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You can change an entitlement's description, subscribers, and encryption. If you
// change the subscribers, the service will remove the outputs that are are used by
// the subscribers that are removed.
func (c *Client) UpdateFlowEntitlement(ctx context.Context, params *UpdateFlowEntitlementInput, optFns ...func(*Options)) (*UpdateFlowEntitlementOutput, error) {
	if params == nil {
		params = &UpdateFlowEntitlementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlowEntitlement", params, optFns, c.addOperationUpdateFlowEntitlementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowEntitlementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The entitlement fields that you want to update.
type UpdateFlowEntitlementInput struct {

	// The ARN of the entitlement that you want to update.
	//
	// This member is required.
	EntitlementArn *string

	// The flow that is associated with the entitlement that you want to update.
	//
	// This member is required.
	FlowArn *string

	// A description of the entitlement. This description appears only on the AWS
	// Elemental MediaConnect console and will not be seen by the subscriber or end
	// user.
	Description *string

	// The type of encryption that will be used on the output associated with this
	// entitlement.
	Encryption *types.UpdateEncryption

	// An indication of whether you want to enable the entitlement to allow access, or
	// disable it to stop streaming content to the subscriber’s flow temporarily. If
	// you don’t specify the entitlementStatus field in your request, MediaConnect
	// leaves the value unchanged.
	EntitlementStatus types.EntitlementStatus

	// The AWS account IDs that you want to share your content with. The receiving
	// accounts (subscribers) will be allowed to create their own flow using your
	// content as the source.
	Subscribers []string
}

type UpdateFlowEntitlementOutput struct {

	// The new configuration of the entitlement that you updated.
	Entitlement *types.Entitlement

	// The ARN of the flow that this entitlement was granted on.
	FlowArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateFlowEntitlementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlowEntitlement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlowEntitlement{}, middleware.After)
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
	if err = addOpUpdateFlowEntitlementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlowEntitlement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFlowEntitlement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "UpdateFlowEntitlement",
	}
}
