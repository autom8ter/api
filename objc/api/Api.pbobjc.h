// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: api.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

@class Access;
@class Address;
@class AttachmentActionOption;
@class AttachmentActionOptionGroup;
@class AttachmentConfirmationField;
@class AttachmentField;
@class EmailAddress;
@class ItemRef;
@class JSON;
@class PubsubMessage;
@class RecipientEmail;
@class Topic;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum CustomerIndex

typedef GPB_ENUM(CustomerIndex) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  CustomerIndex_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  CustomerIndex_Id = 0,
  CustomerIndex_Email = 1,
  CustomerIndex_Phone = 2,
};

GPBEnumDescriptor *CustomerIndex_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL CustomerIndex_IsValidValue(int32_t value);

#pragma mark - Enum Grant

typedef GPB_ENUM(Grant) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  Grant_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  Grant_Twilio = 0,
  Grant_Sendgrid = 1,
  Grant_Stripe = 2,
  Grant_Slack = 3,
  Grant_Gcp = 4,
};

GPBEnumDescriptor *Grant_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL Grant_IsValidValue(int32_t value);

#pragma mark - ApiRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface ApiRoot : GPBRootObject
@end

#pragma mark - Empty

@interface Empty : GPBMessage

@end

#pragma mark - Customer

typedef GPB_ENUM(Customer_FieldNumber) {
  Customer_FieldNumber_UserId = 1,
  Customer_FieldNumber_Plan = 2,
  Customer_FieldNumber_Name = 3,
  Customer_FieldNumber_Email = 4,
  Customer_FieldNumber_Description_p = 5,
  Customer_FieldNumber_Phone = 6,
  Customer_FieldNumber_Address = 8,
  Customer_FieldNumber_Metadata = 9,
  Customer_FieldNumber_Deleted = 10,
  Customer_FieldNumber_CreateDate = 20,
};

/**
 * User is a user of the application
 **/
@interface Customer : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *description_p;

@property(nonatomic, readwrite, copy, null_resettable) NSString *phone;

@property(nonatomic, readwrite, strong, null_resettable) Address *address;
/** Test to see if @c address has been set. */
@property(nonatomic, readwrite) BOOL hasAddress;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *metadata;
/** The number of items in @c metadata without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger metadata_Count;

@property(nonatomic, readwrite) BOOL deleted;

@property(nonatomic, readwrite) int64_t createDate;

@end

#pragma mark - AddCustomerRequest

typedef GPB_ENUM(AddCustomerRequest_FieldNumber) {
  AddCustomerRequest_FieldNumber_Email = 1,
  AddCustomerRequest_FieldNumber_Plan = 2,
  AddCustomerRequest_FieldNumber_Phone = 3,
  AddCustomerRequest_FieldNumber_Name = 4,
  AddCustomerRequest_FieldNumber_Description_p = 7,
  AddCustomerRequest_FieldNumber_Address = 8,
};

@interface AddCustomerRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, copy, null_resettable) NSString *phone;

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *description_p;

@property(nonatomic, readwrite, strong, null_resettable) Address *address;
/** Test to see if @c address has been set. */
@property(nonatomic, readwrite) BOOL hasAddress;

@end

#pragma mark - SubscribeCustomerRequest

typedef GPB_ENUM(SubscribeCustomerRequest_FieldNumber) {
  SubscribeCustomerRequest_FieldNumber_Email = 1,
  SubscribeCustomerRequest_FieldNumber_Plan = 2,
  SubscribeCustomerRequest_FieldNumber_CardNumber = 3,
  SubscribeCustomerRequest_FieldNumber_ExpMonth = 4,
  SubscribeCustomerRequest_FieldNumber_ExpYear = 5,
  SubscribeCustomerRequest_FieldNumber_Cvc = 6,
};

@interface SubscribeCustomerRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, copy, null_resettable) NSString *cardNumber;

@property(nonatomic, readwrite, copy, null_resettable) NSString *expMonth;

@property(nonatomic, readwrite, copy, null_resettable) NSString *expYear;

@property(nonatomic, readwrite, copy, null_resettable) NSString *cvc;

@end

#pragma mark - Address

typedef GPB_ENUM(Address_FieldNumber) {
  Address_FieldNumber_City = 1,
  Address_FieldNumber_Country = 2,
  Address_FieldNumber_Line1 = 3,
  Address_FieldNumber_Line2 = 4,
  Address_FieldNumber_PostalCode = 5,
  Address_FieldNumber_State = 6,
};

@interface Address : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *city;

@property(nonatomic, readwrite, copy, null_resettable) NSString *country;

@property(nonatomic, readwrite, copy, null_resettable) NSString *line1;

@property(nonatomic, readwrite, copy, null_resettable) NSString *line2;

@property(nonatomic, readwrite, copy, null_resettable) NSString *postalCode;

@property(nonatomic, readwrite, copy, null_resettable) NSString *state;

@end

#pragma mark - SubscribeCustomerResponse

typedef GPB_ENUM(SubscribeCustomerResponse_FieldNumber) {
  SubscribeCustomerResponse_FieldNumber_SubscriptionId = 1,
};

@interface SubscribeCustomerResponse : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *subscriptionId;

@end

#pragma mark - CreatePlanResponse

typedef GPB_ENUM(CreatePlanResponse_FieldNumber) {
  CreatePlanResponse_FieldNumber_PlanId = 1,
};

@interface CreatePlanResponse : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *planId;

@end

#pragma mark - CancelSubscriptionRequest

typedef GPB_ENUM(CancelSubscriptionRequest_FieldNumber) {
  CancelSubscriptionRequest_FieldNumber_Email = 1,
};

@interface CancelSubscriptionRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@end

#pragma mark - CreatePlanRequest

typedef GPB_ENUM(CreatePlanRequest_FieldNumber) {
  CreatePlanRequest_FieldNumber_PlanId = 1,
  CreatePlanRequest_FieldNumber_Amount = 2,
  CreatePlanRequest_FieldNumber_ServiceId = 3,
  CreatePlanRequest_FieldNumber_ServiceName = 4,
  CreatePlanRequest_FieldNumber_FriendlyName = 5,
};

@interface CreatePlanRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *planId;

@property(nonatomic, readwrite) int64_t amount;

@property(nonatomic, readwrite, copy, null_resettable) NSString *serviceId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *serviceName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *friendlyName;

@end

#pragma mark - SMSRequest

typedef GPB_ENUM(SMSRequest_FieldNumber) {
  SMSRequest_FieldNumber_UserId = 1,
  SMSRequest_FieldNumber_Body = 2,
};

@interface SMSRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *body;

@end

#pragma mark - CallRequest

typedef GPB_ENUM(CallRequest_FieldNumber) {
  CallRequest_FieldNumber_UserId = 1,
  CallRequest_FieldNumber_CallbackURL = 2,
};

@interface CallRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *callbackURL;

@end

#pragma mark - MMSRequest

typedef GPB_ENUM(MMSRequest_FieldNumber) {
  MMSRequest_FieldNumber_UserId = 1,
  MMSRequest_FieldNumber_Body = 2,
  MMSRequest_FieldNumber_MediaURL = 3,
};

@interface MMSRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *body;

@property(nonatomic, readwrite, copy, null_resettable) NSString *mediaURL;

@end

#pragma mark - EmailRequest

typedef GPB_ENUM(EmailRequest_FieldNumber) {
  EmailRequest_FieldNumber_UserId = 1,
  EmailRequest_FieldNumber_Subject = 2,
  EmailRequest_FieldNumber_PlainText = 3,
  EmailRequest_FieldNumber_HtmlAlt = 4,
};

@interface EmailRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *subject;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plainText;

@property(nonatomic, readwrite, copy, null_resettable) NSString *htmlAlt;

@end

#pragma mark - ChannelReminder

typedef GPB_ENUM(ChannelReminder_FieldNumber) {
  ChannelReminder_FieldNumber_ChannelId = 1,
  ChannelReminder_FieldNumber_Text = 2,
  ChannelReminder_FieldNumber_Time = 3,
};

@interface ChannelReminder : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *channelId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, copy, null_resettable) NSString *time;

@end

#pragma mark - UserReminder

typedef GPB_ENUM(UserReminder_FieldNumber) {
  UserReminder_FieldNumber_UserId = 1,
  UserReminder_FieldNumber_Text = 2,
  UserReminder_FieldNumber_Time = 3,
  UserReminder_FieldNumber_Item = 4,
};

@interface UserReminder : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, copy, null_resettable) NSString *time;

@property(nonatomic, readwrite, strong, null_resettable) ItemRef *item;
/** Test to see if @c item has been set. */
@property(nonatomic, readwrite) BOOL hasItem;

@end

#pragma mark - ItemRef

typedef GPB_ENUM(ItemRef_FieldNumber) {
  ItemRef_FieldNumber_Channel = 1,
  ItemRef_FieldNumber_File = 2,
  ItemRef_FieldNumber_Comment = 3,
};

@interface ItemRef : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *channel;

@property(nonatomic, readwrite, copy, null_resettable) NSString *file;

@property(nonatomic, readwrite, copy, null_resettable) NSString *comment;

@end

#pragma mark - Star

typedef GPB_ENUM(Star_FieldNumber) {
  Star_FieldNumber_Text = 1,
  Star_FieldNumber_Item = 4,
};

@interface Star : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, strong, null_resettable) ItemRef *item;
/** Test to see if @c item has been set. */
@property(nonatomic, readwrite) BOOL hasItem;

@end

#pragma mark - Pin

typedef GPB_ENUM(Pin_FieldNumber) {
  Pin_FieldNumber_Text = 1,
  Pin_FieldNumber_Item = 4,
};

@interface Pin : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, strong, null_resettable) ItemRef *item;
/** Test to see if @c item has been set. */
@property(nonatomic, readwrite) BOOL hasItem;

@end

#pragma mark - Access

typedef GPB_ENUM(Access_FieldNumber) {
  Access_FieldNumber_TwilioAccount = 1,
  Access_FieldNumber_TwilioKey = 2,
  Access_FieldNumber_SendgridAccount = 3,
  Access_FieldNumber_SendgridKey = 4,
  Access_FieldNumber_StripeAccount = 5,
  Access_FieldNumber_StripeKey = 6,
  Access_FieldNumber_SlackAccount = 7,
  Access_FieldNumber_SlackKey = 8,
  Access_FieldNumber_GcpProject = 9,
  Access_FieldNumber_GcpKey = 10,
};

@interface Access : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *twilioAccount;

@property(nonatomic, readwrite, copy, null_resettable) NSString *twilioKey;

@property(nonatomic, readwrite, copy, null_resettable) NSString *sendgridAccount;

@property(nonatomic, readwrite, copy, null_resettable) NSString *sendgridKey;

@property(nonatomic, readwrite, copy, null_resettable) NSString *stripeAccount;

@property(nonatomic, readwrite, copy, null_resettable) NSString *stripeKey;

@property(nonatomic, readwrite, copy, null_resettable) NSString *slackAccount;

@property(nonatomic, readwrite, copy, null_resettable) NSString *slackKey;

@property(nonatomic, readwrite, copy, null_resettable) NSString *gcpProject;

@property(nonatomic, readwrite, copy, null_resettable) NSString *gcpKey;

@end

#pragma mark - Token

typedef GPB_ENUM(Token_FieldNumber) {
  Token_FieldNumber_Access = 1,
  Token_FieldNumber_Audience = 2,
  Token_FieldNumber_Subject = 3,
  Token_FieldNumber_ExpiresAt = 4,
  Token_FieldNumber_Id_p = 5,
  Token_FieldNumber_IssuedAt = 6,
  Token_FieldNumber_NotBefore = 7,
  Token_FieldNumber_GrantsArray = 8,
};

@interface Token : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Access *access;
/** Test to see if @c access has been set. */
@property(nonatomic, readwrite) BOOL hasAccess;

@property(nonatomic, readwrite, copy, null_resettable) NSString *audience;

@property(nonatomic, readwrite, copy, null_resettable) NSString *subject;

@property(nonatomic, readwrite) int64_t expiresAt;

@property(nonatomic, readwrite, copy, null_resettable) NSString *id_p;

@property(nonatomic, readwrite) int64_t issuedAt;

@property(nonatomic, readwrite) int64_t notBefore;

// |grantsArray| contains |Grant|
@property(nonatomic, readwrite, strong, null_resettable) GPBEnumArray *grantsArray;
/** The number of items in @c grantsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger grantsArray_Count;

@end

#pragma mark - LogConfig

typedef GPB_ENUM(LogConfig_FieldNumber) {
  LogConfig_FieldNumber_Username = 1,
  LogConfig_FieldNumber_Channel = 2,
};

@interface LogConfig : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *username;

@property(nonatomic, readwrite, copy, null_resettable) NSString *channel;

@end

#pragma mark - EmailAddress

typedef GPB_ENUM(EmailAddress_FieldNumber) {
  EmailAddress_FieldNumber_Name = 1,
  EmailAddress_FieldNumber_Address = 2,
};

@interface EmailAddress : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *address;

@end

#pragma mark - Email

typedef GPB_ENUM(Email_FieldNumber) {
  Email_FieldNumber_From = 1,
  Email_FieldNumber_Recipient = 2,
};

@interface Email : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) EmailAddress *from;
/** Test to see if @c from has been set. */
@property(nonatomic, readwrite) BOOL hasFrom;

@property(nonatomic, readwrite, strong, null_resettable) RecipientEmail *recipient;
/** Test to see if @c recipient has been set. */
@property(nonatomic, readwrite) BOOL hasRecipient;

@end

#pragma mark - RecipientEmail

typedef GPB_ENUM(RecipientEmail_FieldNumber) {
  RecipientEmail_FieldNumber_To = 2,
  RecipientEmail_FieldNumber_Subject = 3,
  RecipientEmail_FieldNumber_PlainText = 4,
  RecipientEmail_FieldNumber_Html = 5,
};

@interface RecipientEmail : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) EmailAddress *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, copy, null_resettable) NSString *subject;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plainText;

@property(nonatomic, readwrite, copy, null_resettable) NSString *html;

@end

#pragma mark - SMS

typedef GPB_ENUM(SMS_FieldNumber) {
  SMS_FieldNumber_To = 1,
  SMS_FieldNumber_From = 2,
  SMS_FieldNumber_Body = 3,
  SMS_FieldNumber_MediaURL = 4,
  SMS_FieldNumber_Callback = 5,
  SMS_FieldNumber_App = 6,
};

@interface SMS : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *to;

@property(nonatomic, readwrite, copy, null_resettable) NSString *from;

@property(nonatomic, readwrite, copy, null_resettable) NSString *body;

@property(nonatomic, readwrite, copy, null_resettable) NSString *mediaURL;

@property(nonatomic, readwrite, copy, null_resettable) NSString *callback;

@property(nonatomic, readwrite, copy, null_resettable) NSString *app;

@end

#pragma mark - Call

typedef GPB_ENUM(Call_FieldNumber) {
  Call_FieldNumber_To = 1,
  Call_FieldNumber_From = 2,
  Call_FieldNumber_Callback = 5,
};

@interface Call : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *to;

@property(nonatomic, readwrite, copy, null_resettable) NSString *from;

@property(nonatomic, readwrite, copy, null_resettable) NSString *callback;

@end

#pragma mark - Fax

typedef GPB_ENUM(Fax_FieldNumber) {
  Fax_FieldNumber_To = 1,
  Fax_FieldNumber_From = 2,
  Fax_FieldNumber_MediaURL = 3,
  Fax_FieldNumber_Quality = 4,
  Fax_FieldNumber_Callback = 5,
  Fax_FieldNumber_StoreMedia = 6,
};

@interface Fax : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *to;

@property(nonatomic, readwrite, copy, null_resettable) NSString *from;

@property(nonatomic, readwrite, copy, null_resettable) NSString *mediaURL;

@property(nonatomic, readwrite, copy, null_resettable) NSString *quality;

@property(nonatomic, readwrite, copy, null_resettable) NSString *callback;

@property(nonatomic, readwrite) BOOL storeMedia;

@end

#pragma mark - LogHook

typedef GPB_ENUM(LogHook_FieldNumber) {
  LogHook_FieldNumber_Author = 1,
  LogHook_FieldNumber_Icon = 2,
  LogHook_FieldNumber_Title = 3,
};

@interface LogHook : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *author;

@property(nonatomic, readwrite, copy, null_resettable) NSString *icon;

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@end

#pragma mark - Attachment

typedef GPB_ENUM(Attachment_FieldNumber) {
  Attachment_FieldNumber_Color = 1,
  Attachment_FieldNumber_Fallback = 2,
  Attachment_FieldNumber_CallbackId = 3,
  Attachment_FieldNumber_Id_p = 4,
  Attachment_FieldNumber_AuthorId = 5,
  Attachment_FieldNumber_AuthorName = 6,
  Attachment_FieldNumber_AuthorLink = 7,
  Attachment_FieldNumber_AuthorIcon = 8,
  Attachment_FieldNumber_Title = 9,
  Attachment_FieldNumber_TitlePrefix = 10,
  Attachment_FieldNumber_Pretext = 11,
  Attachment_FieldNumber_Text = 12,
  Attachment_FieldNumber_ImageURL = 13,
  Attachment_FieldNumber_ThumbURL = 14,
  Attachment_FieldNumber_FieldsArray = 15,
};

@interface Attachment : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *color;

@property(nonatomic, readwrite, copy, null_resettable) NSString *fallback;

@property(nonatomic, readwrite, copy, null_resettable) NSString *callbackId;

@property(nonatomic, readwrite) int64_t id_p;

@property(nonatomic, readwrite, copy, null_resettable) NSString *authorId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *authorName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *authorLink;

@property(nonatomic, readwrite, copy, null_resettable) NSString *authorIcon;

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@property(nonatomic, readwrite, copy, null_resettable) NSString *titlePrefix;

@property(nonatomic, readwrite, copy, null_resettable) NSString *pretext;

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, copy, null_resettable) NSString *imageURL;

@property(nonatomic, readwrite, copy, null_resettable) NSString *thumbURL;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<AttachmentField*> *fieldsArray;
/** The number of items in @c fieldsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger fieldsArray_Count;

@end

#pragma mark - AttachmentAction

typedef GPB_ENUM(AttachmentAction_FieldNumber) {
  AttachmentAction_FieldNumber_Name = 1,
  AttachmentAction_FieldNumber_Text = 2,
  AttachmentAction_FieldNumber_Style = 3,
  AttachmentAction_FieldNumber_Type = 4,
  AttachmentAction_FieldNumber_Value = 5,
  AttachmentAction_FieldNumber_DataSource = 6,
  AttachmentAction_FieldNumber_MinQueryLength = 7,
  AttachmentAction_FieldNumber_OptionsArray = 8,
  AttachmentAction_FieldNumber_SelectedOptionsArray = 9,
  AttachmentAction_FieldNumber_OptionGroupsArray = 10,
  AttachmentAction_FieldNumber_Confirm = 11,
  AttachmentAction_FieldNumber_URL = 12,
};

@interface AttachmentAction : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, copy, null_resettable) NSString *style;

@property(nonatomic, readwrite, copy, null_resettable) NSString *type;

@property(nonatomic, readwrite, copy, null_resettable) NSString *value;

@property(nonatomic, readwrite, copy, null_resettable) NSString *dataSource;

@property(nonatomic, readwrite) int64_t minQueryLength;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<AttachmentActionOption*> *optionsArray;
/** The number of items in @c optionsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger optionsArray_Count;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<AttachmentActionOption*> *selectedOptionsArray;
/** The number of items in @c selectedOptionsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger selectedOptionsArray_Count;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<AttachmentActionOptionGroup*> *optionGroupsArray;
/** The number of items in @c optionGroupsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger optionGroupsArray_Count;

@property(nonatomic, readwrite, strong, null_resettable) AttachmentConfirmationField *confirm;
/** Test to see if @c confirm has been set. */
@property(nonatomic, readwrite) BOOL hasConfirm;

@property(nonatomic, readwrite, copy, null_resettable) NSString *URL;

@end

#pragma mark - AttachmentConfirmationField

typedef GPB_ENUM(AttachmentConfirmationField_FieldNumber) {
  AttachmentConfirmationField_FieldNumber_Title = 1,
  AttachmentConfirmationField_FieldNumber_Text = 2,
  AttachmentConfirmationField_FieldNumber_OkText = 3,
  AttachmentConfirmationField_FieldNumber_DismissText = 4,
};

@interface AttachmentConfirmationField : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, copy, null_resettable) NSString *okText;

@property(nonatomic, readwrite, copy, null_resettable) NSString *dismissText;

@end

#pragma mark - AttachmentActionOptionGroup

typedef GPB_ENUM(AttachmentActionOptionGroup_FieldNumber) {
  AttachmentActionOptionGroup_FieldNumber_Text = 1,
  AttachmentActionOptionGroup_FieldNumber_OptionsArray = 2,
};

@interface AttachmentActionOptionGroup : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<AttachmentActionOption*> *optionsArray;
/** The number of items in @c optionsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger optionsArray_Count;

@end

#pragma mark - AttachmentActionOption

typedef GPB_ENUM(AttachmentActionOption_FieldNumber) {
  AttachmentActionOption_FieldNumber_Title = 1,
  AttachmentActionOption_FieldNumber_Value = 2,
  AttachmentActionOption_FieldNumber_Description_p = 3,
};

@interface AttachmentActionOption : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@property(nonatomic, readwrite, copy, null_resettable) NSString *value;

@property(nonatomic, readwrite, copy, null_resettable) NSString *description_p;

@end

#pragma mark - AttachmentField

typedef GPB_ENUM(AttachmentField_FieldNumber) {
  AttachmentField_FieldNumber_Title = 1,
  AttachmentField_FieldNumber_Value = 2,
  AttachmentField_FieldNumber_Short_p = 3,
};

@interface AttachmentField : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *title;

@property(nonatomic, readwrite, copy, null_resettable) NSString *value;

@property(nonatomic, readwrite) BOOL short_p;

@end

#pragma mark - JSONMap

typedef GPB_ENUM(JSONMap_FieldNumber) {
  JSONMap_FieldNumber_JsonMap = 1,
};

@interface JSONMap : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, JSON*> *jsonMap;
/** The number of items in @c jsonMap without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger jsonMap_Count;

@end

#pragma mark - JSON

typedef GPB_ENUM(JSON_FieldNumber) {
  JSON_FieldNumber_Data_p = 1,
  JSON_FieldNumber_Size = 2,
};

@interface JSON : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *data_p;

@property(nonatomic, readwrite) int64_t size;

@end

#pragma mark - SignedKey

typedef GPB_ENUM(SignedKey_FieldNumber) {
  SignedKey_FieldNumber_SignedKey = 1,
};

@interface SignedKey : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *signedKey;

@end

#pragma mark - PubSubMessage

typedef GPB_ENUM(PubSubMessage_FieldNumber) {
  PubSubMessage_FieldNumber_Message = 1,
};

@interface PubSubMessage : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) PubsubMessage *message;
/** Test to see if @c message has been set. */
@property(nonatomic, readwrite) BOOL hasMessage;

@end

#pragma mark - PubSubTopic

typedef GPB_ENUM(PubSubTopic_FieldNumber) {
  PubSubTopic_FieldNumber_Topic = 1,
};

@interface PubSubTopic : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Topic *topic;
/** Test to see if @c topic has been set. */
@property(nonatomic, readwrite) BOOL hasTopic;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
