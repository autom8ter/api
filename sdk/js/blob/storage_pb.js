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

var os_os_pb = require('../os/os_pb.js');
var auth_auth_pb = require('../auth/auth_pb.js');
goog.exportSymbol('proto.blob.Bucket', null, global);
goog.exportSymbol('proto.blob.CreateObjectRequest', null, global);
goog.exportSymbol('proto.blob.Object', null, global);

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
proto.blob.CreateObjectRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.blob.CreateObjectRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.blob.CreateObjectRequest.displayName = 'proto.blob.CreateObjectRequest';
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
proto.blob.CreateObjectRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.blob.CreateObjectRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.blob.CreateObjectRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blob.CreateObjectRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    bucket: jspb.Message.getFieldWithDefault(msg, 1, ""),
    file: (f = msg.getFile()) && os_os_pb.File.toObject(includeInstance, f)
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
 * @return {!proto.blob.CreateObjectRequest}
 */
proto.blob.CreateObjectRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.blob.CreateObjectRequest;
  return proto.blob.CreateObjectRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.blob.CreateObjectRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.blob.CreateObjectRequest}
 */
proto.blob.CreateObjectRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setBucket(value);
      break;
    case 2:
      var value = new os_os_pb.File;
      reader.readMessage(value,os_os_pb.File.deserializeBinaryFromReader);
      msg.setFile(value);
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
proto.blob.CreateObjectRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.blob.CreateObjectRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.blob.CreateObjectRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blob.CreateObjectRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBucket();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFile();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      os_os_pb.File.serializeBinaryToWriter
    );
  }
};


/**
 * optional string bucket = 1;
 * @return {string}
 */
proto.blob.CreateObjectRequest.prototype.getBucket = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.blob.CreateObjectRequest.prototype.setBucket = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional os.File file = 2;
 * @return {?proto.os.File}
 */
proto.blob.CreateObjectRequest.prototype.getFile = function() {
  return /** @type{?proto.os.File} */ (
    jspb.Message.getWrapperField(this, os_os_pb.File, 2));
};


/** @param {?proto.os.File|undefined} value */
proto.blob.CreateObjectRequest.prototype.setFile = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.blob.CreateObjectRequest.prototype.clearFile = function() {
  this.setFile(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.blob.CreateObjectRequest.prototype.hasFile = function() {
  return jspb.Message.getField(this, 2) != null;
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
proto.blob.Object = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.blob.Object, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.blob.Object.displayName = 'proto.blob.Object';
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
proto.blob.Object.prototype.toObject = function(opt_includeInstance) {
  return proto.blob.Object.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.blob.Object} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blob.Object.toObject = function(includeInstance, msg) {
  var f, obj = {
    bucket: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    contentType: jspb.Message.getFieldWithDefault(msg, 3, ""),
    contentLanguage: jspb.Message.getFieldWithDefault(msg, 4, ""),
    cacheControl: jspb.Message.getFieldWithDefault(msg, 5, ""),
    owner: jspb.Message.getFieldWithDefault(msg, 6, ""),
    size: jspb.Message.getFieldWithDefault(msg, 7, 0),
    metaMap: (f = msg.getMetaMap()) ? f.toObject(includeInstance, undefined) : [],
    created: jspb.Message.getFieldWithDefault(msg, 9, 0),
    updated: jspb.Message.getFieldWithDefault(msg, 10, 0),
    deleted: jspb.Message.getFieldWithDefault(msg, 11, 0)
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
 * @return {!proto.blob.Object}
 */
proto.blob.Object.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.blob.Object;
  return proto.blob.Object.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.blob.Object} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.blob.Object}
 */
proto.blob.Object.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setBucket(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setContentType(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setContentLanguage(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setCacheControl(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setOwner(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSize(value);
      break;
    case 8:
      var value = msg.getMetaMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCreated(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUpdated(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDeleted(value);
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
proto.blob.Object.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.blob.Object.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.blob.Object} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blob.Object.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBucket();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getContentType();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getContentLanguage();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getCacheControl();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getOwner();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getSize();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getMetaMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(8, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreated();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
  f = message.getUpdated();
  if (f !== 0) {
    writer.writeInt64(
      10,
      f
    );
  }
  f = message.getDeleted();
  if (f !== 0) {
    writer.writeInt64(
      11,
      f
    );
  }
};


/**
 * optional string bucket = 1;
 * @return {string}
 */
proto.blob.Object.prototype.getBucket = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.blob.Object.prototype.setBucket = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.blob.Object.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.blob.Object.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string content_type = 3;
 * @return {string}
 */
proto.blob.Object.prototype.getContentType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.blob.Object.prototype.setContentType = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string content_language = 4;
 * @return {string}
 */
proto.blob.Object.prototype.getContentLanguage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.blob.Object.prototype.setContentLanguage = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string cache_control = 5;
 * @return {string}
 */
proto.blob.Object.prototype.getCacheControl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.blob.Object.prototype.setCacheControl = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string owner = 6;
 * @return {string}
 */
proto.blob.Object.prototype.getOwner = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.blob.Object.prototype.setOwner = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional int64 size = 7;
 * @return {number}
 */
proto.blob.Object.prototype.getSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.blob.Object.prototype.setSize = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * map<string, string> meta = 8;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.blob.Object.prototype.getMetaMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 8, opt_noLazyCreate,
      null));
};


proto.blob.Object.prototype.clearMetaMap = function() {
  this.getMetaMap().clear();
};


/**
 * optional int64 created = 9;
 * @return {number}
 */
proto.blob.Object.prototype.getCreated = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.blob.Object.prototype.setCreated = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int64 updated = 10;
 * @return {number}
 */
proto.blob.Object.prototype.getUpdated = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.blob.Object.prototype.setUpdated = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int64 deleted = 11;
 * @return {number}
 */
proto.blob.Object.prototype.getDeleted = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.blob.Object.prototype.setDeleted = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
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
proto.blob.Bucket = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.blob.Bucket.repeatedFields_, null);
};
goog.inherits(proto.blob.Bucket, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.blob.Bucket.displayName = 'proto.blob.Bucket';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.blob.Bucket.repeatedFields_ = [7];



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
proto.blob.Bucket.prototype.toObject = function(opt_includeInstance) {
  return proto.blob.Bucket.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.blob.Bucket} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blob.Bucket.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    location: jspb.Message.getFieldWithDefault(msg, 2, ""),
    encryptionKey: jspb.Message.getFieldWithDefault(msg, 3, ""),
    versioning: jspb.Message.getFieldWithDefault(msg, 4, false),
    pb_class: jspb.Message.getFieldWithDefault(msg, 5, ""),
    cors: (f = msg.getCors()) && auth_auth_pb.CORS.toObject(includeInstance, f),
    rulesList: jspb.Message.toObjectList(msg.getRulesList(),
    auth_auth_pb.ACLRule.toObject, includeInstance),
    metaMap: (f = msg.getMetaMap()) ? f.toObject(includeInstance, undefined) : [],
    created: jspb.Message.getFieldWithDefault(msg, 9, 0),
    updated: jspb.Message.getFieldWithDefault(msg, 10, 0),
    deleted: jspb.Message.getFieldWithDefault(msg, 11, 0)
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
 * @return {!proto.blob.Bucket}
 */
proto.blob.Bucket.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.blob.Bucket;
  return proto.blob.Bucket.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.blob.Bucket} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.blob.Bucket}
 */
proto.blob.Bucket.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLocation(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setEncryptionKey(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setVersioning(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setClass(value);
      break;
    case 6:
      var value = new auth_auth_pb.CORS;
      reader.readMessage(value,auth_auth_pb.CORS.deserializeBinaryFromReader);
      msg.setCors(value);
      break;
    case 7:
      var value = new auth_auth_pb.ACLRule;
      reader.readMessage(value,auth_auth_pb.ACLRule.deserializeBinaryFromReader);
      msg.addRules(value);
      break;
    case 8:
      var value = msg.getMetaMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "");
         });
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCreated(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUpdated(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setDeleted(value);
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
proto.blob.Bucket.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.blob.Bucket.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.blob.Bucket} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.blob.Bucket.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getLocation();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getEncryptionKey();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVersioning();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getClass();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getCors();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      auth_auth_pb.CORS.serializeBinaryToWriter
    );
  }
  f = message.getRulesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      7,
      f,
      auth_auth_pb.ACLRule.serializeBinaryToWriter
    );
  }
  f = message.getMetaMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(8, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreated();
  if (f !== 0) {
    writer.writeInt64(
      9,
      f
    );
  }
  f = message.getUpdated();
  if (f !== 0) {
    writer.writeInt64(
      10,
      f
    );
  }
  f = message.getDeleted();
  if (f !== 0) {
    writer.writeInt64(
      11,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.blob.Bucket.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.blob.Bucket.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string location = 2;
 * @return {string}
 */
proto.blob.Bucket.prototype.getLocation = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.blob.Bucket.prototype.setLocation = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string encryption_key = 3;
 * @return {string}
 */
proto.blob.Bucket.prototype.getEncryptionKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.blob.Bucket.prototype.setEncryptionKey = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional bool versioning = 4;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.blob.Bucket.prototype.getVersioning = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 4, false));
};


/** @param {boolean} value */
proto.blob.Bucket.prototype.setVersioning = function(value) {
  jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional string class = 5;
 * @return {string}
 */
proto.blob.Bucket.prototype.getClass = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.blob.Bucket.prototype.setClass = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional auth.CORS cors = 6;
 * @return {?proto.auth.CORS}
 */
proto.blob.Bucket.prototype.getCors = function() {
  return /** @type{?proto.auth.CORS} */ (
    jspb.Message.getWrapperField(this, auth_auth_pb.CORS, 6));
};


/** @param {?proto.auth.CORS|undefined} value */
proto.blob.Bucket.prototype.setCors = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.blob.Bucket.prototype.clearCors = function() {
  this.setCors(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.blob.Bucket.prototype.hasCors = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * repeated auth.ACLRule rules = 7;
 * @return {!Array<!proto.auth.ACLRule>}
 */
proto.blob.Bucket.prototype.getRulesList = function() {
  return /** @type{!Array<!proto.auth.ACLRule>} */ (
    jspb.Message.getRepeatedWrapperField(this, auth_auth_pb.ACLRule, 7));
};


/** @param {!Array<!proto.auth.ACLRule>} value */
proto.blob.Bucket.prototype.setRulesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 7, value);
};


/**
 * @param {!proto.auth.ACLRule=} opt_value
 * @param {number=} opt_index
 * @return {!proto.auth.ACLRule}
 */
proto.blob.Bucket.prototype.addRules = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 7, opt_value, proto.auth.ACLRule, opt_index);
};


proto.blob.Bucket.prototype.clearRulesList = function() {
  this.setRulesList([]);
};


/**
 * map<string, string> meta = 8;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.blob.Bucket.prototype.getMetaMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 8, opt_noLazyCreate,
      null));
};


proto.blob.Bucket.prototype.clearMetaMap = function() {
  this.getMetaMap().clear();
};


/**
 * optional int64 created = 9;
 * @return {number}
 */
proto.blob.Bucket.prototype.getCreated = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.blob.Bucket.prototype.setCreated = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int64 updated = 10;
 * @return {number}
 */
proto.blob.Bucket.prototype.getUpdated = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.blob.Bucket.prototype.setUpdated = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int64 deleted = 11;
 * @return {number}
 */
proto.blob.Bucket.prototype.getDeleted = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {number} value */
proto.blob.Bucket.prototype.setDeleted = function(value) {
  jspb.Message.setProto3IntField(this, 11, value);
};


goog.object.extend(exports, proto.blob);
