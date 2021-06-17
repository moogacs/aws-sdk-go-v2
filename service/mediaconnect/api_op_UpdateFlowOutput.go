// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing flow output.
func (c *Client) UpdateFlowOutput(ctx context.Context, params *UpdateFlowOutputInput, optFns ...func(*Options)) (*UpdateFlowOutputOutput, error) {
	if params == nil {
		params = &UpdateFlowOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlowOutput", params, optFns, c.addOperationUpdateFlowOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The fields that you want to update in the output.
type UpdateFlowOutputInput struct {

	// The flow that is associated with the output that you want to update.
	//
	// This member is required.
	FlowArn *string

	// The ARN of the output that you want to update.
	//
	// This member is required.
	OutputArn *string

	// The range of IP addresses that should be allowed to initiate output requests to
	// this flow. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList []string

	// A description of the output. This description appears only on the AWS Elemental
	// MediaConnect console and will not be seen by the end user.
	Description *string

	// The IP address where you want to send the output.
	Destination *string

	// The type of key used for the encryption. If no keyType is provided, the service
	// will use the default setting (static-key).
	Encryption *types.UpdateEncryption

	// The maximum latency in milliseconds for Zixi-based streams.
	MaxLatency int32

	// The media streams that are associated with the output, and the parameters for
	// those associations.
	MediaStreamOutputConfigurations []types.MediaStreamOutputConfigurationRequest

	// The minimum latency in milliseconds for SRT-based streams. In streams that use
	// the SRT protocol, this value that you set on your MediaConnect source or output
	// represents the minimal potential latency of that connection. The latency of the
	// stream is set to the highest number between the sender’s minimum latency and the
	// receiver’s minimum latency.
	MinLatency int32

	// The port to use when content is distributed to this output.
	Port int32

	// The protocol to use for the output.
	Protocol types.Protocol

	// The remote ID for the Zixi-pull stream.
	RemoteId *string

	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency int32

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string

	// The name of the VPC interface attachment to use for this output.
	VpcInterfaceAttachment *types.VpcInterfaceAttachment
}

type UpdateFlowOutputOutput struct {

	// The ARN of the flow that is associated with the updated output.
	FlowArn *string

	// The new settings of the output that you updated.
	Output *types.Output

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateFlowOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlowOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlowOutput{}, middleware.After)
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
	if err = addOpUpdateFlowOutputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlowOutput(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFlowOutput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "UpdateFlowOutput",
	}
}
