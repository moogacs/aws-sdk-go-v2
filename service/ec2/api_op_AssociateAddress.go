// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an Elastic IP address, or carrier IP address (for instances that are
// in subnets in Wavelength Zones) with an instance or a network interface. Before
// you can use an Elastic IP address, you must allocate it to your account. An
// Elastic IP address is for use in either the EC2-Classic platform or in a VPC.
// For more information, see Elastic IP Addresses
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
// in the Amazon Elastic Compute Cloud User Guide. [EC2-Classic, VPC in an
// EC2-VPC-only account] If the Elastic IP address is already associated with a
// different instance, it is disassociated from that instance and associated with
// the specified instance. If you associate an Elastic IP address with an instance
// that has an existing Elastic IP address, the existing address is disassociated
// from the instance, but remains allocated to your account. [VPC in an EC2-Classic
// account] If you don't specify a private IP address, the Elastic IP address is
// associated with the primary IP address. If the Elastic IP address is already
// associated with a different instance or a network interface, you get an error
// unless you allow reassociation. You cannot associate an Elastic IP address with
// an instance or network interface that has an existing Elastic IP address.
// [Subnets in Wavelength Zones] You can associate an IP address from the
// telecommunication carrier to the instance or network interface. You cannot
// associate an Elastic IP address with an interface in a different network border
// group. This is an idempotent operation. If you perform the operation more than
// once, Amazon EC2 doesn't return an error, and you may be charged for each time
// the Elastic IP address is remapped to the same instance. For more information,
// see the Elastic IP Addresses section of Amazon EC2 Pricing
// (http://aws.amazon.com/ec2/pricing/).
func (c *Client) AssociateAddress(ctx context.Context, params *AssociateAddressInput, optFns ...func(*Options)) (*AssociateAddressOutput, error) {
	if params == nil {
		params = &AssociateAddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateAddress", params, optFns, c.addOperationAssociateAddressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateAddressInput struct {

	// [EC2-VPC] The allocation ID. This is required for EC2-VPC.
	AllocationId *string

	// [EC2-VPC] For a VPC in an EC2-Classic account, specify true to allow an Elastic
	// IP address that is already associated with an instance or network interface to
	// be reassociated with the specified instance or network interface. Otherwise, the
	// operation fails. In a VPC in an EC2-VPC-only account, reassociation is
	// automatic, therefore you can specify false to ensure the operation fails if the
	// Elastic IP address is already associated with another resource.
	AllowReassociation *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the instance. The instance must have exactly one attached network
	// interface. For EC2-VPC, you can specify either the instance ID or the network
	// interface ID, but not both. For EC2-Classic, you must specify an instance ID and
	// the instance must be in the running state.
	InstanceId *string

	// [EC2-VPC] The ID of the network interface. If the instance has more than one
	// network interface, you must specify a network interface ID. For EC2-VPC, you can
	// specify either the instance ID or the network interface ID, but not both.
	NetworkInterfaceId *string

	// [EC2-VPC] The primary or secondary private IP address to associate with the
	// Elastic IP address. If no private IP address is specified, the Elastic IP
	// address is associated with the primary private IP address.
	PrivateIpAddress *string

	// [EC2-Classic] The Elastic IP address to associate with the instance. This is
	// required for EC2-Classic.
	PublicIp *string
}

type AssociateAddressOutput struct {

	// [EC2-VPC] The ID that represents the association of the Elastic IP address with
	// an instance.
	AssociationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAssociateAddressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateAddress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateAddress{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateAddress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateAddress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateAddress",
	}
}
