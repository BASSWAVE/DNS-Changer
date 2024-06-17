package main

import (
	"context"
	"fmt"
	"github.com/BASSWAVE/DNS-Changer/api/dns"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	dns.UnimplementedDNSServiceServer
}

func (s *server) SetHostname(ctx context.Context, req *dns.SetHostnameRequest) (*dns.SetHostnameResponse, error) {
	if !validateHostName(req.GetHostname()) {
		return &dns.SetHostnameResponse{Success: false, Message: "Not valid hostname"}, status.Errorf(codes.Internal, "Not valid hostname")
	}
	cmd := exec.Command("hostnamectl", "set-hostname", req.GetHostname())
	if err := cmd.Run(); err != nil {
		return &dns.SetHostnameResponse{Success: false, Message: err.Error()}, status.Errorf(codes.Internal, "Failed to set hostname: %v", err)
	}
	return &dns.SetHostnameResponse{Success: true, Message: "Hostname set successfully"}, nil
}

func (s *server) GetDNS(ctx context.Context, req *dns.GetDNSRequest) (*dns.GetDNSResponse, error) {
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to read resolv.conf: %v", err)
	}
	dnsServers := parseDNSServers(string(content))
	return &dns.GetDNSResponse{DnsServers: dnsServers}, nil
}

func (s *server) AddDNS(ctx context.Context, req *dns.AddDNSRequest) (*dns.AddDNSResponse, error) {
	if ip := net.ParseIP(req.GetDnsServer()); ip == nil {
		return &dns.AddDNSResponse{Success: false, Message: "Failed to parse IP address"}, status.Errorf(codes.Internal, "Failed to parse IP address")
	}
	f, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return &dns.AddDNSResponse{Success: false, Message: err.Error()}, status.Errorf(codes.Internal, "Failed to open resolv.conf: %v", err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("nameserver %s\n", req.GetDnsServer())); err != nil {
		return &dns.AddDNSResponse{Success: false, Message: err.Error()}, status.Errorf(codes.Internal, "Failed to write to resolv.conf: %v", err)
	}
	return &dns.AddDNSResponse{Success: true, Message: "DNS server added successfully"}, nil
}

func (s *server) RemoveDNS(ctx context.Context, req *dns.RemoveDNSRequest) (*dns.RemoveDNSResponse, error) {
	if ip := net.ParseIP(req.GetDnsServer()); ip == nil {
		return &dns.RemoveDNSResponse{Success: false, Message: "Failed to parse IP address"}, status.Errorf(codes.Internal, "Failed to parse IP address")
	}
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return &dns.RemoveDNSResponse{Success: false, Message: err.Error()}, status.Errorf(codes.Internal, "Failed to read resolv.conf: %v", err)
	}
	updatedContent, dnsWasRemoved := removeDNSServer(string(content), req.GetDnsServer())
	if err := os.WriteFile("/etc/resolv.conf", []byte(updatedContent), 0644); err != nil {
		return &dns.RemoveDNSResponse{Success: false, Message: err.Error()}, status.Errorf(codes.Internal, "Failed to write to resolv.conf: %v", err)
	}
	if dnsWasRemoved {
		return &dns.RemoveDNSResponse{Success: true, Message: "DNS server removed successfully"}, nil
	}
	return &dns.RemoveDNSResponse{Success: false, Message: "DNS server was not at resolv.conf"}, nil
}

func parseDNSServers(content string) []string {
	var dnsServers []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver") {
			dnsServers = append(dnsServers, strings.Fields(line)[1])
		}
	}
	return dnsServers
}

func removeDNSServer(content, dnsServer string) (string, bool) {
	lines := strings.Split(content, "\n")
	var updatedLines []string
	dnsWasRemoved := false
	for _, line := range lines {
		if strings.HasPrefix(line, "nameserver") && strings.Fields(line)[1] == dnsServer {
			dnsWasRemoved = true
			continue
		}
		updatedLines = append(updatedLines, line)
	}

	return strings.Join(updatedLines, "\n"), dnsWasRemoved
}

func validateHostName(s string) bool {
	validStringPattern := `^[a-zA-Z0-9-]+$`
	match, _ := regexp.MatchString(validStringPattern, s)
	return match
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dns.RegisterDNSServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = dns.RegisterDNSServiceHandlerFromEndpoint(ctx, mux, ":50051", opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	log.Printf("Serving gRPC-Gateway on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve HTTP gateway: %v", err)
	}
}
