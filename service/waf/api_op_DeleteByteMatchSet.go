// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Permanently deletes a ByteMatchSet. You can't delete a ByteMatchSet
// if it's still used in any Rules or if it still includes any ByteMatchTuple
// objects (any filters). If you just want to remove a ByteMatchSet from a Rule,
// use UpdateRule. To permanently delete a ByteMatchSet, perform the following
// steps:
//
// * Update the ByteMatchSet to remove filters, if any. For more
// information, see UpdateByteMatchSet.
//
// * Use GetChangeToken to get the change
// token that you provide in the ChangeToken parameter of a DeleteByteMatchSet
// request.
//
// * Submit a DeleteByteMatchSet request.
func (c *Client) DeleteByteMatchSet(ctx context.Context, params *DeleteByteMatchSetInput, optFns ...func(*Options)) (*DeleteByteMatchSetOutput, error) {
	if params == nil {
		params = &DeleteByteMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteByteMatchSet", params, optFns, c.addOperationDeleteByteMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteByteMatchSetInput struct {

	// The ByteMatchSetId of the ByteMatchSet that you want to delete. ByteMatchSetId
	// is returned by CreateByteMatchSet and by ListByteMatchSets.
	//
	// This member is required.
	ByteMatchSetId *string

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string
}

type DeleteByteMatchSetOutput struct {

	// The ChangeToken that you used to submit the DeleteByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteByteMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteByteMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteByteMatchSet{}, middleware.After)
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
	if err = addOpDeleteByteMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteByteMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteByteMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "DeleteByteMatchSet",
	}
}
