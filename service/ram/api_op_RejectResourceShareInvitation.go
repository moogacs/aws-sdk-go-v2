// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Rejects an invitation to a resource share from another AWS account.
func (c *Client) RejectResourceShareInvitation(ctx context.Context, params *RejectResourceShareInvitationInput, optFns ...func(*Options)) (*RejectResourceShareInvitationOutput, error) {
	if params == nil {
		params = &RejectResourceShareInvitationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectResourceShareInvitation", params, optFns, c.addOperationRejectResourceShareInvitationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectResourceShareInvitationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectResourceShareInvitationInput struct {

	// The Amazon Resource Name (ARN) of the invitation.
	//
	// This member is required.
	ResourceShareInvitationArn *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
}

type RejectResourceShareInvitationOutput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// Information about the invitation.
	ResourceShareInvitation *types.ResourceShareInvitation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationRejectResourceShareInvitationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRejectResourceShareInvitation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRejectResourceShareInvitation{}, middleware.After)
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
	if err = addOpRejectResourceShareInvitationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRejectResourceShareInvitation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRejectResourceShareInvitation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "RejectResourceShareInvitation",
	}
}
