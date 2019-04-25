# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
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
    optional :name, :string, 6
    optional :given_name, :string, 7
    optional :family_name, :string, 8
    optional :gender, :string, 9
    optional :birthdate, :string, 10
    optional :email, :string, 11
    optional :picture, :string, 12
    optional :user_metadata, :message, 13, "api.UserMetadata"
    optional :app_metadata, :message, 14, "api.AppMetadata"
  end
  add_message "api.UserMetadata" do
    optional :phone, :string, 1
    optional :preferred_contact, :string, 2
    optional :status, :string, 3
    repeated :tags, :string, 4
  end
  add_message "api.AppMetadata" do
    optional :plan, :string, 1
    optional :pay_token, :string, 2
    optional :delinquent, :string, 3
    repeated :tags, :string, 4
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
end

module Api
  Identifier = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Identifier").msgclass
  SMSStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMSStatus").msgclass
  SMS = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMS").msgclass
  EmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailRequest").msgclass
  Email = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Email").msgclass
  Call = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Call").msgclass
  Message = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Message").msgclass
  UserInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserInfo").msgclass
  UserMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserMetadata").msgclass
  AppMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AppMetadata").msgclass
  Auth = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Auth").msgclass
  Bytes = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Bytes").msgclass
  Template = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Template").msgclass
end
