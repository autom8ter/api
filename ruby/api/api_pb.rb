# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "api.SubscribeRequest" do
    optional :email, :string, 1
    optional :plan, :string, 2
    optional :card, :message, 3, "api.Card"
  end
  add_message "api.UnSubscribeRequest" do
    optional :email, :string, 1
    optional :plan, :string, 2
  end
  add_message "api.Card" do
    optional :number, :string, 1
    optional :exp_month, :string, 2
    optional :exp_year, :string, 3
    optional :cvc, :string, 4
  end
  add_message "api.Empty" do
  end
  add_message "api.ManagementToken" do
    optional :token, :string, 1
  end
  add_message "api.UserRequest" do
    optional :string, :message, 1, "api.ManagementToken"
    optional :user, :message, 2, "api.User"
  end
  add_message "api.UserByEmailRequest" do
    optional :token, :message, 1, "api.ManagementToken"
    optional :email, :string, 2
  end
  add_message "api.Identifier" do
    optional :id, :string, 1
  end
  add_message "api.SMS" do
    optional :service, :string, 1
    optional :to, :string, 2
    optional :message, :message, 3, "api.Message"
    optional :mediaURL, :string, 4
    optional :callback, :string, 5
    optional :app, :string, 6
  end
  add_message "api.SMSBlast" do
    optional :service, :string, 1
    repeated :to, :string, 2
    optional :message, :message, 3, "api.Message"
    optional :mediaURL, :string, 4
    optional :callback, :string, 5
    optional :app, :string, 6
  end
  add_message "api.EmailRequest" do
    optional :from_name, :string, 1
    optional :from_email, :string, 2
    optional :email, :message, 3, "api.Email"
  end
  add_message "api.EmailBlastRequest" do
    optional :from_name, :string, 1
    optional :from_email, :string, 2
    optional :blast, :message, 3, "api.EmailBlast"
  end
  add_message "api.EmailBlast" do
    map :name_address, :string, :string, 1
    optional :subject, :string, 2
    optional :plain, :string, 3
    optional :html, :string, 4
  end
  add_message "api.Email" do
    optional :name, :string, 1
    optional :address, :string, 2
    optional :subject, :string, 3
    optional :plain, :string, 4
    optional :html, :string, 5
  end
  add_message "api.Call" do
    optional :from, :string, 1
    optional :to, :string, 2
    optional :app, :string, 3
  end
  add_message "api.CallBlast" do
    optional :from, :string, 1
    repeated :to, :string, 2
    optional :app, :string, 3
  end
  add_message "api.Message" do
    optional :value, :string, 1
  end
  add_message "api.User" do
    optional :user_id, :string, 1
    optional :name, :string, 2
    optional :given_name, :string, 3
    optional :family_name, :string, 4
    optional :gender, :string, 5
    optional :birthdate, :string, 6
    optional :email, :string, 7
    optional :phone_number, :string, 8
    optional :picture, :string, 9
    optional :user_metadata, :message, 10, "api.UserMetadata"
    optional :app_metadata, :message, 11, "api.AppMetadata"
    optional :last_ip, :string, 12
    optional :blocked, :bool, 13
    optional :nickname, :string, 14
    repeated :multifactor, :string, 15
    optional :created_at, :string, 17
    optional :updated_at, :string, 18
    optional :phone_verified, :bool, 19
    repeated :identities, :message, 20, "api.Identity"
  end
  add_message "api.Identity" do
    optional :connection, :string, 1
    optional :user_id, :string, 2
    optional :provider, :string, 3
    optional :isSocial, :string, 4
  end
  add_message "api.UserMetadata" do
    map :metadata, :string, :string, 1
  end
  add_message "api.AppMetadata" do
    map :metadata, :string, :string, 1
  end
  add_message "api.Auth" do
    optional :domain, :string, 1
    optional :client_id, :string, 2
    optional :client_secret, :string, 3
    optional :redirect, :string, 4
    optional :audience, :string, 5
    repeated :scopes, :enum, 6, "api.Scope"
  end
  add_message "api.Bytes" do
    optional :bits, :bytes, 1
  end
  add_message "api.Template" do
    optional :name, :string, 1
    optional :text, :string, 2
  end
  add_message "api.JSONWebKeys" do
    optional :kty, :string, 1
    optional :kid, :string, 2
    optional :use, :string, 3
    optional :n, :string, 4
    optional :e, :string, 5
    repeated :x5c, :string, 6
  end
  add_message "api.Jwks" do
    repeated :keys, :message, 1, "api.JSONWebKeys"
  end
  add_message "api.HTTPRequest" do
    optional :method, :enum, 1, "api.HTTPMethod"
    optional :url, :string, 2
    optional :token, :string, 3
    optional :account, :string, 4
    optional :contentType, :string, 5
    map :headers, :string, :string, 6
    map :form, :string, :string, 7
    map :cookies, :string, :string, 8
    optional :body, :message, 9, "api.Bytes"
  end
  add_message "api.RenderRequest" do
    optional :template, :message, 1, "api.Template"
    optional :data, :message, 2, "api.Bytes"
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
  add_enum "api.HTTPMethod" do
    value :GET, 0
    value :POST, 1
  end
end

module Api
  SubscribeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscribeRequest").msgclass
  UnSubscribeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UnSubscribeRequest").msgclass
  Card = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Card").msgclass
  Empty = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Empty").msgclass
  ManagementToken = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ManagementToken").msgclass
  UserRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserRequest").msgclass
  UserByEmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserByEmailRequest").msgclass
  Identifier = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Identifier").msgclass
  SMS = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMS").msgclass
  SMSBlast = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMSBlast").msgclass
  EmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailRequest").msgclass
  EmailBlastRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailBlastRequest").msgclass
  EmailBlast = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailBlast").msgclass
  Email = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Email").msgclass
  Call = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Call").msgclass
  CallBlast = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CallBlast").msgclass
  Message = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Message").msgclass
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.User").msgclass
  Identity = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Identity").msgclass
  UserMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserMetadata").msgclass
  AppMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AppMetadata").msgclass
  Auth = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Auth").msgclass
  Bytes = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Bytes").msgclass
  Template = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Template").msgclass
  JSONWebKeys = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.JSONWebKeys").msgclass
  Jwks = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Jwks").msgclass
  HTTPRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.HTTPRequest").msgclass
  RenderRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.RenderRequest").msgclass
  Scope = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Scope").enummodule
  HTTPMethod = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.HTTPMethod").enummodule
end
