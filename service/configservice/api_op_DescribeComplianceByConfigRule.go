// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Indicates whether the specified AWS Config rules are compliant. If a rule is
// noncompliant, this action returns the number of AWS resources that do not comply
// with the rule. A rule is compliant if all of the evaluated resources comply with
// it. It is noncompliant if any of these resources do not comply. If AWS Config
// has no current evaluation results for the rule, it returns INSUFFICIENT_DATA.
// This result might indicate one of the following conditions:
//
// * AWS Config has
// never invoked an evaluation for the rule. To check whether it has, use the
// DescribeConfigRuleEvaluationStatus action to get the
// LastSuccessfulInvocationTime and LastFailedInvocationTime.
//
// * The rule's AWS
// Lambda function is failing to send evaluation results to AWS Config. Verify that
// the role you assigned to your configuration recorder includes the
// config:PutEvaluations permission. If the rule is a custom rule, verify that the
// AWS Lambda execution role includes the config:PutEvaluations permission.
//
// * The
// rule's AWS Lambda function has returned NOT_APPLICABLE for all evaluation
// results. This can occur if the resources were deleted or removed from the rule's
// scope.
func (c *Client) DescribeComplianceByConfigRule(ctx context.Context, params *DescribeComplianceByConfigRuleInput, optFns ...func(*Options)) (*DescribeComplianceByConfigRuleOutput, error) {
	if params == nil {
		params = &DescribeComplianceByConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeComplianceByConfigRule", params, optFns, c.addOperationDescribeComplianceByConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeComplianceByConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeComplianceByConfigRuleInput struct {

	// Filters the results by compliance. The allowed values are COMPLIANT and
	// NON_COMPLIANT.
	ComplianceTypes []types.ComplianceType

	// Specify one or more AWS Config rule names to filter the results by rule.
	ConfigRuleNames []string

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string
}

//
type DescribeComplianceByConfigRuleOutput struct {

	// Indicates whether each of the specified AWS Config rules is compliant.
	ComplianceByConfigRules []types.ComplianceByConfigRule

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeComplianceByConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeComplianceByConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeComplianceByConfigRule{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComplianceByConfigRule(options.Region), middleware.Before); err != nil {
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

// DescribeComplianceByConfigRuleAPIClient is a client that implements the
// DescribeComplianceByConfigRule operation.
type DescribeComplianceByConfigRuleAPIClient interface {
	DescribeComplianceByConfigRule(context.Context, *DescribeComplianceByConfigRuleInput, ...func(*Options)) (*DescribeComplianceByConfigRuleOutput, error)
}

var _ DescribeComplianceByConfigRuleAPIClient = (*Client)(nil)

// DescribeComplianceByConfigRulePaginatorOptions is the paginator options for
// DescribeComplianceByConfigRule
type DescribeComplianceByConfigRulePaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeComplianceByConfigRulePaginator is a paginator for
// DescribeComplianceByConfigRule
type DescribeComplianceByConfigRulePaginator struct {
	options   DescribeComplianceByConfigRulePaginatorOptions
	client    DescribeComplianceByConfigRuleAPIClient
	params    *DescribeComplianceByConfigRuleInput
	nextToken *string
	firstPage bool
}

// NewDescribeComplianceByConfigRulePaginator returns a new
// DescribeComplianceByConfigRulePaginator
func NewDescribeComplianceByConfigRulePaginator(client DescribeComplianceByConfigRuleAPIClient, params *DescribeComplianceByConfigRuleInput, optFns ...func(*DescribeComplianceByConfigRulePaginatorOptions)) *DescribeComplianceByConfigRulePaginator {
	if params == nil {
		params = &DescribeComplianceByConfigRuleInput{}
	}

	options := DescribeComplianceByConfigRulePaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeComplianceByConfigRulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeComplianceByConfigRulePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeComplianceByConfigRule page.
func (p *DescribeComplianceByConfigRulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeComplianceByConfigRuleOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeComplianceByConfigRule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeComplianceByConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeComplianceByConfigRule",
	}
}
