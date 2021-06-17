// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the settings of query suggestions for an index. This is used to check
// the current settings applied to query suggestions.
func (c *Client) DescribeQuerySuggestionsConfig(ctx context.Context, params *DescribeQuerySuggestionsConfigInput, optFns ...func(*Options)) (*DescribeQuerySuggestionsConfigOutput, error) {
	if params == nil {
		params = &DescribeQuerySuggestionsConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeQuerySuggestionsConfig", params, optFns, c.addOperationDescribeQuerySuggestionsConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeQuerySuggestionsConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQuerySuggestionsConfigInput struct {

	// The identifier of the index you want to describe query suggestions settings for.
	//
	// This member is required.
	IndexId *string
}

type DescribeQuerySuggestionsConfigOutput struct {

	// Shows whether Amazon Kendra uses all queries or only uses queries that include
	// user information to generate query suggestions.
	IncludeQueriesWithoutUserInformation *bool

	// Shows the date-time query suggestions for an index was last cleared. After you
	// clear suggestions, Amazon Kendra learns new suggestions based on new queries
	// added to the query log from the time you cleared suggestions. Amazon Kendra only
	// considers re-occurences of a query from the time you cleared suggestions.
	LastClearTime *time.Time

	// Shows the date-time query suggestions for an index was last updated.
	LastSuggestionsBuildTime *time.Time

	// Shows the minimum number of unique users who must search a query in order for
	// the query to be eligible to suggest to your users.
	MinimumNumberOfQueryingUsers *int32

	// Shows the minimum number of times a query must be searched in order for the
	// query to be eligible to suggest to your users.
	MinimumQueryCount *int32

	// Shows whether query suggestions are currently in ENABLED mode or LEARN_ONLY
	// mode. By default, Amazon Kendra enables query suggestions.LEARN_ONLY turns off
	// query suggestions for your users. You can change the mode using the
	// UpdateQuerySuggestionsConfig
	// (https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateQuerySuggestionsConfig.html)
	// operation.
	Mode types.Mode

	// Shows how recent your queries are in your query log time window (in days).
	QueryLogLookBackWindowInDays *int32

	// Shows whether the status of query suggestions settings is currently Active or
	// Updating. Active means the current settings apply and Updating means your
	// changed settings are in the process of applying.
	Status types.QuerySuggestionsStatus

	// Shows the current total count of query suggestions for an index. This count can
	// change when you update your query suggestions settings, if you filter out
	// certain queries from suggestions using a block list, and as the query log
	// accumulates more queries for Amazon Kendra to learn from.
	TotalSuggestionsCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeQuerySuggestionsConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQuerySuggestionsConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQuerySuggestionsConfig{}, middleware.After)
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
	if err = addOpDescribeQuerySuggestionsConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQuerySuggestionsConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeQuerySuggestionsConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DescribeQuerySuggestionsConfig",
	}
}
