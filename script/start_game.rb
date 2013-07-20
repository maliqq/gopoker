$:.unshift(File.dirname(__FILE__))

id = '0'
require 'lib/rpc'

rpc = RPC::Client.new "localhost", 8081

rpc.call "NodeRPC.StartRoom", {
  Id: id,
  Mode: "cash"
}
