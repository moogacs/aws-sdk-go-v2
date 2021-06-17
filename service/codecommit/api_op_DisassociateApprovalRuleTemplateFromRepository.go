// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the association between a template and a repository so that approval
// rules based on the template are not automatically created when pull requests are
// created in the specified repository. This does not delete any approval rules
// previously created for pull requests through the template association.
func (c *Client) DisassociateApprovalRuleTemplateFromRepository(ctx context.Context, params *DisassociateApprovalRuleTemplateFromRepositoryInput, optFns ...func(*Options)) (*DisassociateApprovalRuleTemplateFromRepositoryOutput, error) {
	if params == nil {
		params = &DisassociateApprovalRuleTemplateFromRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateApprovalRuleTemplateFromRepository", params, optFns, c.addOperationDisassociateApprovalRuleTemplateFromRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateApprovalRuleTemplateFromRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateApprovalRuleTemplateFromRepositoryInput struct {

	// The name of the approval rule template to disassociate from a specified
	// repository.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The name of the repository you want to disassociate from the template.
	//
	// This member is required.
	RepositoryName *string
}

type DisassociateApprovalRuleTemplateFromRepositoryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDisassociateApprovalRuleTemplateFromRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateApprovalRuleTemplateFromRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateApprovalRuleTemplateFromRepository{}, middleware.After)
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
	if err = addOpDisassociateApprovalRuleTemplateFromRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateApprovalRuleTemplateFromRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateApprovalRuleTemplateFromRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "DisassociateApprovalRuleTemplateFromRepository",
	}
}
