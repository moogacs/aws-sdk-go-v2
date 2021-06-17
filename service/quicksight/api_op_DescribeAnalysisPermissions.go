// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides the read and write permissions for an analysis.
func (c *Client) DescribeAnalysisPermissions(ctx context.Context, params *DescribeAnalysisPermissionsInput, optFns ...func(*Options)) (*DescribeAnalysisPermissionsOutput, error) {
	if params == nil {
		params = &DescribeAnalysisPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAnalysisPermissions", params, optFns, c.addOperationDescribeAnalysisPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAnalysisPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAnalysisPermissionsInput struct {

	// The ID of the analysis whose permissions you're describing. The ID is part of
	// the analysis URL.
	//
	// This member is required.
	AnalysisId *string

	// The ID of the AWS account that contains the analysis whose permissions you're
	// describing. You must be using the AWS account that the analysis is in.
	//
	// This member is required.
	AwsAccountId *string
}

type DescribeAnalysisPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the analysis whose permissions you're
	// describing.
	AnalysisArn *string

	// The ID of the analysis whose permissions you're describing.
	AnalysisId *string

	// A structure that describes the principals and the resource-level permissions on
	// an analysis.
	Permissions []types.ResourcePermission

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeAnalysisPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAnalysisPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAnalysisPermissions{}, middleware.After)
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
	if err = addOpDescribeAnalysisPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAnalysisPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAnalysisPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DescribeAnalysisPermissions",
	}
}
