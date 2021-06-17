// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed information about a given Amazon QLDB journal stream. The
// output includes the Amazon Resource Name (ARN), stream name, current status,
// creation time, and the parameters of the original stream creation request. This
// action does not return any expired journal streams. For more information, see
// Expiration for terminal streams
// (https://docs.aws.amazon.com/qldb/latest/developerguide/streams.create.html#streams.create.states.expiration)
// in the Amazon QLDB Developer Guide.
func (c *Client) DescribeJournalKinesisStream(ctx context.Context, params *DescribeJournalKinesisStreamInput, optFns ...func(*Options)) (*DescribeJournalKinesisStreamOutput, error) {
	if params == nil {
		params = &DescribeJournalKinesisStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJournalKinesisStream", params, optFns, c.addOperationDescribeJournalKinesisStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJournalKinesisStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJournalKinesisStreamInput struct {

	// The name of the ledger.
	//
	// This member is required.
	LedgerName *string

	// The UUID (represented in Base62-encoded text) of the QLDB journal stream to
	// describe.
	//
	// This member is required.
	StreamId *string
}

type DescribeJournalKinesisStreamOutput struct {

	// Information about the QLDB journal stream returned by a DescribeJournalS3Export
	// request.
	Stream *types.JournalKinesisStreamDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeJournalKinesisStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJournalKinesisStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJournalKinesisStream{}, middleware.After)
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
	if err = addOpDescribeJournalKinesisStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJournalKinesisStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeJournalKinesisStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "DescribeJournalKinesisStream",
	}
}
