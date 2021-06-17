// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a detector version. The detector version attributes that you can update
// include models, external model endpoints, rules, rule execution mode, and
// description. You can only update a DRAFT detector version.
func (c *Client) UpdateDetectorVersion(ctx context.Context, params *UpdateDetectorVersionInput, optFns ...func(*Options)) (*UpdateDetectorVersionOutput, error) {
	if params == nil {
		params = &UpdateDetectorVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDetectorVersion", params, optFns, c.addOperationUpdateDetectorVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDetectorVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDetectorVersionInput struct {

	// The parent detector ID for the detector version you want to update.
	//
	// This member is required.
	DetectorId *string

	// The detector version ID.
	//
	// This member is required.
	DetectorVersionId *string

	// The Amazon SageMaker model endpoints to include in the detector version.
	//
	// This member is required.
	ExternalModelEndpoints []string

	// The rules to include in the detector version.
	//
	// This member is required.
	Rules []types.Rule

	// The detector version description.
	Description *string

	// The model versions to include in the detector version.
	ModelVersions []types.ModelVersion

	// The rule execution mode to add to the detector. If you specify FIRST_MATCHED,
	// Amazon Fraud Detector evaluates rules sequentially, first to last, stopping at
	// the first matched rule. Amazon Fraud dectector then provides the outcomes for
	// that single rule. If you specifiy ALL_MATCHED, Amazon Fraud Detector evaluates
	// all rules and returns the outcomes for all matched rules. You can define and
	// edit the rule mode at the detector version level, when it is in draft status.
	// The default behavior is FIRST_MATCHED.
	RuleExecutionMode types.RuleExecutionMode
}

type UpdateDetectorVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateDetectorVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDetectorVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDetectorVersion{}, middleware.After)
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
	if err = addOpUpdateDetectorVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDetectorVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDetectorVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "UpdateDetectorVersion",
	}
}
