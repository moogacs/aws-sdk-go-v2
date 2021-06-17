// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a description of the certificate.
func (c *Client) DescribeCertificates(ctx context.Context, params *DescribeCertificatesInput, optFns ...func(*Options)) (*DescribeCertificatesOutput, error) {
	if params == nil {
		params = &DescribeCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCertificates", params, optFns, c.addOperationDescribeCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCertificatesInput struct {

	// Filters applied to the certificates described in the form of key-value pairs.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 10
	MaxRecords *int32
}

type DescribeCertificatesOutput struct {

	// The Secure Sockets Layer (SSL) certificates associated with the replication
	// instance.
	Certificates []types.Certificate

	// The pagination token.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCertificates{}, middleware.After)
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
	if err = addOpDescribeCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCertificates(options.Region), middleware.Before); err != nil {
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

// DescribeCertificatesAPIClient is a client that implements the
// DescribeCertificates operation.
type DescribeCertificatesAPIClient interface {
	DescribeCertificates(context.Context, *DescribeCertificatesInput, ...func(*Options)) (*DescribeCertificatesOutput, error)
}

var _ DescribeCertificatesAPIClient = (*Client)(nil)

// DescribeCertificatesPaginatorOptions is the paginator options for
// DescribeCertificates
type DescribeCertificatesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 10
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCertificatesPaginator is a paginator for DescribeCertificates
type DescribeCertificatesPaginator struct {
	options   DescribeCertificatesPaginatorOptions
	client    DescribeCertificatesAPIClient
	params    *DescribeCertificatesInput
	nextToken *string
	firstPage bool
}

// NewDescribeCertificatesPaginator returns a new DescribeCertificatesPaginator
func NewDescribeCertificatesPaginator(client DescribeCertificatesAPIClient, params *DescribeCertificatesInput, optFns ...func(*DescribeCertificatesPaginatorOptions)) *DescribeCertificatesPaginator {
	if params == nil {
		params = &DescribeCertificatesInput{}
	}

	options := DescribeCertificatesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCertificatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCertificatesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeCertificates page.
func (p *DescribeCertificatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCertificatesOutput, error) {
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

	result, err := p.client.DescribeCertificates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeCertificates",
	}
}
