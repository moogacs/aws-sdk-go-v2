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

// Creates a detector version. The detector version starts in a DRAFT status.
func (c *Client) CreateDetectorVersion(ctx context.Context, params *CreateDetectorVersionInput, optFns ...func(*Options)) (*CreateDetectorVersionOutput, error) {
	if params == nil {
		params = &CreateDetectorVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDetectorVersion", params, optFns, c.addOperationCreateDetectorVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDetectorVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDetectorVersionInput struct {

	// The ID of the detector under which you want to create a new version.
	//
	// This member is required.
	DetectorId *string

	// The rules to include in the detector version.
	//
	// This member is required.
	Rules []types.Rule

	// The description of the detector version.
	Description *string

	// The Amazon Sagemaker model endpoints to include in the detector version.
	ExternalModelEndpoints []string

	// The model versions to include in the detector version.
	ModelVersions []types.ModelVersion

	// The rule execution mode for the rules included in the detector version. You can
	// define and edit the rule mode at the detector version level, when it is in draft
	// status. If you specify FIRST_MATCHED, Amazon Fraud Detector evaluates rules
	// sequentially, first to last, stopping at the first matched rule. Amazon Fraud
	// dectector then provides the outcomes for that single rule. If you specifiy
	// ALL_MATCHED, Amazon Fraud Detector evaluates all rules and returns the outcomes
	// for all matched rules. The default behavior is FIRST_MATCHED.
	RuleExecutionMode types.RuleExecutionMode

	// A collection of key and value pairs.
	Tags []types.Tag
}

type CreateDetectorVersionOutput struct {

	// The ID for the created version's parent detector.
	DetectorId *string

	// The ID for the created detector.
	DetectorVersionId *string

	// The status of the detector version.
	Status types.DetectorVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateDetectorVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDetectorVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDetectorVersion{}, middleware.After)
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
	if err = addOpCreateDetectorVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDetectorVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDetectorVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "CreateDetectorVersion",
	}
}
