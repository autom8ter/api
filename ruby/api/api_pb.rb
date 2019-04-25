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
  add_message "api.Secret" do
    optional :text, :string, 1
  end
  add_message "api.Empty" do
  end
  add_message "api.Dashboard" do
    optional :users, :message, 1, "api.UsersWidget"
    optional :customers, :message, 2, "api.CustomersWidget"
    optional :plans, :message, 3, "api.PlansWidget"
    optional :subscriptions, :message, 4, "api.SubscriptionsWidget"
    optional :charges, :message, 5, "api.ChargesWidget"
  end
  add_message "api.CustomersWidget" do
    optional :count, :int64, 1
  end
  add_message "api.PlansWidget" do
    optional :count, :int64, 1
  end
  add_message "api.SubscriptionsWidget" do
    optional :count, :int64, 1
  end
  add_message "api.ChargesWidget" do
    optional :count, :int64, 1
    optional :total, :double, 2
    optional :dollars_per_charge, :double, 3
  end
  add_message "api.UsersWidget" do
    optional :count, :int64, 1
  end
  add_message "api.Identifier" do
    optional :id, :string, 1
  end
  add_message "api.SMSStatus" do
    optional :id, :message, 1, "api.Identifier"
    optional :sms, :message, 2, "api.SMS"
    optional :status, :string, 3
    optional :uri, :string, 4
  end
  add_message "api.SMS" do
    optional :service, :string, 1
    optional :to, :string, 2
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
  add_message "api.Message" do
    optional :value, :string, 1
  end
  add_message "api.UserInfo" do
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
    repeated :scopes, :string, 6
  end
  add_message "api.Bytes" do
    optional :bits, :bytes, 1
  end
  add_message "api.Template" do
    optional :text, :string, 1
    optional :data, :bytes, 2
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
end

module Api
  SubscribeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscribeRequest").msgclass
  UnSubscribeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UnSubscribeRequest").msgclass
  Card = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Card").msgclass
  Secret = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Secret").msgclass
  Empty = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Empty").msgclass
  Dashboard = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Dashboard").msgclass
  CustomersWidget = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CustomersWidget").msgclass
  PlansWidget = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.PlansWidget").msgclass
  SubscriptionsWidget = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscriptionsWidget").msgclass
  ChargesWidget = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ChargesWidget").msgclass
  UsersWidget = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UsersWidget").msgclass
  Identifier = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Identifier").msgclass
  SMSStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMSStatus").msgclass
  SMS = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMS").msgclass
  EmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailRequest").msgclass
  Email = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Email").msgclass
  Call = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Call").msgclass
  Message = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Message").msgclass
  UserInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserInfo").msgclass
  Identity = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Identity").msgclass
  UserMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserMetadata").msgclass
  AppMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AppMetadata").msgclass
  Auth = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Auth").msgclass
  Bytes = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Bytes").msgclass
  Template = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Template").msgclass
  JSONWebKeys = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.JSONWebKeys").msgclass
  Jwks = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Jwks").msgclass
end
