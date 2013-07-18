require 'socket'
require 'json'

module RPC
  class Client

    def initialize(host, port)
      @socket = TCPSocket.new(host, port)
      @counter = 0
    end

    def encode(name, params)
      @counter += 1
      {
        id: @counter,
        method: name,
        params: params
      }
    end

    def call(name, *params)
      msg = encode(name, params)
      id = msg[:id]
      
      puts "sending #{msg.to_json} ==>"
      @socket.write(msg.to_json)
      
      resp = JSON.parse(@socket.gets)
      puts "==> received #{resp}"

      if resp['id'] != id
        raise "expected id=#{id}, got id=#{resp['id']}"
      end

      if resp['error']
        raise resp['error']
      end

      resp['result']
    end

    def close
      @socket.close
    end

  end
end
