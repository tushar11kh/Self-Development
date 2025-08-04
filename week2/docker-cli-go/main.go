package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/docker/docker/api/types/container"
	// "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func formatSize(bytes *int64) string {
	if bytes == nil {
		return "n/a"
	}
	size := float64(*bytes)
	switch {
	case size > 1024*1024*1024:
		return fmt.Sprintf("%.2f GB", size/(1024*1024*1024))
	case size > 1024*1024:
		return fmt.Sprintf("%.2f MB", size/(1024*1024))
	case size > 1024:
		return fmt.Sprintf("%.2f KB", size/1024)
	default:
		return fmt.Sprintf("%d B", int(size))
	}
}

func main() {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	cli.NegotiateAPIVersion(context.Background())

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true, Size: true})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}

	if len(containers) == 0 {
		fmt.Println("No containers found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "CONTAINER ID\tNAME\tIMAGE\tSIZE\tSTATUS")

	for _, container := range containers {

		sizeRw := formatSize(&container.SizeRw)
		sizeRootFs := formatSize(&container.SizeRootFs)
		fullSize := fmt.Sprintf("%s (virtual %s)", sizeRw, sizeRootFs)

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",

			container.ID[:12],
			container.Names[0][1:],
			container.Image,
			fullSize,
			container.Status,
		)

		w.Flush()

	}

}
