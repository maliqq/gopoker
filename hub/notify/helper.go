package notify

import (
  "gopoker/hub"
)

func All(message interface{}) Notification {
  return Notification{
    Route: ExchangeRoute{
      All: true,
    },
    Message: message,
  }
}

func One()