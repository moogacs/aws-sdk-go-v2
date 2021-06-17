// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Simple AD directory. For more information, see Simple Active Directory
// (https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_simple_ad.html)
// in the AWS Directory Service Admin Guide. Before you call CreateDirectory,
// ensure that all of the required permissions have been explicitly granted through
// a policy. For details about what permissions are required to run the
// CreateDirectory operation, see AWS Directory Service API Permissions: Actions,
// Resources, and Conditions Reference
// (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
func (c *Client) CreateDirectory(ctx context.Context, params *CreateDirectoryInput, optFns ...func(*Options)) (*CreateDirectoryOutput, error) {
	if params == nil {
		params = &CreateDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDirectory", params, optFns, c.addOperationCreateDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the CreateDirectory operation.
type CreateDirectoryInput struct {

	// The fully qualified name for the directory, such as corp.example.com.
	//
	// This member is required.
	Name *string

	// The password for the directory administrator. The directory creation process
	// creates a directory administrator account with the user name Administrator and
	// this password. If you need to change the password for the administrator account,
	// you can use the ResetUserPassword API call. The regex pattern for this string is
	// made up of the following conditions:
	//
	// * Length (?=^.{8,64}$) – Must be between 8
	// and 64 characters
	//
	// AND any 3 of the following password complexity rules required
	// by Active Directory:
	//
	// * Numbers and upper case and lowercase
	// (?=.*\d)(?=.*[A-Z])(?=.*[a-z])
	//
	// * Numbers and special characters and lower case
	// (?=.*\d)(?=.*[^A-Za-z0-9\s])(?=.*[a-z])
	//
	// * Special characters and upper case and
	// lower case (?=.*[^A-Za-z0-9\s])(?=.*[A-Z])(?=.*[a-z])
	//
	// * Numbers and upper case
	// and special characters (?=.*\d)(?=.*[A-Z])(?=.*[^A-Za-z0-9\s])
	//
	// For additional
	// information about how Active Directory passwords are enforced, see Password must
	// meet complexity requirements
	// (https://docs.microsoft.com/en-us/windows/security/threat-protection/security-policy-settings/password-must-meet-complexity-requirements)
	// on the Microsoft website.
	//
	// This member is required.
	Password *string

	// The size of the directory.
	//
	// This member is required.
	Size types.DirectorySize

	// A description for the directory.
	Description *string

	// The NetBIOS name of the directory, such as CORP.
	ShortName *string

	// The tags to be assigned to the Simple AD directory.
	Tags []types.Tag

	// A DirectoryVpcSettings object that contains additional information for the
	// operation.
	VpcSettings *types.DirectoryVpcSettings
}

// Contains the results of the CreateDirectory operation.
type CreateDirectoryOutput struct {

	// The identifier of the directory that was created.
	DirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDirectory{}, middleware.After)
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
	if err = addOpCreateDirectoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDirectory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CreateDirectory",
	}
}
