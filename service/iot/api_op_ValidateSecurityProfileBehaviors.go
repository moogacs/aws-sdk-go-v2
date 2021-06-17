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

// Validates a Device Defender security profile behaviors specification.
func (c *Client) ValidateSecurityProfileBehaviors(ctx context.Context, params *ValidateSecurityProfileBehaviorsInput, optFns ...func(*Options)) (*ValidateSecurityProfileBehaviorsOutput, error) {
	if params == nil {
		params = &ValidateSecurityProfileBehaviorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateSecurityProfileBehaviors", params, optFns, c.addOperationValidateSecurityProfileBehaviorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateSecurityProfileBehaviorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidateSecurityProfileBehaviorsInput struct {

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	//
	// This member is required.
	Behaviors []types.Behavior
}

type ValidateSecurityProfileBehaviorsOutput struct {

	// True if the behaviors were valid.
	Valid bool

	// The list of any errors found in the behaviors.
	ValidationErrors []types.ValidationError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationValidateSecurityProfileBehaviorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidateSecurityProfileBehaviors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidateSecurityProfileBehaviors{}, middleware.After)
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
	if err = addOpValidateSecurityProfileBehaviorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidateSecurityProfileBehaviors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opValidateSecurityProfileBehaviors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ValidateSecurityProfileBehaviors",
	}
}
