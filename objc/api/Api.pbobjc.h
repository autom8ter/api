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
@class Email;
@class Identifier;
@class Message;
@class SMS;
@class UserMetadata;

NS_ASSUME_NONNULL_BEGIN

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

#pragma mark - Dashboard

typedef GPB_ENUM(Dashboard_FieldNumber) {
  Dashboard_FieldNumber_TotalUsers = 1,
  Dashboard_FieldNumber_TotalCustomers = 2,
  Dashboard_FieldNumber_TotalPlans = 3,
  Dashboard_FieldNumber_TotalSubscriptions = 4,
  Dashboard_FieldNumber_TotalCharges = 5,
};

@interface Dashboard : GPBMessage

@property(nonatomic, readwrite) int64_t totalUsers;

@property(nonatomic, readwrite) int64_t totalCustomers;

@property(nonatomic, readwrite) int64_t totalPlans;

@property(nonatomic, readwrite) int64_t totalSubscriptions;

@property(nonatomic, readwrite) int64_t totalCharges;

@end

#pragma mark - Identifier

typedef GPB_ENUM(Identifier_FieldNumber) {
  Identifier_FieldNumber_Id_p = 1,
};

@interface Identifier : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *id_p;

@end

#pragma mark - SMSStatus

typedef GPB_ENUM(SMSStatus_FieldNumber) {
  SMSStatus_FieldNumber_Id_p = 1,
  SMSStatus_FieldNumber_Sms = 2,
  SMSStatus_FieldNumber_Status = 3,
  SMSStatus_FieldNumber_Uri = 4,
};

@interface SMSStatus : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) Identifier *id_p;
/** Test to see if @c id_p has been set. */
@property(nonatomic, readwrite) BOOL hasId_p;

@property(nonatomic, readwrite, strong, null_resettable) SMS *sms;
/** Test to see if @c sms has been set. */
@property(nonatomic, readwrite) BOOL hasSms;

@property(nonatomic, readwrite, copy, null_resettable) NSString *status;

@property(nonatomic, readwrite, copy, null_resettable) NSString *uri;

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

#pragma mark - Message

typedef GPB_ENUM(Message_FieldNumber) {
  Message_FieldNumber_Value = 1,
};

@interface Message : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *value;

@end

#pragma mark - UserInfo

typedef GPB_ENUM(UserInfo_FieldNumber) {
  UserInfo_FieldNumber_Name = 6,
  UserInfo_FieldNumber_GivenName = 7,
  UserInfo_FieldNumber_FamilyName = 8,
  UserInfo_FieldNumber_Gender = 9,
  UserInfo_FieldNumber_Birthdate = 10,
  UserInfo_FieldNumber_Email = 11,
  UserInfo_FieldNumber_Picture = 12,
  UserInfo_FieldNumber_UserMetadata = 13,
  UserInfo_FieldNumber_AppMetadata = 14,
};

@interface UserInfo : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *givenName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *familyName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *gender;

@property(nonatomic, readwrite, copy, null_resettable) NSString *birthdate;

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *picture;

@property(nonatomic, readwrite, strong, null_resettable) UserMetadata *userMetadata;
/** Test to see if @c userMetadata has been set. */
@property(nonatomic, readwrite) BOOL hasUserMetadata;

@property(nonatomic, readwrite, strong, null_resettable) AppMetadata *appMetadata;
/** Test to see if @c appMetadata has been set. */
@property(nonatomic, readwrite) BOOL hasAppMetadata;

@end

#pragma mark - UserMetadata

typedef GPB_ENUM(UserMetadata_FieldNumber) {
  UserMetadata_FieldNumber_Phone = 1,
  UserMetadata_FieldNumber_PreferredContact = 2,
  UserMetadata_FieldNumber_Status = 3,
  UserMetadata_FieldNumber_TagsArray = 4,
};

@interface UserMetadata : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *phone;

@property(nonatomic, readwrite, copy, null_resettable) NSString *preferredContact;

@property(nonatomic, readwrite, copy, null_resettable) NSString *status;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *tagsArray;
/** The number of items in @c tagsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger tagsArray_Count;

@end

#pragma mark - AppMetadata

typedef GPB_ENUM(AppMetadata_FieldNumber) {
  AppMetadata_FieldNumber_Plan = 1,
  AppMetadata_FieldNumber_PayToken = 2,
  AppMetadata_FieldNumber_Delinquent = 3,
  AppMetadata_FieldNumber_TagsArray = 4,
};

@interface AppMetadata : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, copy, null_resettable) NSString *payToken;

@property(nonatomic, readwrite, copy, null_resettable) NSString *delinquent;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *tagsArray;
/** The number of items in @c tagsArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger tagsArray_Count;

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

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *scopesArray;
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
  Template_FieldNumber_Text = 1,
  Template_FieldNumber_Data_p = 2,
};

@interface Template : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *text;

@property(nonatomic, readwrite, copy, null_resettable) NSData *data_p;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
