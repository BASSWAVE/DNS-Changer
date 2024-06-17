package main

import (
	"github.com/BASSWAVE/DNS-Changer/api/dns"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var rootCmd = &cobra.Command{Use: "dnscli"}

	var hostname string
	var dnsServer string

	var setHostnameCmd = &cobra.Command{
		Use:   "set-hostname",
		Short: "Set the system hostname",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Failed to connect to gRPC server: %v", err)
			}
			defer conn.Close()

			client := dns.NewDNSServiceClient(conn)
			req := &dns.SetHostnameRequest{Hostname: hostname}
			res, err := client.SetHostname(context.Background(), req)
			if err != nil {
				log.Fatalf("Error calling SetHostname: %v", err)
			}

			log.Printf("Response: %v", res.GetMessage())
		},
	}
	setHostnameCmd.Flags().StringVarP(&hostname, "hostname", "n", "", "Hostname to set")
	setHostnameCmd.MarkFlagRequired("hostname")

	var getDNSCmd = &cobra.Command{
		Use:   "get-dns",
		Short: "Get the list of DNS servers",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Failed to connect to gRPC server: %v", err)
			}
			defer conn.Close()

			client := dns.NewDNSServiceClient(conn)
			req := &dns.GetDNSRequest{}
			res, err := client.GetDNS(context.Background(), req)
			if err != nil {
				log.Fatalf("Error calling GetDNS: %v", err)
			}

			log.Printf("DNS Servers: %v", res.GetDnsServers())
		},
	}

	var addDNSCmd = &cobra.Command{
		Use:   "add-dns",
		Short: "Add a DNS server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Failed to connect to gRPC server: %v", err)
			}
			defer conn.Close()

			client := dns.NewDNSServiceClient(conn)
			req := &dns.AddDNSRequest{DnsServer: dnsServer}
			res, err := client.AddDNS(context.Background(), req)
			if err != nil {
				log.Fatalf("Error calling AddDNS: %v", err)
			}

			log.Printf("Response: %v", res.GetMessage())
		},
	}
	addDNSCmd.Flags().StringVarP(&dnsServer, "dns-server", "d", "", "DNS server to add")
	addDNSCmd.MarkFlagRequired("dns-server")

	var removeDNSCmd = &cobra.Command{
		Use:   "remove-dns",
		Short: "Remove a DNS server",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Failed to connect to gRPC server: %v", err)
			}
			defer conn.Close()

			client := dns.NewDNSServiceClient(conn)
			req := &dns.RemoveDNSRequest{DnsServer: dnsServer}
			res, err := client.RemoveDNS(context.Background(), req)
			if err != nil {
				log.Fatalf("Error calling RemoveDNS: %v", err)
			}

			log.Printf("Response: %v", res.GetMessage())
		},
	}
	removeDNSCmd.Flags().StringVarP(&dnsServer, "dns-server", "d", "", "DNS server to remove")
	removeDNSCmd.MarkFlagRequired("dns-server")

	rootCmd.AddCommand(setHostnameCmd, getDNSCmd, addDNSCmd, removeDNSCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
