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

// Creates a new thing type.
func (c *Client) CreateThingType(ctx context.Context, params *CreateThingTypeInput, optFns ...func(*Options)) (*CreateThingTypeOutput, error) {
	if params == nil {
		params = &CreateThingTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateThingType", params, optFns, c.addOperationCreateThingTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateThingTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the CreateThingType operation.
type CreateThingTypeInput struct {

	// The name of the thing type.
	//
	// This member is required.
	ThingTypeName *string

	// Metadata which can be used to manage the thing type.
	Tags []types.Tag

	// The ThingTypeProperties for the thing type to create. It contains information
	// about the new thing type including a description, and a list of searchable thing
	// attribute names.
	ThingTypeProperties *types.ThingTypeProperties
}

// The output of the CreateThingType operation.
type CreateThingTypeOutput struct {

	// The Amazon Resource Name (ARN) of the thing type.
	ThingTypeArn *string

	// The thing type ID.
	ThingTypeId *string

	// The name of the thing type.
	ThingTypeName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateThingTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateThingType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThingType{}, middleware.After)
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
	if err = addOpCreateThingTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThingType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateThingType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateThingType",
	}
}
