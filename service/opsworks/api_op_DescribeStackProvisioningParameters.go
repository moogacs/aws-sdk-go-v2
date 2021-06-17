// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests a description of a stack's provisioning parameters. Required
// Permissions: To use this action, an IAM user must have a Show, Deploy, or Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information about user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribeStackProvisioningParameters(ctx context.Context, params *DescribeStackProvisioningParametersInput, optFns ...func(*Options)) (*DescribeStackProvisioningParametersOutput, error) {
	if params == nil {
		params = &DescribeStackProvisioningParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStackProvisioningParameters", params, optFns, c.addOperationDescribeStackProvisioningParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStackProvisioningParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStackProvisioningParametersInput struct {

	// The stack ID.
	//
	// This member is required.
	StackId *string
}

// Contains the response to a DescribeStackProvisioningParameters request.
type DescribeStackProvisioningParametersOutput struct {

	// The AWS OpsWorks Stacks agent installer's URL.
	AgentInstallerUrl *string

	// An embedded object that contains the provisioning parameters.
	Parameters map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeStackProvisioningParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStackProvisioningParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStackProvisioningParameters{}, middleware.After)
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
	if err = addOpDescribeStackProvisioningParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackProvisioningParameters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStackProvisioningParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribeStackProvisioningParameters",
	}
}
