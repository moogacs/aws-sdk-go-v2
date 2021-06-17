// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// AddJobFlowSteps adds new steps to a running cluster. A maximum of 256 steps are
// allowed in each job flow. If your cluster is long-running (such as a Hive data
// warehouse) or complex, you may require more than 256 steps to process your data.
// You can bypass the 256-step limitation in various ways, including using SSH to
// connect to the master node and submitting queries directly to the software
// running on the master node, such as Hive and Hadoop. For more information on how
// to do this, see Add More than 256 Steps to a Cluster
// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/AddMoreThan256Steps.html)
// in the Amazon EMR Management Guide. A step specifies the location of a JAR file
// stored either on the master node of the cluster or in Amazon S3. Each step is
// performed by the main function of the main class of the JAR file. The main class
// can be specified either in the manifest of the JAR or by using the MainFunction
// parameter of the step. Amazon EMR executes each step in the order listed. For a
// step to be considered complete, the main function must exit with a zero exit
// code and all Hadoop jobs started while the step was running must have completed
// and run successfully. You can only add steps to a cluster that is in one of the
// following states: STARTING, BOOTSTRAPPING, RUNNING, or WAITING.
func (c *Client) AddJobFlowSteps(ctx context.Context, params *AddJobFlowStepsInput, optFns ...func(*Options)) (*AddJobFlowStepsOutput, error) {
	if params == nil {
		params = &AddJobFlowStepsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddJobFlowSteps", params, optFns, c.addOperationAddJobFlowStepsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddJobFlowStepsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input argument to the AddJobFlowSteps operation.
type AddJobFlowStepsInput struct {

	// A string that uniquely identifies the job flow. This identifier is returned by
	// RunJobFlow and can also be obtained from ListClusters.
	//
	// This member is required.
	JobFlowId *string

	// A list of StepConfig to be executed by the job flow.
	//
	// This member is required.
	Steps []types.StepConfig
}

// The output for the AddJobFlowSteps operation.
type AddJobFlowStepsOutput struct {

	// The identifiers of the list of steps added to the job flow.
	StepIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAddJobFlowStepsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddJobFlowSteps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddJobFlowSteps{}, middleware.After)
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
	if err = addOpAddJobFlowStepsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddJobFlowSteps(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddJobFlowSteps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "AddJobFlowSteps",
	}
}
