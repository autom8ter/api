# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "api.EchoMessage" do
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
  end
  add_message "api.AppMetadata" do
    optional :plan, :string, 1
    optional :pay_token, :string, 2
  end
  add_message "api.Auth" do
    optional :domain, :string, 1
    optional :client_id, :string, 2
    optional :client_secret, :string, 3
    optional :redirect, :string, 4
    optional :audience, :string, 5
    repeated :scopes, :string, 6
  end
  add_message "api.Template" do
    optional :name, :string, 1
    optional :text, :string, 2
  end
  add_message "api.Bytes" do
    optional :bits, :bytes, 1
  end
end

module Api
  EchoMessage = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EchoMessage").msgclass
  UserInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserInfo").msgclass
  UserMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserMetadata").msgclass
  AppMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AppMetadata").msgclass
  Auth = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Auth").msgclass
  Template = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Template").msgclass
  Bytes = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Bytes").msgclass
end
