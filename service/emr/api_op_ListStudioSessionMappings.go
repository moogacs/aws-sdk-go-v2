// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all user or group session mappings for the Amazon EMR Studio
// specified by StudioId.
func (c *Client) ListStudioSessionMappings(ctx context.Context, params *ListStudioSessionMappingsInput, optFns ...func(*Options)) (*ListStudioSessionMappingsOutput, error) {
	if params == nil {
		params = &ListStudioSessionMappingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStudioSessionMappings", params, optFns, c.addOperationListStudioSessionMappingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStudioSessionMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStudioSessionMappingsInput struct {

	// Specifies whether to return session mappings for users or groups. If not
	// specified, the results include session mapping details for both users and
	// groups.
	IdentityType types.IdentityType

	// The pagination token that indicates the set of results to retrieve.
	Marker *string

	// The ID of the Amazon EMR Studio.
	StudioId *string
}

type ListStudioSessionMappingsOutput struct {

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// A list of session mapping summary objects. Each object includes session mapping
	// details such as creation time, identity type (user or group), and Amazon EMR
	// Studio ID.
	SessionMappings []types.SessionMappingSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListStudioSessionMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStudioSessionMappings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStudioSessionMappings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStudioSessionMappings(options.Region), middleware.Before); err != nil {
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

// ListStudioSessionMappingsAPIClient is a client that implements the
// ListStudioSessionMappings operation.
type ListStudioSessionMappingsAPIClient interface {
	ListStudioSessionMappings(context.Context, *ListStudioSessionMappingsInput, ...func(*Options)) (*ListStudioSessionMappingsOutput, error)
}

var _ ListStudioSessionMappingsAPIClient = (*Client)(nil)

// ListStudioSessionMappingsPaginatorOptions is the paginator options for
// ListStudioSessionMappings
type ListStudioSessionMappingsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStudioSessionMappingsPaginator is a paginator for ListStudioSessionMappings
type ListStudioSessionMappingsPaginator struct {
	options   ListStudioSessionMappingsPaginatorOptions
	client    ListStudioSessionMappingsAPIClient
	params    *ListStudioSessionMappingsInput
	nextToken *string
	firstPage bool
}

// NewListStudioSessionMappingsPaginator returns a new
// ListStudioSessionMappingsPaginator
func NewListStudioSessionMappingsPaginator(client ListStudioSessionMappingsAPIClient, params *ListStudioSessionMappingsInput, optFns ...func(*ListStudioSessionMappingsPaginatorOptions)) *ListStudioSessionMappingsPaginator {
	if params == nil {
		params = &ListStudioSessionMappingsInput{}
	}

	options := ListStudioSessionMappingsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStudioSessionMappingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStudioSessionMappingsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListStudioSessionMappings page.
func (p *ListStudioSessionMappingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStudioSessionMappingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListStudioSessionMappings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStudioSessionMappings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListStudioSessionMappings",
	}
}
