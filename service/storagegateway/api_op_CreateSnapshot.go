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

// Initiates a snapshot of a volume. AWS Storage Gateway provides the ability to
// back up point-in-time snapshots of your data to Amazon Simple Storage (Amazon
// S3) for durable off-site recovery, and also import the data to an Amazon Elastic
// Block Store (EBS) volume in Amazon Elastic Compute Cloud (EC2). You can take
// snapshots of your gateway volume on a scheduled or ad hoc basis. This API
// enables you to take an ad hoc snapshot. For more information, see Editing a
// snapshot schedule
// (https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-volumes.html#SchedulingSnapshot).
// In the CreateSnapshot request, you identify the volume by providing its Amazon
// Resource Name (ARN). You must also provide description for the snapshot. When
// AWS Storage Gateway takes the snapshot of specified volume, the snapshot and
// description appears in the AWS Storage Gateway console. In response, AWS Storage
// Gateway returns you a snapshot ID. You can use this snapshot ID to check the
// snapshot progress or later use it when you want to create a volume from a
// snapshot. This operation is only supported in stored and cached volume gateway
// type. To list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, see DescribeSnapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html)
// or DeleteSnapshot
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html)
// in the Amazon Elastic Compute Cloud API Reference. Volume and snapshot IDs are
// changing to a longer length ID format. For more information, see the important
// note on the Welcome
// (https://docs.aws.amazon.com/storagegateway/latest/APIReference/Welcome.html)
// page.
func (c *Client) CreateSnapshot(ctx context.Context, params *CreateSnapshotInput, optFns ...func(*Options)) (*CreateSnapshotOutput, error) {
	if params == nil {
		params = &CreateSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSnapshot", params, optFns, c.addOperationCreateSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:
//
// *
// CreateSnapshotInput$SnapshotDescription
//
// * CreateSnapshotInput$VolumeARN
type CreateSnapshotInput struct {

	// Textual description of the snapshot that appears in the Amazon EC2 console,
	// Elastic Block Store snapshots panel in the Description field, and in the AWS
	// Storage Gateway snapshot Details pane, Description field.
	//
	// This member is required.
	SnapshotDescription *string

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation to
	// return a list of gateway volumes.
	//
	// This member is required.
	VolumeARN *string

	// A list of up to 50 tags that can be assigned to a snapshot. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag
}

// A JSON object containing the following fields:
type CreateSnapshotOutput struct {

	// The snapshot ID that is used to refer to the snapshot in future operations such
	// as describing snapshots (Amazon Elastic Compute Cloud API DescribeSnapshots) or
	// creating a volume from a snapshot (CreateStorediSCSIVolume).
	SnapshotId *string

	// The Amazon Resource Name (ARN) of the volume of which the snapshot was taken.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSnapshot{}, middleware.After)
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
	if err = addOpCreateSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateSnapshot",
	}
}
