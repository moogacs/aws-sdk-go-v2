// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of all CloudFront functions in your AWS account. You can optionally
// apply a filter to return only the functions that are in the specified stage,
// either DEVELOPMENT or LIVE. You can optionally specify the maximum number of
// items to receive in the response. If the total number of items in the list
// exceeds the maximum that you specify, or the default maximum, the response is
// paginated. To get the next page of items, send a subsequent request that
// specifies the NextMarker value from the current response as the Marker value in
// the subsequent request.
func (c *Client) ListFunctions(ctx context.Context, params *ListFunctionsInput, optFns ...func(*Options)) (*ListFunctionsOutput, error) {
	if params == nil {
		params = &ListFunctionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFunctions", params, optFns, c.addOperationListFunctionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFunctionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionsInput struct {

	// Use this field when paginating results to indicate where to begin in your list
	// of functions. The response includes functions in the list that occur after the
	// marker. To get the next page of the list, set this field’s value to the value of
	// NextMarker from the current page’s response.
	Marker *string

	// The maximum number of functions that you want in the response.
	MaxItems *int32

	// An optional filter to return only the functions that are in the specified stage,
	// either DEVELOPMENT or LIVE.
	Stage types.FunctionStage
}

type ListFunctionsOutput struct {

	// A list of CloudFront functions.
	FunctionList *types.FunctionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListFunctionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListFunctions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListFunctions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFunctions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListFunctions",
	}
}
