// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Instantiates an auto-scaling virtual server based on the selected file transfer
// protocol in AWS. When you make updates to your file transfer protocol-enabled
// server or when you work with users, use the service-generated ServerId property
// that is assigned to the newly created server.
func (c *Client) CreateServer(ctx context.Context, params *CreateServerInput, optFns ...func(*Options)) (*CreateServerOutput, error) {
	if params == nil {
		params = &CreateServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServer", params, optFns, c.addOperationCreateServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServerInput struct {

	// The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate.
	// Required when Protocols is set to FTPS. To request a new public certificate, see
	// Request a public certificate
	// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html) in
	// the AWS Certificate Manager User Guide. To import an existing certificate into
	// ACM, see Importing certificates into ACM
	// (https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in
	// the AWS Certificate Manager User Guide. To request a private certificate to use
	// FTPS through private IP addresses, see Request a private certificate
	// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-private.html)
	// in the AWS Certificate Manager User Guide. Certificates with the following
	// cryptographic algorithms and key sizes are supported:
	//
	// * 2048-bit RSA
	// (RSA_2048)
	//
	// * 4096-bit RSA (RSA_4096)
	//
	// * Elliptic Prime Curve 256 bit
	// (EC_prime256v1)
	//
	// * Elliptic Prime Curve 384 bit (EC_secp384r1)
	//
	// * Elliptic Prime
	// Curve 521 bit (EC_secp521r1)
	//
	// The certificate must be a valid SSL/TLS X.509
	// version 3 certificate with FQDN or IP address specified and information about
	// the issuer.
	Certificate *string

	// The domain of the storage system that is used for file transfers. There are two
	// domains available: Amazon Simple Storage Service (Amazon S3) and Amazon Elastic
	// File System (Amazon EFS). The default value is S3. After the server is created,
	// the domain cannot be changed.
	Domain types.Domain

	// The virtual private cloud (VPC) endpoint settings that are configured for your
	// server. When you host your endpoint within your VPC, you can make it accessible
	// only to resources within your VPC, or you can attach Elastic IP addresses and
	// make it accessible to clients over the internet. Your VPC's default security
	// groups are automatically assigned to your endpoint.
	EndpointDetails *types.EndpointDetails

	// The type of endpoint that you want your server to use. You can choose to make
	// your server's endpoint publicly accessible (PUBLIC) or host it inside your VPC.
	// With an endpoint that is hosted in a VPC, you can restrict access to your server
	// and resources only within your VPC or choose to make it internet facing by
	// attaching Elastic IP addresses directly to it. After May 19, 2021, you won't be
	// able to create a server using EndpointType=VPC_ENDPOINT in your AWS account if
	// your account hasn't already done so before May 19, 2021. If you have already
	// created servers with EndpointType=VPC_ENDPOINT in your AWS account on or before
	// May 19, 2021, you will not be affected. After this date, use EndpointType=VPC.
	// For more information, see
	// https://docs.aws.amazon.com/transfer/latest/userguide/create-server-in-vpc.html#deprecate-vpc-endpoint.
	// It is recommended that you use VPC as the EndpointType. With this endpoint type,
	// you have the option to directly associate up to three Elastic IPv4 addresses
	// (BYO IP included) with your server's endpoint and use VPC security groups to
	// restrict traffic by the client's public IP address. This is not possible with
	// EndpointType set to VPC_ENDPOINT.
	EndpointType types.EndpointType

	// The RSA private key as generated by the ssh-keygen -N "" -m PEM -f
	// my-new-server-key command. If you aren't planning to migrate existing users from
	// an existing SFTP-enabled server to a new server, don't update the host key.
	// Accidentally changing a server's host key can be disruptive. For more
	// information, see Change the host key for your SFTP-enabled server
	// (https://docs.aws.amazon.com/transfer/latest/userguide/edit-server-config.html#configuring-servers-change-host-key)
	// in the AWS Transfer Family User Guide.
	HostKey *string

	// Required when IdentityProviderType is set to AWS_DIRECTORY_SERVICE or
	// API_GATEWAY. Accepts an array containing all of the information required to use
	// a directory in AWS_DIRECTORY_SERVICE or invoke a customer-supplied
	// authentication API, including the API Gateway URL. Not required when
	// IdentityProviderType is set to SERVICE_MANAGED.
	IdentityProviderDetails *types.IdentityProviderDetails

	// Specifies the mode of authentication for a server. The default value is
	// SERVICE_MANAGED, which allows you to store and access user credentials within
	// the AWS Transfer Family service. Use AWS_DIRECTORY_SERVICE to provide access to
	// Active Directory groups in AWS Managed Active Directory or Microsoft Active
	// Directory in your on-premises environment or in AWS using AD Connectors. This
	// option also requires you to provide a Directory ID using the
	// IdentityProviderDetails parameter. Use the API_GATEWAY value to integrate with
	// an identity provider of your choosing. The API_GATEWAY setting requires you to
	// provide an API Gateway endpoint URL to call for authentication using the
	// IdentityProviderDetails parameter.
	IdentityProviderType types.IdentityProviderType

	// Allows the service to write your users' activity to your Amazon CloudWatch logs
	// for monitoring and auditing purposes.
	LoggingRole *string

	// Specifies the file transfer protocol or protocols over which your file transfer
	// protocol client can connect to your server's endpoint. The available protocols
	// are:
	//
	// * SFTP (Secure Shell (SSH) File Transfer Protocol): File transfer over
	// SSH
	//
	// * FTPS (File Transfer Protocol Secure): File transfer with TLS
	// encryption
	//
	// * FTP (File Transfer Protocol): Unencrypted file transfer
	//
	// If you
	// select FTPS, you must choose a certificate stored in AWS Certificate Manager
	// (ACM) which will be used to identify your server when clients connect to it over
	// FTPS. If Protocol includes either FTP or FTPS, then the EndpointType must be VPC
	// and the IdentityProviderType must be AWS_DIRECTORY_SERVICE or API_GATEWAY. If
	// Protocol includes FTP, then AddressAllocationIds cannot be associated. If
	// Protocol is set only to SFTP, the EndpointType can be set to PUBLIC and the
	// IdentityProviderType can be set to SERVICE_MANAGED.
	Protocols []types.Protocol

	// Specifies the name of the security policy that is attached to the server.
	SecurityPolicyName *string

	// Key-value pairs that can be used to group and search for servers.
	Tags []types.Tag
}

type CreateServerOutput struct {

	// The service-assigned ID of the server that is created.
	//
	// This member is required.
	ServerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateServer{}, middleware.After)
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
	if err = addOpCreateServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateServer",
	}
}
