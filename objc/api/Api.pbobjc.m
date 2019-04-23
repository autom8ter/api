// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: api.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

#import "Api.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - ApiRoot

@implementation ApiRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - ApiRoot_FileDescriptor

static GPBFileDescriptor *ApiRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"api"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - UserInfo

@implementation UserInfo

@dynamic name;
@dynamic givenName;
@dynamic familyName;
@dynamic gender;
@dynamic birthdate;
@dynamic email;
@dynamic picture;
@dynamic hasUserMetadata, userMetadata;
@dynamic hasAppMetadata, appMetadata;

typedef struct UserInfo__storage_ {
  uint32_t _has_storage_[1];
  NSString *name;
  NSString *givenName;
  NSString *familyName;
  NSString *gender;
  NSString *birthdate;
  NSString *email;
  NSString *picture;
  UserMetadata *userMetadata;
  AppMetadata *appMetadata;
} UserInfo__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "name",
        .dataTypeSpecific.className = NULL,
        .number = UserInfo_FieldNumber_Name,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(UserInfo__storage_, name),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "givenName",
        .dataTypeSpecific.className = NULL,
        .number = UserInfo_FieldNumber_GivenName,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(UserInfo__storage_, givenName),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "familyName",
        .dataTypeSpecific.className = NULL,
        .number = UserInfo_FieldNumber_FamilyName,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(UserInfo__storage_, familyName),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "gender",
        .dataTypeSpecific.className = NULL,
        .number = UserInfo_FieldNumber_Gender,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(UserInfo__storage_, gender),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "birthdate",
        .dataTypeSpecific.className = NULL,
        .number = UserInfo_FieldNumber_Birthdate,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(UserInfo__storage_, birthdate),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "email",
        .dataTypeSpecific.className = NULL,
        .number = UserInfo_FieldNumber_Email,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(UserInfo__storage_, email),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "picture",
        .dataTypeSpecific.className = NULL,
        .number = UserInfo_FieldNumber_Picture,
        .hasIndex = 6,
        .offset = (uint32_t)offsetof(UserInfo__storage_, picture),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "userMetadata",
        .dataTypeSpecific.className = GPBStringifySymbol(UserMetadata),
        .number = UserInfo_FieldNumber_UserMetadata,
        .hasIndex = 7,
        .offset = (uint32_t)offsetof(UserInfo__storage_, userMetadata),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "appMetadata",
        .dataTypeSpecific.className = GPBStringifySymbol(AppMetadata),
        .number = UserInfo_FieldNumber_AppMetadata,
        .hasIndex = 8,
        .offset = (uint32_t)offsetof(UserInfo__storage_, appMetadata),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[UserInfo class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(UserInfo__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - UserMetadata

@implementation UserMetadata

@dynamic phone;
@dynamic preferredContact;

typedef struct UserMetadata__storage_ {
  uint32_t _has_storage_[1];
  NSString *phone;
  NSString *preferredContact;
} UserMetadata__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "phone",
        .dataTypeSpecific.className = NULL,
        .number = UserMetadata_FieldNumber_Phone,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(UserMetadata__storage_, phone),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "preferredContact",
        .dataTypeSpecific.className = NULL,
        .number = UserMetadata_FieldNumber_PreferredContact,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(UserMetadata__storage_, preferredContact),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[UserMetadata class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(UserMetadata__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - AppMetadata

@implementation AppMetadata

@dynamic plan;
@dynamic payToken;

typedef struct AppMetadata__storage_ {
  uint32_t _has_storage_[1];
  NSString *plan;
  NSString *payToken;
} AppMetadata__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "plan",
        .dataTypeSpecific.className = NULL,
        .number = AppMetadata_FieldNumber_Plan,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(AppMetadata__storage_, plan),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "payToken",
        .dataTypeSpecific.className = NULL,
        .number = AppMetadata_FieldNumber_PayToken,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(AppMetadata__storage_, payToken),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[AppMetadata class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(AppMetadata__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - AuthToken

@implementation AuthToken

@dynamic token;

typedef struct AuthToken__storage_ {
  uint32_t _has_storage_[1];
  NSString *token;
} AuthToken__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "token",
        .dataTypeSpecific.className = NULL,
        .number = AuthToken_FieldNumber_Token,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(AuthToken__storage_, token),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[AuthToken class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(AuthToken__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - IDToken

@implementation IDToken

@dynamic token;

typedef struct IDToken__storage_ {
  uint32_t _has_storage_[1];
  NSString *token;
} IDToken__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "token",
        .dataTypeSpecific.className = NULL,
        .number = IDToken_FieldNumber_Token,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(IDToken__storage_, token),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[IDToken class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(IDToken__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
