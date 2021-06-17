// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Obtains information about the conditional forwarders for this account. If no
// input parameters are provided for RemoteDomainNames, this request describes all
// conditional forwarders for the specified directory ID.
func (c *Client) DescribeConditionalForwarders(ctx context.Context, params *DescribeConditionalForwardersInput, optFns ...func(*Options)) (*DescribeConditionalForwardersOutput, error) {
	if params == nil {
		params = &DescribeConditionalForwardersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConditionalForwarders", params, optFns, c.addOperationDescribeConditionalForwardersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConditionalForwardersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Describes a conditional forwarder.
type DescribeConditionalForwardersInput struct {

	// The directory ID for which to get the list of associated conditional forwarders.
	//
	// This member is required.
	DirectoryId *string

	// The fully qualified domain names (FQDN) of the remote domains for which to get
	// the list of associated conditional forwarders. If this member is null, all
	// conditional forwarders are returned.
	RemoteDomainNames []string
}

// The result of a DescribeConditionalForwarder request.
type DescribeConditionalForwardersOutput struct {

	// The list of conditional forwarders that have been created.
	ConditionalForwarders []types.ConditionalForwarder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeConditionalForwardersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConditionalForwarders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConditionalForwarders{}, middleware.After)
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
	if err = addOpDescribeConditionalForwardersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConditionalForwarders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConditionalForwarders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeConditionalForwarders",
	}
}
