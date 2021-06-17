// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a deployment.
func (c *Client) StartDeployment(ctx context.Context, params *StartDeploymentInput, optFns ...func(*Options)) (*StartDeploymentOutput, error) {
	if params == nil {
		params = &StartDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDeployment", params, optFns, c.addOperationStartDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDeploymentInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The configuration profile ID.
	//
	// This member is required.
	ConfigurationProfileId *string

	// The configuration version to deploy.
	//
	// This member is required.
	ConfigurationVersion *string

	// The deployment strategy ID.
	//
	// This member is required.
	DeploymentStrategyId *string

	// The environment ID.
	//
	// This member is required.
	EnvironmentId *string

	// A description of the deployment.
	Description *string

	// Metadata to assign to the deployment. Tags help organize and categorize your
	// AppConfig resources. Each tag consists of a key and an optional value, both of
	// which you define.
	Tags map[string]string
}

type StartDeploymentOutput struct {

	// The ID of the application that was deployed.
	ApplicationId *string

	// The time the deployment completed.
	CompletedAt *time.Time

	// Information about the source location of the configuration.
	ConfigurationLocationUri *string

	// The name of the configuration.
	ConfigurationName *string

	// The ID of the configuration profile that was deployed.
	ConfigurationProfileId *string

	// The configuration version that was deployed.
	ConfigurationVersion *string

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes int32

	// The sequence number of the deployment.
	DeploymentNumber int32

	// The ID of the deployment strategy that was deployed.
	DeploymentStrategyId *string

	// The description of the deployment.
	Description *string

	// The ID of the environment that was deployed.
	EnvironmentId *string

	// A list containing all events related to a deployment. The most recent events are
	// displayed first.
	EventLog []types.DeploymentEvent

	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes int32

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor float32

	// The algorithm used to define how percentage grew over time.
	GrowthType types.GrowthType

	// The percentage of targets for which the deployment is available.
	PercentageComplete float32

	// The time the deployment started.
	StartedAt *time.Time

	// The state of the deployment.
	State types.DeploymentState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationStartDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartDeployment{}, middleware.After)
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
	if err = addOpStartDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "StartDeployment",
	}
}
