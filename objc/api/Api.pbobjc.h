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

@class AppMetadata;
@class Bytes;
@class Card;
@class Email;
@class EmailBlast;
@class Identity;
@class JSONWebKeys;
@class ManagementToken;
@class Message;
@class Template;
@class User;
@class UserMetadata;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum Scope

typedef GPB_ENUM(Scope) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  Scope_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  Scope_Openid = 0,
  Scope_Profile = 1,
  Scope_Email = 2,
  Scope_ReadUsers = 3,
  Scope_ReadUserIdpTokens = 4,
  Scope_CreateUsers = 5,
  Scope_ReadStats = 6,
  Scope_ReadEmailTemplates = 7,
  Scope_UpdateEmailTemplates = 8,
  Scope_CreateEmailTemplates = 9,
  Scope_ReadRules = 10,
  Scope_UpdateRules = 11,
  Scope_CreateRules = 12,
  Scope_DeleteRules = 13,
  Scope_ReadRoles = 14,
  Scope_UpdateRoles = 15,
  Scope_CreateRoles = 16,
  Scope_DeleteRoles = 17,
  Scope_ReadLogs = 18,
};

GPBEnumDescriptor *Scope_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL Scope_IsValidValue(int32_t value);

#pragma mark - Enum HTTPMethod

typedef GPB_ENUM(HTTPMethod) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  HTTPMethod_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  HTTPMethod_Get = 0,
  HTTPMethod_Post = 1,
};

GPBEnumDescriptor *HTTPMethod_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL HTTPMethod_IsValidValue(int32_t value);

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

#pragma mark - SubscribeRequest

typedef GPB_ENUM(SubscribeRequest_FieldNumber) {
  SubscribeRequest_FieldNumber_Email = 1,
  SubscribeRequest_FieldNumber_Plan = 2,
  SubscribeRequest_FieldNumber_Card = 3,
};

@interface SubscribeRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, strong, null_resettable) Card *card;
/** Test to see if @c card has been set. */
@property(nonatomic, readwrite) BOOL hasCard;

@end

#pragma mark - UnSubscribeRequest

typedef GPB_ENUM(UnSubscribeRequest_FieldNumber) {
  UnSubscribeRequest_FieldNumber_Email = 1,
  UnSubscribeRequest_FieldNumber_Plan = 2,
};

@interface UnSubscribeRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@end

#pragma mark - Card

typedef GPB_ENUM(Card_FieldNumber) {
  Card_FieldNumber_Number = 1,
  Card_FieldNumber_ExpMonth = 2,
  Card_FieldNumber_ExpYear = 3,
  Card_FieldNumber_Cvc = 4,
};

@interface Card : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *number;

@property(nonatomic, readwrite, copy, null_resettable) NSString *expMonth;

@property(nonatomic, readwrite, copy, null_resettable) NSString *expYear;

@property(nonatomic, readwrite, copy, null_resettable) NSString *cvc;

@end

#pragma mark - Empty

@interface Empty : GPBMessage

@end

#pragma mark - ManagementToken

typedef GPB_ENUM(ManagementToken_FieldNumber) {
  ManagementToken_FieldNumber_Token = 1,
};

@interface ManagementToken : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *token;

@end

#pragma mark - UserRequest

typedef GPB_ENUM(UserRequest_FieldNumber) {
  UserRequest_FieldNumber_String = 1,
  UserRequest_FieldNumber_User = 2,
};

@interface UserRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) ManagementToken *string;
/** Test to see if @c string has been set. */
@property(nonatomic, readwrite) BOOL hasString;

@property(nonatomic, readwrite, strong, null_resettable) User *user;
/** Test to see if @c user has been set. */
@property(nonatomic, readwrite) BOOL hasUser;

@end

#pragma mark - UserByEmailRequest

typedef GPB_ENUM(UserByEmailRequest_FieldNumber) {
  UserByEmailRequest_FieldNumber_Token = 1,
  UserByEmailRequest_FieldNumber_Email = 2,
};

@interface UserByEmailRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) ManagementToken *token;
/** Test to see if @c token has been set. */
@property(nonatomic, readwrite) BOOL hasToken;

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@end

#pragma mark - Identifier

typedef GPB_ENUM(Identifier_FieldNumber) {
  Identifier_FieldNumber_Id_p = 1,
};

@interface Identifier : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *id_p;

@end

#pragma mark - SMS

typedef GPB_ENUM(SMS_FieldNumber) {
  SMS_FieldNumber_Service = 1,
  SMS_FieldNumber_To = 2,
  SMS_FieldNumber_Message = 3,
  SMS_FieldNumber_MediaURL = 4,
  SMS_FieldNumber_Callback = 5,
  SMS_FieldNumber_App = 6,
};

@interface SMS : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *service;

@property(nonatomic, readwrite, copy, null_resettable) NSString *to;

@property(nonatomic, readwrite, strong, null_resettable) Message *message;
/** Test to see if @c message has been set. */
@property(nonatomic, readwrite) BOOL hasMessage;

@property(nonatomic, readwrite, copy, null_resettable) NSString *mediaURL;

@property(nonatomic, readwrite, copy, null_resettable) NSString *callback;

@property(nonatomic, readwrite, copy, null_resettable) NSString *app;

@end

#pragma mark - SMSBlast

typedef GPB_ENUM(SMSBlast_FieldNumber) {
  SMSBlast_FieldNumber_Service = 1,
  SMSBlast_FieldNumber_ToArray = 2,
  SMSBlast_FieldNumber_Message = 3,
  SMSBlast_FieldNumber_MediaURL = 4,
  SMSBlast_FieldNumber_Callback = 5,
  SMSBlast_FieldNumber_App = 6,
};

@interface SMSBlast : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *service;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *toArray;
/** The number of items in @c toArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger toArray_Count;

@property(nonatomic, readwrite, strong, null_resettable) Message *message;
/** Test to see if @c message has been set. */
@property(nonatomic, readwrite) BOOL hasMessage;

@property(nonatomic, readwrite, copy, null_resettable) NSString *mediaURL;

@property(nonatomic, readwrite, copy, null_resettable) NSString *callback;

@property(nonatomic, readwrite, copy, null_resettable) NSString *app;

@end

#pragma mark - EmailRequest

typedef GPB_ENUM(EmailRequest_FieldNumber) {
  EmailRequest_FieldNumber_FromName = 1,
  EmailRequest_FieldNumber_FromEmail = 2,
  EmailRequest_FieldNumber_Email = 3,
};

@interface EmailRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *fromName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *fromEmail;

@property(nonatomic, readwrite, strong, null_resettable) Email *email;
/** Test to see if @c email has been set. */
@property(nonatomic, readwrite) BOOL hasEmail;

@end

#pragma mark - EmailBlastRequest

typedef GPB_ENUM(EmailBlastRequest_FieldNumber) {
  EmailBlastRequest_FieldNumber_FromName = 1,
  EmailBlastRequest_FieldNumber_FromEmail = 2,
  EmailBlastRequest_FieldNumber_Blast = 3,
};

@interface EmailBlastRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *fromName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *fromEmail;

@property(nonatomic, readwrite, strong, null_resettable) EmailBlast *blast;
/** Test to see if @c blast has been set. */
@property(nonatomic, readwrite) BOOL hasBlast;

@end

#pragma mark - EmailBlast

typedef GPB_ENUM(EmailBlast_FieldNumber) {
  EmailBlast_FieldNumber_NameAddress = 1,
  EmailBlast_FieldNumber_Subject = 2,
  EmailBlast_FieldNumber_Plain = 3,
  EmailBlast_FieldNumber_Html = 4,
};

@interface EmailBlast : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *nameAddress;
/** The number of items in @c nameAddress without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger nameAddress_Count;

@property(nonatomic, readwrite, copy, null_resettable) NSString *subject;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plain;

@property(nonatomic, readwrite, copy, null_resettable) NSString *html;

@end

#pragma mark - Email

typedef GPB_ENUM(Email_FieldNumber) {
  Email_FieldNumber_Name = 1,
  Email_FieldNumber_Address = 2,
  Email_FieldNumber_Subject = 3,
  Email_FieldNumber_Plain = 4,
  Email_FieldNumber_Html = 5,
};

@interface Email : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *address;

@property(nonatomic, readwrite, copy, null_resettable) NSString *subject;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plain;

@property(nonatomic, readwrite, copy, null_resettable) NSString *html;

@end

#pragma mark - Call

typedef GPB_ENUM(Call_FieldNumber) {
  Call_FieldNumber_From = 1,
  Call_FieldNumber_To = 2,
  Call_FieldNumber_App = 3,
};

@interface Call : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *from;

@property(nonatomic, readwrite, copy, null_resettable) NSString *to;

@property(nonatomic, readwrite, copy, null_resettable) NSString *app;

@end

#pragma mark - CallBlast

typedef GPB_ENUM(CallBlast_FieldNumber) {
  CallBlast_FieldNumber_From = 1,
  CallBlast_FieldNumber_ToArray = 2,
  CallBlast_FieldNumber_App = 3,
};

@interface CallBlast : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *from;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *toArray;
/** The number of items in @c toArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger toArray_Count;

@property(nonatomic, readwrite, copy, null_resettable) NSString *app;

@end

#pragma mark - Message

typedef GPB_ENUM(Message_FieldNumber) {
  Message_FieldNumber_Value = 1,
};

@interface Message : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *value;

@end

#pragma mark - User

typedef GPB_ENUM(User_FieldNumber) {
  User_FieldNumber_UserId = 1,
  User_FieldNumber_Name = 2,
  User_FieldNumber_GivenName = 3,
  User_FieldNumber_FamilyName = 4,
  User_FieldNumber_Gender = 5,
  User_FieldNumber_Birthdate = 6,
  User_FieldNumber_Email = 7,
  User_FieldNumber_PhoneNumber = 8,
  User_FieldNumber_Picture = 9,
  User_FieldNumber_UserMetadata = 10,
  User_FieldNumber_AppMetadata = 11,
  User_FieldNumber_LastIp = 12,
  User_FieldNumber_Blocked = 13,
  User_FieldNumber_Nickname = 14,
  User_FieldNumber_MultifactorArray = 15,
  User_FieldNumber_CreatedAt = 17,
  User_FieldNumber_UpdatedAt = 18,
  User_FieldNumber_PhoneVerified = 19,
  User_FieldNumber_IdentitiesArray = 20,
};

@interface User : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *givenName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *familyName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *gender;

@property(nonatomic, readwrite, copy, null_resettable) NSString *birthdate;

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *phoneNumber;

@property(nonatomic, readwrite, copy, null_resettable) NSString *picture;

@property(nonatomic, readwrite, strong, null_resettable) UserMetadata *userMetadata;
/** Test to see if @c userMetadata has been set. */
@property(nonatomic, readwrite) BOOL hasUserMetadata;

@property(nonatomic, readwrite, strong, null_resettable) AppMetadata *appMetadata;
/** Test to see if @c appMetadata has been set. */
@property(nonatomic, readwrite) BOOL hasAppMetadata;

@property(nonatomic, readwrite, copy, null_resettable) NSString *lastIp;

@property(nonatomic, readwrite) BOOL blocked;

@property(nonatomic, readwrite, copy, null_resettable) NSString *nickname;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *multifactorArray;
/** The number of items in @c multifactorArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger multifactorArray_Count;

@property(nonatomic, readwrite, copy, null_resettable) NSString *createdAt;

@property(nonatomic, readwrite, copy, null_resettable) NSString *updatedAt;

@property(nonatomic, readwrite) BOOL phoneVerified;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<Identity*> *identitiesArray;
/** The number of items in @c identitiesArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger identitiesArray_Count;

@end

#pragma mark - Identity

typedef GPB_ENUM(Identity_FieldNumber) {
  Identity_FieldNumber_Connection = 1,
  Identity_FieldNumber_UserId = 2,
  Identity_FieldNumber_Provider = 3,
  Identity_FieldNumber_IsSocial = 4,
};

@interface Identity : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *connection;

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *provider;

@property(nonatomic, readwrite, copy, null_resettable) NSString *isSocial;

@end

#pragma mark - UserMetadata

typedef GPB_ENUM(UserMetadata_FieldNumber) {
  UserMetadata_FieldNumber_Metadata = 1,
};

@interface UserMetadata : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *metadata;
/** The number of items in @c metadata without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger metadata_Count;

@end

#pragma mark - AppMetadata

typedef GPB_ENUM(AppMetadata_FieldNumber) {
  AppMetadata_FieldNumber_Metadata = 1,
};

@interface AppMetadata : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *metadata;
/** The number of items in @c metadata without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger metadata_Count;

@end

#pragma mark - Auth

typedef GPB_ENUM(Auth_FieldNumber) {
  Auth_FieldNumber_Domain = 1,
  Auth_FieldNumber_ClientId = 2,
  Auth_FieldNumber_ClientSecret = 3,
  Auth_FieldNumber_Redirect = 4,
  Auth_FieldNumber_Audience = 5,
  Auth_FieldNumber_ScopesArray = 6,
};

@interface Auth : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *domain;

@property(nonatomic, readwrite, copy, null_resettable) NSString *clientId;

@property(nonatomic, readwrite, copy, null_resettable) NSString *clientSecret;

@property(nonatomic, readwrite, copy, null_resettable) NSString *redirect;

@property(nonatomic, readwrite, copy, null_resettable) NSString *audience;

// |scopesArray| contains |Scope|
@property(nonatomic, readwrite, strong, null_resettable) GPBEnumArray *scopesArray;
/** The number of items in @c scopesArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger scopesArray_Count;

@end

#pragma mark - Bytes

typedef GPB_ENUM(Bytes_FieldNumber) {
  Bytes_FieldNumber_Bits = 1,
};

@interface Bytes : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSData *bits;

@end

#pragma mark - Template

typedef GPB_ENUM(Template_FieldNumber) {
  Template_FieldNumber_Name = 1,
  Template_FieldNumber_Text = 2,
};

@interface Template : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@end

#pragma mark - JSONWebKeys

typedef GPB_ENUM(JSONWebKeys_FieldNumber) {
  JSONWebKeys_FieldNumber_Kty = 1,
  JSONWebKeys_FieldNumber_Kid = 2,
  JSONWebKeys_FieldNumber_Use = 3,
  JSONWebKeys_FieldNumber_N = 4,
  JSONWebKeys_FieldNumber_E = 5,
  JSONWebKeys_FieldNumber_X5CArray = 6,
};

@interface JSONWebKeys : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *kty;

@property(nonatomic, readwrite, copy, null_resettable) NSString *kid;

@property(nonatomic, readwrite, copy, null_resettable) NSString *use;

@property(nonatomic, readwrite, copy, null_resettable) NSString *n;

@property(nonatomic, readwrite, copy, null_resettable) NSString *e;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *x5CArray;
/** The number of items in @c x5CArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger x5CArray_Count;

@end

#pragma mark - Jwks

typedef GPB_ENUM(Jwks_FieldNumber) {
  Jwks_FieldNumber_KeysArray = 1,
};

@interface Jwks : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<JSONWebKeys*> *keysArray;
/** The number of items in @c keysArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger keysArray_Count;

@end

#pragma mark - HTTPRequest

typedef GPB_ENUM(HTTPRequest_FieldNumber) {
  HTTPRequest_FieldNumber_Method = 1,
  HTTPRequest_FieldNumber_URL = 2,
  HTTPRequest_FieldNumber_Token = 3,
  HTTPRequest_FieldNumber_Account = 4,
  HTTPRequest_FieldNumber_ContentType = 5,
  HTTPRequest_FieldNumber_Headers = 6,
  HTTPRequest_FieldNumber_Form = 7,
  HTTPRequest_FieldNumber_Cookies = 8,
  HTTPRequest_FieldNumber_Body = 9,
};

@interface HTTPRequest : GPBMessage

@property(nonatomic, readwrite) HTTPMethod method;

@property(nonatomic, readwrite, copy, null_resettable) NSString *URL;

@property(nonatomic, readwrite, copy, null_resettable) NSString *token;

@property(nonatomic, readwrite, copy, null_resettable) NSString *account;

@property(nonatomic, readwrite, copy, null_resettable) NSString *contentType;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *headers;
/** The number of items in @c headers without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger headers_Count;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *form;
/** The number of items in @c form without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger form_Count;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *cookies;
/** The number of items in @c cookies without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger cookies_Count;

@property(nonatomic, readwrite, strong, null_resettable) Bytes *body;
/** Test to see if @c body has been set. */
@property(nonatomic, readwrite) BOOL hasBody;

@end

/**
 * Fetches the raw value of a @c HTTPRequest's @c method property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t HTTPRequest_Method_RawValue(HTTPRequest *message);
/**
 * Sets the raw value of an @c HTTPRequest's @c method property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetHTTPRequest_Method_RawValue(HTTPRequest *message, int32_t value);

#pragma mark - RenderRequest

typedef GPB_ENUM(RenderRequest_FieldNumber) {
  RenderRequest_FieldNumber_Template_p = 1,
  RenderRequest_FieldNumber_Data_p = 2,
};

@interface RenderRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Template *template_p;
/** Test to see if @c template_p has been set. */
@property(nonatomic, readwrite) BOOL hasTemplate_p;

@property(nonatomic, readwrite, strong, null_resettable) Bytes *data_p;
/** Test to see if @c data_p has been set. */
@property(nonatomic, readwrite) BOOL hasData_p;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
