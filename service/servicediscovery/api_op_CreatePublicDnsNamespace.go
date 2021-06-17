// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a public namespace based on DNS, which is visible on the internet. The
// namespace defines your service naming scheme. For example, if you name your
// namespace example.com and name your service backend, the resulting DNS name for
// the service is backend.example.com. You can discover instances that were
// registered with a public DNS namespace by using either a DiscoverInstances
// request or using DNS. For the current quota on the number of namespaces that you
// can create using the same AWS account, see AWS Cloud Map Quotas
// (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)in the
// AWS Cloud Map Developer Guide.
func (c *Client) CreatePublicDnsNamespace(ctx context.Context, params *CreatePublicDnsNamespaceInput, optFns ...func(*Options)) (*CreatePublicDnsNamespaceOutput, error) {
	if params == nil {
		params = &CreatePublicDnsNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePublicDnsNamespace", params, optFns, c.addOperationCreatePublicDnsNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePublicDnsNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePublicDnsNamespaceInput struct {

	// The name that you want to assign to this namespace.
	//
	// This member is required.
	Name *string

	// A unique string that identifies the request and that allows failed
	// CreatePublicDnsNamespace requests to be retried without the risk of running the
	// operation twice. CreatorRequestId can be any unique string (for example, a
	// date/timestamp).
	CreatorRequestId *string

	// A description for the namespace.
	Description *string

	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value that you define. Tags keys can be up to 128 characters in length, and tag
	// values can be up to 256 characters in length.
	Tags []types.Tag
}

type CreatePublicDnsNamespaceOutput struct {

	// A value that you can use to determine whether the request completed
	// successfully. To get the status of the operation, see GetOperation
	// (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreatePublicDnsNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePublicDnsNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePublicDnsNamespace{}, middleware.After)
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
	if err = addIdempotencyToken_opCreatePublicDnsNamespaceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePublicDnsNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePublicDnsNamespace(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePublicDnsNamespace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePublicDnsNamespace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePublicDnsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePublicDnsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePublicDnsNamespaceInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePublicDnsNamespaceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePublicDnsNamespace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePublicDnsNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicediscovery",
		OperationName: "CreatePublicDnsNamespace",
	}
}
