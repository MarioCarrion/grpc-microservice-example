#!/usr/bin/env ruby
# frozen_string_literal: true

# XXX: In the real world a Ruby gem would be generated and required instead
lib_dir = File.expand_path(File.join('..', '..', '..', 'gen', 'ruby'))
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'user/v1/user_service_services_pb'

stub = User::V1::UserService::Stub.new('localhost:9879', :this_channel_is_insecure, timeout: 2) # 2 seconds timeout
response = stub.get_user(User::V1::GetUserRequest.new(uuid: '1-2-3-4'))

puts "Message #{response}"
