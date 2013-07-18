$:.unshift(File.dirname(__FILE__))

require 'lib/rpc'

Signal.trap(:INT) { exit }

id = '0'

rpc = RPC::Client.new "localhost", 8081
rpc.call "NodeRPC.CreateRoom", {
  Id: id,
  TableSize: 9,
  BetSize: 10.0,
  Mix: nil,
  Game: {
    Type: "texas",
    Limit: "no-limit"
  }
}

rpc.call "NodeRPC.RoomSend", {
}

rpc.close
