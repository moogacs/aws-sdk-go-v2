// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the details for the input device
func (c *Client) DescribeInputDevice(ctx context.Context, params *DescribeInputDeviceInput, optFns ...func(*Options)) (*DescribeInputDeviceOutput, error) {
	if params == nil {
		params = &DescribeInputDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInputDevice", params, optFns, c.addOperationDescribeInputDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInputDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeInputDeviceRequest
type DescribeInputDeviceInput struct {

	// The unique ID of this input device. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string
}

// Placeholder documentation for DescribeInputDeviceResponse
type DescribeInputDeviceOutput struct {

	// The unique ARN of the input device.
	Arn *string

	// The state of the connection between the input device and AWS.
	ConnectionState types.InputDeviceConnectionState

	// The status of the action to synchronize the device configuration. If you change
	// the configuration of the input device (for example, the maximum bitrate),
	// MediaLive sends the new data to the device. The device might not update itself
	// immediately. SYNCED means the device has updated its configuration. SYNCING
	// means that it has not updated its configuration.
	DeviceSettingsSyncState types.DeviceSettingsSyncState

	// The status of software on the input device.
	DeviceUpdateStatus types.DeviceUpdateStatus

	// Settings that describe an input device that is type HD.
	HdDeviceSettings *types.InputDeviceHdSettings

	// The unique ID of the input device.
	Id *string

	// The network MAC address of the input device.
	MacAddress *string

	// A name that you specify for the input device.
	Name *string

	// The network settings for the input device.
	NetworkSettings *types.InputDeviceNetworkSettings

	// The unique serial number of the input device.
	SerialNumber *string

	// The type of the input device.
	Type types.InputDeviceType

	// Settings that describe an input device that is type UHD.
	UhdDeviceSettings *types.InputDeviceUhdSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeInputDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInputDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInputDevice{}, middleware.After)
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
	if err = addOpDescribeInputDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInputDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInputDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeInputDevice",
	}
}
