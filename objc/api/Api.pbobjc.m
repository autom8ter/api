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
#import "google/api/Annotations.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - ApiRoot

@implementation ApiRoot

+ (GPBExtensionRegistry*)extensionRegistry {
  // This is called by +initialize so there is no need to worry
  // about thread safety and initialization of registry.
  static GPBExtensionRegistry* registry = nil;
  if (!registry) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    registry = [[GPBExtensionRegistry alloc] init];
    // Merge in the imports (direct or indirect) that defined extensions.
    [registry addExtensions:[GAPIAnnotationsRoot extensionRegistry]];
  }
  return registry;
}

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

#pragma mark - GetProfileByEmail

@implementation GetProfileByEmail

@dynamic email;

typedef struct GetProfileByEmail__storage_ {
  uint32_t _has_storage_[1];
  NSString *email;
} GetProfileByEmail__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "email",
        .dataTypeSpecific.className = NULL,
        .number = GetProfileByEmail_FieldNumber_Email,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(GetProfileByEmail__storage_, email),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[GetProfileByEmail class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(GetProfileByEmail__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - Profile

@implementation Profile

@dynamic email;
@dynamic emailVerified;
@dynamic name;
@dynamic givenName;
@dynamic familyName;
@dynamic picture;
@dynamic locale;
@dynamic userId;
@dynamic nickname;
@dynamic connection;
@dynamic identitiesArray, identitiesArray_Count;
@dynamic lastIp;
@dynamic loginCount;
@dynamic updatedAt;
@dynamic createdAt;
@dynamic hasUserMetadata, userMetadata;

typedef struct Profile__storage_ {
  uint32_t _has_storage_[1];
  NSString *email;
  NSString *name;
  NSString *givenName;
  NSString *familyName;
  NSString *picture;
  NSString *locale;
  NSString *userId;
  NSString *nickname;
  NSString *connection;
  NSMutableArray *identitiesArray;
  NSString *lastIp;
  NSString *updatedAt;
  NSString *createdAt;
  UserMetadata *userMetadata;
  int64_t loginCount;
} Profile__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "email",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_Email,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Profile__storage_, email),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "emailVerified",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_EmailVerified,
        .hasIndex = 1,
        .offset = 2,  // Stored in _has_storage_ to save space.
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeBool,
      },
      {
        .name = "name",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_Name,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(Profile__storage_, name),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "givenName",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_GivenName,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(Profile__storage_, givenName),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "familyName",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_FamilyName,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(Profile__storage_, familyName),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "picture",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_Picture,
        .hasIndex = 6,
        .offset = (uint32_t)offsetof(Profile__storage_, picture),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "locale",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_Locale,
        .hasIndex = 7,
        .offset = (uint32_t)offsetof(Profile__storage_, locale),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "userId",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_UserId,
        .hasIndex = 8,
        .offset = (uint32_t)offsetof(Profile__storage_, userId),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "nickname",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_Nickname,
        .hasIndex = 9,
        .offset = (uint32_t)offsetof(Profile__storage_, nickname),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "connection",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_Connection,
        .hasIndex = 10,
        .offset = (uint32_t)offsetof(Profile__storage_, connection),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "identitiesArray",
        .dataTypeSpecific.className = GPBStringifySymbol(Identity),
        .number = Profile_FieldNumber_IdentitiesArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(Profile__storage_, identitiesArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "lastIp",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_LastIp,
        .hasIndex = 11,
        .offset = (uint32_t)offsetof(Profile__storage_, lastIp),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "loginCount",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_LoginCount,
        .hasIndex = 12,
        .offset = (uint32_t)offsetof(Profile__storage_, loginCount),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "updatedAt",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_UpdatedAt,
        .hasIndex = 13,
        .offset = (uint32_t)offsetof(Profile__storage_, updatedAt),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "createdAt",
        .dataTypeSpecific.className = NULL,
        .number = Profile_FieldNumber_CreatedAt,
        .hasIndex = 14,
        .offset = (uint32_t)offsetof(Profile__storage_, createdAt),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "userMetadata",
        .dataTypeSpecific.className = GPBStringifySymbol(UserMetadata),
        .number = Profile_FieldNumber_UserMetadata,
        .hasIndex = 15,
        .offset = (uint32_t)offsetof(Profile__storage_, userMetadata),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Profile class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Profile__storage_)
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
@dynamic plan;
@dynamic payToken;
@dynamic lastContact;

typedef struct UserMetadata__storage_ {
  uint32_t _has_storage_[1];
  NSString *phone;
  NSString *plan;
  NSString *payToken;
  NSString *lastContact;
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
        .name = "plan",
        .dataTypeSpecific.className = NULL,
        .number = UserMetadata_FieldNumber_Plan,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(UserMetadata__storage_, plan),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "payToken",
        .dataTypeSpecific.className = NULL,
        .number = UserMetadata_FieldNumber_PayToken,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(UserMetadata__storage_, payToken),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "lastContact",
        .dataTypeSpecific.className = NULL,
        .number = UserMetadata_FieldNumber_LastContact,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(UserMetadata__storage_, lastContact),
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

#pragma mark - Identity

@implementation Identity

@dynamic provider;
@dynamic userId;
@dynamic connection;
@dynamic isSocial;

typedef struct Identity__storage_ {
  uint32_t _has_storage_[1];
  NSString *provider;
  NSString *userId;
  NSString *connection;
} Identity__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "provider",
        .dataTypeSpecific.className = NULL,
        .number = Identity_FieldNumber_Provider,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Identity__storage_, provider),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "userId",
        .dataTypeSpecific.className = NULL,
        .number = Identity_FieldNumber_UserId,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(Identity__storage_, userId),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "connection",
        .dataTypeSpecific.className = NULL,
        .number = Identity_FieldNumber_Connection,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(Identity__storage_, connection),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "isSocial",
        .dataTypeSpecific.className = NULL,
        .number = Identity_FieldNumber_IsSocial,
        .hasIndex = 3,
        .offset = 4,  // Stored in _has_storage_ to save space.
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeBool,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Identity class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Identity__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\001\004\010\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
