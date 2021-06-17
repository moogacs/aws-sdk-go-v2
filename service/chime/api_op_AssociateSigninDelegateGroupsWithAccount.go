// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified sign-in delegate groups with the specified Amazon Chime
// account.
func (c *Client) AssociateSigninDelegateGroupsWithAccount(ctx context.Context, params *AssociateSigninDelegateGroupsWithAccountInput, optFns ...func(*Options)) (*AssociateSigninDelegateGroupsWithAccountOutput, error) {
	if params == nil {
		params = &AssociateSigninDelegateGroupsWithAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSigninDelegateGroupsWithAccount", params, optFns, c.addOperationAssociateSigninDelegateGroupsWithAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSigninDelegateGroupsWithAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSigninDelegateGroupsWithAccountInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The sign-in delegate groups.
	//
	// This member is required.
	SigninDelegateGroups []types.SigninDelegateGroup
}

type AssociateSigninDelegateGroupsWithAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAssociateSigninDelegateGroupsWithAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateSigninDelegateGroupsWithAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateSigninDelegateGroupsWithAccount{}, middleware.After)
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
	if err = addOpAssociateSigninDelegateGroupsWithAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSigninDelegateGroupsWithAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateSigninDelegateGroupsWithAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "AssociateSigninDelegateGroupsWithAccount",
	}
}
