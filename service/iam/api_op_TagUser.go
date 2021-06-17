// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds one or more tags to an IAM user. If a tag with the same key name already
// exists, then that tag is overwritten with the new value. A tag consists of a key
// name and an associated value. By assigning tags to your resources, you can do
// the following:
//
// * Administrative grouping and discovery - Attach tags to
// resources to aid in organization and search. For example, you could search for
// all resources with the key name Project and the value MyImportantProject. Or
// search for all resources with the key name Cost Center and the value 41200.
//
// *
// Access control - Include tags in IAM user-based and resource-based policies. You
// can use tags to restrict access to only an IAM requesting user that has a
// specified tag attached. You can also restrict access to only those resources
// that have a certain tag attached. For examples of policies that show how to use
// tags to control access, see Control access using IAM tags
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html) in the IAM
// User Guide.
//
// * Cost allocation - Use tags to help track which individuals and
// teams are using which AWS resources.
//
// * If any one of the tags is invalid or if
// you exceed the allowed maximum number of tags, then the entire request fails and
// the resource is not created. For more information about tagging, see Tagging IAM
// resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the
// IAM User Guide.
//
// * AWS always interprets the tag Value as a single string. If
// you need to store an array, you can store comma-separated values in the string.
// However, you must interpret the value in your code.
//
// For more information about
// tagging, see Tagging IAM identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User
// Guide.
func (c *Client) TagUser(ctx context.Context, params *TagUserInput, optFns ...func(*Options)) (*TagUserOutput, error) {
	if params == nil {
		params = &TagUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagUser", params, optFns, c.addOperationTagUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagUserInput struct {

	// The list of tags that you want to attach to the IAM user. Each tag consists of a
	// key name and an associated value.
	//
	// This member is required.
	Tags []types.Tag

	// The name of the IAM user to which you want to add tags. This parameter accepts
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters that consist of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: =,.@-
	//
	// This member is required.
	UserName *string
}

type TagUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationTagUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTagUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTagUser{}, middleware.After)
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
	if err = addOpTagUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "TagUser",
	}
}
