package node_response

type Status string

const (
  Ok Status = "ok"
  Error Status = "error"
)

// CallResult - RPC call result
type CallResult struct {
  Status  Status
  Message string
}

