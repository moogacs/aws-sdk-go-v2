// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks out the specified license.
func (c *Client) CheckoutLicense(ctx context.Context, params *CheckoutLicenseInput, optFns ...func(*Options)) (*CheckoutLicenseOutput, error) {
	if params == nil {
		params = &CheckoutLicenseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckoutLicense", params, optFns, c.addOperationCheckoutLicenseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckoutLicenseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CheckoutLicenseInput struct {

	// Checkout type.
	//
	// This member is required.
	CheckoutType types.CheckoutType

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientToken *string

	// License entitlements.
	//
	// This member is required.
	Entitlements []types.EntitlementData

	// Key fingerprint identifying the license.
	//
	// This member is required.
	KeyFingerprint *string

	// Product SKU.
	//
	// This member is required.
	ProductSKU *string

	// License beneficiary.
	Beneficiary *string

	// Node ID.
	NodeId *string
}

type CheckoutLicenseOutput struct {

	// Checkout type.
	CheckoutType types.CheckoutType

	// Allowed license entitlements.
	EntitlementsAllowed []types.EntitlementData

	// Date and time at which the license checkout expires.
	Expiration *string

	// Date and time at which the license checkout is issued.
	IssuedAt *string

	// License consumption token.
	LicenseConsumptionToken *string

	// Node ID.
	NodeId *string

	// Signed token.
	SignedToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCheckoutLicenseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCheckoutLicense{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCheckoutLicense{}, middleware.After)
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
	if err = addOpCheckoutLicenseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckoutLicense(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckoutLicense(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "CheckoutLicense",
	}
}
