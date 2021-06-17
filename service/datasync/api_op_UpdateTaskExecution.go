// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates execution of a task. You can modify bandwidth throttling for a task
// execution that is running or queued. For more information, see Adjusting
// Bandwidth Throttling for a Task Execution
// (https://docs.aws.amazon.com/datasync/latest/userguide/working-with-task-executions.html#adjust-bandwidth-throttling).
// The only Option that can be modified by UpdateTaskExecution is BytesPerSecond
// (https://docs.aws.amazon.com/datasync/latest/userguide/API_Options.html#DataSync-Type-Options-BytesPerSecond).
func (c *Client) UpdateTaskExecution(ctx context.Context, params *UpdateTaskExecutionInput, optFns ...func(*Options)) (*UpdateTaskExecutionOutput, error) {
	if params == nil {
		params = &UpdateTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTaskExecution", params, optFns, c.addOperationUpdateTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTaskExecutionInput struct {

	// Represents the options that are available to control the behavior of a
	// StartTaskExecution
	// (https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html)
	// operation. Behavior includes preserving metadata such as user ID (UID), group ID
	// (GID), and file permissions, and also overwriting files in the destination, data
	// integrity verification, and so on. A task has a set of default options
	// associated with it. If you don't specify an option in StartTaskExecution
	// (https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html),
	// the default value is used. You can override the defaults options on each task
	// execution by specifying an overriding Options value to StartTaskExecution
	// (https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html).
	//
	// This member is required.
	Options *types.Options

	// The Amazon Resource Name (ARN) of the specific task execution that is being
	// updated.
	//
	// This member is required.
	TaskExecutionArn *string
}

type UpdateTaskExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateTaskExecution{}, middleware.After)
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
	if err = addOpUpdateTaskExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTaskExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTaskExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "UpdateTaskExecution",
	}
}
