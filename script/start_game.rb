$:.unshift(File.dirname(__FILE__))

require 'yaml'
require 'lib/rpc'

Signal.trap(:INT) { exit }

id = '0'
size = 9

rpc = RPC::Client.new "localhost", 8081
rpc.call "NodeRPC.CreateRoom", {
  Id: id,
  TableSize: size,
  BetSize: 10.0,
  Mix: nil,
  Game: {
    Type: "texas",
    Limit: "no-limit"
  }
}

size.times { |i|
  rpc.call "NodeRPC.NotifyRoom", {
    Id: id,
    Message: {
      Type: "JoinTable",
      Timestamp: Time.now.to_i,
      Envelope: {
        JoinTable: {
          Pos: i,
          Player: "player-#{i}",
          Amount: (rand() * 80 + 20).round(2)
        }
      }
    }
  }
}

rpc.close
