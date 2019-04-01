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

goog.exportSymbol('proto.sms.GetSMSRequest', null, global);
goog.exportSymbol('proto.sms.SMSResponse', null, global);
goog.exportSymbol('proto.sms.SendMMSRequest', null, global);
goog.exportSymbol('proto.sms.SendSMSRequest', null, global);
goog.exportSymbol('proto.sms.SendSMSWithCopilotRequest', null, global);

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
proto.sms.SendSMSRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.sms.SendSMSRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.sms.SendSMSRequest.displayName = 'proto.sms.SendSMSRequest';
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
proto.sms.SendSMSRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.sms.SendSMSRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.sms.SendSMSRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SendSMSRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    to: jspb.Message.getFieldWithDefault(msg, 2, ""),
    from: jspb.Message.getFieldWithDefault(msg, 3, ""),
    body: jspb.Message.getFieldWithDefault(msg, 4, ""),
    statusCallback: jspb.Message.getFieldWithDefault(msg, 5, ""),
    applicationSid: jspb.Message.getFieldWithDefault(msg, 6, "")
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
 * @return {!proto.sms.SendSMSRequest}
 */
proto.sms.SendSMSRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.sms.SendSMSRequest;
  return proto.sms.SendSMSRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.sms.SendSMSRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.sms.SendSMSRequest}
 */
proto.sms.SendSMSRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTo(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setFrom(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setBody(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setStatusCallback(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationSid(value);
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
proto.sms.SendSMSRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.sms.SendSMSRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.sms.SendSMSRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SendSMSRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTo();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getFrom();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getBody();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getStatusCallback();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getApplicationSid();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional string to = 2;
 * @return {string}
 */
proto.sms.SendSMSRequest.prototype.getTo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.sms.SendSMSRequest.prototype.setTo = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string from = 3;
 * @return {string}
 */
proto.sms.SendSMSRequest.prototype.getFrom = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.sms.SendSMSRequest.prototype.setFrom = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string body = 4;
 * @return {string}
 */
proto.sms.SendSMSRequest.prototype.getBody = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.sms.SendSMSRequest.prototype.setBody = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string status_callback = 5;
 * @return {string}
 */
proto.sms.SendSMSRequest.prototype.getStatusCallback = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.sms.SendSMSRequest.prototype.setStatusCallback = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string application_sid = 6;
 * @return {string}
 */
proto.sms.SendSMSRequest.prototype.getApplicationSid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.sms.SendSMSRequest.prototype.setApplicationSid = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
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
proto.sms.SendMMSRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.sms.SendMMSRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.sms.SendMMSRequest.displayName = 'proto.sms.SendMMSRequest';
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
proto.sms.SendMMSRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.sms.SendMMSRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.sms.SendMMSRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SendMMSRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    to: jspb.Message.getFieldWithDefault(msg, 2, ""),
    from: jspb.Message.getFieldWithDefault(msg, 3, ""),
    body: jspb.Message.getFieldWithDefault(msg, 4, ""),
    mediaUrl: jspb.Message.getFieldWithDefault(msg, 5, ""),
    statusCallback: jspb.Message.getFieldWithDefault(msg, 6, ""),
    applicationSid: jspb.Message.getFieldWithDefault(msg, 7, "")
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
 * @return {!proto.sms.SendMMSRequest}
 */
proto.sms.SendMMSRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.sms.SendMMSRequest;
  return proto.sms.SendMMSRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.sms.SendMMSRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.sms.SendMMSRequest}
 */
proto.sms.SendMMSRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTo(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setFrom(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setBody(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setMediaUrl(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setStatusCallback(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationSid(value);
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
proto.sms.SendMMSRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.sms.SendMMSRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.sms.SendMMSRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SendMMSRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTo();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getFrom();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getBody();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getMediaUrl();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getStatusCallback();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getApplicationSid();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
};


/**
 * optional string to = 2;
 * @return {string}
 */
proto.sms.SendMMSRequest.prototype.getTo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.sms.SendMMSRequest.prototype.setTo = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string from = 3;
 * @return {string}
 */
proto.sms.SendMMSRequest.prototype.getFrom = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.sms.SendMMSRequest.prototype.setFrom = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string body = 4;
 * @return {string}
 */
proto.sms.SendMMSRequest.prototype.getBody = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.sms.SendMMSRequest.prototype.setBody = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string media_url = 5;
 * @return {string}
 */
proto.sms.SendMMSRequest.prototype.getMediaUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.sms.SendMMSRequest.prototype.setMediaUrl = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string status_callback = 6;
 * @return {string}
 */
proto.sms.SendMMSRequest.prototype.getStatusCallback = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.sms.SendMMSRequest.prototype.setStatusCallback = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string application_sid = 7;
 * @return {string}
 */
proto.sms.SendMMSRequest.prototype.getApplicationSid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.sms.SendMMSRequest.prototype.setApplicationSid = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
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
proto.sms.SendSMSWithCopilotRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.sms.SendSMSWithCopilotRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.sms.SendSMSWithCopilotRequest.displayName = 'proto.sms.SendSMSWithCopilotRequest';
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
proto.sms.SendSMSWithCopilotRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.sms.SendSMSWithCopilotRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.sms.SendSMSWithCopilotRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SendSMSWithCopilotRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    messagingServiceSid: jspb.Message.getFieldWithDefault(msg, 1, ""),
    to: jspb.Message.getFieldWithDefault(msg, 2, ""),
    body: jspb.Message.getFieldWithDefault(msg, 3, ""),
    statusCallback: jspb.Message.getFieldWithDefault(msg, 4, ""),
    applicationSid: jspb.Message.getFieldWithDefault(msg, 5, "")
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
 * @return {!proto.sms.SendSMSWithCopilotRequest}
 */
proto.sms.SendSMSWithCopilotRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.sms.SendSMSWithCopilotRequest;
  return proto.sms.SendSMSWithCopilotRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.sms.SendSMSWithCopilotRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.sms.SendSMSWithCopilotRequest}
 */
proto.sms.SendSMSWithCopilotRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessagingServiceSid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTo(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setBody(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setStatusCallback(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationSid(value);
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
proto.sms.SendSMSWithCopilotRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.sms.SendSMSWithCopilotRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.sms.SendSMSWithCopilotRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SendSMSWithCopilotRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMessagingServiceSid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTo();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getBody();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getStatusCallback();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getApplicationSid();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
};


/**
 * optional string messaging_service_sid = 1;
 * @return {string}
 */
proto.sms.SendSMSWithCopilotRequest.prototype.getMessagingServiceSid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.sms.SendSMSWithCopilotRequest.prototype.setMessagingServiceSid = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string to = 2;
 * @return {string}
 */
proto.sms.SendSMSWithCopilotRequest.prototype.getTo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.sms.SendSMSWithCopilotRequest.prototype.setTo = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string body = 3;
 * @return {string}
 */
proto.sms.SendSMSWithCopilotRequest.prototype.getBody = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.sms.SendSMSWithCopilotRequest.prototype.setBody = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string status_callback = 4;
 * @return {string}
 */
proto.sms.SendSMSWithCopilotRequest.prototype.getStatusCallback = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.sms.SendSMSWithCopilotRequest.prototype.setStatusCallback = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string application_sid = 5;
 * @return {string}
 */
proto.sms.SendSMSWithCopilotRequest.prototype.getApplicationSid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.sms.SendSMSWithCopilotRequest.prototype.setApplicationSid = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
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
proto.sms.GetSMSRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.sms.GetSMSRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.sms.GetSMSRequest.displayName = 'proto.sms.GetSMSRequest';
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
proto.sms.GetSMSRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.sms.GetSMSRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.sms.GetSMSRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.GetSMSRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    sid: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.sms.GetSMSRequest}
 */
proto.sms.GetSMSRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.sms.GetSMSRequest;
  return proto.sms.GetSMSRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.sms.GetSMSRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.sms.GetSMSRequest}
 */
proto.sms.GetSMSRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setSid(value);
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
proto.sms.GetSMSRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.sms.GetSMSRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.sms.GetSMSRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.GetSMSRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string sid = 1;
 * @return {string}
 */
proto.sms.GetSMSRequest.prototype.getSid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.sms.GetSMSRequest.prototype.setSid = function(value) {
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
proto.sms.SMSResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.sms.SMSResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.sms.SMSResponse.displayName = 'proto.sms.SMSResponse';
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
proto.sms.SMSResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.sms.SMSResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.sms.SMSResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SMSResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    sid: jspb.Message.getFieldWithDefault(msg, 1, ""),
    dateCreated: jspb.Message.getFieldWithDefault(msg, 2, ""),
    dateUpdated: jspb.Message.getFieldWithDefault(msg, 3, ""),
    dateSent: jspb.Message.getFieldWithDefault(msg, 4, ""),
    accoundSid: jspb.Message.getFieldWithDefault(msg, 5, ""),
    to: jspb.Message.getFieldWithDefault(msg, 6, ""),
    from: jspb.Message.getFieldWithDefault(msg, 7, ""),
    mediaUrl: jspb.Message.getFieldWithDefault(msg, 8, ""),
    body: jspb.Message.getFieldWithDefault(msg, 10, ""),
    status: jspb.Message.getFieldWithDefault(msg, 11, ""),
    direction: jspb.Message.getFieldWithDefault(msg, 12, ""),
    apiVersion: jspb.Message.getFieldWithDefault(msg, 13, ""),
    price: jspb.Message.getFieldWithDefault(msg, 17, ""),
    uri: jspb.Message.getFieldWithDefault(msg, 20, "")
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
 * @return {!proto.sms.SMSResponse}
 */
proto.sms.SMSResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.sms.SMSResponse;
  return proto.sms.SMSResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.sms.SMSResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.sms.SMSResponse}
 */
proto.sms.SMSResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setSid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setDateCreated(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDateUpdated(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setDateSent(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccoundSid(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setTo(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setFrom(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setMediaUrl(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setBody(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setStatus(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setDirection(value);
      break;
    case 13:
      var value = /** @type {string} */ (reader.readString());
      msg.setApiVersion(value);
      break;
    case 17:
      var value = /** @type {string} */ (reader.readString());
      msg.setPrice(value);
      break;
    case 20:
      var value = /** @type {string} */ (reader.readString());
      msg.setUri(value);
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
proto.sms.SMSResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.sms.SMSResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.sms.SMSResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.sms.SMSResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getDateCreated();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDateUpdated();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getDateSent();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getAccoundSid();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getTo();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getFrom();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getMediaUrl();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getBody();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getStatus();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getDirection();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getApiVersion();
  if (f.length > 0) {
    writer.writeString(
      13,
      f
    );
  }
  f = message.getPrice();
  if (f.length > 0) {
    writer.writeString(
      17,
      f
    );
  }
  f = message.getUri();
  if (f.length > 0) {
    writer.writeString(
      20,
      f
    );
  }
};


/**
 * optional string sid = 1;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getSid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setSid = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string date_created = 2;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getDateCreated = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setDateCreated = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string date_updated = 3;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getDateUpdated = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setDateUpdated = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string date_sent = 4;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getDateSent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setDateSent = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string accound_sid = 5;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getAccoundSid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setAccoundSid = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string to = 6;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getTo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setTo = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string from = 7;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getFrom = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setFrom = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string media_url = 8;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getMediaUrl = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setMediaUrl = function(value) {
  jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string body = 10;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getBody = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setBody = function(value) {
  jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string status = 11;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getStatus = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setStatus = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional string direction = 12;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getDirection = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setDirection = function(value) {
  jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional string api_version = 13;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getApiVersion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 13, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setApiVersion = function(value) {
  jspb.Message.setProto3StringField(this, 13, value);
};


/**
 * optional string price = 17;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getPrice = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 17, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setPrice = function(value) {
  jspb.Message.setProto3StringField(this, 17, value);
};


/**
 * optional string uri = 20;
 * @return {string}
 */
proto.sms.SMSResponse.prototype.getUri = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 20, ""));
};


/** @param {string} value */
proto.sms.SMSResponse.prototype.setUri = function(value) {
  jspb.Message.setProto3StringField(this, 20, value);
};


goog.object.extend(exports, proto.sms);
