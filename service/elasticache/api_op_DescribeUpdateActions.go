// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details of the update actions
func (c *Client) DescribeUpdateActions(ctx context.Context, params *DescribeUpdateActionsInput, optFns ...func(*Options)) (*DescribeUpdateActionsOutput, error) {
	if params == nil {
		params = &DescribeUpdateActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUpdateActions", params, optFns, c.addOperationDescribeUpdateActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUpdateActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUpdateActionsInput struct {

	// The cache cluster IDs
	CacheClusterIds []string

	// The Elasticache engine to which the update applies. Either Redis or Memcached
	Engine *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response
	MaxRecords *int32

	// The replication group IDs
	ReplicationGroupIds []string

	// The unique ID of the service update
	ServiceUpdateName *string

	// The status of the service update
	ServiceUpdateStatus []types.ServiceUpdateStatus

	// The range of time specified to search for service updates that are in available
	// status
	ServiceUpdateTimeRange *types.TimeRangeFilter

	// Dictates whether to include node level update status in the response
	ShowNodeLevelUpdateStatus *bool

	// The status of the update action.
	UpdateActionStatus []types.UpdateActionStatus
}

type DescribeUpdateActionsOutput struct {

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// Returns a list of update actions
	UpdateActions []types.UpdateAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeUpdateActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeUpdateActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeUpdateActions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUpdateActions(options.Region), middleware.Before); err != nil {
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

// DescribeUpdateActionsAPIClient is a client that implements the
// DescribeUpdateActions operation.
type DescribeUpdateActionsAPIClient interface {
	DescribeUpdateActions(context.Context, *DescribeUpdateActionsInput, ...func(*Options)) (*DescribeUpdateActionsOutput, error)
}

var _ DescribeUpdateActionsAPIClient = (*Client)(nil)

// DescribeUpdateActionsPaginatorOptions is the paginator options for
// DescribeUpdateActions
type DescribeUpdateActionsPaginatorOptions struct {
	// The maximum number of records to include in the response
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeUpdateActionsPaginator is a paginator for DescribeUpdateActions
type DescribeUpdateActionsPaginator struct {
	options   DescribeUpdateActionsPaginatorOptions
	client    DescribeUpdateActionsAPIClient
	params    *DescribeUpdateActionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeUpdateActionsPaginator returns a new DescribeUpdateActionsPaginator
func NewDescribeUpdateActionsPaginator(client DescribeUpdateActionsAPIClient, params *DescribeUpdateActionsInput, optFns ...func(*DescribeUpdateActionsPaginatorOptions)) *DescribeUpdateActionsPaginator {
	if params == nil {
		params = &DescribeUpdateActionsInput{}
	}

	options := DescribeUpdateActionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeUpdateActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeUpdateActionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeUpdateActions page.
func (p *DescribeUpdateActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeUpdateActionsOutput, error) {
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

	result, err := p.client.DescribeUpdateActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeUpdateActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeUpdateActions",
	}
}
