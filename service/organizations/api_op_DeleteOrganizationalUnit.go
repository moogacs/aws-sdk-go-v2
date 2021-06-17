// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an organizational unit (OU) from a root or another OU. You must first
// remove all accounts and child OUs from the OU that you want to delete. This
// operation can be called only from the organization's management account.
func (c *Client) DeleteOrganizationalUnit(ctx context.Context, params *DeleteOrganizationalUnitInput, optFns ...func(*Options)) (*DeleteOrganizationalUnitOutput, error) {
	if params == nil {
		params = &DeleteOrganizationalUnitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOrganizationalUnit", params, optFns, c.addOperationDeleteOrganizationalUnitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOrganizationalUnitInput struct {

	// The unique identifier (ID) of the organizational unit that you want to delete.
	// You can get the ID from the ListOrganizationalUnitsForParent operation. The
	// regex pattern (http://wikipedia.org/wiki/regex) for an organizational unit ID
	// string requires "ou-" followed by from 4 to 32 lowercase letters or digits (the
	// ID of the root that contains the OU). This string is followed by a second "-"
	// dash and from 8 to 32 additional lowercase letters or digits.
	//
	// This member is required.
	OrganizationalUnitId *string
}

type DeleteOrganizationalUnitOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteOrganizationalUnitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteOrganizationalUnit{}, middleware.After)
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
	if err = addOpDeleteOrganizationalUnitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOrganizationalUnit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOrganizationalUnit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DeleteOrganizationalUnit",
	}
}
