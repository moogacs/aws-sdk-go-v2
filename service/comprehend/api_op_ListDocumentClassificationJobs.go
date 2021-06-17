// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the documentation classification jobs that you have submitted.
func (c *Client) ListDocumentClassificationJobs(ctx context.Context, params *ListDocumentClassificationJobsInput, optFns ...func(*Options)) (*ListDocumentClassificationJobsOutput, error) {
	if params == nil {
		params = &ListDocumentClassificationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDocumentClassificationJobs", params, optFns, c.addOperationListDocumentClassificationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDocumentClassificationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentClassificationJobsInput struct {

	// Filters the jobs that are returned. You can filter jobs on their names, status,
	// or the date and time that they were submitted. You can only set one filter at a
	// time.
	Filter *types.DocumentClassificationJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string
}

type ListDocumentClassificationJobsOutput struct {

	// A list containing the properties of each job returned.
	DocumentClassificationJobPropertiesList []types.DocumentClassificationJobProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListDocumentClassificationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDocumentClassificationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocumentClassificationJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDocumentClassificationJobs(options.Region), middleware.Before); err != nil {
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

// ListDocumentClassificationJobsAPIClient is a client that implements the
// ListDocumentClassificationJobs operation.
type ListDocumentClassificationJobsAPIClient interface {
	ListDocumentClassificationJobs(context.Context, *ListDocumentClassificationJobsInput, ...func(*Options)) (*ListDocumentClassificationJobsOutput, error)
}

var _ ListDocumentClassificationJobsAPIClient = (*Client)(nil)

// ListDocumentClassificationJobsPaginatorOptions is the paginator options for
// ListDocumentClassificationJobs
type ListDocumentClassificationJobsPaginatorOptions struct {
	// The maximum number of results to return in each page. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDocumentClassificationJobsPaginator is a paginator for
// ListDocumentClassificationJobs
type ListDocumentClassificationJobsPaginator struct {
	options   ListDocumentClassificationJobsPaginatorOptions
	client    ListDocumentClassificationJobsAPIClient
	params    *ListDocumentClassificationJobsInput
	nextToken *string
	firstPage bool
}

// NewListDocumentClassificationJobsPaginator returns a new
// ListDocumentClassificationJobsPaginator
func NewListDocumentClassificationJobsPaginator(client ListDocumentClassificationJobsAPIClient, params *ListDocumentClassificationJobsInput, optFns ...func(*ListDocumentClassificationJobsPaginatorOptions)) *ListDocumentClassificationJobsPaginator {
	if params == nil {
		params = &ListDocumentClassificationJobsInput{}
	}

	options := ListDocumentClassificationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDocumentClassificationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDocumentClassificationJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDocumentClassificationJobs page.
func (p *ListDocumentClassificationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDocumentClassificationJobsOutput, error) {
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

	result, err := p.client.ListDocumentClassificationJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDocumentClassificationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListDocumentClassificationJobs",
	}
}
