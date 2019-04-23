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

#pragma mark - GetByEmail

typedef GPB_ENUM(GetByEmail_FieldNumber) {
  GetByEmail_FieldNumber_Email = 1,
};

@interface GetByEmail : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@end

#pragma mark - IDToken

typedef GPB_ENUM(IDToken_FieldNumber) {
  IDToken_FieldNumber_Iss = 1,
  IDToken_FieldNumber_Sub = 2,
  IDToken_FieldNumber_Aud = 3,
  IDToken_FieldNumber_Exp = 4,
  IDToken_FieldNumber_Iat = 5,
  IDToken_FieldNumber_Name = 6,
  IDToken_FieldNumber_GivenName = 7,
  IDToken_FieldNumber_FamilyName = 8,
  IDToken_FieldNumber_Gender = 9,
  IDToken_FieldNumber_Birthdate = 10,
  IDToken_FieldNumber_Email = 11,
  IDToken_FieldNumber_Picture = 12,
};

@interface IDToken : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *iss;

@property(nonatomic, readwrite, copy, null_resettable) NSString *sub;

@property(nonatomic, readwrite, copy, null_resettable) NSString *aud;

@property(nonatomic, readwrite) int64_t exp;

@property(nonatomic, readwrite) int64_t iat;

@property(nonatomic, readwrite, copy, null_resettable) NSString *name;

@property(nonatomic, readwrite, copy, null_resettable) NSString *givenName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *familyName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *gender;

@property(nonatomic, readwrite, copy, null_resettable) NSString *birthdate;

@property(nonatomic, readwrite, copy, null_resettable) NSString *email;

@property(nonatomic, readwrite, copy, null_resettable) NSString *picture;

@end

#pragma mark - UserMetadata

typedef GPB_ENUM(UserMetadata_FieldNumber) {
  UserMetadata_FieldNumber_Phone = 1,
  UserMetadata_FieldNumber_Plan = 2,
  UserMetadata_FieldNumber_PayToken = 3,
  UserMetadata_FieldNumber_LastContact = 4,
};

@interface UserMetadata : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *phone;

@property(nonatomic, readwrite, copy, null_resettable) NSString *plan;

@property(nonatomic, readwrite, copy, null_resettable) NSString *payToken;

@property(nonatomic, readwrite, copy, null_resettable) NSString *lastContact;

@end

#pragma mark - AccessToken

typedef GPB_ENUM(AccessToken_FieldNumber) {
  AccessToken_FieldNumber_Iss = 1,
  AccessToken_FieldNumber_Sub = 2,
  AccessToken_FieldNumber_AudArray = 3,
  AccessToken_FieldNumber_Azp = 4,
  AccessToken_FieldNumber_Exp = 5,
  AccessToken_FieldNumber_Iat = 6,
  AccessToken_FieldNumber_Scope = 7,
};

@interface AccessToken : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *iss;

@property(nonatomic, readwrite, copy, null_resettable) NSString *sub;

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<NSString*> *audArray;
/** The number of items in @c audArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger audArray_Count;

@property(nonatomic, readwrite, copy, null_resettable) NSString *azp;

@property(nonatomic, readwrite) int64_t exp;

@property(nonatomic, readwrite) int64_t iat;

@property(nonatomic, readwrite, copy, null_resettable) NSString *scope;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
