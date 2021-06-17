// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a snapshot schedule that can be associated to a cluster and which
// overrides the default system backup schedule.
func (c *Client) CreateSnapshotSchedule(ctx context.Context, params *CreateSnapshotScheduleInput, optFns ...func(*Options)) (*CreateSnapshotScheduleOutput, error) {
	if params == nil {
		params = &CreateSnapshotScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSnapshotSchedule", params, optFns, c.addOperationCreateSnapshotScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSnapshotScheduleInput struct {

	//
	DryRun *bool

	//
	NextInvocations *int32

	// The definition of the snapshot schedule. The definition is made up of schedule
	// expressions, for example "cron(30 12 *)" or "rate(12 hours)".
	ScheduleDefinitions []string

	// The description of the snapshot schedule.
	ScheduleDescription *string

	// A unique identifier for a snapshot schedule. Only alphanumeric characters are
	// allowed for the identifier.
	ScheduleIdentifier *string

	// An optional set of tags you can use to search for the schedule.
	Tags []types.Tag
}

// Describes a snapshot schedule. You can set a regular interval for creating
// snapshots of a cluster. You can also schedule snapshots for specific dates.
type CreateSnapshotScheduleOutput struct {

	// The number of clusters associated with the schedule.
	AssociatedClusterCount *int32

	// A list of clusters associated with the schedule. A maximum of 100 clusters is
	// returned.
	AssociatedClusters []types.ClusterAssociatedToSchedule

	//
	NextInvocations []time.Time

	// A list of ScheduleDefinitions.
	ScheduleDefinitions []string

	// The description of the schedule.
	ScheduleDescription *string

	// A unique identifier for the schedule.
	ScheduleIdentifier *string

	// An optional set of tags describing the schedule.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateSnapshotScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateSnapshotSchedule{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshotSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSnapshotSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateSnapshotSchedule",
	}
}
