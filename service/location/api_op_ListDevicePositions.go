// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the latest device positions for requested devices.
func (c *Client) ListDevicePositions(ctx context.Context, params *ListDevicePositionsInput, optFns ...func(*Options)) (*ListDevicePositionsOutput, error) {
	if params == nil {
		params = &ListDevicePositionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevicePositions", params, optFns, c.addOperationListDevicePositionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevicePositionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevicePositionsInput struct {

	// The tracker resource containing the requested devices.
	//
	// This member is required.
	TrackerName *string

	// An optional limit for the number of entries returned in a single call. Default
	// value: 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the response.
	// If no token is provided, the default page is the first page. Default value: null
	NextToken *string
}

type ListDevicePositionsOutput struct {

	// Contains details about each device's last known position. These details includes
	// the device ID, the time when the position was sampled on the device, the time
	// that the service received the update, and the most recent coordinates.
	//
	// This member is required.
	Entries []types.ListDevicePositionsResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListDevicePositionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDevicePositions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevicePositions{}, middleware.After)
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
	if err = addOpListDevicePositionsValidationMiddleware(stack); err != nil {
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

// ListDevicePositionsAPIClient is a client that implements the ListDevicePositions
// operation.
type ListDevicePositionsAPIClient interface {
	ListDevicePositions(context.Context, *ListDevicePositionsInput, ...func(*Options)) (*ListDevicePositionsOutput, error)
}

var _ ListDevicePositionsAPIClient = (*Client)(nil)

// ListDevicePositionsPaginatorOptions is the paginator options for
// ListDevicePositions
type ListDevicePositionsPaginatorOptions struct {
	// An optional limit for the number of entries returned in a single call. Default
	// value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDevicePositionsPaginator is a paginator for ListDevicePositions
type ListDevicePositionsPaginator struct {
	options   ListDevicePositionsPaginatorOptions
	client    ListDevicePositionsAPIClient
	params    *ListDevicePositionsInput
	nextToken *string
	firstPage bool
}

// NewListDevicePositionsPaginator returns a new ListDevicePositionsPaginator
func NewListDevicePositionsPaginator(client ListDevicePositionsAPIClient, params *ListDevicePositionsInput, optFns ...func(*ListDevicePositionsPaginatorOptions)) *ListDevicePositionsPaginator {
	if params == nil {
		params = &ListDevicePositionsInput{}
	}

	options := ListDevicePositionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDevicePositionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDevicePositionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDevicePositions page.
func (p *ListDevicePositionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDevicePositionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListDevicePositions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}
