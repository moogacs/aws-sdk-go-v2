// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Moves the specified instances into the standby state. If you choose to decrement
// the desired capacity of the Auto Scaling group, the instances can enter standby
// as long as the desired capacity of the Auto Scaling group after the instances
// are placed into standby is equal to or greater than the minimum capacity of the
// group. If you choose not to decrement the desired capacity of the Auto Scaling
// group, the Auto Scaling group launches new instances to replace the instances on
// standby. For more information, see Temporarily removing instances from your Auto
// Scaling group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) EnterStandby(ctx context.Context, params *EnterStandbyInput, optFns ...func(*Options)) (*EnterStandbyOutput, error) {
	if params == nil {
		params = &EnterStandbyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnterStandby", params, optFns, c.addOperationEnterStandbyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnterStandbyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnterStandbyInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Indicates whether to decrement the desired capacity of the Auto Scaling group by
	// the number of instances moved to Standby mode.
	//
	// This member is required.
	ShouldDecrementDesiredCapacity *bool

	// The IDs of the instances. You can specify up to 20 instances.
	InstanceIds []string
}

type EnterStandbyOutput struct {

	// The activities related to moving instances into Standby mode.
	Activities []types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationEnterStandbyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnterStandby{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnterStandby{}, middleware.After)
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
	if err = addOpEnterStandbyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnterStandby(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnterStandby(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "EnterStandby",
	}
}
