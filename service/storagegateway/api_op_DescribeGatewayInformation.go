// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns metadata about a gateway such as its name, network interfaces,
// configured time zone, and the state (whether the gateway is running or not). To
// specify which gateway to describe, use the Amazon Resource Name (ARN) of the
// gateway in your request.
func (c *Client) DescribeGatewayInformation(ctx context.Context, params *DescribeGatewayInformationInput, optFns ...func(*Options)) (*DescribeGatewayInformationOutput, error) {
	if params == nil {
		params = &DescribeGatewayInformationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGatewayInformation", params, optFns, c.addOperationDescribeGatewayInformationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGatewayInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the ID of the gateway.
type DescribeGatewayInformationInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// A JSON object containing the following fields:
type DescribeGatewayInformationOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is used
	// to monitor events in the gateway.
	CloudWatchLogGroupARN *string

	// Date after which this gateway will not receive software updates for new features
	// and bug fixes.
	DeprecationDate *string

	// The ID of the Amazon EC2 instance that was used to launch the gateway.
	Ec2InstanceId *string

	// The AWS Region where the Amazon EC2 instance is located.
	Ec2InstanceRegion *string

	// The type of endpoint for your gateway. Valid Values: STANDARD | FIPS
	EndpointType *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// The unique identifier assigned to your gateway during activation. This ID
	// becomes part of the gateway Amazon Resource Name (ARN), which you use as input
	// for other operations.
	GatewayId *string

	// The name you configured for your gateway.
	GatewayName *string

	// A NetworkInterface array that contains descriptions of the gateway network
	// interfaces.
	GatewayNetworkInterfaces []types.NetworkInterface

	// A value that indicates the operating state of the gateway.
	GatewayState *string

	// A value that indicates the time zone configured for the gateway.
	GatewayTimezone *string

	// The type of the gateway.
	GatewayType *string

	// The type of hypervisor environment used by the host.
	HostEnvironment types.HostEnvironment

	// The date on which the last software update was applied to the gateway. If the
	// gateway has never been updated, this field does not return a value in the
	// response.
	LastSoftwareUpdate *string

	// The date on which an update to the gateway is available. This date is in the
	// time zone of the gateway. If the gateway is not available for an update this
	// field is not returned in the response.
	NextUpdateAvailabilityDate *string

	// Date after which this gateway will not receive software updates for new
	// features.
	SoftwareUpdatesEndDate *string

	// A list of up to 50 tags assigned to the gateway, sorted alphabetically by key
	// name. Each tag is a key-value pair. For a gateway with more than 10 tags
	// assigned, you can view all tags using the ListTagsForResource API operation.
	Tags []types.Tag

	// The configuration settings for the virtual private cloud (VPC) endpoint for your
	// gateway.
	VPCEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeGatewayInformationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGatewayInformation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGatewayInformation{}, middleware.After)
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
	if err = addOpDescribeGatewayInformationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGatewayInformation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGatewayInformation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeGatewayInformation",
	}
}
