// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the source AWS Regions where the current AWS Region can create
// a read replica, copy a DB snapshot from, or replicate automated backups from.
// This API action supports pagination.
func (c *Client) DescribeSourceRegions(ctx context.Context, params *DescribeSourceRegionsInput, optFns ...func(*Options)) (*DescribeSourceRegionsOutput, error) {
	if params == nil {
		params = &DescribeSourceRegionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSourceRegions", params, optFns, c.addOperationDescribeSourceRegionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSourceRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeSourceRegionsInput struct {

	// This parameter isn't currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeSourceRegions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The source AWS Region name. For example, us-east-1. Constraints:
	//
	// * Must specify
	// a valid AWS Region name.
	RegionName *string
}

// Contains the result of a successful invocation of the DescribeSourceRegions
// action.
type DescribeSourceRegionsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// A list of SourceRegion instances that contains each source AWS Region that the
	// current AWS Region can get a read replica or a DB snapshot from.
	SourceRegions []types.SourceRegion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeSourceRegionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSourceRegions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSourceRegions{}, middleware.After)
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
	if err = addOpDescribeSourceRegionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSourceRegions(options.Region), middleware.Before); err != nil {
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

// DescribeSourceRegionsAPIClient is a client that implements the
// DescribeSourceRegions operation.
type DescribeSourceRegionsAPIClient interface {
	DescribeSourceRegions(context.Context, *DescribeSourceRegionsInput, ...func(*Options)) (*DescribeSourceRegionsOutput, error)
}

var _ DescribeSourceRegionsAPIClient = (*Client)(nil)

// DescribeSourceRegionsPaginatorOptions is the paginator options for
// DescribeSourceRegions
type DescribeSourceRegionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSourceRegionsPaginator is a paginator for DescribeSourceRegions
type DescribeSourceRegionsPaginator struct {
	options   DescribeSourceRegionsPaginatorOptions
	client    DescribeSourceRegionsAPIClient
	params    *DescribeSourceRegionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSourceRegionsPaginator returns a new DescribeSourceRegionsPaginator
func NewDescribeSourceRegionsPaginator(client DescribeSourceRegionsAPIClient, params *DescribeSourceRegionsInput, optFns ...func(*DescribeSourceRegionsPaginatorOptions)) *DescribeSourceRegionsPaginator {
	if params == nil {
		params = &DescribeSourceRegionsInput{}
	}

	options := DescribeSourceRegionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSourceRegionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSourceRegionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeSourceRegions page.
func (p *DescribeSourceRegionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSourceRegionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeSourceRegions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeSourceRegions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeSourceRegions",
	}
}
