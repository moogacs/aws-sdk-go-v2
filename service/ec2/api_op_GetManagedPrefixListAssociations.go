// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the resources that are associated with the specified
// managed prefix list.
func (c *Client) GetManagedPrefixListAssociations(ctx context.Context, params *GetManagedPrefixListAssociationsInput, optFns ...func(*Options)) (*GetManagedPrefixListAssociationsOutput, error) {
	if params == nil {
		params = &GetManagedPrefixListAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetManagedPrefixListAssociations", params, optFns, c.addOperationGetManagedPrefixListAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetManagedPrefixListAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetManagedPrefixListAssociationsInput struct {

	// The ID of the prefix list.
	//
	// This member is required.
	PrefixListId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type GetManagedPrefixListAssociationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the associations.
	PrefixListAssociations []types.PrefixListAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetManagedPrefixListAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetManagedPrefixListAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetManagedPrefixListAssociations{}, middleware.After)
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
	if err = addOpGetManagedPrefixListAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetManagedPrefixListAssociations(options.Region), middleware.Before); err != nil {
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

// GetManagedPrefixListAssociationsAPIClient is a client that implements the
// GetManagedPrefixListAssociations operation.
type GetManagedPrefixListAssociationsAPIClient interface {
	GetManagedPrefixListAssociations(context.Context, *GetManagedPrefixListAssociationsInput, ...func(*Options)) (*GetManagedPrefixListAssociationsOutput, error)
}

var _ GetManagedPrefixListAssociationsAPIClient = (*Client)(nil)

// GetManagedPrefixListAssociationsPaginatorOptions is the paginator options for
// GetManagedPrefixListAssociations
type GetManagedPrefixListAssociationsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetManagedPrefixListAssociationsPaginator is a paginator for
// GetManagedPrefixListAssociations
type GetManagedPrefixListAssociationsPaginator struct {
	options   GetManagedPrefixListAssociationsPaginatorOptions
	client    GetManagedPrefixListAssociationsAPIClient
	params    *GetManagedPrefixListAssociationsInput
	nextToken *string
	firstPage bool
}

// NewGetManagedPrefixListAssociationsPaginator returns a new
// GetManagedPrefixListAssociationsPaginator
func NewGetManagedPrefixListAssociationsPaginator(client GetManagedPrefixListAssociationsAPIClient, params *GetManagedPrefixListAssociationsInput, optFns ...func(*GetManagedPrefixListAssociationsPaginatorOptions)) *GetManagedPrefixListAssociationsPaginator {
	if params == nil {
		params = &GetManagedPrefixListAssociationsInput{}
	}

	options := GetManagedPrefixListAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetManagedPrefixListAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetManagedPrefixListAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetManagedPrefixListAssociations page.
func (p *GetManagedPrefixListAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetManagedPrefixListAssociationsOutput, error) {
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

	result, err := p.client.GetManagedPrefixListAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetManagedPrefixListAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetManagedPrefixListAssociations",
	}
}
