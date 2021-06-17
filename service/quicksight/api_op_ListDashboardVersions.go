// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the versions of the dashboards in the QuickSight subscription.
func (c *Client) ListDashboardVersions(ctx context.Context, params *ListDashboardVersionsInput, optFns ...func(*Options)) (*ListDashboardVersionsOutput, error) {
	if params == nil {
		params = &ListDashboardVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDashboardVersions", params, optFns, c.addOperationListDashboardVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDashboardVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDashboardVersionsInput struct {

	// The ID of the AWS account that contains the dashboard that you're listing
	// versions for.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard.
	//
	// This member is required.
	DashboardId *string

	// The maximum number of results to be returned per request.
	MaxResults int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type ListDashboardVersionsOutput struct {

	// A structure that contains information about each version of the dashboard.
	DashboardVersionSummaryList []types.DashboardVersionSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListDashboardVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDashboardVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDashboardVersions{}, middleware.After)
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
	if err = addOpListDashboardVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDashboardVersions(options.Region), middleware.Before); err != nil {
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

// ListDashboardVersionsAPIClient is a client that implements the
// ListDashboardVersions operation.
type ListDashboardVersionsAPIClient interface {
	ListDashboardVersions(context.Context, *ListDashboardVersionsInput, ...func(*Options)) (*ListDashboardVersionsOutput, error)
}

var _ ListDashboardVersionsAPIClient = (*Client)(nil)

// ListDashboardVersionsPaginatorOptions is the paginator options for
// ListDashboardVersions
type ListDashboardVersionsPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDashboardVersionsPaginator is a paginator for ListDashboardVersions
type ListDashboardVersionsPaginator struct {
	options   ListDashboardVersionsPaginatorOptions
	client    ListDashboardVersionsAPIClient
	params    *ListDashboardVersionsInput
	nextToken *string
	firstPage bool
}

// NewListDashboardVersionsPaginator returns a new ListDashboardVersionsPaginator
func NewListDashboardVersionsPaginator(client ListDashboardVersionsAPIClient, params *ListDashboardVersionsInput, optFns ...func(*ListDashboardVersionsPaginatorOptions)) *ListDashboardVersionsPaginator {
	if params == nil {
		params = &ListDashboardVersionsInput{}
	}

	options := ListDashboardVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDashboardVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDashboardVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDashboardVersions page.
func (p *ListDashboardVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDashboardVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDashboardVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDashboardVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListDashboardVersions",
	}
}
