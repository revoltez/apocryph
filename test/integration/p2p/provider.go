package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	ipfs_utils "github.com/comrade-coop/trusted-pods/pkg/ipfs-utils"
	podmanagement "github.com/comrade-coop/trusted-pods/pkg/pod-management"
	pb "github.com/comrade-coop/trusted-pods/pkg/proto"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 6000, "The server port")
)

type server struct {
	pb.UnimplementedSampleServer
}

func (s *server) SendPod(ctx context.Context, in *pb.SampleProvisionPodRequest) (*pb.SampleProvisionPodReply, error) {
	println("PROVIDER: pod package cid:", in.Cid)
	// cleanup unpinned files
	exec.Command("ipfs", "repo", "gc")
	provider, _ := podmanagement.CreateIpfsUploader()
	ipfs_utils.RetreiveFile(provider.Node, in.Cid, "/tmp/pod-package")
	cmd := exec.Command("ls", "-al", "/tmp/pod-package")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("PROVIDER: could not list directory files:", err)
		fmt.Println("PROVIDER: Command output:", string(output))
		return nil, err
	}
	fmt.Println("PROVIDER: Pod Package:", string(output))
	return &pb.SampleProvisionPodReply{Endpoint: "http://provider.com/mypod"}, nil
}

func main() {
	// Create a channel to receive the interrupt signal
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interruptChan
		os.Exit(0)
	}()

	node, err := rpc.NewLocalApi()
	if err != nil {
		println("PROVIDER: could not connect to local node")
		return
	}

	// route all ipfs p2p connections of the provios-pod protocol to the grpc server
	lis, err := ipfs_utils.NewP2pApi(node, multiaddr.StringCast("/ip4/127.0.0.1/")).Listen(pb.ProvisionPod)
	if err != nil {
		log.Fatalf("PROVIDER: failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterSampleServer(s, &server{})
	log.Printf("PROVIDER: server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("PROVIDER: failed to serve: %v", err)
	}
}