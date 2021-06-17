// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Analyzes and accumulates test report values for the specified test reports.
func (c *Client) GetReportGroupTrend(ctx context.Context, params *GetReportGroupTrendInput, optFns ...func(*Options)) (*GetReportGroupTrendOutput, error) {
	if params == nil {
		params = &GetReportGroupTrendInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReportGroupTrend", params, optFns, c.addOperationGetReportGroupTrendMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReportGroupTrendOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReportGroupTrendInput struct {

	// The ARN of the report group that contains the reports to analyze.
	//
	// This member is required.
	ReportGroupArn *string

	// The test report value to accumulate. This must be one of the following values:
	// Test reports: DURATION Accumulate the test run times for the specified reports.
	// PASS_RATE Accumulate the percentage of tests that passed for the specified test
	// reports. TOTAL Accumulate the total number of tests for the specified test
	// reports. Code coverage reports: BRANCH_COVERAGE Accumulate the branch coverage
	// percentages for the specified test reports. BRANCHES_COVERED Accumulate the
	// branches covered values for the specified test reports. BRANCHES_MISSED
	// Accumulate the branches missed values for the specified test reports.
	// LINE_COVERAGE Accumulate the line coverage percentages for the specified test
	// reports. LINES_COVERED Accumulate the lines covered values for the specified
	// test reports. LINES_MISSED Accumulate the lines not covered values for the
	// specified test reports.
	//
	// This member is required.
	TrendField types.ReportGroupTrendFieldType

	// The number of reports to analyze. This operation always retrieves the most
	// recent reports. If this parameter is omitted, the most recent 100 reports are
	// analyzed.
	NumOfReports *int32
}

type GetReportGroupTrendOutput struct {

	// An array that contains the raw data for each report.
	RawData []types.ReportWithRawData

	// Contains the accumulated trend data.
	Stats *types.ReportGroupTrendStats

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetReportGroupTrendMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetReportGroupTrend{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReportGroupTrend{}, middleware.After)
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
	if err = addOpGetReportGroupTrendValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReportGroupTrend(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetReportGroupTrend(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "GetReportGroupTrend",
	}
}
