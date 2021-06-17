// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the shards in a stream and provides information about each shard. This
// operation has a limit of 100 transactions per second per data stream. This API
// is a new operation that is used by the Amazon Kinesis Client Library (KCL). If
// you have a fine-grained IAM policy that only allows specific operations, you
// must update your policy to allow calls to this API. For more information, see
// Controlling Access to Amazon Kinesis Data Streams Resources Using IAM
// (https://docs.aws.amazon.com/streams/latest/dev/controlling-access.html).
func (c *Client) ListShards(ctx context.Context, params *ListShardsInput, optFns ...func(*Options)) (*ListShardsOutput, error) {
	if params == nil {
		params = &ListShardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListShards", params, optFns, c.addOperationListShardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListShardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListShardsInput struct {

	// Specify this parameter to indicate that you want to list the shards starting
	// with the shard whose ID immediately follows ExclusiveStartShardId. If you don't
	// specify this parameter, the default behavior is for ListShards to list the
	// shards starting with the first one in the stream. You cannot specify this
	// parameter if you specify NextToken.
	ExclusiveStartShardId *string

	// The maximum number of shards to return in a single call to ListShards. The
	// minimum value you can specify for this parameter is 1, and the maximum is
	// 10,000, which is also the default. When the number of shards to be listed is
	// greater than the value of MaxResults, the response contains a NextToken value
	// that you can use in a subsequent call to ListShards to list the next set of
	// shards.
	MaxResults *int32

	// When the number of shards in the data stream is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of shards in the data stream, the
	// response includes a pagination token named NextToken. You can specify this
	// NextToken value in a subsequent call to ListShards to list the next set of
	// shards. Don't specify StreamName or StreamCreationTimestamp if you specify
	// NextToken because the latter unambiguously identifies the stream. You can
	// optionally specify a value for the MaxResults parameter when you specify
	// NextToken. If you specify a MaxResults value that is less than the number of
	// shards that the operation returns if you don't specify MaxResults, the response
	// will contain a new NextToken value. You can use the new NextToken value in a
	// subsequent call to the ListShards operation. Tokens expire after 300 seconds.
	// When you obtain a value for NextToken in the response to a call to ListShards,
	// you have 300 seconds to use that value. If you specify an expired token in a
	// call to ListShards, you get ExpiredNextTokenException.
	NextToken *string

	ShardFilter *types.ShardFilter

	// Specify this input parameter to distinguish data streams that have the same
	// name. For example, if you create a data stream and then delete it, and you later
	// create another data stream with the same name, you can use this input parameter
	// to specify which of the two streams you want to list the shards for. You cannot
	// specify this parameter if you specify the NextToken parameter.
	StreamCreationTimestamp *time.Time

	// The name of the data stream whose shards you want to list. You cannot specify
	// this parameter if you specify the NextToken parameter.
	StreamName *string
}

type ListShardsOutput struct {

	// When the number of shards in the data stream is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of shards in the data stream, the
	// response includes a pagination token named NextToken. You can specify this
	// NextToken value in a subsequent call to ListShards to list the next set of
	// shards. For more information about the use of this pagination token when calling
	// the ListShards operation, see ListShardsInput$NextToken. Tokens expire after 300
	// seconds. When you obtain a value for NextToken in the response to a call to
	// ListShards, you have 300 seconds to use that value. If you specify an expired
	// token in a call to ListShards, you get ExpiredNextTokenException.
	NextToken *string

	// An array of JSON objects. Each object represents one shard and specifies the IDs
	// of the shard, the shard's parent, and the shard that's adjacent to the shard's
	// parent. Each object also contains the starting and ending hash keys and the
	// starting and ending sequence numbers for the shard.
	Shards []types.Shard

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListShardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListShards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListShards{}, middleware.After)
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
	if err = addOpListShardsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListShards(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListShards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "ListShards",
	}
}
