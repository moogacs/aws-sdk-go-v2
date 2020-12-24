// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Edits an AWS Audit Manager assessment.
func (c *Client) UpdateAssessment(ctx context.Context, params *UpdateAssessmentInput, optFns ...func(*Options)) (*UpdateAssessmentOutput, error) {
	if params == nil {
		params = &UpdateAssessmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAssessment", params, optFns, addOperationUpdateAssessmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssessmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssessmentInput struct {

	// The identifier for the specified assessment.
	//
	// This member is required.
	AssessmentId *string

	// The scope of the specified assessment.
	//
	// This member is required.
	Scope *types.Scope

	// The description of the specified assessment.
	AssessmentDescription *string

	// The name of the specified assessment to be updated.
	AssessmentName *string

	// The assessment report storage destination for the specified assessment that is
	// being updated.
	AssessmentReportsDestination *types.AssessmentReportsDestination

	// The list of roles for the specified assessment.
	Roles []types.Role
}

type UpdateAssessmentOutput struct {

	// The response object (name of the updated assessment) for the
	// UpdateAssessmentRequest API.
	Assessment *types.Assessment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAssessmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAssessment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAssessment{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpUpdateAssessmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAssessment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAssessment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "UpdateAssessment",
	}
}