$:.unshift(File.dirname(__FILE__))

require 'yaml'
require 'lib/rpc'

data = YAML.load($stdin.read)

host = 'localhost'
port = 8081
rpc = RPC::Client.new host, port

for call in data
  method, args = call
  rpc.call method, args
end

rpc.close
