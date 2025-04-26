package netcheck

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"time"
)

func NewNetCheckCmd() *cobra.Command {
	var host string
	var port int
	var timeout time.Duration

	cmd := &cobra.Command{
		Use:   "nw",
		Short: "Check network reachability to a host:port",
		Run: func(cmd *cobra.Command, args []string) {
			address := fmt.Sprintf("%s:%d", host, port)
			fmt.Printf("Checking %s ...\\n", address)
			start := time.Now()
			conn, err := net.DialTimeout("tcp", address, timeout)
			duration := time.Since(start)

			if err != nil {
				fmt.Printf("❌ Connection failed: %s\\n", err)
				return
			}
			defer func(conn net.Conn) {
				err := conn.Close()
				if err != nil {
					fmt.Printf("❌ Failed to close connection: %s\\n", err)
				} else {
					fmt.Printf("✅ Connection closed successfully.\\n")
				}
			}(conn)
			fmt.Printf("✅ Success! Response time: %v \\n", duration)
		},
	}

	cmd.Flags().StringVar(&host, "host", "", "Target hostname (e.g., google.com)")
	cmd.Flags().IntVar(&port, "port", 80, "Target port")
	cmd.Flags().DurationVar(&timeout, "timeout", 2*time.Second, "Timeout duration")
	err := cmd.MarkFlagRequired("host")
	if err != nil {
		return nil
	}

	return cmd
}
