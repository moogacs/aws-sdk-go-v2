// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an empty receipt rule set. For information about setting up receipt rule
// sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html).
// You can execute this operation no more than once per second.
func (c *Client) CreateReceiptRuleSet(ctx context.Context, params *CreateReceiptRuleSetInput, optFns ...func(*Options)) (*CreateReceiptRuleSetOutput, error) {
	if params == nil {
		params = &CreateReceiptRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReceiptRuleSet", params, optFns, c.addOperationCreateReceiptRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReceiptRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create an empty receipt rule set. You use receipt rule
// sets to receive email with Amazon SES. For more information, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type CreateReceiptRuleSetInput struct {

	// The name of the rule set to create. The name must:
	//
	// * This value can only
	// contain ASCII letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes
	// (-).
	//
	// * Start and end with a letter or number.
	//
	// * Contain less than 64
	// characters.
	//
	// This member is required.
	RuleSetName *string
}

// An empty element returned on a successful request.
type CreateReceiptRuleSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateReceiptRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateReceiptRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateReceiptRuleSet{}, middleware.After)
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
	if err = addOpCreateReceiptRuleSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReceiptRuleSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReceiptRuleSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateReceiptRuleSet",
	}
}
