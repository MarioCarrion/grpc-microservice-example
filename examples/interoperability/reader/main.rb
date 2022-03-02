#!/usr/bin/env ruby

# XXX: In the real world a Ruby gem would be generated and required instead
require_relative File.join("..", "..", "..", "gen", "ruby", "user", "v1", "user_pb")

def main
  message = User::V1::User.decode(File.read("user.bin"))
  puts message.has_salary?
  puts message.inspect
end

main
