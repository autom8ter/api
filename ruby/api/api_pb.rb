# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
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
  add_message "api.AccessToken" do
    optional :token, :string, 1
  end
  add_message "api.IDToken" do
    optional :token, :string, 1
  end
  add_message "api.RefreshToken" do
    optional :token, :string, 1
  end
  add_message "api.Tokens" do
    optional :id, :message, 1, "api.IDToken"
    optional :access, :message, 2, "api.AccessToken"
    optional :refresh, :message, 3, "api.RefreshToken"
  end
  add_message "api.Paths" do
    optional :home, :string, 1
    optional :dashboard, :string, 2
    optional :settings, :string, 3
    optional :logout, :string, 4
    optional :callback, :string, 5
    optional :login, :string, 6
    optional :subscribe, :string, 7
    optional :unsubscribe, :string, 8
    optional :faq, :string, 9
    optional :support, :string, 10
    optional :terms, :string, 11
    optional :privacy, :string, 12
    optional :debug, :string, 13
    optional :blog, :string, 14
  end
  add_message "api.CreateCustomerRequest" do
    optional :user_info, :message, 1, "api.UserInfo"
  end
  add_message "api.CreateCustomerResponse" do
    optional :id, :string, 1
  end
end

module Api
  UserInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserInfo").msgclass
  UserMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserMetadata").msgclass
  AppMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AppMetadata").msgclass
  AccessToken = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AccessToken").msgclass
  IDToken = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.IDToken").msgclass
  RefreshToken = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.RefreshToken").msgclass
  Tokens = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Tokens").msgclass
  Paths = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Paths").msgclass
  CreateCustomerRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CreateCustomerRequest").msgclass
  CreateCustomerResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CreateCustomerResponse").msgclass
end
