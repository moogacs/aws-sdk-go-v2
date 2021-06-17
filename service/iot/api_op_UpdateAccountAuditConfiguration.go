// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures or reconfigures the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks are
// enabled or disabled.
func (c *Client) UpdateAccountAuditConfiguration(ctx context.Context, params *UpdateAccountAuditConfigurationInput, optFns ...func(*Options)) (*UpdateAccountAuditConfigurationOutput, error) {
	if params == nil {
		params = &UpdateAccountAuditConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountAuditConfiguration", params, optFns, c.addOperationUpdateAccountAuditConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountAuditConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountAuditConfigurationInput struct {

	// Specifies which audit checks are enabled and disabled for this account. Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are currently enabled. Some data collection might start immediately when
	// certain checks are enabled. When a check is disabled, any data collected so far
	// in relation to the check is deleted. You cannot disable a check if it's used by
	// any scheduled audit. You must first delete the check from the scheduled audit or
	// delete the scheduled audit itself. On the first call to
	// UpdateAccountAuditConfiguration, this parameter is required and must specify at
	// least one enabled check.
	AuditCheckConfigurations map[string]types.AuditCheckConfiguration

	// Information about the targets to which audit notifications are sent.
	AuditNotificationTargetConfigurations map[string]types.AuditNotificationTarget

	// The Amazon Resource Name (ARN) of the role that grants permission to AWS IoT to
	// access information about your devices, policies, certificates, and other items
	// as required when performing an audit.
	RoleArn *string
}

type UpdateAccountAuditConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateAccountAuditConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccountAuditConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccountAuditConfiguration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountAuditConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAccountAuditConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateAccountAuditConfiguration",
	}
}
