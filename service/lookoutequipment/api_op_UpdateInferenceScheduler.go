// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an inference scheduler.
func (c *Client) UpdateInferenceScheduler(ctx context.Context, params *UpdateInferenceSchedulerInput, optFns ...func(*Options)) (*UpdateInferenceSchedulerOutput, error) {
	if params == nil {
		params = &UpdateInferenceSchedulerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInferenceScheduler", params, optFns, c.addOperationUpdateInferenceSchedulerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInferenceSchedulerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInferenceSchedulerInput struct {

	// The name of the inference scheduler to be updated.
	//
	// This member is required.
	InferenceSchedulerName *string

	// > A period of time (in minutes) by which inference on the data is delayed after
	// the data starts. For instance, if you select an offset delay time of five
	// minutes, inference will not begin on the data until the first data measurement
	// after the five minute mark. For example, if five minutes is selected, the
	// inference scheduler will wake up at the configured frequency with the additional
	// five minute delay time to check the customer S3 bucket. The customer can upload
	// data at the same frequency and they don't need to stop and restart the scheduler
	// when uploading new data.
	DataDelayOffsetInMinutes *int64

	// Specifies information for the input data for the inference scheduler, including
	// delimiter, format, and dataset location.
	DataInputConfiguration *types.InferenceInputConfiguration

	// Specifies information for the output results from the inference scheduler,
	// including the output S3 location.
	DataOutputConfiguration *types.InferenceOutputConfiguration

	// How often data is uploaded to the source S3 bucket for the input data. The value
	// chosen is the length of time between data uploads. For instance, if you select 5
	// minutes, Amazon Lookout for Equipment will upload the real-time data to the
	// source bucket once every 5 minutes. This frequency also determines how often
	// Amazon Lookout for Equipment starts a scheduled inference on your data. In this
	// example, it starts once every 5 minutes.
	DataUploadFrequency types.DataUploadFrequency

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source for the inference scheduler.
	RoleArn *string
}

type UpdateInferenceSchedulerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateInferenceSchedulerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateInferenceScheduler{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateInferenceScheduler{}, middleware.After)
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
	if err = addOpUpdateInferenceSchedulerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInferenceScheduler(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInferenceScheduler(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "UpdateInferenceScheduler",
	}
}
