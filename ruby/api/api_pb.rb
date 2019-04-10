# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/protobuf/empty_pb'
require 'google/protobuf/duration_pb'
require 'google/protobuf/field_mask_pb'
require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "api.Empty" do
  end
  add_message "api.UserMap" do
    map :users, :string, :message, 1, "api.User"
  end
  add_message "api.User" do
    optional :user_id, :string, 1
    optional :plan, :string, 2
    optional :name, :string, 3
    optional :email, :string, 4
    optional :description, :string, 5
    optional :phone, :string, 6
    optional :address, :message, 8, "api.Address"
    map :metadata, :string, :string, 9
    optional :deleted, :bool, 10
    optional :create_date, :int64, 20
  end
  add_message "api.AddUserRequest" do
    optional :email, :string, 1
    optional :plan, :string, 2
    optional :phone, :string, 3
    optional :name, :string, 4
    optional :password, :string, 5
    optional :trial_end, :int64, 6
    optional :description, :string, 7
    optional :address, :message, 8, "api.Address"
  end
  add_message "api.SubscribeUserRequest" do
    optional :email, :string, 1
    optional :plan, :string, 2
    optional :card_number, :string, 3
    optional :exp_month, :string, 4
    optional :exp_year, :string, 5
    optional :cvc, :string, 6
  end
  add_message "api.AddUserMetadataRequest" do
    optional :user_id, :string, 1
    map :metadata, :string, :string, 2
  end
  add_message "api.Address" do
    optional :city, :string, 1
    optional :country, :string, 2
    optional :line1, :string, 3
    optional :line2, :string, 4
    optional :postal_code, :string, 5
    optional :state, :string, 6
  end
  add_message "api.SubscribeUserResponse" do
    optional :subscription_id, :string, 1
  end
  add_message "api.CreatePlanResponse" do
    optional :plan_id, :string, 1
  end
  add_message "api.CancelSubscriptionRequest" do
    optional :email, :string, 1
  end
  add_message "api.CreatePlanRequest" do
    optional :plan_id, :string, 1
    optional :amount, :int64, 2
    optional :service_id, :string, 3
    optional :service_name, :string, 4
    optional :friendly_name, :string, 5
  end
  add_message "api.SMSRequest" do
    optional :user_id, :string, 1
    optional :body, :string, 2
  end
  add_message "api.CallRequest" do
    optional :user_id, :string, 1
    optional :callback_url, :string, 2
  end
  add_message "api.MMSRequest" do
    optional :user_id, :string, 1
    optional :body, :string, 2
    optional :media_url, :string, 3
  end
  add_message "api.EmailRequest" do
    optional :user_id, :string, 1
    optional :subject, :string, 2
    optional :plain_text, :string, 3
    optional :html_alt, :string, 4
  end
  add_message "api.ChannelReminder" do
    optional :channel_id, :string, 1
    optional :text, :string, 2
    optional :time, :string, 3
  end
  add_message "api.UserReminder" do
    optional :user_id, :string, 1
    optional :text, :string, 2
    optional :time, :string, 3
    optional :item, :message, 4, "api.ItemRef"
  end
  add_message "api.ItemRef" do
    optional :channel, :string, 1
    optional :file, :string, 2
    optional :comment, :string, 3
  end
  add_message "api.Star" do
    optional :text, :string, 1
    optional :item, :message, 4, "api.ItemRef"
  end
end

module Api
  Empty = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Empty").msgclass
  UserMap = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserMap").msgclass
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.User").msgclass
  AddUserRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AddUserRequest").msgclass
  SubscribeUserRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscribeUserRequest").msgclass
  AddUserMetadataRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AddUserMetadataRequest").msgclass
  Address = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Address").msgclass
  SubscribeUserResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscribeUserResponse").msgclass
  CreatePlanResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CreatePlanResponse").msgclass
  CancelSubscriptionRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CancelSubscriptionRequest").msgclass
  CreatePlanRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CreatePlanRequest").msgclass
  SMSRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMSRequest").msgclass
  CallRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CallRequest").msgclass
  MMSRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.MMSRequest").msgclass
  EmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailRequest").msgclass
  ChannelReminder = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ChannelReminder").msgclass
  UserReminder = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserReminder").msgclass
  ItemRef = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ItemRef").msgclass
  Star = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Star").msgclass
end
