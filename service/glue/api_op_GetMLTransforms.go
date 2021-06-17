// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a sortable, filterable list of existing AWS Glue machine learning
// transforms. Machine learning transforms are a special type of transform that use
// machine learning to learn the details of the transformation to be performed by
// learning from examples provided by humans. These transformations are then saved
// by AWS Glue, and you can retrieve their metadata by calling GetMLTransforms.
func (c *Client) GetMLTransforms(ctx context.Context, params *GetMLTransformsInput, optFns ...func(*Options)) (*GetMLTransformsOutput, error) {
	if params == nil {
		params = &GetMLTransformsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLTransforms", params, optFns, c.addOperationGetMLTransformsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLTransformsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLTransformsInput struct {

	// The filter transformation criteria.
	Filter *types.TransformFilterCriteria

	// The maximum number of results to return.
	MaxResults *int32

	// A paginated token to offset the results.
	NextToken *string

	// The sorting criteria.
	Sort *types.TransformSortCriteria
}

type GetMLTransformsOutput struct {

	// A list of machine learning transforms.
	//
	// This member is required.
	Transforms []types.MLTransform

	// A pagination token, if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetMLTransformsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMLTransforms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMLTransforms{}, middleware.After)
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
	if err = addOpGetMLTransformsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMLTransforms(options.Region), middleware.Before); err != nil {
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

// GetMLTransformsAPIClient is a client that implements the GetMLTransforms
// operation.
type GetMLTransformsAPIClient interface {
	GetMLTransforms(context.Context, *GetMLTransformsInput, ...func(*Options)) (*GetMLTransformsOutput, error)
}

var _ GetMLTransformsAPIClient = (*Client)(nil)

// GetMLTransformsPaginatorOptions is the paginator options for GetMLTransforms
type GetMLTransformsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetMLTransformsPaginator is a paginator for GetMLTransforms
type GetMLTransformsPaginator struct {
	options   GetMLTransformsPaginatorOptions
	client    GetMLTransformsAPIClient
	params    *GetMLTransformsInput
	nextToken *string
	firstPage bool
}

// NewGetMLTransformsPaginator returns a new GetMLTransformsPaginator
func NewGetMLTransformsPaginator(client GetMLTransformsAPIClient, params *GetMLTransformsInput, optFns ...func(*GetMLTransformsPaginatorOptions)) *GetMLTransformsPaginator {
	if params == nil {
		params = &GetMLTransformsInput{}
	}

	options := GetMLTransformsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetMLTransformsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetMLTransformsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetMLTransforms page.
func (p *GetMLTransformsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetMLTransformsOutput, error) {
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

	result, err := p.client.GetMLTransforms(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetMLTransforms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetMLTransforms",
	}
}
