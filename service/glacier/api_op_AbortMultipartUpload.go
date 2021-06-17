// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation aborts a multipart upload identified by the upload ID. After the
// Abort Multipart Upload request succeeds, you cannot upload any more parts to the
// multipart upload or complete the multipart upload. Aborting a completed upload
// fails. However, aborting an already-aborted upload will succeed, for a short
// time. For more information about uploading a part and completing a multipart
// upload, see UploadMultipartPart and CompleteMultipartUpload. This operation is
// idempotent. An AWS account has full permission to perform all operations
// (actions). However, AWS Identity and Access Management (IAM) users don't have
// any permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see Access Control Using AWS Identity
// and Access Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and underlying REST API, see Working with Archives in
// Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html)
// and Abort Multipart Upload
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-abort-upload.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) AbortMultipartUpload(ctx context.Context, params *AbortMultipartUploadInput, optFns ...func(*Options)) (*AbortMultipartUploadOutput, error) {
	if params == nil {
		params = &AbortMultipartUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AbortMultipartUpload", params, optFns, c.addOperationAbortMultipartUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AbortMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to abort a multipart upload identified by the upload ID. For
// information about the underlying REST API, see Abort Multipart Upload
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-abort-upload.html).
// For conceptual information, see Working with Archives in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html).
type AbortMultipartUploadInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The upload ID of the multipart upload to delete.
	//
	// This member is required.
	UploadId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

type AbortMultipartUploadOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAbortMultipartUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAbortMultipartUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAbortMultipartUpload{}, middleware.After)
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
	if err = addOpAbortMultipartUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAbortMultipartUpload(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAbortMultipartUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "AbortMultipartUpload",
	}
}
