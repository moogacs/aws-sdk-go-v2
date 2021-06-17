// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon Cognito user pool and sets the password policy for the
// pool. This action might generate an SMS text message. Starting June 1, 2021,
// U.S. telecom carriers require that you register an origination phone number
// before you can send SMS messages to U.S. phone numbers. If you use SMS text
// messages in Amazon Cognito, you must register a phone number with Amazon
// Pinpoint (https://console.aws.amazon.com/pinpoint/home/). Cognito will use the
// the registered number automatically. Otherwise, Cognito users that must receive
// SMS messages might be unable to sign up, activate their accounts, or sign in. If
// you have never used SMS text messages with Amazon Cognito or any other AWS
// service, Amazon SNS might place your account in SMS sandbox. In sandbox mode
// (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html) , you’ll have
// limitations, such as sending messages to only verified phone numbers. After
// testing in the sandbox environment, you can move out of the SMS sandbox and into
// production. For more information, see  SMS message settings for Cognito User
// Pools
// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-sms-userpool-settings.html)
// in the Amazon Cognito Developer Guide.
func (c *Client) CreateUserPool(ctx context.Context, params *CreateUserPoolInput, optFns ...func(*Options)) (*CreateUserPoolOutput, error) {
	if params == nil {
		params = &CreateUserPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserPool", params, optFns, c.addOperationCreateUserPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to create a user pool.
type CreateUserPoolInput struct {

	// A string used to name the user pool.
	//
	// This member is required.
	PoolName *string

	// Use this setting to define which verified available method a user can use to
	// recover their password when they call ForgotPassword. It allows you to define a
	// preferred method when a user has more than one method available. With this
	// setting, SMS does not qualify for a valid password recovery mechanism if the
	// user also has SMS MFA enabled. In the absence of this setting, Cognito uses the
	// legacy behavior to determine the recovery method where SMS is preferred over
	// email.
	AccountRecoverySetting *types.AccountRecoverySettingType

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *types.AdminCreateUserConfigType

	// Attributes supported as an alias for this user pool. Possible values:
	// phone_number, email, or preferred_username.
	AliasAttributes []types.AliasAttributeType

	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes []types.VerifiedAttributeType

	// The device configuration.
	DeviceConfiguration *types.DeviceConfigurationType

	// The email configuration.
	EmailConfiguration *types.EmailConfigurationType

	// A string representing the email verification message. EmailVerificationMessage
	// is allowed only if EmailSendingAccount
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount)
	// is DEVELOPER.
	EmailVerificationMessage *string

	// A string representing the email verification subject. EmailVerificationSubject
	// is allowed only if EmailSendingAccount
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_EmailConfigurationType.html#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount)
	// is DEVELOPER.
	EmailVerificationSubject *string

	// The Lambda trigger configuration information for the new user pool. In a push
	// model, event sources (such as Amazon S3 and custom applications) need permission
	// to invoke a function. So you will need to make an extra call to add permission
	// for these event sources to invoke your Lambda function. For more information on
	// using the Lambda API to add permission, see  AddPermission
	// (https://docs.aws.amazon.com/lambda/latest/dg/API_AddPermission.html). For
	// adding permission using the AWS CLI, see  add-permission
	// (https://docs.aws.amazon.com/cli/latest/reference/lambda/add-permission.html).
	LambdaConfig *types.LambdaConfigType

	// Specifies MFA configuration details.
	MfaConfiguration types.UserPoolMfaType

	// The policies associated with the new user pool.
	Policies *types.UserPoolPolicyType

	// An array of schema attributes for the new user pool. These attributes can be
	// standard or custom attributes.
	Schema []types.SchemaAttributeType

	// A string representing the SMS authentication message.
	SmsAuthenticationMessage *string

	// The SMS configuration.
	SmsConfiguration *types.SmsConfigurationType

	// A string representing the SMS verification message.
	SmsVerificationMessage *string

	// Used to enable advanced security risk detection. Set the key
	// AdvancedSecurityMode to the value "AUDIT".
	UserPoolAddOns *types.UserPoolAddOnsType

	// The tag keys and values to assign to the user pool. A tag is a label that you
	// can use to categorize and manage user pools in different ways, such as by
	// purpose, owner, environment, or other criteria.
	UserPoolTags map[string]string

	// Specifies whether email addresses or phone numbers can be specified as usernames
	// when a user signs up.
	UsernameAttributes []types.UsernameAttributeType

	// You can choose to set case sensitivity on the username input for the selected
	// sign-in option. For example, when this is set to False, users will be able to
	// sign in using either "username" or "Username". This configuration is immutable
	// once it has been set. For more information, see UsernameConfigurationType
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UsernameConfigurationType.html).
	UsernameConfiguration *types.UsernameConfigurationType

	// The template for the verification message that the user sees when the app
	// requests permission to access the user's information.
	VerificationMessageTemplate *types.VerificationMessageTemplateType
}

// Represents the response from the server for the request to create a user pool.
type CreateUserPoolOutput struct {

	// A container for the user pool details.
	UserPool *types.UserPoolType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateUserPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserPool{}, middleware.After)
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
	if err = addOpCreateUserPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUserPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateUserPool",
	}
}
