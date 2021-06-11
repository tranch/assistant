package rpcclient

import (
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"sync"
)

var (
	todoOnce   sync.Once
	todoClient pb.TodoClient
)

func GetTodoClient(client *rpc.Client) pb.TodoClient {
	todoOnce.Do(func() {
		conn, err := client.Dial("todo")
		if err != nil {
			fmt.Println(err)
			return
		}
		todoClient = pb.NewTodoClient(conn)
	})

	return todoClient
}