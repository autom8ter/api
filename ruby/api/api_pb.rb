# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
require 'common/common_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "api.FaxRequest" do
    optional :to, :message, 1, "common.String"
    optional :from, :message, 2, "common.String"
    optional :media_url, :message, 3, "common.String"
    optional :quality, :message, 4, "common.String"
    optional :callback, :message, 5, "common.String"
    optional :store_media, :bool, 6
  end
  add_message "api.SubscribeRequest" do
    optional :email, :message, 1, "common.String"
    optional :plan, :enum, 2, "api.Plan"
    optional :card, :message, 3, "api.Card"
  end
  add_message "api.UnSubscribeRequest" do
    optional :email, :message, 1, "common.String"
    optional :plan, :enum, 2, "api.Plan"
  end
  add_message "api.Card" do
    optional :number, :message, 1, "common.String"
    optional :exp_month, :message, 2, "common.String"
    optional :exp_year, :message, 3, "common.String"
    optional :cvc, :message, 4, "common.String"
  end
  add_message "api.SMS" do
    optional :service, :message, 1, "common.String"
    optional :to, :message, 2, "common.String"
    optional :message, :message, 3, "common.String"
    optional :mediaURL, :message, 4, "common.String"
    optional :callback, :message, 5, "common.String"
    optional :app, :message, 6, "common.String"
  end
  add_message "api.SMSBlast" do
    optional :service, :message, 1, "common.String"
    optional :to, :message, 2, "common.StringArray"
    optional :message, :message, 3, "common.String"
    optional :mediaURL, :message, 4, "common.String"
    optional :callback, :message, 5, "common.String"
    optional :app, :message, 6, "common.String"
  end
  add_message "api.EmailRequest" do
    optional :from_name, :message, 1, "common.String"
    optional :from_email, :message, 2, "common.String"
    optional :email, :message, 3, "api.Email"
  end
  add_message "api.EmailBlastRequest" do
    optional :from_name, :message, 1, "common.String"
    optional :from_email, :message, 2, "common.String"
    optional :blast, :message, 3, "api.EmailBlast"
  end
  add_message "api.EmailBlast" do
    optional :name_address, :message, 1, "common.StringMap"
    optional :subject, :message, 2, "common.String"
    optional :plain, :message, 3, "common.String"
    optional :html, :message, 4, "common.String"
  end
  add_message "api.Email" do
    optional :name, :message, 1, "common.String"
    optional :address, :message, 2, "common.String"
    optional :subject, :message, 3, "common.String"
    optional :plain, :message, 4, "common.String"
    optional :html, :message, 5, "common.String"
  end
  add_message "api.Call" do
    optional :from, :message, 1, "common.String"
    optional :to, :message, 2, "common.String"
    optional :app, :message, 3, "common.String"
  end
  add_message "api.CallBlast" do
    optional :from, :message, 1, "common.String"
    optional :to, :message, 2, "common.StringArray"
    optional :app, :message, 3, "common.String"
  end
  add_message "api.User" do
    optional :user_id, :message, 1, "common.Identifier"
    optional :name, :message, 2, "common.String"
    optional :given_name, :message, 3, "common.String"
    optional :family_name, :message, 4, "common.String"
    optional :gender, :message, 5, "common.String"
    optional :birthdate, :message, 6, "common.String"
    optional :email, :message, 7, "common.Identifier"
    optional :phone_number, :message, 8, "common.Identifier"
    optional :picture, :message, 9, "common.String"
    optional :user_metadata, :message, 10, "common.StringMap"
    optional :app_metadata, :message, 11, "common.StringMap"
    optional :last_ip, :message, 12, "common.String"
    optional :blocked, :bool, 13
    optional :nickname, :message, 14, "common.String"
    optional :multifactor, :message, 15, "common.StringArray"
    optional :created_at, :message, 17, "common.String"
    optional :updated_at, :message, 18, "common.String"
    optional :phone_verified, :bool, 19
    optional :email_verified, :bool, 20
    optional :password, :message, 21, "common.String"
    repeated :identities, :message, 22, "api.Identity"
  end
  add_message "api.Identity" do
    optional :connection, :message, 1, "common.String"
    optional :user_id, :message, 2, "common.Identifier"
    optional :provider, :message, 3, "common.String"
    optional :isSocial, :bool, 4
  end
  add_message "api.Auth" do
    optional :domain, :message, 1, "common.String"
    optional :client_id, :message, 2, "common.String"
    optional :client_secret, :message, 3, "common.String"
    optional :redirect, :message, 4, "common.String"
    repeated :scopes, :enum, 5, "api.Scope"
  end
  add_message "api.JSONWebKeys" do
    optional :kty, :message, 1, "common.String"
    optional :kid, :message, 2, "common.Identifier"
    optional :use, :message, 3, "common.String"
    optional :n, :message, 4, "common.String"
    optional :e, :message, 5, "common.String"
    optional :x5c, :message, 6, "common.StringArray"
  end
  add_message "api.Jwks" do
    repeated :keys, :message, 1, "api.JSONWebKeys"
  end
  add_message "api.RenderRequest" do
    optional :name, :message, 1, "common.String"
    optional :text, :message, 2, "common.String"
    optional :data, :message, 3, "common.Bytes"
  end
  add_message "api.SearchPhoneNumberRequest" do
    optional :state, :message, 1, "common.String"
    optional :capabilities, :message, 2, "api.NumberCapabilities"
    optional :total_results, :message, 3, "common.Int64"
  end
  add_message "api.PhoneNumber" do
    optional :friendly_name, :message, 1, "common.String"
    optional :phone_number, :message, 2, "common.String"
    optional :region, :message, 3, "common.String"
    optional :capabilities, :message, 4, "api.NumberCapabilities"
  end
  add_message "api.NumberCapabilities" do
    optional :voice, :bool, 1
    optional :sms, :bool, 2
    optional :mms, :bool, 3
  end
  add_message "api.PhoneNumberResource" do
    optional :number, :message, 1, "api.PhoneNumber"
    optional :id, :message, 2, "common.Identifier"
    optional :uri, :message, 3, "common.String"
  end
  add_message "api.TokenQuery" do
    optional :token, :message, 1, "common.Token"
    optional :query, :message, 2, "common.String"
  end
  add_message "api.IDBody" do
    optional :id, :message, 1, "common.Identifier"
    optional :body, :message, 2, "common.Bytes"
  end
  add_message "api.IDStrings" do
    optional :id, :message, 1, "common.Identifier"
    optional :strings, :message, 2, "common.StringArray"
  end
  add_message "api.Role" do
    optional :id, :message, 1, "common.Identifier"
    optional :name, :message, 2, "common.String"
    optional :description, :message, 3, "common.String"
  end
  add_message "api.CallResponse" do
    optional :id, :message, 1, "common.Identifier"
    optional :to, :message, 5, "common.String"
    optional :from, :message, 6, "common.String"
    optional :status, :message, 9, "common.String"
    optional :answered_by, :message, 10, "common.String"
    optional :forwarded_from, :message, 11, "common.String"
    optional :caller_name, :message, 12, "common.String"
    optional :annotations, :message, 13, "common.StringMap"
  end
  add_message "api.SMSResponse" do
    optional :id, :message, 1, "common.Identifier"
    optional :to, :message, 5, "common.String"
    optional :from, :message, 6, "common.String"
    optional :media_url, :message, 7, "common.String"
    optional :body, :message, 8, "common.String"
    optional :status, :message, 9, "common.String"
    optional :annotations, :message, 10, "common.StringMap"
  end
  add_message "api.SubscriptionResponse" do
    optional :id, :message, 1, "common.Identifier"
    optional :monthly_charge, :message, 2, "common.Int64"
    optional :next_charge, :message, 3, "common.String"
    optional :annotations, :message, 10, "common.StringMap"
    optional :plan, :enum, 4, "api.Plan"
    optional :user, :message, 5, "api.User"
    optional :status, :message, 6, "common.String"
  end
  add_message "api.FaxResponse" do
    optional :id, :message, 1, "common.Identifier"
    optional :media_url, :message, 3, "common.String"
    optional :to, :message, 4, "common.String"
    optional :from, :message, 5, "common.String"
    optional :status, :message, 6, "common.String"
    optional :annotations, :message, 10, "common.StringMap"
  end
  add_enum "api.Scope" do
    value :OPENID, 0
    value :PROFILE, 1
    value :EMAIL, 2
    value :READ_USERS, 3
    value :READ_USER_IDP_TOKENS, 4
    value :CREATE_USERS, 5
    value :READ_STATS, 6
    value :READ_EMAIL_TEMPLATES, 7
    value :UPDATE_EMAIL_TEMPLATES, 8
    value :CREATE_EMAIL_TEMPLATES, 9
    value :READ_RULES, 10
    value :UPDATE_RULES, 11
    value :CREATE_RULES, 12
    value :DELETE_RULES, 13
    value :READ_ROLES, 14
    value :UPDATE_ROLES, 15
    value :CREATE_ROLES, 16
    value :DELETE_ROLES, 17
    value :READ_LOGS, 18
  end
  add_enum "api.URL" do
    value :USER_INFOURL, 0
    value :TOKENURL, 1
    value :AUTHORIZEURL, 2
    value :USERSURL, 3
    value :CLIENTSURL, 4
    value :GRANTSURL, 5
    value :RULESURL, 6
    value :ROLESURL, 7
    value :LOGSURL, 8
    value :STATSURL, 9
    value :CONNECTIONSURL, 10
    value :TENANTSURL, 11
    value :EMAIL_TEMPLATEURL, 12
    value :EMAILURL, 13
    value :SEARCH_USERSURL, 14
    value :DEVICEURL, 18
    value :JWKSURL, 19
    value :CLIENT_GRANTSURL, 20
  end
  add_enum "api.Plan" do
    value :FREE, 0
    value :BASIC, 1
    value :PREMIUM, 2
  end
end

module Api
  FaxRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.FaxRequest").msgclass
  SubscribeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscribeRequest").msgclass
  UnSubscribeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UnSubscribeRequest").msgclass
  Card = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Card").msgclass
  SMS = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMS").msgclass
  SMSBlast = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMSBlast").msgclass
  EmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailRequest").msgclass
  EmailBlastRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailBlastRequest").msgclass
  EmailBlast = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailBlast").msgclass
  Email = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Email").msgclass
  Call = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Call").msgclass
  CallBlast = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CallBlast").msgclass
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.User").msgclass
  Identity = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Identity").msgclass
  Auth = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Auth").msgclass
  JSONWebKeys = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.JSONWebKeys").msgclass
  Jwks = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Jwks").msgclass
  RenderRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.RenderRequest").msgclass
  SearchPhoneNumberRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SearchPhoneNumberRequest").msgclass
  PhoneNumber = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.PhoneNumber").msgclass
  NumberCapabilities = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.NumberCapabilities").msgclass
  PhoneNumberResource = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.PhoneNumberResource").msgclass
  TokenQuery = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.TokenQuery").msgclass
  IDBody = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.IDBody").msgclass
  IDStrings = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.IDStrings").msgclass
  Role = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Role").msgclass
  CallResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CallResponse").msgclass
  SMSResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMSResponse").msgclass
  SubscriptionResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscriptionResponse").msgclass
  FaxResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.FaxResponse").msgclass
  Scope = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Scope").enummodule
  URL = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.URL").enummodule
  Plan = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Plan").enummodule
end
