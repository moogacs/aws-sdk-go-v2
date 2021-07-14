// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Lightsail bucket. A bucket is a cloud storage resource
// available in the Lightsail object storage service. Use buckets to store objects
// such as data and its descriptive metadata. For more information about buckets,
// see Buckets in Amazon Lightsail
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/buckets-in-amazon-lightsail)
// in the Amazon Lightsail Developer Guide.
func (c *Client) CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) {
	if params == nil {
		params = &CreateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucket", params, optFns, c.addOperationCreateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketInput struct {

	// The name for the bucket. For more information about bucket names, see Bucket
	// naming rules in Amazon Lightsail
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/bucket-naming-rules-in-amazon-lightsail)
	// in the Amazon Lightsail Developer Guide.
	//
	// This member is required.
	BucketName *string

	// The ID of the bundle to use for the bucket. A bucket bundle specifies the
	// monthly cost, storage space, and data transfer quota for a bucket. Use the
	// GetBucketBundles action to get a list of bundle IDs that you can specify. Use
	// the UpdateBucketBundle action to change the bundle after the bucket is created.
	//
	// This member is required.
	BundleId *string

	// A Boolean value that indicates whether to enable versioning of objects in the
	// bucket. For more information about versioning, see Enabling and suspending
	// bucket object versioning in Amazon Lightsail
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-managing-bucket-object-versioning)
	// in the Amazon Lightsail Developer Guide.
	EnableObjectVersioning *bool

	// The tag keys and optional values to add to the bucket during creation. Use the
	// TagResource action to tag the bucket after it's created.
	Tags []types.Tag
}

type CreateBucketOutput struct {

	// An object that describes the bucket that is created.
	Bucket *types.Bucket

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBucket{}, middleware.After)
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
	if err = addOpCreateBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucket(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateBucket",
	}
}