// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replaces an IAM instance profile for the specified running instance. You can use
// this action to change the IAM instance profile that's associated with an
// instance without having to disassociate the existing IAM instance profile first.
// Use DescribeIamInstanceProfileAssociations to get the association ID.
func (c *Client) ReplaceIamInstanceProfileAssociation(ctx context.Context, params *ReplaceIamInstanceProfileAssociationInput, optFns ...func(*Options)) (*ReplaceIamInstanceProfileAssociationOutput, error) {
	if params == nil {
		params = &ReplaceIamInstanceProfileAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReplaceIamInstanceProfileAssociation", params, optFns, c.addOperationReplaceIamInstanceProfileAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReplaceIamInstanceProfileAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReplaceIamInstanceProfileAssociationInput struct {

	// The ID of the existing IAM instance profile association.
	//
	// This member is required.
	AssociationId *string

	// The IAM instance profile.
	//
	// This member is required.
	IamInstanceProfile *types.IamInstanceProfileSpecification
}

type ReplaceIamInstanceProfileAssociationOutput struct {

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *types.IamInstanceProfileAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationReplaceIamInstanceProfileAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpReplaceIamInstanceProfileAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReplaceIamInstanceProfileAssociation{}, middleware.After)
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
	if err = addOpReplaceIamInstanceProfileAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReplaceIamInstanceProfileAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReplaceIamInstanceProfileAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReplaceIamInstanceProfileAssociation",
	}
}
