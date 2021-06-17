// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes one or more alarms to the snooze mode. The alarms change to the
// SNOOZE_DISABLED state after you set them to the snooze mode.
func (c *Client) BatchSnoozeAlarm(ctx context.Context, params *BatchSnoozeAlarmInput, optFns ...func(*Options)) (*BatchSnoozeAlarmOutput, error) {
	if params == nil {
		params = &BatchSnoozeAlarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchSnoozeAlarm", params, optFns, c.addOperationBatchSnoozeAlarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchSnoozeAlarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchSnoozeAlarmInput struct {

	// The list of snooze action requests. You can specify up to 10 requests per
	// operation.
	//
	// This member is required.
	SnoozeActionRequests []types.SnoozeAlarmActionRequest
}

type BatchSnoozeAlarmOutput struct {

	// A list of errors associated with the request, or null if there are no errors.
	// Each error entry contains an entry ID that helps you identify the entry that
	// failed.
	ErrorEntries []types.BatchAlarmActionErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationBatchSnoozeAlarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchSnoozeAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchSnoozeAlarm{}, middleware.After)
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
	if err = addOpBatchSnoozeAlarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchSnoozeAlarm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchSnoozeAlarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ioteventsdata",
		OperationName: "BatchSnoozeAlarm",
	}
}
