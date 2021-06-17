// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new script record for your Realtime Servers script. Realtime scripts
// are JavaScript that provide configuration settings and optional custom game
// logic for your game. The script is deployed when you create a Realtime Servers
// fleet to host your game sessions. Script logic is executed during an active game
// session. To create a new script record, specify a script name and provide the
// script file(s). The script files and all dependencies must be zipped into a
// single file. You can pull the zip file from either of these locations:
//
// * A
// locally available directory. Use the ZipFile parameter for this option.
//
// * An
// Amazon Simple Storage Service (Amazon S3) bucket under your AWS account. Use the
// StorageLocation parameter for this option. You'll need to have an Identity
// Access Management (IAM) role that allows the Amazon GameLift service to access
// your S3 bucket.
//
// If the call is successful, a new script record is created with
// a unique script ID. If the script file is provided as a local file, the file is
// uploaded to an Amazon GameLift-owned S3 bucket and the script record's storage
// location reflects this location. If the script file is provided as an S3 bucket,
// Amazon GameLift accesses the file at this storage location as needed for
// deployment. Learn more Amazon GameLift Realtime Servers
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html)Set
// Up a Role for Amazon GameLift Access
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/setting-up-role.html)
// Related actions CreateScript | ListScripts | DescribeScript | UpdateScript |
// DeleteScript | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreateScript(ctx context.Context, params *CreateScriptInput, optFns ...func(*Options)) (*CreateScriptOutput, error) {
	if params == nil {
		params = &CreateScriptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateScript", params, optFns, c.addOperationCreateScriptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScriptInput struct {

	// A descriptive label that is associated with a script. Script names do not need
	// to be unique. You can use UpdateScript to change this value later.
	Name *string

	// The location of the Amazon S3 bucket where a zipped file containing your
	// Realtime scripts is stored. The storage location must specify the Amazon S3
	// bucket name, the zip file name (the "key"), and a role ARN that allows Amazon
	// GameLift to access the Amazon S3 storage location. The S3 bucket must be in the
	// same Region where you want to create a new script. By default, Amazon GameLift
	// uploads the latest version of the zip file; if you have S3 object versioning
	// turned on, you can use the ObjectVersion parameter to specify an earlier
	// version.
	StorageLocation *types.S3Location

	// A list of labels to assign to the new script resource. Tags are
	// developer-defined key-value pairs. Tagging AWS resources are useful for resource
	// management, access management and cost allocation. For more information, see
	// Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource,
	// UntagResource, and ListTagsForResource to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []types.Tag

	// Version information that is associated with a build or script. Version strings
	// do not need to be unique. You can use UpdateScript to change this value later.
	Version *string

	// A data object containing your Realtime scripts and dependencies as a zip file.
	// The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	// When using the AWS CLI tool to create a script, this parameter is set to the zip
	// file name. It must be prepended with the string "fileb://" to indicate that the
	// file data is a binary object. For example: --zip-file
	// fileb://myRealtimeScript.zip.
	ZipFile []byte
}

type CreateScriptOutput struct {

	// The newly created script record with a unique script ID and ARN. The new
	// script's storage location reflects an Amazon S3 location: (1) If the script was
	// uploaded from an S3 bucket under your account, the storage location reflects the
	// information that was provided in the CreateScript request; (2) If the script
	// file was uploaded from a local zip file, the storage location reflects an S3
	// location controls by the Amazon GameLift service.
	Script *types.Script

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateScriptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateScript{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateScript{}, middleware.After)
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
	if err = addOpCreateScriptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScript(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateScript(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateScript",
	}
}
