// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists detectors (the instances of a detector model).
func (c *Client) ListDetectors(ctx context.Context, params *ListDetectorsInput, optFns ...func(*Options)) (*ListDetectorsOutput, error) {
	if params == nil {
		params = &ListDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDetectors", params, optFns, c.addOperationListDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDetectorsInput struct {

	// The name of the detector model whose detectors (instances) are listed.
	//
	// This member is required.
	DetectorModelName *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token that you can use to return the next set of results.
	NextToken *string

	// A filter that limits results to those detectors (instances) in the given state.
	StateName *string
}

type ListDetectorsOutput struct {

	// A list of summary information about the detectors (instances).
	DetectorSummaries []types.DetectorSummary

	// The token that you can use to return the next set of results, or null if there
	// are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDetectors{}, middleware.After)
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
	if err = addOpListDetectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDetectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDetectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ioteventsdata",
		OperationName: "ListDetectors",
	}
}
