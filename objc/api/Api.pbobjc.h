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

@class Address;
@class User;

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

#pragma mark - UserMap

typedef GPB_ENUM(UserMap_FieldNumber) {
  UserMap_FieldNumber_Users = 1,
};

/**
 * UserMap is a map of users with a key of either user id, email, or phone number
 **/
@interface UserMap : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, User*> *users;
/** The number of items in @c users without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger users_Count;

@end

#pragma mark - User

typedef GPB_ENUM(User_FieldNumber) {
  User_FieldNumber_UserId = 1,
  User_FieldNumber_Plan = 2,
  User_FieldNumber_Name = 3,
  User_FieldNumber_Email = 4,
  User_FieldNumber_Description_p = 5,
  User_FieldNumber_Phone = 6,
  User_FieldNumber_Address = 8,
  User_FieldNumber_Metadata = 9,
  User_FieldNumber_Deleted = 10,
  User_FieldNumber_CreateDate = 20,
};

/**
 * User is a user of the application
 **/
@interface User : GPBMessage

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

#pragma mark - AddUserRequest

typedef GPB_ENUM(AddUserRequest_FieldNumber) {
  AddUserRequest_FieldNumber_Email = 1,
  AddUserRequest_FieldNumber_Plan = 2,
  AddUserRequest_FieldNumber_Phone = 3,
  AddUserRequest_FieldNumber_Name = 4,
  AddUserRequest_FieldNumber_Password = 5,
  AddUserRequest_FieldNumber_TrialEnd = 6,
  AddUserRequest_FieldNumber_Description_p = 7,
  AddUserRequest_FieldNumber_Address = 8,
};

@interface AddUserRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, copy, null_resettable) NSString *phone;

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *password;

@property(nonatomic, readwrite) int64_t trialEnd;

@property(nonatomic, readwrite, copy, null_resettable) NSString *description_p;

@property(nonatomic, readwrite, strong, null_resettable) Address *address;
/** Test to see if @c address has been set. */
@property(nonatomic, readwrite) BOOL hasAddress;

@end

#pragma mark - SubscribeUserRequest

typedef GPB_ENUM(SubscribeUserRequest_FieldNumber) {
  SubscribeUserRequest_FieldNumber_Email = 1,
  SubscribeUserRequest_FieldNumber_Plan = 2,
  SubscribeUserRequest_FieldNumber_CardNumber = 3,
  SubscribeUserRequest_FieldNumber_ExpMonth = 4,
  SubscribeUserRequest_FieldNumber_ExpYear = 5,
  SubscribeUserRequest_FieldNumber_Cvc = 6,
};

@interface SubscribeUserRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, copy, null_resettable) NSString *cardNumber;

@property(nonatomic, readwrite, copy, null_resettable) NSString *expMonth;

@property(nonatomic, readwrite, copy, null_resettable) NSString *expYear;

@property(nonatomic, readwrite, copy, null_resettable) NSString *cvc;

@end

#pragma mark - AddUserMetadataRequest

typedef GPB_ENUM(AddUserMetadataRequest_FieldNumber) {
  AddUserMetadataRequest_FieldNumber_UserId = 1,
  AddUserMetadataRequest_FieldNumber_Metadata = 2,
};

@interface AddUserMetadataRequest : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *userId;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableDictionary<NSString*, NSString*> *metadata;
/** The number of items in @c metadata without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger metadata_Count;

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

#pragma mark - SubscribeUserResponse

typedef GPB_ENUM(SubscribeUserResponse_FieldNumber) {
  SubscribeUserResponse_FieldNumber_SubscriptionId = 1,
};

@interface SubscribeUserResponse : GPBMessage

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

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
