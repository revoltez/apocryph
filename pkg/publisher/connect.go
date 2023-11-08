package publisher

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/comrade-coop/trusted-pods/pkg/ipfs"
	pb "github.com/comrade-coop/trusted-pods/pkg/proto"
	"github.com/libp2p/go-libp2p/core/peer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	tpipfs "github.com/comrade-coop/trusted-pods/pkg/ipfs"
)

func ConnectToProvider(ipfsP2p *ipfs.P2pApi, deployment *pb.Deployment) (*tpipfs.IpfsClientConn, error) {
	providerPeerId, err := peer.Decode(deployment.GetProvider().GetLibp2PAddress())
	if err != nil {
		return nil, fmt.Errorf("Failed to parse provider address: %w", err)
	}

	conn, err := ipfsP2p.ConnectToGrpc(pb.ProvisionPod, providerPeerId, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("Failed to dial provider: %w", err)
	}
	return conn, nil
}

func SendToProvider(ctx context.Context, ipfsP2p *ipfs.P2pApi, pod *pb.Pod, deployment *pb.Deployment) error {
	// tpipfs.NewP2pApi(ipfs, ipfsMultiaddr)
	keys := []*pb.Key{}
	pod = LinkUploadsFromDeployment(pod, &keys, deployment)

	conn, err := ConnectToProvider(ipfsP2p, deployment)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewProvisionPodServiceClient(conn)

	fmt.Fprintf(os.Stderr, "Sending request to provider over IPFS p2p...\n")

	if pod != nil {
		var response *pb.ProvisionPodResponse
		if deployment.Deployed == nil || deployment.Deployed.Error != "" {
			request := &pb.ProvisionPodRequest{
				Pod:  pod,
				Keys: keys,
				Payment: &pb.PaymentChannel{
					ChainID:          deployment.Payment.ChainID,
					ProviderAddress:  deployment.Provider.EthereumAddress,
					ContractAddress:  deployment.Payment.PaymentContractAddress,
					PublisherAddress: deployment.Payment.PublisherAddress,
					PodID:            deployment.Payment.PodID,
				},
			}
			response, err = client.ProvisionPod(ctx, request)
			if err != nil {
				return fmt.Errorf("Failed executing provision pod request: %w", err)
			}
		} else {
			request := &pb.UpdatePodRequest{
				Pod:         pod,
				Keys:        keys,
				Credentials: &pb.Credentials{},
			}
			response, err = client.UpdatePod(ctx, request)
			if err != nil {
				return fmt.Errorf("Failed executing update pod request: %w", err)
			}
		}

		if response.Error != "" {
			return fmt.Errorf("Error from provider: %w", errors.New(response.Error))
		}

		deployment.Deployed = response
		fmt.Fprintf(os.Stderr, "Successfully deployed! %v\n", response)
	} else {
		request := &pb.DeletePodRequest{
			Credentials: &pb.Credentials{},
		}
		response, err := client.DeletePod(ctx, request)
		if err != nil {
			return fmt.Errorf("Failed executing update pod request: %w", err)
		}
		if response.Error != "" {
			return fmt.Errorf("Error from provider: %w", errors.New(response.Error))
		}
		deployment.Deployed = nil
		fmt.Fprintf(os.Stderr, "Successfully undeployed!\n")
	}

	return nil
}