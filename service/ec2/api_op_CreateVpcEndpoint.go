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

// Creates a VPC endpoint for a specified service. An endpoint enables you to
// create a private connection between your VPC and the service. The service may be
// provided by AWS, an AWS Marketplace Partner, or another AWS account. For more
// information, see VPC Endpoints
// (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints.html) in the
// Amazon Virtual Private Cloud User Guide. A gateway endpoint serves as a target
// for a route in your route table for traffic destined for the AWS service. You
// can specify an endpoint policy to attach to the endpoint, which will control
// access to the service from your VPC. You can also specify the VPC route tables
// that use the endpoint. An interface endpoint is a network interface in your
// subnet that serves as an endpoint for communicating with the specified service.
// You can specify the subnets in which to create an endpoint, and the security
// groups to associate with the endpoint network interface. A GatewayLoadBalancer
// endpoint is a network interface in your subnet that serves an endpoint for
// communicating with a Gateway Load Balancer that you've configured as a VPC
// endpoint service. Use DescribeVpcEndpointServices to get a list of supported
// services.
func (c *Client) CreateVpcEndpoint(ctx context.Context, params *CreateVpcEndpointInput, optFns ...func(*Options)) (*CreateVpcEndpointOutput, error) {
	if params == nil {
		params = &CreateVpcEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcEndpoint", params, optFns, c.addOperationCreateVpcEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateVpcEndpoint.
type CreateVpcEndpointInput struct {

	// The service name. To get a list of available services, use the
	// DescribeVpcEndpointServices request, or get the name from the service provider.
	//
	// This member is required.
	ServiceName *string

	// The ID of the VPC in which the endpoint will be used.
	//
	// This member is required.
	VpcId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to Ensure Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// (Interface and gateway endpoints) A policy to attach to the endpoint that
	// controls access to the service. The policy must be in valid JSON format. If this
	// parameter is not specified, we attach a default policy that allows full access
	// to the service.
	PolicyDocument *string

	// (Interface endpoint) Indicates whether to associate a private hosted zone with
	// the specified VPC. The private hosted zone contains a record set for the default
	// public DNS name for the service for the Region (for example,
	// kinesis.us-east-1.amazonaws.com), which resolves to the private IP addresses of
	// the endpoint network interfaces in the VPC. This enables you to make requests to
	// the default public DNS name for the service instead of the public DNS names that
	// are automatically generated by the VPC endpoint service. To use a private hosted
	// zone, you must set the following VPC attributes to true: enableDnsHostnames and
	// enableDnsSupport. Use ModifyVpcAttribute to set the VPC attributes. Default:
	// true
	PrivateDnsEnabled *bool

	// (Gateway endpoint) One or more route table IDs.
	RouteTableIds []string

	// (Interface endpoint) The ID of one or more security groups to associate with the
	// endpoint network interface.
	SecurityGroupIds []string

	// (Interface and Gateway Load Balancer endpoints) The ID of one or more subnets in
	// which to create an endpoint network interface. For a Gateway Load Balancer
	// endpoint, you can specify one subnet only.
	SubnetIds []string

	// The tags to associate with the endpoint.
	TagSpecifications []types.TagSpecification

	// The type of endpoint. Default: Gateway
	VpcEndpointType types.VpcEndpointType
}

// Contains the output of CreateVpcEndpoint.
type CreateVpcEndpointOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// Information about the endpoint.
	VpcEndpoint *types.VpcEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateVpcEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpcEndpoint{}, middleware.After)
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
	if err = addOpCreateVpcEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpcEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVpcEndpoint",
	}
}
