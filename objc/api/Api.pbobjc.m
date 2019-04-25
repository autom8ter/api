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

#pragma mark - Identifier

@implementation Identifier

@dynamic id_p;

typedef struct Identifier__storage_ {
  uint32_t _has_storage_[1];
  NSString *id_p;
} Identifier__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "id_p",
        .dataTypeSpecific.className = NULL,
        .number = Identifier_FieldNumber_Id_p,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Identifier__storage_, id_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Identifier class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Identifier__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - SMSStatus

@implementation SMSStatus

@dynamic hasId_p, id_p;
@dynamic hasSms, sms;
@dynamic status;
@dynamic uri;

typedef struct SMSStatus__storage_ {
  uint32_t _has_storage_[1];
  Identifier *id_p;
  SMS *sms;
  NSString *status;
  NSString *uri;
} SMSStatus__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "id_p",
        .dataTypeSpecific.className = GPBStringifySymbol(Identifier),
        .number = SMSStatus_FieldNumber_Id_p,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(SMSStatus__storage_, id_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "sms",
        .dataTypeSpecific.className = GPBStringifySymbol(SMS),
        .number = SMSStatus_FieldNumber_Sms,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(SMSStatus__storage_, sms),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "status",
        .dataTypeSpecific.className = NULL,
        .number = SMSStatus_FieldNumber_Status,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(SMSStatus__storage_, status),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "uri",
        .dataTypeSpecific.className = NULL,
        .number = SMSStatus_FieldNumber_Uri,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(SMSStatus__storage_, uri),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[SMSStatus class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(SMSStatus__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - SMS

@implementation SMS

@dynamic from;
@dynamic to;
@dynamic hasMessage, message;
@dynamic mediaURL;

typedef struct SMS__storage_ {
  uint32_t _has_storage_[1];
  NSString *from;
  NSString *to;
  Message *message;
  NSString *mediaURL;
} SMS__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "from",
        .dataTypeSpecific.className = NULL,
        .number = SMS_FieldNumber_From,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(SMS__storage_, from),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "to",
        .dataTypeSpecific.className = NULL,
        .number = SMS_FieldNumber_To,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(SMS__storage_, to),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "message",
        .dataTypeSpecific.className = GPBStringifySymbol(Message),
        .number = SMS_FieldNumber_Message,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(SMS__storage_, message),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
      {
        .name = "mediaURL",
        .dataTypeSpecific.className = NULL,
        .number = SMS_FieldNumber_MediaURL,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(SMS__storage_, mediaURL),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[SMS class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(SMS__storage_)
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

#pragma mark - EmailRequest

@implementation EmailRequest

@dynamic fromName;
@dynamic fromEmail;
@dynamic hasEmail, email;

typedef struct EmailRequest__storage_ {
  uint32_t _has_storage_[1];
  NSString *fromName;
  NSString *fromEmail;
  Email *email;
} EmailRequest__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "fromName",
        .dataTypeSpecific.className = NULL,
        .number = EmailRequest_FieldNumber_FromName,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(EmailRequest__storage_, fromName),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "fromEmail",
        .dataTypeSpecific.className = NULL,
        .number = EmailRequest_FieldNumber_FromEmail,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(EmailRequest__storage_, fromEmail),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "email",
        .dataTypeSpecific.className = GPBStringifySymbol(Email),
        .number = EmailRequest_FieldNumber_Email,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(EmailRequest__storage_, email),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[EmailRequest class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(EmailRequest__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - Email

@implementation Email

@dynamic name;
@dynamic address;
@dynamic subject;
@dynamic plain;
@dynamic html;

typedef struct Email__storage_ {
  uint32_t _has_storage_[1];
  NSString *name;
  NSString *address;
  NSString *subject;
  NSString *plain;
  NSString *html;
} Email__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "name",
        .dataTypeSpecific.className = NULL,
        .number = Email_FieldNumber_Name,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Email__storage_, name),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "address",
        .dataTypeSpecific.className = NULL,
        .number = Email_FieldNumber_Address,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(Email__storage_, address),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "subject",
        .dataTypeSpecific.className = NULL,
        .number = Email_FieldNumber_Subject,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(Email__storage_, subject),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "plain",
        .dataTypeSpecific.className = NULL,
        .number = Email_FieldNumber_Plain,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(Email__storage_, plain),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "html",
        .dataTypeSpecific.className = NULL,
        .number = Email_FieldNumber_Html,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(Email__storage_, html),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Email class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Email__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - Call

@implementation Call

@dynamic from;
@dynamic to;
@dynamic callback;

typedef struct Call__storage_ {
  uint32_t _has_storage_[1];
  NSString *from;
  NSString *to;
  NSString *callback;
} Call__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "from",
        .dataTypeSpecific.className = NULL,
        .number = Call_FieldNumber_From,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Call__storage_, from),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "to",
        .dataTypeSpecific.className = NULL,
        .number = Call_FieldNumber_To,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(Call__storage_, to),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "callback",
        .dataTypeSpecific.className = NULL,
        .number = Call_FieldNumber_Callback,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(Call__storage_, callback),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Call class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Call__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - Message

@implementation Message

@dynamic value;

typedef struct Message__storage_ {
  uint32_t _has_storage_[1];
  NSString *value;
} Message__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "value",
        .dataTypeSpecific.className = NULL,
        .number = Message_FieldNumber_Value,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Message__storage_, value),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Message class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Message__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

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
@dynamic status;
@dynamic tagsArray, tagsArray_Count;

typedef struct UserMetadata__storage_ {
  uint32_t _has_storage_[1];
  NSString *phone;
  NSString *preferredContact;
  NSString *status;
  NSMutableArray *tagsArray;
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
      {
        .name = "status",
        .dataTypeSpecific.className = NULL,
        .number = UserMetadata_FieldNumber_Status,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(UserMetadata__storage_, status),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "tagsArray",
        .dataTypeSpecific.className = NULL,
        .number = UserMetadata_FieldNumber_TagsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(UserMetadata__storage_, tagsArray),
        .flags = GPBFieldRepeated,
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
@dynamic delinquent;
@dynamic tagsArray, tagsArray_Count;

typedef struct AppMetadata__storage_ {
  uint32_t _has_storage_[1];
  NSString *plan;
  NSString *payToken;
  NSString *delinquent;
  NSMutableArray *tagsArray;
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
      {
        .name = "delinquent",
        .dataTypeSpecific.className = NULL,
        .number = AppMetadata_FieldNumber_Delinquent,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(AppMetadata__storage_, delinquent),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "tagsArray",
        .dataTypeSpecific.className = NULL,
        .number = AppMetadata_FieldNumber_TagsArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(AppMetadata__storage_, tagsArray),
        .flags = GPBFieldRepeated,
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

#pragma mark - Auth

@implementation Auth

@dynamic domain;
@dynamic clientId;
@dynamic clientSecret;
@dynamic redirect;
@dynamic audience;
@dynamic scopesArray, scopesArray_Count;

typedef struct Auth__storage_ {
  uint32_t _has_storage_[1];
  NSString *domain;
  NSString *clientId;
  NSString *clientSecret;
  NSString *redirect;
  NSString *audience;
  NSMutableArray *scopesArray;
} Auth__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "domain",
        .dataTypeSpecific.className = NULL,
        .number = Auth_FieldNumber_Domain,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Auth__storage_, domain),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "clientId",
        .dataTypeSpecific.className = NULL,
        .number = Auth_FieldNumber_ClientId,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(Auth__storage_, clientId),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "clientSecret",
        .dataTypeSpecific.className = NULL,
        .number = Auth_FieldNumber_ClientSecret,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(Auth__storage_, clientSecret),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "redirect",
        .dataTypeSpecific.className = NULL,
        .number = Auth_FieldNumber_Redirect,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(Auth__storage_, redirect),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "audience",
        .dataTypeSpecific.className = NULL,
        .number = Auth_FieldNumber_Audience,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(Auth__storage_, audience),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "scopesArray",
        .dataTypeSpecific.className = NULL,
        .number = Auth_FieldNumber_ScopesArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(Auth__storage_, scopesArray),
        .flags = GPBFieldRepeated,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Auth class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Auth__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - Bytes

@implementation Bytes

@dynamic bits;

typedef struct Bytes__storage_ {
  uint32_t _has_storage_[1];
  NSData *bits;
} Bytes__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "bits",
        .dataTypeSpecific.className = NULL,
        .number = Bytes_FieldNumber_Bits,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Bytes__storage_, bits),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeBytes,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Bytes class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Bytes__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - Template

@implementation Template

@dynamic text;
@dynamic data_p;

typedef struct Template__storage_ {
  uint32_t _has_storage_[1];
  NSString *text;
  NSData *data_p;
} Template__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "text",
        .dataTypeSpecific.className = NULL,
        .number = Template_FieldNumber_Text,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(Template__storage_, text),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "data_p",
        .dataTypeSpecific.className = NULL,
        .number = Template_FieldNumber_Data_p,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(Template__storage_, data_p),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeBytes,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[Template class]
                                     rootClass:[ApiRoot class]
                                          file:ApiRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(Template__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
