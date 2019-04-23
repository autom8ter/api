/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_api_annotations_pb = require('./google/api/annotations_pb.js');
goog.exportSymbol('proto.api.AccessToken', null, global);
goog.exportSymbol('proto.api.GetByEmail', null, global);
goog.exportSymbol('proto.api.IDToken', null, global);
goog.exportSymbol('proto.api.UserMetadata', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.api.GetByEmail = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.GetByEmail, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.api.GetByEmail.displayName = 'proto.api.GetByEmail';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.api.GetByEmail.prototype.toObject = function(opt_includeInstance) {
  return proto.api.GetByEmail.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.GetByEmail} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.GetByEmail.toObject = function(includeInstance, msg) {
  var f, obj = {
    email: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.api.GetByEmail}
 */
proto.api.GetByEmail.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.GetByEmail;
  return proto.api.GetByEmail.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.GetByEmail} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.GetByEmail}
 */
proto.api.GetByEmail.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.api.GetByEmail.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.GetByEmail.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.GetByEmail} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.GetByEmail.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string email = 1;
 * @return {string}
 */
proto.api.GetByEmail.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.api.GetByEmail.prototype.setEmail = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.api.IDToken = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.IDToken, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.api.IDToken.displayName = 'proto.api.IDToken';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.api.IDToken.prototype.toObject = function(opt_includeInstance) {
  return proto.api.IDToken.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.IDToken} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.IDToken.toObject = function(includeInstance, msg) {
  var f, obj = {
    iss: jspb.Message.getFieldWithDefault(msg, 1, ""),
    sub: jspb.Message.getFieldWithDefault(msg, 2, false),
    aud: jspb.Message.getFieldWithDefault(msg, 3, ""),
    exp: jspb.Message.getFieldWithDefault(msg, 4, 0),
    iat: jspb.Message.getFieldWithDefault(msg, 5, 0),
    name: jspb.Message.getFieldWithDefault(msg, 6, ""),
    givenName: jspb.Message.getFieldWithDefault(msg, 7, ""),
    familyName: jspb.Message.getFieldWithDefault(msg, 8, ""),
    gender: jspb.Message.getFieldWithDefault(msg, 9, ""),
    birthdate: jspb.Message.getFieldWithDefault(msg, 10, ""),
    email: jspb.Message.getFieldWithDefault(msg, 11, ""),
    picture: jspb.Message.getFieldWithDefault(msg, 12, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.api.IDToken}
 */
proto.api.IDToken.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.IDToken;
  return proto.api.IDToken.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.IDToken} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.IDToken}
 */
proto.api.IDToken.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setIss(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSub(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setAud(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setExp(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIat(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setGivenName(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setFamilyName(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setGender(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setBirthdate(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPicture(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.api.IDToken.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.IDToken.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.IDToken} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.IDToken.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIss();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSub();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
  f = message.getAud();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getExp();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getIat();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getGivenName();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getFamilyName();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getGender();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getBirthdate();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getPicture();
  if (f !== 0) {
    writer.writeInt64(
      12,
      f
    );
  }
};


/**
 * optional string iss = 1;
 * @return {string}
 */
proto.api.IDToken.prototype.getIss = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setIss = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional bool sub = 2;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.api.IDToken.prototype.getSub = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 2, false));
};


/** @param {boolean} value */
proto.api.IDToken.prototype.setSub = function(value) {
  jspb.Message.setProto3BooleanField(this, 2, value);
};


/**
 * optional string aud = 3;
 * @return {string}
 */
proto.api.IDToken.prototype.getAud = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setAud = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional int64 exp = 4;
 * @return {number}
 */
proto.api.IDToken.prototype.getExp = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.api.IDToken.prototype.setExp = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 iat = 5;
 * @return {number}
 */
proto.api.IDToken.prototype.getIat = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.api.IDToken.prototype.setIat = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional string name = 6;
 * @return {string}
 */
proto.api.IDToken.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string given_name = 7;
 * @return {string}
 */
proto.api.IDToken.prototype.getGivenName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setGivenName = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string family_name = 8;
 * @return {string}
 */
proto.api.IDToken.prototype.getFamilyName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setFamilyName = function(value) {
  jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string gender = 9;
 * @return {string}
 */
proto.api.IDToken.prototype.getGender = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setGender = function(value) {
  jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional string birthdate = 10;
 * @return {string}
 */
proto.api.IDToken.prototype.getBirthdate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setBirthdate = function(value) {
  jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string email = 11;
 * @return {string}
 */
proto.api.IDToken.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.api.IDToken.prototype.setEmail = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional int64 picture = 12;
 * @return {number}
 */
proto.api.IDToken.prototype.getPicture = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.api.IDToken.prototype.setPicture = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.api.UserMetadata = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.UserMetadata, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.api.UserMetadata.displayName = 'proto.api.UserMetadata';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.api.UserMetadata.prototype.toObject = function(opt_includeInstance) {
  return proto.api.UserMetadata.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.UserMetadata} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.UserMetadata.toObject = function(includeInstance, msg) {
  var f, obj = {
    phone: jspb.Message.getFieldWithDefault(msg, 1, ""),
    plan: jspb.Message.getFieldWithDefault(msg, 2, ""),
    payToken: jspb.Message.getFieldWithDefault(msg, 3, ""),
    lastContact: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.api.UserMetadata}
 */
proto.api.UserMetadata.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.UserMetadata;
  return proto.api.UserMetadata.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.UserMetadata} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.UserMetadata}
 */
proto.api.UserMetadata.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setPhone(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPlan(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPayToken(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastContact(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.api.UserMetadata.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.UserMetadata.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.UserMetadata} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.UserMetadata.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPhone();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getPlan();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPayToken();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getLastContact();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string phone = 1;
 * @return {string}
 */
proto.api.UserMetadata.prototype.getPhone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.api.UserMetadata.prototype.setPhone = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string plan = 2;
 * @return {string}
 */
proto.api.UserMetadata.prototype.getPlan = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.api.UserMetadata.prototype.setPlan = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string pay_token = 3;
 * @return {string}
 */
proto.api.UserMetadata.prototype.getPayToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.api.UserMetadata.prototype.setPayToken = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string last_contact = 4;
 * @return {string}
 */
proto.api.UserMetadata.prototype.getLastContact = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.api.UserMetadata.prototype.setLastContact = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.api.AccessToken = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.api.AccessToken.repeatedFields_, null);
};
goog.inherits(proto.api.AccessToken, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.api.AccessToken.displayName = 'proto.api.AccessToken';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.api.AccessToken.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.api.AccessToken.prototype.toObject = function(opt_includeInstance) {
  return proto.api.AccessToken.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.AccessToken} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.AccessToken.toObject = function(includeInstance, msg) {
  var f, obj = {
    iss: jspb.Message.getFieldWithDefault(msg, 1, ""),
    sub: jspb.Message.getFieldWithDefault(msg, 2, ""),
    audList: jspb.Message.getRepeatedField(msg, 3),
    azp: jspb.Message.getFieldWithDefault(msg, 4, ""),
    exp: jspb.Message.getFieldWithDefault(msg, 5, 0),
    iat: jspb.Message.getFieldWithDefault(msg, 6, 0),
    scope: jspb.Message.getFieldWithDefault(msg, 7, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.api.AccessToken}
 */
proto.api.AccessToken.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.AccessToken;
  return proto.api.AccessToken.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.AccessToken} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.AccessToken}
 */
proto.api.AccessToken.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setIss(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setSub(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addAud(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setAzp(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setExp(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIat(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setScope(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.api.AccessToken.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.AccessToken.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.AccessToken} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.AccessToken.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIss();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSub();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getAudList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
  f = message.getAzp();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getExp();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getIat();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getScope();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
};


/**
 * optional string iss = 1;
 * @return {string}
 */
proto.api.AccessToken.prototype.getIss = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.api.AccessToken.prototype.setIss = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string sub = 2;
 * @return {string}
 */
proto.api.AccessToken.prototype.getSub = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.api.AccessToken.prototype.setSub = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * repeated string aud = 3;
 * @return {!Array<string>}
 */
proto.api.AccessToken.prototype.getAudList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/** @param {!Array<string>} value */
proto.api.AccessToken.prototype.setAudList = function(value) {
  jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.api.AccessToken.prototype.addAud = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


proto.api.AccessToken.prototype.clearAudList = function() {
  this.setAudList([]);
};


/**
 * optional string azp = 4;
 * @return {string}
 */
proto.api.AccessToken.prototype.getAzp = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.api.AccessToken.prototype.setAzp = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int64 exp = 5;
 * @return {number}
 */
proto.api.AccessToken.prototype.getExp = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.api.AccessToken.prototype.setExp = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int64 iat = 6;
 * @return {number}
 */
proto.api.AccessToken.prototype.getIat = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.api.AccessToken.prototype.setIat = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional string scope = 7;
 * @return {string}
 */
proto.api.AccessToken.prototype.getScope = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.api.AccessToken.prototype.setScope = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


goog.object.extend(exports, proto.api);
