$:.unshift(File.dirname(__FILE__))

require 'yaml'
require 'lib/rpc'

Signal.trap(:INT) { exit }

id = '0'
size = 9

avatars = %w(bender2.jpg  bender.jpg  fry.jpg  hermes.jpg  homer.jpg  jake.jpg  labarbara.jpg  leela.jpg  roger.jpg stewie.jpg)
places = YAML.load(File.open(File.join(File.dirname(__FILE__), 'places.yml')).read)

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
      Payload: {
        Pos: i,
        Player: {
          Id: "player-#{i}",
          Name: "Player #{i}",
          NickName: "player_#{i}",
          Place: places[rand(places.size)],
          Avatar: avatars[rand(avatars.size)]
        },
        Amount: (rand() * 80 + 20).round(2)
      }
    }
  }
}

rpc.close
