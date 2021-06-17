// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists private workforce information, including workforce name, Amazon Resource
// Name (ARN), and, if applicable, allowed IP address ranges (CIDRs
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)). Allowable
// IP address ranges are the IP addresses that workers can use to access tasks.
// This operation applies only to private workforces.
func (c *Client) DescribeWorkforce(ctx context.Context, params *DescribeWorkforceInput, optFns ...func(*Options)) (*DescribeWorkforceOutput, error) {
	if params == nil {
		params = &DescribeWorkforceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkforce", params, optFns, c.addOperationDescribeWorkforceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkforceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkforceInput struct {

	// The name of the private workforce whose access you want to restrict.
	// WorkforceName is automatically set to default when a workforce is created and
	// cannot be modified.
	//
	// This member is required.
	WorkforceName *string
}

type DescribeWorkforceOutput struct {

	// A single private workforce, which is automatically created when you create your
	// first private work team. You can create one private work force in each AWS
	// Region. By default, any workforce-related API operation used in a specific
	// region will apply to the workforce created in that region. To learn how to
	// create a private workforce, see Create a Private Workforce
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private.html).
	//
	// This member is required.
	Workforce *types.Workforce

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeWorkforceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkforce{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkforce{}, middleware.After)
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
	if err = addOpDescribeWorkforceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkforce(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorkforce(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeWorkforce",
	}
}
