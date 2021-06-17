// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// DescribeExportConfigurations is deprecated. Use DescribeImportTasks
// (https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeExportTasks.html),
// instead.
//
// Deprecated: This operation has been deprecated.
func (c *Client) DescribeExportConfigurations(ctx context.Context, params *DescribeExportConfigurationsInput, optFns ...func(*Options)) (*DescribeExportConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeExportConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExportConfigurations", params, optFns, c.addOperationDescribeExportConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportConfigurationsInput struct {

	// A list of continuous export IDs to search for.
	ExportIds []string

	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	MaxResults int32

	// The token from the previous call to describe-export-tasks.
	NextToken *string
}

type DescribeExportConfigurationsOutput struct {

	//
	ExportsInfo []types.ExportInfo

	// The token from the previous call to describe-export-tasks.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeExportConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExportConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExportConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExportConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeExportConfigurations",
	}
}
