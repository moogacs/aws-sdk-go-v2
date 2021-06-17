// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a policy to a container image. We recommend that you call the RAM API
// CreateResourceShare
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html)
// to share resources. If you call the Image Builder API PutContainerImagePolicy,
// you must also call the RAM API PromoteResourceShareCreatedFromPolicy
// (https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html)
// in order for the resource to be visible to all principals with whom the resource
// is shared.
func (c *Client) PutContainerRecipePolicy(ctx context.Context, params *PutContainerRecipePolicyInput, optFns ...func(*Options)) (*PutContainerRecipePolicyOutput, error) {
	if params == nil {
		params = &PutContainerRecipePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutContainerRecipePolicy", params, optFns, c.addOperationPutContainerRecipePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutContainerRecipePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutContainerRecipePolicyInput struct {

	// The Amazon Resource Name (ARN) of the container recipe that this policy should
	// be applied to.
	//
	// This member is required.
	ContainerRecipeArn *string

	// The policy to apply to the container recipe.
	//
	// This member is required.
	Policy *string
}

type PutContainerRecipePolicyOutput struct {

	// The Amazon Resource Name (ARN) of the container recipe that this policy was
	// applied to.
	ContainerRecipeArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutContainerRecipePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutContainerRecipePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutContainerRecipePolicy{}, middleware.After)
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
	if err = addOpPutContainerRecipePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutContainerRecipePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutContainerRecipePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "PutContainerRecipePolicy",
	}
}
