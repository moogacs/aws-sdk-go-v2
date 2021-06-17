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

// Describes the available CloudWatch metrics for Amazon EC2 Auto Scaling. The
// GroupStandbyInstances metric is not returned by default. You must explicitly
// request this metric when calling the EnableMetricsCollection API.
func (c *Client) DescribeMetricCollectionTypes(ctx context.Context, params *DescribeMetricCollectionTypesInput, optFns ...func(*Options)) (*DescribeMetricCollectionTypesOutput, error) {
	if params == nil {
		params = &DescribeMetricCollectionTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMetricCollectionTypes", params, optFns, c.addOperationDescribeMetricCollectionTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMetricCollectionTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMetricCollectionTypesInput struct {
}

type DescribeMetricCollectionTypesOutput struct {

	// The granularities for the metrics.
	Granularities []types.MetricGranularityType

	// One or more metrics.
	Metrics []types.MetricCollectionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeMetricCollectionTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeMetricCollectionTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeMetricCollectionTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMetricCollectionTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMetricCollectionTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeMetricCollectionTypes",
	}
}
