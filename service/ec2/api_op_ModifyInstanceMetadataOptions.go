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

// Modify the instance metadata parameters on a running or stopped instance. When
// you modify the parameters on a stopped instance, they are applied when the
// instance is started. When you modify the parameters on a running instance, the
// API responds with a state of “pending”. After the parameter modifications are
// successfully applied to the instance, the state of the modifications changes
// from “pending” to “applied” in subsequent describe-instances API calls. For more
// information, see Instance metadata and user data
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)
// in the Amazon EC2 User Guide.
func (c *Client) ModifyInstanceMetadataOptions(ctx context.Context, params *ModifyInstanceMetadataOptionsInput, optFns ...func(*Options)) (*ModifyInstanceMetadataOptionsOutput, error) {
	if params == nil {
		params = &ModifyInstanceMetadataOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceMetadataOptions", params, optFns, c.addOperationModifyInstanceMetadataOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceMetadataOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceMetadataOptionsInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// This parameter enables or disables the HTTP metadata endpoint on your instances.
	// If the parameter is not specified, the existing state is maintained. If you
	// specify a value of disabled, you will not be able to access your instance
	// metadata.
	HttpEndpoint types.InstanceMetadataEndpointState

	// The desired HTTP PUT response hop limit for instance metadata requests. The
	// larger the number, the further instance metadata requests can travel. If no
	// parameter is specified, the existing state is maintained. Possible values:
	// Integers from 1 to 64
	HttpPutResponseHopLimit *int32

	// The state of token usage for your instance metadata requests. If the parameter
	// is not specified in the request, the default state is optional. If the state is
	// optional, you can choose to retrieve instance metadata with or without a signed
	// token header on your request. If you retrieve the IAM role credentials without a
	// token, the version 1.0 role credentials are returned. If you retrieve the IAM
	// role credentials using a valid signed token, the version 2.0 role credentials
	// are returned. If the state is required, you must send a signed token header with
	// any instance metadata retrieval requests. In this state, retrieving the IAM role
	// credential always returns the version 2.0 credentials; the version 1.0
	// credentials are not available.
	HttpTokens types.HttpTokensState
}

type ModifyInstanceMetadataOptionsOutput struct {

	// The ID of the instance.
	InstanceId *string

	// The metadata options for the instance.
	InstanceMetadataOptions *types.InstanceMetadataOptionsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationModifyInstanceMetadataOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceMetadataOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceMetadataOptions{}, middleware.After)
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
	if err = addOpModifyInstanceMetadataOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceMetadataOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceMetadataOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyInstanceMetadataOptions",
	}
}
