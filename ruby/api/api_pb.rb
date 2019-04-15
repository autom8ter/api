# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/protobuf/empty_pb'
require 'google/protobuf/duration_pb'
require 'google/protobuf/field_mask_pb'
require 'google/api/annotations_pb'
require 'google/api/auth_pb'
require 'google/pubsub/v1/pubsub_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "api.String" do
    optional :text, :string, 1
  end
  add_message "api.Id" do
    optional :id, :string, 1
  end
  add_message "api.ActionHookRequest" do
    optional :attachment, :message, 1, "api.Attachment"
    optional :actions, :message, 2, "api.AttachmentAction"
  end
  add_message "api.MessageUserRequest" do
    optional :id, :string, 1
    optional :message, :string, 2
  end
  add_message "api.RefundRequest" do
    optional :id, :string, 1
    optional :reason, :string, 2
    optional :amount, :int64, 3
    optional :reverse_transfer, :bool, 4
    optional :status, :string, 5
  end
  add_message "api.ChargeRequest" do
    optional :product, :message, 1, "api.Product"
    optional :id, :string, 2
  end
  add_message "api.CancelSubscriptionRequest" do
    optional :id, :string, 1
  end
  add_message "api.CreatePlanRequest" do
    optional :plan_id, :string, 1
    optional :amount, :int64, 2
    optional :service_id, :string, 3
    optional :service_name, :string, 4
    optional :friendly_name, :string, 5
  end
  add_message "api.SMSRequest" do
    optional :id, :string, 1
    optional :body, :string, 2
  end
  add_message "api.CallRequest" do
    optional :id, :string, 1
    optional :callback_url, :string, 2
  end
  add_message "api.MMSRequest" do
    optional :sms, :message, 1, "api.SMSRequest"
    optional :media_url, :string, 3
  end
  add_message "api.EmailRequest" do
    optional :id, :string, 1
    optional :subject, :string, 2
    optional :plain_text, :string, 3
    optional :html_alt, :string, 4
  end
  add_message "api.CustomerRequest" do
    optional :email, :string, 1
    optional :plan, :string, 2
    optional :phone, :string, 3
    optional :name, :string, 4
    optional :description, :string, 7
    optional :address, :message, 8, "api.Address"
  end
  add_message "api.UpdateCustomerRequest" do
    optional :id, :string, 1
    optional :customer, :message, 2, "api.CustomerRequest"
  end
  add_message "api.SubscribeCustomerRequest" do
    optional :id, :string, 1
    optional :plan, :string, 2
    optional :card_number, :string, 3
    optional :exp_month, :string, 4
    optional :exp_year, :string, 5
    optional :cvc, :string, 6
  end
  add_message "api.CreateAccountRequest" do
    optional :customer, :message, 1, "api.CustomerRequest"
    optional :access, :message, 2, "api.Access"
  end
  add_message "api.Account" do
    optional :customer, :message, 1, "api.Customer"
    optional :access, :message, 2, "api.Access"
  end
  add_message "api.User" do
    optional :id, :string, 1
    optional :team_id, :string, 2
    optional :name, :string, 3
    optional :phone, :string, 4
    optional :profile, :message, 5, "api.Profile"
    optional :deleted, :bool, 6
    optional :admin, :bool, 7
    optional :ownder, :bool, 8
    optional :primary_owner, :bool, 9
    optional :restricted, :bool, 10
    optional :ultra_restricted, :bool, 11
    optional :stranger, :bool, 12
    optional :bot, :bool, 13
    optional :has2fa, :bool, 14
    optional :locale, :string, 15
  end
  add_message "api.Profile" do
    optional :avatar_hash, :string, 1
    optional :status, :string, 2
    optional :status_emoji, :string, 3
    optional :display_name, :string, 4
    optional :name, :string, 5
    optional :email, :string, 6
    repeated :image_urls, :string, 7
    optional :team, :string, 8
  end
  add_message "api.Empty" do
  end
  add_message "api.Customer" do
    optional :id, :string, 1
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
  add_message "api.Card" do
    optional :card_type, :enum, 1, "api.CardType"
    optional :card_number, :string, 3
    optional :exp_month, :string, 4
    optional :exp_year, :string, 5
    optional :cvc, :string, 6
    optional :debit, :bool, 7
  end
  add_message "api.BankAccount" do
    optional :account_number, :string, 1
    optional :routing_number, :string, 2
  end
  add_message "api.Address" do
    optional :city, :string, 1
    optional :country, :string, 2
    optional :line1, :string, 3
    optional :line2, :string, 4
    optional :postal_code, :string, 5
    optional :state, :string, 6
  end
  add_message "api.ChannelReminder" do
    optional :channel_id, :string, 1
    optional :text, :string, 2
    optional :time, :string, 3
  end
  add_message "api.UserReminder" do
    optional :id, :string, 1
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
  add_message "api.Pin" do
    optional :text, :string, 1
    optional :item, :message, 4, "api.ItemRef"
  end
  add_message "api.SignedKey" do
    optional :signed_key, :string, 1
  end
  add_message "api.Access" do
    optional :autom8ter_account, :string, 1
    optional :autom8ter_key, :string, 2
    optional :twilio_account, :string, 3
    optional :twilio_key, :string, 4
    optional :sendgrid_account, :string, 5
    optional :sendgrid_key, :string, 6
    optional :stripe_account, :string, 7
    optional :stripe_key, :string, 8
    optional :slack_account, :string, 9
    optional :slack_key, :string, 10
    optional :gcp_project, :string, 11
    optional :gcp_key, :string, 12
  end
  add_message "api.StandardClaims" do
    optional :access, :message, 1, "api.Access"
    optional :audience, :string, 2
    optional :subject, :string, 3
    optional :expires_at, :int64, 4
    optional :id, :string, 5
    optional :issued_at, :int64, 6
    optional :not_before, :int64, 7
  end
  add_message "api.LogConfig" do
    optional :username, :string, 1
    optional :channel, :string, 2
  end
  add_message "api.EmailAddress" do
    optional :name, :string, 1
    optional :address, :string, 2
  end
  add_message "api.Email" do
    optional :from, :message, 1, "api.EmailAddress"
    optional :recipient, :message, 2, "api.RecipientEmail"
  end
  add_message "api.RecipientEmail" do
    optional :to, :message, 2, "api.EmailAddress"
    optional :subject, :string, 3
    optional :plain_text, :string, 4
    optional :html, :string, 5
  end
  add_message "api.SMS" do
    optional :to, :string, 1
    optional :from, :string, 2
    optional :body, :string, 3
    optional :media_url, :string, 4
    optional :callback, :string, 5
    optional :app, :string, 6
  end
  add_message "api.Call" do
    optional :to, :string, 1
    optional :from, :string, 2
    optional :callback, :string, 5
  end
  add_message "api.Fax" do
    optional :to, :string, 1
    optional :from, :string, 2
    optional :media_url, :string, 3
    optional :quality, :string, 4
    optional :callback, :string, 5
    optional :store_media, :bool, 6
  end
  add_message "api.SlashCommand" do
    optional :token, :string, 1
    optional :team_id, :string, 2
    optional :team_domain, :string, 3
    optional :enterprise_id, :string, 4
    optional :enterprise_name, :string, 6
    optional :channel_id, :string, 7
    optional :channel_name, :string, 8
    optional :user_id, :string, 9
    optional :user_name, :string, 10
    optional :command, :string, 11
    optional :text, :string, 12
    optional :response_url, :string, 13
    optional :trigger_id, :string, 14
  end
  add_message "api.LogHook" do
    optional :author, :string, 1
    optional :icon, :string, 2
    optional :title, :string, 3
  end
  add_message "api.Attachment" do
    optional :color, :string, 1
    optional :fallback, :string, 2
    optional :callback_id, :string, 3
    optional :id, :int64, 4
    optional :author_id, :string, 5
    optional :author_name, :string, 6
    optional :author_link, :string, 7
    optional :author_icon, :string, 8
    optional :title, :string, 9
    optional :title_prefix, :string, 10
    optional :pretext, :string, 11
    optional :text, :string, 12
    optional :image_url, :string, 13
    optional :thumb_url, :string, 14
    repeated :fields, :message, 15, "api.AttachmentField"
  end
  add_message "api.AttachmentAction" do
    optional :name, :string, 1
    optional :text, :string, 2
    optional :style, :string, 3
    optional :type, :string, 4
    optional :value, :string, 5
    optional :data_source, :string, 6
    optional :min_query_length, :int64, 7
    repeated :options, :message, 8, "api.AttachmentActionOption"
    repeated :selected_options, :message, 9, "api.AttachmentActionOption"
    repeated :option_groups, :message, 10, "api.AttachmentActionOptionGroup"
    optional :confirm, :message, 11, "api.AttachmentConfirmationField"
    optional :url, :string, 12
  end
  add_message "api.AttachmentConfirmationField" do
    optional :title, :string, 1
    optional :text, :string, 2
    optional :ok_text, :string, 3
    optional :dismiss_text, :string, 4
  end
  add_message "api.AttachmentActionOptionGroup" do
    optional :text, :string, 1
    repeated :options, :message, 2, "api.AttachmentActionOption"
  end
  add_message "api.AttachmentActionOption" do
    optional :title, :string, 1
    optional :value, :string, 2
    optional :description, :string, 3
  end
  add_message "api.AttachmentField" do
    optional :title, :string, 1
    optional :value, :string, 2
    optional :short, :bool, 3
  end
  add_message "api.JSON" do
    optional :data, :bytes, 1
    optional :size, :int64, 2
  end
  add_message "api.File" do
    optional :data, :bytes, 1
    optional :size, :int64, 2
    optional :name, :string, 3
    map :tags, :string, :string, 4
  end
  add_message "api.SlackHook" do
    optional :username, :string, 1
    optional :channel, :string, 2
  end
  add_message "api.Product" do
    optional :name, :string, 1
    optional :amount, :int64, 2
    optional :description, :string, 3
    repeated :files, :message, 4, "api.File"
    map :tags, :string, :string, 5
    optional :available, :bool, 6
  end
  add_message "api.StringMapString" do
    map :map, :string, :string, 1
  end
  add_message "api.Msg" do
    optional :id, :string, 1
    map :meta, :string, :string, 2
    optional :data, :bytes, 3
    optional :publish_time, :string, 4
  end
  add_enum "api.CustomerIndex" do
    value :ID, 0
    value :EMAIL, 1
    value :PHONE, 2
  end
  add_enum "api.Claim" do
    value :TWILIO, 0
    value :SENDGRID, 1
    value :STRIPE, 2
    value :SLACK, 3
    value :GCP, 4
    value :AUTOM8TER, 5
  end
  add_enum "api.SigningMethod" do
    value :HMAC, 0
    value :ECDSA, 1
    value :RSA, 2
    value :RSAPPS, 3
  end
  add_enum "api.CardType" do
    value :VISA, 0
    value :MASTERCARD, 1
    value :DISCOVER, 2
    value :AMEX, 3
  end
  add_enum "api.Topic" do
    value :USER, 0
    value :ACCOUNT, 1
    value :CUSTOMER, 2
    value :OTHER, 3
  end
end

module Api
  String = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.String").msgclass
  Id = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Id").msgclass
  ActionHookRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ActionHookRequest").msgclass
  MessageUserRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.MessageUserRequest").msgclass
  RefundRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.RefundRequest").msgclass
  ChargeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ChargeRequest").msgclass
  CancelSubscriptionRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CancelSubscriptionRequest").msgclass
  CreatePlanRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CreatePlanRequest").msgclass
  SMSRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMSRequest").msgclass
  CallRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CallRequest").msgclass
  MMSRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.MMSRequest").msgclass
  EmailRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailRequest").msgclass
  CustomerRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CustomerRequest").msgclass
  UpdateCustomerRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UpdateCustomerRequest").msgclass
  SubscribeCustomerRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SubscribeCustomerRequest").msgclass
  CreateAccountRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CreateAccountRequest").msgclass
  Account = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Account").msgclass
  User = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.User").msgclass
  Profile = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Profile").msgclass
  Empty = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Empty").msgclass
  Customer = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Customer").msgclass
  Card = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Card").msgclass
  BankAccount = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.BankAccount").msgclass
  Address = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Address").msgclass
  ChannelReminder = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ChannelReminder").msgclass
  UserReminder = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.UserReminder").msgclass
  ItemRef = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.ItemRef").msgclass
  Star = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Star").msgclass
  Pin = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Pin").msgclass
  SignedKey = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SignedKey").msgclass
  Access = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Access").msgclass
  StandardClaims = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.StandardClaims").msgclass
  LogConfig = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.LogConfig").msgclass
  EmailAddress = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.EmailAddress").msgclass
  Email = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Email").msgclass
  RecipientEmail = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.RecipientEmail").msgclass
  SMS = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SMS").msgclass
  Call = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Call").msgclass
  Fax = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Fax").msgclass
  SlashCommand = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SlashCommand").msgclass
  LogHook = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.LogHook").msgclass
  Attachment = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Attachment").msgclass
  AttachmentAction = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AttachmentAction").msgclass
  AttachmentConfirmationField = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AttachmentConfirmationField").msgclass
  AttachmentActionOptionGroup = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AttachmentActionOptionGroup").msgclass
  AttachmentActionOption = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AttachmentActionOption").msgclass
  AttachmentField = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.AttachmentField").msgclass
  JSON = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.JSON").msgclass
  File = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.File").msgclass
  SlackHook = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SlackHook").msgclass
  Product = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Product").msgclass
  StringMapString = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.StringMapString").msgclass
  Msg = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Msg").msgclass
  CustomerIndex = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CustomerIndex").enummodule
  Claim = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Claim").enummodule
  SigningMethod = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.SigningMethod").enummodule
  CardType = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.CardType").enummodule
  Topic = Google::Protobuf::DescriptorPool.generated_pool.lookup("api.Topic").enummodule
end
