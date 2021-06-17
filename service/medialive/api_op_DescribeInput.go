// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Produces details about an input
func (c *Client) DescribeInput(ctx context.Context, params *DescribeInputInput, optFns ...func(*Options)) (*DescribeInputOutput, error) {
	if params == nil {
		params = &DescribeInputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInput", params, optFns, c.addOperationDescribeInputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeInputRequest
type DescribeInputInput struct {

	// Unique ID of the input
	//
	// This member is required.
	InputId *string
}

// Placeholder documentation for DescribeInputResponse
type DescribeInputOutput struct {

	// The Unique ARN of the input (generated, immutable).
	Arn *string

	// A list of channel IDs that that input is attached to (currently an input can
	// only be attached to one channel).
	AttachedChannels []string

	// A list of the destinations of the input (PUSH-type).
	Destinations []types.InputDestination

	// The generated ID of the input (unique for user account, immutable).
	Id *string

	// STANDARD - MediaLive expects two sources to be connected to this input. If the
	// channel is also STANDARD, both sources will be ingested. If the channel is
	// SINGLE_PIPELINE, only the first source will be ingested; the second source will
	// always be ignored, even if the first source fails. SINGLE_PIPELINE - You can
	// connect only one source to this input. If the ChannelClass is also
	// SINGLE_PIPELINE, this value is valid. If the ChannelClass is STANDARD, this
	// value is not valid because the channel requires two sources in the input.
	InputClass types.InputClass

	// Settings for the input devices.
	InputDevices []types.InputDeviceSettings

	// A list of IDs for all Inputs which are partners of this one.
	InputPartnerIds []string

	// Certain pull input sources can be dynamic, meaning that they can have their
	// URL's dynamically changes during input switch actions. Presently, this
	// functionality only works with MP4_FILE inputs.
	InputSourceType types.InputSourceType

	// A list of MediaConnect Flows for this input.
	MediaConnectFlows []types.MediaConnectFlow

	// The user-assigned name (This is a mutable value).
	Name *string

	// The Amazon Resource Name (ARN) of the role this input assumes during and after
	// creation.
	RoleArn *string

	// A list of IDs for all the Input Security Groups attached to the input.
	SecurityGroups []string

	// A list of the sources of the input (PULL-type).
	Sources []types.InputSource

	// Placeholder documentation for InputState
	State types.InputState

	// A collection of key-value pairs.
	Tags map[string]string

	// Placeholder documentation for InputType
	Type types.InputType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeInputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInput{}, middleware.After)
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
	if err = addOpDescribeInputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInput(options.Region), middleware.Before); err != nil {
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

// DescribeInputAPIClient is a client that implements the DescribeInput operation.
type DescribeInputAPIClient interface {
	DescribeInput(context.Context, *DescribeInputInput, ...func(*Options)) (*DescribeInputOutput, error)
}

var _ DescribeInputAPIClient = (*Client)(nil)

// InputAttachedWaiterOptions are waiter options for InputAttachedWaiter
type InputAttachedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InputAttachedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, InputAttachedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeInputInput, *DescribeInputOutput, error) (bool, error)
}

// InputAttachedWaiter defines the waiters for InputAttached
type InputAttachedWaiter struct {
	client DescribeInputAPIClient

	options InputAttachedWaiterOptions
}

// NewInputAttachedWaiter constructs a InputAttachedWaiter.
func NewInputAttachedWaiter(client DescribeInputAPIClient, optFns ...func(*InputAttachedWaiterOptions)) *InputAttachedWaiter {
	options := InputAttachedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = inputAttachedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InputAttachedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InputAttached waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *InputAttachedWaiter) Wait(ctx context.Context, params *DescribeInputInput, maxWaitDur time.Duration, optFns ...func(*InputAttachedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeInput(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for InputAttached waiter")
}

func inputAttachedStateRetryable(ctx context.Context, input *DescribeInputInput, output *DescribeInputOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ATTACHED"
		value, ok := pathValue.(types.InputState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.InputState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DETACHED"
		value, ok := pathValue.(types.InputState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.InputState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// InputDeletedWaiterOptions are waiter options for InputDeletedWaiter
type InputDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InputDeletedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, InputDeletedWaiter will use default max delay of 120 seconds. Note that
	// MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeInputInput, *DescribeInputOutput, error) (bool, error)
}

// InputDeletedWaiter defines the waiters for InputDeleted
type InputDeletedWaiter struct {
	client DescribeInputAPIClient

	options InputDeletedWaiterOptions
}

// NewInputDeletedWaiter constructs a InputDeletedWaiter.
func NewInputDeletedWaiter(client DescribeInputAPIClient, optFns ...func(*InputDeletedWaiterOptions)) *InputDeletedWaiter {
	options := InputDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = inputDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InputDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InputDeleted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *InputDeletedWaiter) Wait(ctx context.Context, params *DescribeInputInput, maxWaitDur time.Duration, optFns ...func(*InputDeletedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeInput(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for InputDeleted waiter")
}

func inputDeletedStateRetryable(ctx context.Context, input *DescribeInputInput, output *DescribeInputOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETED"
		value, ok := pathValue.(types.InputState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.InputState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETING"
		value, ok := pathValue.(types.InputState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.InputState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

// InputDetachedWaiterOptions are waiter options for InputDetachedWaiter
type InputDetachedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InputDetachedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, InputDetachedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeInputInput, *DescribeInputOutput, error) (bool, error)
}

// InputDetachedWaiter defines the waiters for InputDetached
type InputDetachedWaiter struct {
	client DescribeInputAPIClient

	options InputDetachedWaiterOptions
}

// NewInputDetachedWaiter constructs a InputDetachedWaiter.
func NewInputDetachedWaiter(client DescribeInputAPIClient, optFns ...func(*InputDetachedWaiterOptions)) *InputDetachedWaiter {
	options := InputDetachedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = inputDetachedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InputDetachedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InputDetached waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *InputDetachedWaiter) Wait(ctx context.Context, params *DescribeInputInput, maxWaitDur time.Duration, optFns ...func(*InputDetachedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeInput(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for InputDetached waiter")
}

func inputDetachedStateRetryable(ctx context.Context, input *DescribeInputInput, output *DescribeInputOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DETACHED"
		value, ok := pathValue.(types.InputState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.InputState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATING"
		value, ok := pathValue.(types.InputState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.InputState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ATTACHED"
		value, ok := pathValue.(types.InputState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.InputState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err != nil {
		var errorType *types.InternalServerErrorException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeInput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeInput",
	}
}
