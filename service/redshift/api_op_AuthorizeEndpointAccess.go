// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants access to a cluster.
func (c *Client) AuthorizeEndpointAccess(ctx context.Context, params *AuthorizeEndpointAccessInput, optFns ...func(*Options)) (*AuthorizeEndpointAccessOutput, error) {
	if params == nil {
		params = &AuthorizeEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeEndpointAccess", params, optFns, c.addOperationAuthorizeEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeEndpointAccessInput struct {

	// The AWS account ID to grant access to.
	//
	// This member is required.
	Account *string

	// The cluster identifier of the cluster to grant access to.
	ClusterIdentifier *string

	// The virtual private cloud (VPC) identifiers to grant access to.
	VpcIds []string
}

// Describes an endpoint authorization for authorizing Redshift-managed VPC
// endpoint access to a cluster across AWS accounts.
type AuthorizeEndpointAccessOutput struct {

	// Indicates whether all VPCs in the grantee account are allowed access to the
	// cluster.
	AllowedAllVPCs bool

	// The VPCs allowed access to the cluster.
	AllowedVPCs []string

	// The time (UTC) when the authorization was created.
	AuthorizeTime *time.Time

	// The cluster identifier.
	ClusterIdentifier *string

	// The status of the cluster.
	ClusterStatus *string

	// The number of Redshift-managed VPC endpoints created for the authorization.
	EndpointCount int32

	// The AWS account ID of the grantee of the cluster.
	Grantee *string

	// The AWS account ID of the cluster owner.
	Grantor *string

	// The status of the authorization action.
	Status types.AuthorizationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAuthorizeEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAuthorizeEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAuthorizeEndpointAccess{}, middleware.After)
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
	if err = addOpAuthorizeEndpointAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeEndpointAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "AuthorizeEndpointAccess",
	}
}
