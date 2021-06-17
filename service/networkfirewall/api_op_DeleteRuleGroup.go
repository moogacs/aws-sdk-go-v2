// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified RuleGroup.
func (c *Client) DeleteRuleGroup(ctx context.Context, params *DeleteRuleGroupInput, optFns ...func(*Options)) (*DeleteRuleGroupOutput, error) {
	if params == nil {
		params = &DeleteRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRuleGroup", params, optFns, c.addOperationDeleteRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRuleGroupInput struct {

	// The Amazon Resource Name (ARN) of the rule group. You must specify the ARN or
	// the name, and you can specify both.
	RuleGroupArn *string

	// The descriptive name of the rule group. You can't change the name of a rule
	// group after you create it. You must specify the ARN or the name, and you can
	// specify both.
	RuleGroupName *string

	// Indicates whether the rule group is stateless or stateful. If the rule group is
	// stateless, it contains stateless rules. If it is stateful, it contains stateful
	// rules. This setting is required for requests that do not include the
	// RuleGroupARN.
	Type types.RuleGroupType
}

type DeleteRuleGroupOutput struct {

	// The high-level properties of a rule group. This, along with the RuleGroup,
	// define the rule group. You can retrieve all objects for a rule group by calling
	// DescribeRuleGroup.
	//
	// This member is required.
	RuleGroupResponse *types.RuleGroupResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteRuleGroup{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "DeleteRuleGroup",
	}
}
