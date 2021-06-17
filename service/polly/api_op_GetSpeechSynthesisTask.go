// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a specific SpeechSynthesisTask object based on its TaskID. This object
// contains information about the given speech synthesis task, including the status
// of the task, and a link to the S3 bucket containing the output of the task.
func (c *Client) GetSpeechSynthesisTask(ctx context.Context, params *GetSpeechSynthesisTaskInput, optFns ...func(*Options)) (*GetSpeechSynthesisTaskOutput, error) {
	if params == nil {
		params = &GetSpeechSynthesisTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSpeechSynthesisTask", params, optFns, c.addOperationGetSpeechSynthesisTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSpeechSynthesisTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSpeechSynthesisTaskInput struct {

	// The Amazon Polly generated identifier for a speech synthesis task.
	//
	// This member is required.
	TaskId *string
}

type GetSpeechSynthesisTaskOutput struct {

	// SynthesisTask object that provides information from the requested task,
	// including output format, creation time, task status, and so on.
	SynthesisTask *types.SynthesisTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetSpeechSynthesisTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSpeechSynthesisTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSpeechSynthesisTask{}, middleware.After)
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
	if err = addOpGetSpeechSynthesisTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSpeechSynthesisTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSpeechSynthesisTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "GetSpeechSynthesisTask",
	}
}
