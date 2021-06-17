// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a dashboard from a template. To first create a template, see the
// CreateTemplate API operation. A dashboard is an entity in QuickSight that
// identifies QuickSight reports, created from analyses. You can share QuickSight
// dashboards. With the right permissions, you can create scheduled email reports
// from them. If you have the correct permissions, you can create a dashboard from
// a template that exists in a different AWS account.
func (c *Client) CreateDashboard(ctx context.Context, params *CreateDashboardInput, optFns ...func(*Options)) (*CreateDashboardOutput, error) {
	if params == nil {
		params = &CreateDashboardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDashboard", params, optFns, c.addOperationCreateDashboardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDashboardInput struct {

	// The ID of the AWS account where you want to create the dashboard.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard, also added to the IAM policy.
	//
	// This member is required.
	DashboardId *string

	// The display name of the dashboard.
	//
	// This member is required.
	Name *string

	// The entity that you are using as a source when you create the dashboard. In
	// SourceEntity, you specify the type of object you're using as source. You can
	// only create a dashboard from a template, so you use a SourceTemplate entity. If
	// you need to create a dashboard from an analysis, first convert the analysis to a
	// template by using the CreateTemplate API operation. For SourceTemplate, specify
	// the Amazon Resource Name (ARN) of the source template. The SourceTemplateARN can
	// contain any AWS Account and any QuickSight-supported AWS Region. Use the
	// DataSetReferences entity within SourceTemplate to list the replacement datasets
	// for the placeholders listed in the original. The schema in each dataset must
	// match its placeholder.
	//
	// This member is required.
	SourceEntity *types.DashboardSourceEntity

	// Options for publishing the dashboard when you create it:
	//
	// * AvailabilityStatus
	// for AdHocFilteringOption - This status can be either ENABLED or DISABLED. When
	// this is set to DISABLED, QuickSight disables the left filter pane on the
	// published dashboard, which can be used for ad hoc (one-time) filtering. This
	// option is ENABLED by default.
	//
	// * AvailabilityStatus for ExportToCSVOption - This
	// status can be either ENABLED or DISABLED. The visual option to export data to
	// .CSV format isn't enabled when this is set to DISABLED. This option is ENABLED
	// by default.
	//
	// * VisibilityState for SheetControlsOption - This visibility state
	// can be either COLLAPSED or EXPANDED. This option is COLLAPSED by default.
	DashboardPublishOptions *types.DashboardPublishOptions

	// The parameters for the creation of the dashboard, which you want to use to
	// override the default settings. A dashboard can have any type of parameters, and
	// some parameters might accept multiple values.
	Parameters *types.Parameters

	// A structure that contains the permissions of the dashboard. You can use this
	// structure for granting permissions by providing a list of IAM action information
	// for each principal ARN. To specify no permissions, omit the permissions list.
	Permissions []types.ResourcePermission

	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the dashboard.
	Tags []types.Tag

	// The Amazon Resource Name (ARN) of the theme that is being used for this
	// dashboard. If you add a value for this field, it overrides the value that is
	// used in the source entity. The theme ARN must exist in the same AWS account
	// where you create the dashboard.
	ThemeArn *string

	// A description for the first version of the dashboard being created.
	VersionDescription *string
}

type CreateDashboardOutput struct {

	// The ARN of the dashboard.
	Arn *string

	// The status of the dashboard creation request.
	CreationStatus types.ResourceStatus

	// The ID for the dashboard.
	DashboardId *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ARN of the dashboard, including the version number of the first version that
	// is created.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateDashboardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDashboard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDashboard{}, middleware.After)
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
	if err = addOpCreateDashboardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDashboard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDashboard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateDashboard",
	}
}
