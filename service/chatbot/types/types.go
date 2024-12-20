// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Preferences related to AWS Chatbot usage in the calling AWS account.
type AccountPreferences struct {

	// Turns on training data collection.
	//
	// This helps improve the AWS Chatbot experience by allowing AWS Chatbot to store
	// and use your customer information, such as AWS Chatbot configurations,
	// notifications, user inputs, AWS Chatbot generated responses, and interaction
	// data. This data helps us to continuously improve and develop Artificial
	// Intelligence (AI) technologies. Your data is not shared with any third parties
	// and is protected using sophisticated controls to prevent unauthorized access and
	// misuse. AWS Chatbot does not store or use interactions in chat channels with
	// Amazon Q for training AI technologies for AWS Chatbot.
	TrainingDataCollectionEnabled *bool

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired *bool

	noSmithyDocumentSerde
}

// A listing of an association with a channel configuration.
type AssociationListing struct {

	// The Amazon Resource Name (ARN) of the resource (for example, a custom action).
	//
	// This member is required.
	Resource *string

	noSmithyDocumentSerde
}

// An AWS Chatbot configuration for Amazon Chime.
type ChimeWebhookConfiguration struct {

	// The Amazon Resource Name (ARN) of the ChimeWebhookConfiguration.
	//
	// This member is required.
	ChatConfigurationArn *string

	// A user-defined role that AWS Chatbot assumes. This is not the service-linked
	// role.
	//
	// For more information, see [IAM policies for AWS Chatbot] in the AWS Chatbot Administrator Guide.
	//
	// [IAM policies for AWS Chatbot]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html
	//
	// This member is required.
	IamRoleArn *string

	// The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications
	// to AWS Chatbot.
	//
	// This member is required.
	SnsTopicArns []string

	// A description of the webhook. We recommend using the convention
	// RoomName/WebhookName .
	//
	// For more information, see [Tutorial: Get started with Amazon Chime] in the AWS Chatbot Administrator Guide.
	//
	// [Tutorial: Get started with Amazon Chime]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chime-setup.html
	//
	// This member is required.
	WebhookDescription *string

	// The name of the configuration.
	ConfigurationName *string

	// Logging levels include ERROR , INFO , or NONE .
	LoggingLevel *string

	// Either ENABLED or DISABLED . The resource returns DISABLED if the
	// organization's AWS Chatbot policy has explicitly denied that configuration. For
	// example, if Amazon Chime is disabled.
	State *string

	// Provided if State is DISABLED . Provides context as to why the resource is
	// disabled.
	StateReason *string

	// A map of tags assigned to a resource. A tag is a string-to-string map of
	// key-value pairs.
	Tags []Tag

	noSmithyDocumentSerde
}

// A Microsoft Teams team that is authorized with AWS Chatbot.
type ConfiguredTeam struct {

	//  The ID of the Microsoft Teams authorized with AWS Chatbot.
	//
	// To get the team ID, you must perform the initial authorization flow with
	// Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the team
	// ID from the console. For more information, see [Step 1: Configure a Microsoft Teams client]in the AWS Chatbot Administrator
	// Guide.
	//
	// [Step 1: Configure a Microsoft Teams client]: https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup
	//
	// This member is required.
	TeamId *string

	// The ID of the Microsoft Teams tenant.
	//
	// This member is required.
	TenantId *string

	// Either ENABLED or DISABLED . The resource returns DISABLED if the
	// organization's AWS Chatbot policy has explicitly denied that configuration. For
	// example, if Amazon Chime is disabled.
	State *string

	// Provided if State is DISABLED . Provides context as to why the resource is
	// disabled.
	StateReason *string

	// The name of the Microsoft Teams Team.
	TeamName *string

	noSmithyDocumentSerde
}

// Represents a parameterized command that can be invoked as an alias or as a
// notification button in the chat client.
type CustomAction struct {

	// The fully defined Amazon Resource Name (ARN) of the custom action.
	//
	// This member is required.
	CustomActionArn *string

	// The definition of the command to run when invoked an alias or as an action
	// button.
	//
	// This member is required.
	Definition *CustomActionDefinition

	// The name of the custom action that is included in the ARN.
	ActionName *string

	// The name used to invoke this action in the chat channel. For example, @aws run
	// my-alias .
	AliasName *string

	// Defines when this custom action button should be attached to a notification.
	Attachments []CustomActionAttachment

	noSmithyDocumentSerde
}

// Defines when a custom action button should be attached to a notification.
type CustomActionAttachment struct {

	// The text of the button that appears on the notification.
	ButtonText *string

	// The criteria for when a button should be shown based on values in the
	// notification.
	Criteria []CustomActionAttachmentCriteria

	// The type of notification that the custom action should be attached to.
	NotificationType *string

	// The variables to extract from the notification.
	Variables map[string]string

	noSmithyDocumentSerde
}

// A criteria for when a button should be shown based on values in the notification
type CustomActionAttachmentCriteria struct {

	// The operation to perform on the named variable.
	//
	// This member is required.
	Operator CustomActionAttachmentCriteriaOperator

	// The name of the variable to operate on.
	//
	// This member is required.
	VariableName *string

	// A value that is compared with the actual value of the variable based on the
	// behavior of the operator.
	Value *string

	noSmithyDocumentSerde
}

// The definition of the command to run when invoked as an alias or as an action
// button.
type CustomActionDefinition struct {

	// The command string to run which may include variables by prefixing with a
	// dollar sign ($).
	//
	// This member is required.
	CommandText *string

	noSmithyDocumentSerde
}

// An AWS Chatbot configuration for Slack.
type SlackChannelConfiguration struct {

	// The Amazon Resource Name (ARN) of the SlackChannelConfiguration.
	//
	// This member is required.
	ChatConfigurationArn *string

	// A user-defined role that AWS Chatbot assumes. This is not the service-linked
	// role.
	//
	// For more information, see [IAM policies for AWS Chatbot] in the AWS Chatbot Administrator Guide.
	//
	// [IAM policies for AWS Chatbot]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html
	//
	// This member is required.
	IamRoleArn *string

	// The ID of the Slack channel.
	//
	// To get this ID, open Slack, right click on the channel name in the left pane,
	// then choose Copy Link. The channel ID is the 9-character string at the end of
	// the URL. For example, ABCBBLZZZ.
	//
	// This member is required.
	SlackChannelId *string

	// The name of the Slack channel.
	//
	// This member is required.
	SlackChannelName *string

	// The ID of the Slack workspace authorized with Amazon Chime.
	//
	// This member is required.
	SlackTeamId *string

	// Name of the Slack workspace.
	//
	// This member is required.
	SlackTeamName *string

	// The ARNs of the SNS topics that deliver notifications to AWS Chatbot.
	//
	// This member is required.
	SnsTopicArns []string

	// The name of the configuration.
	ConfigurationName *string

	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS
	// managed AdministratorAccess policy is applied by default if this is not set.
	GuardrailPolicyArns []string

	// Logging levels include ERROR , INFO , or NONE .
	LoggingLevel *string

	// Either ENABLED or DISABLED . The resource returns DISABLED if the
	// organization's AWS Chatbot policy has explicitly denied that configuration. For
	// example, if Amazon Chime is disabled.
	State *string

	// Provided if State is DISABLED . Provides context as to why the resource is
	// disabled.
	StateReason *string

	// A map of tags assigned to a resource. A tag is a string-to-string map of
	// key-value pairs.
	Tags []Tag

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired *bool

	noSmithyDocumentSerde
}

// Identifes a user level permission for a channel configuration.
type SlackUserIdentity struct {

	// The Amazon Resource Name (ARN) of the SlackChannelConfiguration associated with
	// the user identity to delete.
	//
	// This member is required.
	ChatConfigurationArn *string

	// A user-defined role that AWS Chatbot assumes. This is not the service-linked
	// role.
	//
	// For more information, see [IAM policies for AWS Chatbot] in the AWS Chatbot Administrator Guide.
	//
	// [IAM policies for AWS Chatbot]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html
	//
	// This member is required.
	IamRoleArn *string

	// The ID of the Slack workspace authorized with AWS Chatbot.
	//
	// This member is required.
	SlackTeamId *string

	// The ID of the user in Slack
	//
	// This member is required.
	SlackUserId *string

	// The AWS user identity ARN used to associate a Slack user ID with an IAM Role.
	AwsUserIdentity *string

	noSmithyDocumentSerde
}

// A Slack workspace.
type SlackWorkspace struct {

	// The ID of the Slack workspace authorized with AWS Chatbot.
	//
	// This member is required.
	SlackTeamId *string

	// The name of the Slack workspace.
	//
	// This member is required.
	SlackTeamName *string

	// Either ENABLED or DISABLED . The resource returns DISABLED if the
	// organization's AWS Chatbot policy has explicitly denied that configuration. For
	// example, if Amazon Chime is disabled.
	State *string

	// Provided if State is DISABLED . Provides context as to why the resource is
	// disabled.
	StateReason *string

	noSmithyDocumentSerde
}

// A key-value pair. A tag consists of a tag key and a tag value. Tag keys and tag
// values are both required, but tag values can be empty (null) strings.
//
// Do not include confidential or sensitive information in this field.
//
// For more information, see [User-Defined Tag Restrictions] in the AWS Billing and Cost Management User Guide.
//
// [User-Defined Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	TagKey *string

	// The value of the tag.
	//
	// This member is required.
	TagValue *string

	noSmithyDocumentSerde
}

// An AWS Chatbot configuration for Microsoft Teams.
type TeamsChannelConfiguration struct {

	// The ID of the Microsoft Teams channel.
	//
	// This member is required.
	ChannelId *string

	// The Amazon Resource Name (ARN) of the MicrosoftTeamsChannelConfiguration
	// associated with the user identity to delete.
	//
	// This member is required.
	ChatConfigurationArn *string

	// A user-defined role that AWS Chatbot assumes. This is not the service-linked
	// role.
	//
	// For more information, see [IAM policies for AWS Chatbot] in the AWS Chatbot Administrator Guide.
	//
	// [IAM policies for AWS Chatbot]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html
	//
	// This member is required.
	IamRoleArn *string

	// The Amazon Resource Names (ARNs) of the SNS topics that deliver notifications
	// to AWS Chatbot.
	//
	// This member is required.
	SnsTopicArns []string

	//  The ID of the Microsoft Teams authorized with AWS Chatbot.
	//
	// To get the team ID, you must perform the initial authorization flow with
	// Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the team
	// ID from the console. For more information, see [Step 1: Configure a Microsoft Teams client]in the AWS Chatbot Administrator
	// Guide.
	//
	// [Step 1: Configure a Microsoft Teams client]: https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup
	//
	// This member is required.
	TeamId *string

	// The ID of the Microsoft Teams tenant.
	//
	// This member is required.
	TenantId *string

	// The name of the Microsoft Teams channel.
	ChannelName *string

	// The name of the configuration.
	ConfigurationName *string

	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS
	// managed AdministratorAccess policy is applied by default if this is not set.
	GuardrailPolicyArns []string

	// Logging levels include ERROR , INFO , or NONE .
	LoggingLevel *string

	// Either ENABLED or DISABLED . The resource returns DISABLED if the
	// organization's AWS Chatbot policy has explicitly denied that configuration. For
	// example, if Amazon Chime is disabled.
	State *string

	// Provided if State is DISABLED . Provides context as to why the resource is
	// disabled.
	StateReason *string

	// A map of tags assigned to a resource. A tag is a string-to-string map of
	// key-value pairs.
	Tags []Tag

	// The name of the Microsoft Teams Team.
	TeamName *string

	// Enables use of a user role requirement in your chat configuration.
	UserAuthorizationRequired *bool

	noSmithyDocumentSerde
}

// Identifes a user level permission for a channel configuration.
type TeamsUserIdentity struct {

	// The Amazon Resource Name (ARN) of the MicrosoftTeamsChannelConfiguration
	// associated with the user identity to delete.
	//
	// This member is required.
	ChatConfigurationArn *string

	// A user-defined role that AWS Chatbot assumes. This is not the service-linked
	// role.
	//
	// For more information, see [IAM policies for AWS Chatbot] in the AWS Chatbot Administrator Guide.
	//
	// [IAM policies for AWS Chatbot]: https://docs.aws.amazon.com/chatbot/latest/adminguide/chatbot-iam-policies.html
	//
	// This member is required.
	IamRoleArn *string

	//  The ID of the Microsoft Teams authorized with AWS Chatbot.
	//
	// To get the team ID, you must perform the initial authorization flow with
	// Microsoft Teams in the AWS Chatbot console. Then you can copy and paste the team
	// ID from the console. For more information, see [Step 1: Configure a Microsoft Teams client]in the AWS Chatbot Administrator
	// Guide.
	//
	// [Step 1: Configure a Microsoft Teams client]: https://docs.aws.amazon.com/chatbot/latest/adminguide/teams-setup.html#teams-client-setup
	//
	// This member is required.
	TeamId *string

	// The AWS user identity ARN used to associate a Microsoft Teams user Identity
	// with an IAM Role.
	AwsUserIdentity *string

	// The ID of the Microsoft Teams channel.
	TeamsChannelId *string

	// The ID of the Microsoft Teams tenant.
	TeamsTenantId *string

	// The Microsoft Teams user ID.
	UserId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
