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

@class Card;
@class Config;
@class Email;
@class EmailBlast;
@class Identifier;
@class Identity;
@class Int64;
@class JSONWebKeys;
@class NumberCapabilities;
@class PhoneNumber;
@class String;
@class StringArray;
@class StringMap;
@class Token;
@class TokenSet;
@class User;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum Plan

typedef GPB_ENUM(Plan) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  Plan_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  Plan_Free = 0,
  Plan_Basic = 1,
  Plan_Premium = 2,
};

GPBEnumDescriptor *Plan_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL Plan_IsValidValue(int32_t value);

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

#pragma mark - FaxRequest

typedef GPB_ENUM(FaxRequest_FieldNumber) {
  FaxRequest_FieldNumber_To = 1,
  FaxRequest_FieldNumber_From = 2,
  FaxRequest_FieldNumber_MediaURL = 3,
  FaxRequest_FieldNumber_Quality = 4,
  FaxRequest_FieldNumber_Callback = 5,
  FaxRequest_FieldNumber_StoreMedia = 6,
};

@interface FaxRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *from;
/** Test to see if @c from has been set. */
@property(nonatomic, readwrite) BOOL hasFrom;

@property(nonatomic, readwrite, strong, null_resettable) String *mediaURL;
/** Test to see if @c mediaURL has been set. */
@property(nonatomic, readwrite) BOOL hasMediaURL;

@property(nonatomic, readwrite, strong, null_resettable) String *quality;
/** Test to see if @c quality has been set. */
@property(nonatomic, readwrite) BOOL hasQuality;

@property(nonatomic, readwrite, strong, null_resettable) String *callback;
/** Test to see if @c callback has been set. */
@property(nonatomic, readwrite) BOOL hasCallback;

@property(nonatomic, readwrite) BOOL storeMedia;

@end

#pragma mark - SubscribeRequest

typedef GPB_ENUM(SubscribeRequest_FieldNumber) {
  SubscribeRequest_FieldNumber_Email = 1,
  SubscribeRequest_FieldNumber_Plan = 2,
  SubscribeRequest_FieldNumber_Card = 3,
};

@interface SubscribeRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *email;
/** Test to see if @c email has been set. */
@property(nonatomic, readwrite) BOOL hasEmail;

@property(nonatomic, readwrite) Plan plan;

@property(nonatomic, readwrite, strong, null_resettable) Card *card;
/** Test to see if @c card has been set. */
@property(nonatomic, readwrite) BOOL hasCard;

@end

/**
 * Fetches the raw value of a @c SubscribeRequest's @c plan property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t SubscribeRequest_Plan_RawValue(SubscribeRequest *message);
/**
 * Sets the raw value of an @c SubscribeRequest's @c plan property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetSubscribeRequest_Plan_RawValue(SubscribeRequest *message, int32_t value);

#pragma mark - UnSubscribeRequest

typedef GPB_ENUM(UnSubscribeRequest_FieldNumber) {
  UnSubscribeRequest_FieldNumber_Email = 1,
  UnSubscribeRequest_FieldNumber_Plan = 2,
};

@interface UnSubscribeRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *email;
/** Test to see if @c email has been set. */
@property(nonatomic, readwrite) BOOL hasEmail;

@property(nonatomic, readwrite) Plan plan;

@end

/**
 * Fetches the raw value of a @c UnSubscribeRequest's @c plan property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t UnSubscribeRequest_Plan_RawValue(UnSubscribeRequest *message);
/**
 * Sets the raw value of an @c UnSubscribeRequest's @c plan property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetUnSubscribeRequest_Plan_RawValue(UnSubscribeRequest *message, int32_t value);

#pragma mark - Card

typedef GPB_ENUM(Card_FieldNumber) {
  Card_FieldNumber_Number = 1,
  Card_FieldNumber_ExpMonth = 2,
  Card_FieldNumber_ExpYear = 3,
  Card_FieldNumber_Cvc = 4,
};

@interface Card : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *number;
/** Test to see if @c number has been set. */
@property(nonatomic, readwrite) BOOL hasNumber;

@property(nonatomic, readwrite, strong, null_resettable) String *expMonth;
/** Test to see if @c expMonth has been set. */
@property(nonatomic, readwrite) BOOL hasExpMonth;

@property(nonatomic, readwrite, strong, null_resettable) String *expYear;
/** Test to see if @c expYear has been set. */
@property(nonatomic, readwrite) BOOL hasExpYear;

@property(nonatomic, readwrite, strong, null_resettable) String *cvc;
/** Test to see if @c cvc has been set. */
@property(nonatomic, readwrite) BOOL hasCvc;

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

@property(nonatomic, readwrite, strong, null_resettable) String *service;
/** Test to see if @c service has been set. */
@property(nonatomic, readwrite) BOOL hasService;

@property(nonatomic, readwrite, strong, null_resettable) String *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *message;
/** Test to see if @c message has been set. */
@property(nonatomic, readwrite) BOOL hasMessage;

@property(nonatomic, readwrite, strong, null_resettable) String *mediaURL;
/** Test to see if @c mediaURL has been set. */
@property(nonatomic, readwrite) BOOL hasMediaURL;

@property(nonatomic, readwrite, strong, null_resettable) String *callback;
/** Test to see if @c callback has been set. */
@property(nonatomic, readwrite) BOOL hasCallback;

@property(nonatomic, readwrite, strong, null_resettable) String *app;
/** Test to see if @c app has been set. */
@property(nonatomic, readwrite) BOOL hasApp;

@end

#pragma mark - SMSBlast

typedef GPB_ENUM(SMSBlast_FieldNumber) {
  SMSBlast_FieldNumber_Service = 1,
  SMSBlast_FieldNumber_To = 2,
  SMSBlast_FieldNumber_Message = 3,
  SMSBlast_FieldNumber_MediaURL = 4,
  SMSBlast_FieldNumber_Callback = 5,
  SMSBlast_FieldNumber_App = 6,
};

@interface SMSBlast : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *service;
/** Test to see if @c service has been set. */
@property(nonatomic, readwrite) BOOL hasService;

@property(nonatomic, readwrite, strong, null_resettable) StringArray *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *message;
/** Test to see if @c message has been set. */
@property(nonatomic, readwrite) BOOL hasMessage;

@property(nonatomic, readwrite, strong, null_resettable) String *mediaURL;
/** Test to see if @c mediaURL has been set. */
@property(nonatomic, readwrite) BOOL hasMediaURL;

@property(nonatomic, readwrite, strong, null_resettable) String *callback;
/** Test to see if @c callback has been set. */
@property(nonatomic, readwrite) BOOL hasCallback;

@property(nonatomic, readwrite, strong, null_resettable) String *app;
/** Test to see if @c app has been set. */
@property(nonatomic, readwrite) BOOL hasApp;

@end

#pragma mark - EmailRequest

typedef GPB_ENUM(EmailRequest_FieldNumber) {
  EmailRequest_FieldNumber_FromName = 1,
  EmailRequest_FieldNumber_FromEmail = 2,
  EmailRequest_FieldNumber_Email = 3,
};

@interface EmailRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *fromName;
/** Test to see if @c fromName has been set. */
@property(nonatomic, readwrite) BOOL hasFromName;

@property(nonatomic, readwrite, strong, null_resettable) String *fromEmail;
/** Test to see if @c fromEmail has been set. */
@property(nonatomic, readwrite) BOOL hasFromEmail;

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

@property(nonatomic, readwrite, strong, null_resettable) String *fromName;
/** Test to see if @c fromName has been set. */
@property(nonatomic, readwrite) BOOL hasFromName;

@property(nonatomic, readwrite, strong, null_resettable) String *fromEmail;
/** Test to see if @c fromEmail has been set. */
@property(nonatomic, readwrite) BOOL hasFromEmail;

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

@property(nonatomic, readwrite, strong, null_resettable) StringMap *nameAddress;
/** Test to see if @c nameAddress has been set. */
@property(nonatomic, readwrite) BOOL hasNameAddress;

@property(nonatomic, readwrite, strong, null_resettable) String *subject;
/** Test to see if @c subject has been set. */
@property(nonatomic, readwrite) BOOL hasSubject;

@property(nonatomic, readwrite, strong, null_resettable) String *plain;
/** Test to see if @c plain has been set. */
@property(nonatomic, readwrite) BOOL hasPlain;

@property(nonatomic, readwrite, strong, null_resettable) String *html;
/** Test to see if @c html has been set. */
@property(nonatomic, readwrite) BOOL hasHtml;

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

@property(nonatomic, readwrite, strong, null_resettable) String *name;
/** Test to see if @c name has been set. */
@property(nonatomic, readwrite) BOOL hasName;

@property(nonatomic, readwrite, strong, null_resettable) String *address;
/** Test to see if @c address has been set. */
@property(nonatomic, readwrite) BOOL hasAddress;

@property(nonatomic, readwrite, strong, null_resettable) String *subject;
/** Test to see if @c subject has been set. */
@property(nonatomic, readwrite) BOOL hasSubject;

@property(nonatomic, readwrite, strong, null_resettable) String *plain;
/** Test to see if @c plain has been set. */
@property(nonatomic, readwrite) BOOL hasPlain;

@property(nonatomic, readwrite, strong, null_resettable) String *html;
/** Test to see if @c html has been set. */
@property(nonatomic, readwrite) BOOL hasHtml;

@end

#pragma mark - Call

typedef GPB_ENUM(Call_FieldNumber) {
  Call_FieldNumber_From = 1,
  Call_FieldNumber_To = 2,
  Call_FieldNumber_App = 3,
};

@interface Call : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *from;
/** Test to see if @c from has been set. */
@property(nonatomic, readwrite) BOOL hasFrom;

@property(nonatomic, readwrite, strong, null_resettable) String *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *app;
/** Test to see if @c app has been set. */
@property(nonatomic, readwrite) BOOL hasApp;

@end

#pragma mark - CallBlast

typedef GPB_ENUM(CallBlast_FieldNumber) {
  CallBlast_FieldNumber_From = 1,
  CallBlast_FieldNumber_To = 2,
  CallBlast_FieldNumber_App = 3,
};

@interface CallBlast : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *from;
/** Test to see if @c from has been set. */
@property(nonatomic, readwrite) BOOL hasFrom;

@property(nonatomic, readwrite, strong, null_resettable) StringArray *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *app;
/** Test to see if @c app has been set. */
@property(nonatomic, readwrite) BOOL hasApp;

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
  User_FieldNumber_Multifactor = 15,
  User_FieldNumber_CreatedAt = 17,
  User_FieldNumber_UpdatedAt = 18,
  User_FieldNumber_PhoneVerified = 19,
  User_FieldNumber_EmailVerified = 20,
  User_FieldNumber_Password = 21,
  User_FieldNumber_IdentitiesArray = 22,
};

@interface User : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *userId;
/** Test to see if @c userId has been set. */
@property(nonatomic, readwrite) BOOL hasUserId;

@property(nonatomic, readwrite, strong, null_resettable) String *name;
/** Test to see if @c name has been set. */
@property(nonatomic, readwrite) BOOL hasName;

@property(nonatomic, readwrite, strong, null_resettable) String *givenName;
/** Test to see if @c givenName has been set. */
@property(nonatomic, readwrite) BOOL hasGivenName;

@property(nonatomic, readwrite, strong, null_resettable) String *familyName;
/** Test to see if @c familyName has been set. */
@property(nonatomic, readwrite) BOOL hasFamilyName;

@property(nonatomic, readwrite, strong, null_resettable) String *gender;
/** Test to see if @c gender has been set. */
@property(nonatomic, readwrite) BOOL hasGender;

@property(nonatomic, readwrite, strong, null_resettable) String *birthdate;
/** Test to see if @c birthdate has been set. */
@property(nonatomic, readwrite) BOOL hasBirthdate;

@property(nonatomic, readwrite, strong, null_resettable) Identifier *email;
/** Test to see if @c email has been set. */
@property(nonatomic, readwrite) BOOL hasEmail;

@property(nonatomic, readwrite, strong, null_resettable) Identifier *phoneNumber;
/** Test to see if @c phoneNumber has been set. */
@property(nonatomic, readwrite) BOOL hasPhoneNumber;

@property(nonatomic, readwrite, strong, null_resettable) String *picture;
/** Test to see if @c picture has been set. */
@property(nonatomic, readwrite) BOOL hasPicture;

@property(nonatomic, readwrite, strong, null_resettable) StringMap *userMetadata;
/** Test to see if @c userMetadata has been set. */
@property(nonatomic, readwrite) BOOL hasUserMetadata;

@property(nonatomic, readwrite, strong, null_resettable) StringMap *appMetadata;
/** Test to see if @c appMetadata has been set. */
@property(nonatomic, readwrite) BOOL hasAppMetadata;

@property(nonatomic, readwrite, strong, null_resettable) String *lastIp;
/** Test to see if @c lastIp has been set. */
@property(nonatomic, readwrite) BOOL hasLastIp;

@property(nonatomic, readwrite) BOOL blocked;

@property(nonatomic, readwrite, strong, null_resettable) String *nickname;
/** Test to see if @c nickname has been set. */
@property(nonatomic, readwrite) BOOL hasNickname;

@property(nonatomic, readwrite, strong, null_resettable) StringArray *multifactor;
/** Test to see if @c multifactor has been set. */
@property(nonatomic, readwrite) BOOL hasMultifactor;

@property(nonatomic, readwrite, strong, null_resettable) String *createdAt;
/** Test to see if @c createdAt has been set. */
@property(nonatomic, readwrite) BOOL hasCreatedAt;

@property(nonatomic, readwrite, strong, null_resettable) String *updatedAt;
/** Test to see if @c updatedAt has been set. */
@property(nonatomic, readwrite) BOOL hasUpdatedAt;

@property(nonatomic, readwrite) BOOL phoneVerified;

@property(nonatomic, readwrite) BOOL emailVerified;

@property(nonatomic, readwrite, strong, null_resettable) String *password;
/** Test to see if @c password has been set. */
@property(nonatomic, readwrite) BOOL hasPassword;

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

@property(nonatomic, readwrite, strong, null_resettable) String *connection;
/** Test to see if @c connection has been set. */
@property(nonatomic, readwrite) BOOL hasConnection;

@property(nonatomic, readwrite, strong, null_resettable) Identifier *userId;
/** Test to see if @c userId has been set. */
@property(nonatomic, readwrite) BOOL hasUserId;

@property(nonatomic, readwrite, strong, null_resettable) String *provider;
/** Test to see if @c provider has been set. */
@property(nonatomic, readwrite) BOOL hasProvider;

@property(nonatomic, readwrite) BOOL isSocial;

@end

#pragma mark - Auth

typedef GPB_ENUM(Auth_FieldNumber) {
  Auth_FieldNumber_Config = 1,
  Auth_FieldNumber_TokenSet = 3,
};

@interface Auth : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Config *config;
/** Test to see if @c config has been set. */
@property(nonatomic, readwrite) BOOL hasConfig;

@property(nonatomic, readwrite, strong, null_resettable) TokenSet *tokenSet;
/** Test to see if @c tokenSet has been set. */
@property(nonatomic, readwrite) BOOL hasTokenSet;

@end

#pragma mark - JSONWebKeys

typedef GPB_ENUM(JSONWebKeys_FieldNumber) {
  JSONWebKeys_FieldNumber_Kty = 1,
  JSONWebKeys_FieldNumber_Kid = 2,
  JSONWebKeys_FieldNumber_Use = 3,
  JSONWebKeys_FieldNumber_N = 4,
  JSONWebKeys_FieldNumber_E = 5,
  JSONWebKeys_FieldNumber_X5C = 6,
};

@interface JSONWebKeys : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *kty;
/** Test to see if @c kty has been set. */
@property(nonatomic, readwrite) BOOL hasKty;

@property(nonatomic, readwrite, strong, null_resettable) Identifier *kid;
/** Test to see if @c kid has been set. */
@property(nonatomic, readwrite) BOOL hasKid;

@property(nonatomic, readwrite, strong, null_resettable) String *use;
/** Test to see if @c use has been set. */
@property(nonatomic, readwrite) BOOL hasUse;

@property(nonatomic, readwrite, strong, null_resettable) String *n;
/** Test to see if @c n has been set. */
@property(nonatomic, readwrite) BOOL hasN;

@property(nonatomic, readwrite, strong, null_resettable) String *e;
/** Test to see if @c e has been set. */
@property(nonatomic, readwrite) BOOL hasE;

@property(nonatomic, readwrite, strong, null_resettable) StringArray *x5C;
/** Test to see if @c x5C has been set. */
@property(nonatomic, readwrite) BOOL hasX5C;

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

#pragma mark - RenderRequest

typedef GPB_ENUM(RenderRequest_FieldNumber) {
  RenderRequest_FieldNumber_Name = 1,
  RenderRequest_FieldNumber_Text = 2,
  RenderRequest_FieldNumber_Data_p = 3,
};

@interface RenderRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *name;
/** Test to see if @c name has been set. */
@property(nonatomic, readwrite) BOOL hasName;

@property(nonatomic, readwrite, strong, null_resettable) String *text;
/** Test to see if @c text has been set. */
@property(nonatomic, readwrite) BOOL hasText;

@property(nonatomic, readwrite, strong, null_resettable) String *data_p;
/** Test to see if @c data_p has been set. */
@property(nonatomic, readwrite) BOOL hasData_p;

@end

#pragma mark - SearchPhoneNumberRequest

typedef GPB_ENUM(SearchPhoneNumberRequest_FieldNumber) {
  SearchPhoneNumberRequest_FieldNumber_State = 1,
  SearchPhoneNumberRequest_FieldNumber_Capabilities = 2,
  SearchPhoneNumberRequest_FieldNumber_TotalResults = 3,
};

@interface SearchPhoneNumberRequest : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *state;
/** Test to see if @c state has been set. */
@property(nonatomic, readwrite) BOOL hasState;

@property(nonatomic, readwrite, strong, null_resettable) NumberCapabilities *capabilities;
/** Test to see if @c capabilities has been set. */
@property(nonatomic, readwrite) BOOL hasCapabilities;

@property(nonatomic, readwrite, strong, null_resettable) Int64 *totalResults;
/** Test to see if @c totalResults has been set. */
@property(nonatomic, readwrite) BOOL hasTotalResults;

@end

#pragma mark - PhoneNumber

typedef GPB_ENUM(PhoneNumber_FieldNumber) {
  PhoneNumber_FieldNumber_FriendlyName = 1,
  PhoneNumber_FieldNumber_PhoneNumber = 2,
  PhoneNumber_FieldNumber_Region = 3,
  PhoneNumber_FieldNumber_Capabilities = 4,
};

@interface PhoneNumber : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) String *friendlyName;
/** Test to see if @c friendlyName has been set. */
@property(nonatomic, readwrite) BOOL hasFriendlyName;

@property(nonatomic, readwrite, strong, null_resettable) String *phoneNumber;
/** Test to see if @c phoneNumber has been set. */
@property(nonatomic, readwrite) BOOL hasPhoneNumber;

@property(nonatomic, readwrite, strong, null_resettable) String *region;
/** Test to see if @c region has been set. */
@property(nonatomic, readwrite) BOOL hasRegion;

@property(nonatomic, readwrite, strong, null_resettable) NumberCapabilities *capabilities;
/** Test to see if @c capabilities has been set. */
@property(nonatomic, readwrite) BOOL hasCapabilities;

@end

#pragma mark - NumberCapabilities

typedef GPB_ENUM(NumberCapabilities_FieldNumber) {
  NumberCapabilities_FieldNumber_Voice = 1,
  NumberCapabilities_FieldNumber_Sms = 2,
  NumberCapabilities_FieldNumber_Mms = 3,
};

@interface NumberCapabilities : GPBMessage

@property(nonatomic, readwrite) BOOL voice;

@property(nonatomic, readwrite) BOOL sms;

@property(nonatomic, readwrite) BOOL mms;

@end

#pragma mark - PhoneNumberResource

typedef GPB_ENUM(PhoneNumberResource_FieldNumber) {
  PhoneNumberResource_FieldNumber_Number = 1,
  PhoneNumberResource_FieldNumber_Id_p = 2,
  PhoneNumberResource_FieldNumber_Uri = 3,
};

@interface PhoneNumberResource : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) PhoneNumber *number;
/** Test to see if @c number has been set. */
@property(nonatomic, readwrite) BOOL hasNumber;

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) String *uri;
/** Test to see if @c uri has been set. */
@property(nonatomic, readwrite) BOOL hasUri;

@end

#pragma mark - TokenQuery

typedef GPB_ENUM(TokenQuery_FieldNumber) {
  TokenQuery_FieldNumber_Token = 1,
  TokenQuery_FieldNumber_Query = 2,
};

@interface TokenQuery : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Token *token;
/** Test to see if @c token has been set. */
@property(nonatomic, readwrite) BOOL hasToken;

@property(nonatomic, readwrite, strong, null_resettable) String *query;
/** Test to see if @c query has been set. */
@property(nonatomic, readwrite) BOOL hasQuery;

@end

#pragma mark - IDBody

typedef GPB_ENUM(IDBody_FieldNumber) {
  IDBody_FieldNumber_Id_p = 1,
  IDBody_FieldNumber_Body = 2,
};

@interface IDBody : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) String *body;
/** Test to see if @c body has been set. */
@property(nonatomic, readwrite) BOOL hasBody;

@end

#pragma mark - IDStrings

typedef GPB_ENUM(IDStrings_FieldNumber) {
  IDStrings_FieldNumber_Id_p = 1,
  IDStrings_FieldNumber_Strings = 2,
};

@interface IDStrings : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) StringArray *strings;
/** Test to see if @c strings has been set. */
@property(nonatomic, readwrite) BOOL hasStrings;

@end

#pragma mark - Role

typedef GPB_ENUM(Role_FieldNumber) {
  Role_FieldNumber_Id_p = 1,
  Role_FieldNumber_Name = 2,
  Role_FieldNumber_Description_p = 3,
};

@interface Role : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) String *name;
/** Test to see if @c name has been set. */
@property(nonatomic, readwrite) BOOL hasName;

@property(nonatomic, readwrite, strong, null_resettable) String *description_p;
/** Test to see if @c description_p has been set. */
@property(nonatomic, readwrite) BOOL hasDescription_p;

@end

#pragma mark - CallResponse

typedef GPB_ENUM(CallResponse_FieldNumber) {
  CallResponse_FieldNumber_Id_p = 1,
  CallResponse_FieldNumber_To = 5,
  CallResponse_FieldNumber_From = 6,
  CallResponse_FieldNumber_Status = 9,
  CallResponse_FieldNumber_AnsweredBy = 10,
  CallResponse_FieldNumber_ForwardedFrom = 11,
  CallResponse_FieldNumber_CallerName = 12,
  CallResponse_FieldNumber_Annotations = 13,
};

@interface CallResponse : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) String *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *from;
/** Test to see if @c from has been set. */
@property(nonatomic, readwrite) BOOL hasFrom;

@property(nonatomic, readwrite, strong, null_resettable) String *status;
/** Test to see if @c status has been set. */
@property(nonatomic, readwrite) BOOL hasStatus;

@property(nonatomic, readwrite, strong, null_resettable) String *answeredBy;
/** Test to see if @c answeredBy has been set. */
@property(nonatomic, readwrite) BOOL hasAnsweredBy;

@property(nonatomic, readwrite, strong, null_resettable) String *forwardedFrom;
/** Test to see if @c forwardedFrom has been set. */
@property(nonatomic, readwrite) BOOL hasForwardedFrom;

@property(nonatomic, readwrite, strong, null_resettable) String *callerName;
/** Test to see if @c callerName has been set. */
@property(nonatomic, readwrite) BOOL hasCallerName;

@property(nonatomic, readwrite, strong, null_resettable) StringMap *annotations;
/** Test to see if @c annotations has been set. */
@property(nonatomic, readwrite) BOOL hasAnnotations;

@end

#pragma mark - SMSResponse

typedef GPB_ENUM(SMSResponse_FieldNumber) {
  SMSResponse_FieldNumber_Id_p = 1,
  SMSResponse_FieldNumber_To = 5,
  SMSResponse_FieldNumber_From = 6,
  SMSResponse_FieldNumber_MediaURL = 7,
  SMSResponse_FieldNumber_Body = 8,
  SMSResponse_FieldNumber_Status = 9,
  SMSResponse_FieldNumber_Annotations = 10,
};

@interface SMSResponse : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) String *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *from;
/** Test to see if @c from has been set. */
@property(nonatomic, readwrite) BOOL hasFrom;

@property(nonatomic, readwrite, strong, null_resettable) String *mediaURL;
/** Test to see if @c mediaURL has been set. */
@property(nonatomic, readwrite) BOOL hasMediaURL;

@property(nonatomic, readwrite, strong, null_resettable) String *body;
/** Test to see if @c body has been set. */
@property(nonatomic, readwrite) BOOL hasBody;

@property(nonatomic, readwrite, strong, null_resettable) String *status;
/** Test to see if @c status has been set. */
@property(nonatomic, readwrite) BOOL hasStatus;

@property(nonatomic, readwrite, strong, null_resettable) StringMap *annotations;
/** Test to see if @c annotations has been set. */
@property(nonatomic, readwrite) BOOL hasAnnotations;

@end

#pragma mark - SubscriptionResponse

typedef GPB_ENUM(SubscriptionResponse_FieldNumber) {
  SubscriptionResponse_FieldNumber_Id_p = 1,
  SubscriptionResponse_FieldNumber_Amount = 2,
  SubscriptionResponse_FieldNumber_DaysUntilDue = 3,
  SubscriptionResponse_FieldNumber_Plan = 4,
  SubscriptionResponse_FieldNumber_User = 5,
  SubscriptionResponse_FieldNumber_Status = 6,
  SubscriptionResponse_FieldNumber_Annotations = 10,
};

@interface SubscriptionResponse : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) Int64 *amount;
/** Test to see if @c amount has been set. */
@property(nonatomic, readwrite) BOOL hasAmount;

@property(nonatomic, readwrite, strong, null_resettable) Int64 *daysUntilDue;
/** Test to see if @c daysUntilDue has been set. */
@property(nonatomic, readwrite) BOOL hasDaysUntilDue;

@property(nonatomic, readwrite, strong, null_resettable) StringMap *annotations;
/** Test to see if @c annotations has been set. */
@property(nonatomic, readwrite) BOOL hasAnnotations;

@property(nonatomic, readwrite) Plan plan;

@property(nonatomic, readwrite, strong, null_resettable) User *user;
/** Test to see if @c user has been set. */
@property(nonatomic, readwrite) BOOL hasUser;

@property(nonatomic, readwrite, strong, null_resettable) String *status;
/** Test to see if @c status has been set. */
@property(nonatomic, readwrite) BOOL hasStatus;

@end

/**
 * Fetches the raw value of a @c SubscriptionResponse's @c plan property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t SubscriptionResponse_Plan_RawValue(SubscriptionResponse *message);
/**
 * Sets the raw value of an @c SubscriptionResponse's @c plan property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetSubscriptionResponse_Plan_RawValue(SubscriptionResponse *message, int32_t value);

#pragma mark - FaxResponse

typedef GPB_ENUM(FaxResponse_FieldNumber) {
  FaxResponse_FieldNumber_Id_p = 1,
  FaxResponse_FieldNumber_MediaURL = 3,
  FaxResponse_FieldNumber_To = 4,
  FaxResponse_FieldNumber_From = 5,
  FaxResponse_FieldNumber_Status = 6,
  FaxResponse_FieldNumber_Annotations = 10,
};

@interface FaxResponse : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) String *mediaURL;
/** Test to see if @c mediaURL has been set. */
@property(nonatomic, readwrite) BOOL hasMediaURL;

@property(nonatomic, readwrite, strong, null_resettable) String *to;
/** Test to see if @c to has been set. */
@property(nonatomic, readwrite) BOOL hasTo;

@property(nonatomic, readwrite, strong, null_resettable) String *from;
/** Test to see if @c from has been set. */
@property(nonatomic, readwrite) BOOL hasFrom;

@property(nonatomic, readwrite, strong, null_resettable) String *status;
/** Test to see if @c status has been set. */
@property(nonatomic, readwrite) BOOL hasStatus;

@property(nonatomic, readwrite, strong, null_resettable) StringMap *annotations;
/** Test to see if @c annotations has been set. */
@property(nonatomic, readwrite) BOOL hasAnnotations;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
