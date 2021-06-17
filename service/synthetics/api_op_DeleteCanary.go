// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Permanently deletes the specified canary. When you delete a canary, resources
// used and created by the canary are not automatically deleted. After you delete a
// canary that you do not intend to use again, you should also delete the
// following:
//
// * The Lambda functions and layers used by this canary. These have
// the prefix cwsyn-MyCanaryName .
//
// * The CloudWatch alarms created for this
// canary. These alarms have a name of Synthetics-SharpDrop-Alarm-MyCanaryName .
//
// *
// Amazon S3 objects and buckets, such as the canary's artifact location.
//
// * IAM
// roles created for the canary. If they were created in the console, these roles
// have the name  role/service-role/CloudWatchSyntheticsRole-MyCanaryName .
//
// *
// CloudWatch Logs log groups created for the canary. These logs groups have the
// name /aws/lambda/cwsyn-MyCanaryName .
//
// Before you delete a canary, you might
// want to use GetCanary to display the information about this canary. Make note of
// the information returned by this operation so that you can delete these
// resources after you delete the canary.
func (c *Client) DeleteCanary(ctx context.Context, params *DeleteCanaryInput, optFns ...func(*Options)) (*DeleteCanaryOutput, error) {
	if params == nil {
		params = &DeleteCanaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCanary", params, optFns, c.addOperationDeleteCanaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCanaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCanaryInput struct {

	// The name of the canary that you want to delete. To find the names of your
	// canaries, use DescribeCanaries
	// (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html).
	//
	// This member is required.
	Name *string
}

type DeleteCanaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteCanaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteCanary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteCanary{}, middleware.After)
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
	if err = addOpDeleteCanaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCanary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCanary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "DeleteCanary",
	}
}
