// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// In an environment account, view the detail data for an environment account
// connection. For more information, see Environment account connections in the AWS
// Proton Administration guide.
func (c *Client) GetEnvironmentAccountConnection(ctx context.Context, params *GetEnvironmentAccountConnectionInput, optFns ...func(*Options)) (*GetEnvironmentAccountConnectionOutput, error) {
	if params == nil {
		params = &GetEnvironmentAccountConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnvironmentAccountConnection", params, optFns, c.addOperationGetEnvironmentAccountConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnvironmentAccountConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnvironmentAccountConnectionInput struct {

	// The ID of the environment account connection.
	//
	// This member is required.
	Id *string
}

type GetEnvironmentAccountConnectionOutput struct {

	// The environment account connection detail data that's returned by AWS Proton.
	//
	// This member is required.
	EnvironmentAccountConnection *types.EnvironmentAccountConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetEnvironmentAccountConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEnvironmentAccountConnection{}, middleware.After)
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
	if err = addOpGetEnvironmentAccountConnectionValidationMiddleware(stack); err != nil {
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
