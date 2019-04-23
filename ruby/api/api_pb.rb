# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "api.GetByEmail" do
    optional :email, :string, 1
  end
  add_message "api.IDToken" do
    optional :iss, :string, 1
    optional :sub, :bool, 2
    optional :aud, :string, 3
    optional :exp, :int64, 4
    optional :iat, :int64, 5
    optional :name, :string, 6
    optional :given_name, :string, 7
    optional :family_name, :string, 8
    optional :gender, :string, 9
    optional :birthdate, :string, 10
    optional :email, :string, 11
    optional :picture, :int64, 12
  end
  add_message "api.UserMetadata" do
    optional :phone, :string, 1
    optional :plan, :string, 2
    optional :pay_token, :string, 3
    optional :last_contact, :string, 4
  end
  add_message "api.AccessToken" do
    optional :iss, :string, 1
    optional :sub, :string, 2
    repeated :aud, :string, 3
    optional :azp, :string, 4
    optional :exp, :int64, 5
    optional :iat, :int64, 6
    optional :scope, :string, 7
  end
end

module Api
  GetByEmail = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.GetByEmail").msgclass
  IDToken = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.IDToken").msgclass
  UserMetadata = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserMetadata").msgclass
  AccessToken = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AccessToken").msgclass
end
