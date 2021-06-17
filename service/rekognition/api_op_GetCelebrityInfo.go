// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the name and additional information about a celebrity based on his or her
// Amazon Rekognition ID. The additional information is returned as an array of
// URLs. If there is no additional information about the celebrity, this list is
// empty. For more information, see Recognizing Celebrities in an Image in the
// Amazon Rekognition Developer Guide. This operation requires permissions to
// perform the rekognition:GetCelebrityInfo action.
func (c *Client) GetCelebrityInfo(ctx context.Context, params *GetCelebrityInfoInput, optFns ...func(*Options)) (*GetCelebrityInfoOutput, error) {
	if params == nil {
		params = &GetCelebrityInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCelebrityInfo", params, optFns, c.addOperationGetCelebrityInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCelebrityInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCelebrityInfoInput struct {

	// The ID for the celebrity. You get the celebrity ID from a call to the
	// RecognizeCelebrities operation, which recognizes celebrities in an image.
	//
	// This member is required.
	Id *string
}

type GetCelebrityInfoOutput struct {

	// The name of the celebrity.
	Name *string

	// An array of URLs pointing to additional celebrity information.
	Urls []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetCelebrityInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCelebrityInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCelebrityInfo{}, middleware.After)
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
	if err = addOpGetCelebrityInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCelebrityInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCelebrityInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetCelebrityInfo",
	}
}
