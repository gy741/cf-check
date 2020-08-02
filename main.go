package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	concurrency := 20
	jobs := make(chan string)
	var wg sync.WaitGroup
	var domainMode bool

	flag.IntVar(&concurrency, "c", 20, "Set the concurrency level")
	flag.BoolVar(&domainMode, "d", false, "Prints domain instead of IP address")
	flag.Parse()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for host := range jobs {
				addr, err := net.LookupIP(strings.TrimSpace(host))
				if err != nil {
					continue
				}

				if !isCloudflare(addr[0]) {
					if domainMode {
						fmt.Println(host)
					} else {
						fmt.Println(addr[0])
					}
				}
			}
			wg.Done()
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		jobs <- sc.Text()
	}

	close(jobs)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	wg.Wait()
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, nil

	default:
		return ips[1 : len(ips)-1], nil
	}
}

func isCloudflare(ip net.IP) bool {
	cidrs := []string{"173.245.48.0/20", "103.21.244.0/22", "103.22.200.0/22", "103.31.4.0/22", "141.101.64.0/18", "108.162.192.0/18", "190.93.240.0/20", "188.114.96.0/20", "197.234.240.0/22", "198.41.128.0/17", "162.158.0.0/15", "104.16.0.0/12", "172.64.0.0/13", "131.0.72.0/22", "1.0.0.0/24", "1.1.1.0/24", "103.21.246.0/24", "103.21.247.0/24", "103.22.202.0/24", "103.22.203.0/24", "104.16.112.0/20", "104.16.128.0/20", "104.16.144.0/20", "104.16.16.0/20", "104.16.160.0/20", "104.16.176.0/20", "104.16.192.0/20", "104.16.208.0/20", "104.16.224.0/20", "104.16.240.0/20", "104.16.32.0/20", "104.16.48.0/20", "104.16.64.0/20", "104.16.80.0/20", "104.16.96.0/20", "104.17.0.0/20", "104.17.112.0/20", "104.17.128.0/20", "104.17.144.0/20", "104.17.16.0/20", "104.17.160.0/20", "104.17.176.0/20", "104.17.192.0/20", "104.17.208.0/20", "104.17.224.0/20", "104.17.240.0/20", "104.17.32.0/20", "104.17.48.0/20", "104.17.64.0/20", "104.17.80.0/20", "104.17.96.0/20", "104.18.0.0/20", "104.18.112.0/20", "104.18.128.0/20", "104.18.144.0/20", "104.18.16.0/20", "104.18.160.0/20", "104.18.176.0/20", "104.18.192.0/20", "104.18.208.0/20", "104.18.224.0/20", "104.18.240.0/20", "104.18.32.0/19", "104.18.48.0/20", "104.18.64.0/20", "104.18.80.0/20", "104.18.96.0/20", "104.19.0.0/20", "104.19.112.0/20", "104.19.128.0/20", "104.19.144.0/20", "104.19.16.0/20", "104.19.160.0/20", "104.19.176.0/20", "104.19.192.0/20", "104.19.208.0/20", "104.19.224.0/20", "104.19.240.0/20", "104.19.32.0/20", "104.19.48.0/20", "104.19.64.0/20", "104.19.80.0/20", "104.19.96.0/20", "104.20.0.0/20", "104.20.112.0/20", "104.20.128.0/20", "104.20.144.0/20", "104.20.16.0/20", "104.20.160.0/20", "104.20.176.0/20", "104.20.192.0/20", "104.20.208.0/20", "104.20.224.0/20", "104.20.240.0/20", "104.20.32.0/20", "104.20.48.0/20", "104.20.64.0/20", "104.20.80.0/20", "104.20.96.0/20", "104.22.0.0/20", "104.22.16.0/20", "104.22.32.0/20", "104.22.48.0/20", "104.22.64.0/20", "104.23.112.0/20", "104.23.128.0/20", "104.23.192.0/20", "104.23.240.0/22", "104.23.96.0/20", "104.24.0.0/20", "104.24.112.0/20", "104.24.128.0/20", "104.24.144.0/20", "104.24.16.0/20", "104.24.160.0/20", "104.24.176.0/20", "104.24.192.0/20", "104.24.208.0/20", "104.24.224.0/20", "104.24.240.0/20", "104.24.32.0/20", "104.24.48.0/20", "104.24.64.0/20", "104.24.80.0/20", "104.24.96.0/19", "104.25.0.0/20", "104.25.112.0/20", "104.25.128.0/20", "104.25.144.0/20", "104.25.16.0/20", "104.25.160.0/20", "104.25.176.0/20", "104.25.192.0/20", "104.25.208.0/20", "104.25.224.0/20", "104.25.240.0/20", "104.25.32.0/20", "104.25.48.0/20", "104.25.64.0/20", "104.25.80.0/20", "104.25.96.0/20", "104.26.0.0/20", "104.27.0.0/20", "104.27.112.0/20", "104.27.128.0/19", "104.27.144.0/20", "104.27.16.0/20", "104.27.160.0/19", "104.27.176.0/20", "104.27.192.0/20", "104.27.208.0/20", "104.27.240.0/22", "104.27.32.0/20", "104.27.48.0/20", "104.27.64.0/20", "104.27.80.0/20", "104.27.96.0/20", "104.28.0.0/19", "104.28.112.0/20", "104.28.128.0/19", "104.28.144.0/20", "104.28.16.0/20", "104.28.160.0/19", "104.28.176.0/20", "104.28.192.0/19", "104.28.208.0/20", "104.28.224.0/19", "104.28.240.0/20", "104.28.32.0/19", "104.28.48.0/20", "104.28.64.0/19", "104.28.80.0/20", "104.28.96.0/19", "104.31.0.0/20", "104.31.112.0/22", "104.31.128.0/20", "104.31.144.0/20", "104.31.16.0/20", "104.31.160.0/19", "104.31.176.0/20", "104.31.192.0/19", "104.31.208.0/20", "104.31.224.0/20", "104.31.240.0/20", "104.31.64.0/19", "104.31.80.0/20", "108.162.193.0/24", "108.162.194.0/24", "108.162.208.0/24", "108.162.210.0/24", "108.162.211.0/24", "108.162.212.0/24", "108.162.213.0/24", "108.162.214.0/24", "108.162.215.0/24", "108.162.216.0/24", "108.162.217.0/24", "108.162.218.0/24", "108.162.219.0/24", "108.162.220.0/24", "108.162.221.0/24", "108.162.222.0/24", "108.162.223.0/24", "108.162.226.0/24", "108.162.227.0/24", "108.162.228.0/24", "108.162.229.0/24", "108.162.235.0/24", "108.162.236.0/24", "108.162.237.0/24", "108.162.238.0/24", "108.162.239.0/24", "108.162.240.0/24", "108.162.241.0/24", "108.162.242.0/24", "108.162.243.0/24", "108.162.244.0/24", "108.162.245.0/24", "108.162.246.0/24", "108.162.247.0/24", "108.162.248.0/24", "108.162.249.0/24", "108.162.250.0/24", "108.162.252.0/24", "108.162.253.0/24", "108.162.255.0/24", "141.101.100.0/22", "141.101.104.0/24", "141.101.105.0/24", "141.101.106.0/24", "141.101.107.0/24", "141.101.108.0/24", "141.101.109.0/24", "141.101.110.0/24", "141.101.111.0/24", "141.101.112.0/20", "141.101.114.0/23", "141.101.120.0/22", "141.101.65.0/24", "141.101.66.0/24", "141.101.67.0/24", "141.101.68.0/24", "141.101.69.0/24", "141.101.70.0/24", "141.101.71.0/24", "141.101.72.0/24", "141.101.73.0/24", "141.101.74.0/24", "141.101.75.0/24", "141.101.76.0/23", "141.101.82.0/24", "141.101.83.0/24", "141.101.84.0/24", "141.101.85.0/24", "141.101.88.0/22", "141.101.94.0/24", "141.101.95.0/24", "141.101.96.0/24", "141.101.97.0/24", "141.101.98.0/24", "141.101.99.0/24", "162.158.10.0/24", "162.158.100.0/24", "162.158.101.0/24", "162.158.102.0/24", "162.158.103.0/24", "162.158.104.0/24", "162.158.105.0/24", "162.158.106.0/24", "162.158.107.0/24", "162.158.108.0/22", "162.158.11.0/24", "162.158.112.0/24", "162.158.113.0/24", "162.158.114.0/24", "162.158.115.0/24", "162.158.116.0/22", "162.158.12.0/22", "162.158.120.0/24", "162.158.121.0/24", "162.158.122.0/24", "162.158.123.0/24", "162.158.124.0/22", "162.158.128.0/22", "162.158.132.0/24", "162.158.133.0/24", "162.158.134.0/24", "162.158.135.0/24", "162.158.136.0/22", "162.158.140.0/24", "162.158.141.0/24", "162.158.142.0/24", "162.158.143.0/24", "162.158.144.0/24", "162.158.145.0/24", "162.158.146.0/24", "162.158.147.0/24", "162.158.148.0/22", "162.158.152.0/22", "162.158.156.0/22", "162.158.16.0/22", "162.158.160.0/20", "162.158.176.0/24", "162.158.177.0/24", "162.158.178.0/24", "162.158.179.0/24", "162.158.180.0/22", "162.158.184.0/24", "162.158.185.0/24", "162.158.186.0/24", "162.158.187.0/24", "162.158.188.0/24", "162.158.189.0/24", "162.158.190.0/24", "162.158.191.0/24", "162.158.192.0/24", "162.158.193.0/24", "162.158.194.0/24", "162.158.195.0/24", "162.158.196.0/24", "162.158.197.0/24", "162.158.198.0/24", "162.158.199.0/24", "162.158.20.0/22", "162.158.200.0/22", "162.158.204.0/23", "162.158.206.0/23", "162.158.207.0/24", "162.158.208.0/22", "162.158.212.0/22", "162.158.216.0/23", "162.158.218.0/23", "162.158.220.0/22", "162.158.224.0/22", "162.158.225.0/24", "162.158.226.0/23", "162.158.228.0/22", "162.158.232.0/22", "162.158.236.0/22", "162.158.24.0/23", "162.158.240.0/22", "162.158.244.0/22", "162.158.248.0/22", "162.158.25.0/24", "162.158.250.0/23", "162.158.252.0/22", "162.158.26.0/24", "162.158.27.0/24", "162.158.28.0/24", "162.158.29.0/24", "162.158.30.0/24", "162.158.31.0/24", "162.158.32.0/22", "162.158.36.0/24", "162.158.37.0/24", "162.158.38.0/24", "162.158.39.0/24", "162.158.4.0/22", "162.158.40.0/24", "162.158.41.0/24", "162.158.42.0/24", "162.158.43.0/24", "162.158.44.0/22", "162.158.45.0/24", "162.158.46.0/24", "162.158.47.0/24", "162.158.48.0/24", "162.158.49.0/24", "162.158.50.0/24", "162.158.51.0/24", "162.158.52.0/22", "162.158.53.0/24", "162.158.54.0/24", "162.158.55.0/24", "162.158.56.0/22", "162.158.59.0/24", "162.158.60.0/22", "162.158.64.0/21", "162.158.72.0/22", "162.158.76.0/22", "162.158.8.0/24", "162.158.80.0/23", "162.158.82.0/23", "162.158.83.0/24", "162.158.84.0/22", "162.158.85.0/24", "162.158.86.0/24", "162.158.87.0/24", "162.158.88.0/21", "162.158.89.0/24", "162.158.9.0/24", "162.158.90.0/24", "162.158.91.0/24", "162.158.92.0/24", "162.158.93.0/24", "162.158.94.0/24", "162.158.95.0/24", "162.158.96.0/24", "162.158.97.0/24", "162.158.98.0/24", "162.158.99.0/24", "162.159.0.0/20", "162.159.1.0/24", "162.159.10.0/24", "162.159.11.0/24", "162.159.12.0/24", "162.159.128.0/17 "162.159.13.0/24", "162.159.132.0/24", "162.159.14.0/24", "162.159.15.0/24", "162.159.16.0/20", "162.159.160.0/24", "162.159.17.0/24", "162.159.18.0/24", "162.159.19.0/24", "162.159.192.0/22", "162.159.2.0/24", "162.159.20.0/24", "162.159.200.0/24", "162.159.208.0/22", "162.159.209.0/24", "162.159.21.0/24", "162.159.210.0/24", "162.159.211.0/24", "162.159.212.0/22", "162.159.216.0/21", "162.159.22.0/24", "162.159.224.0/20", "162.159.23.0/24", "162.159.24.0/24", "162.159.240.0/20", "162.159.25.0/24", "162.159.26.0/24", "162.159.27.0/24", "162.159.28.0/24", "162.159.29.0/24", "162.159.3.0/24", "162.159.30.0/24", "162.159.31.0/24", "162.159.32.0/20", "162.159.34.0/23", "162.159.36.0/24", "162.159.4.0/24", "162.159.40.0/23", "162.159.42.0/23", "162.159.46.0/24", "162.159.48.0/20", "162.159.5.0/24", "162.159.6.0/24", "162.159.64.0/20", "162.159.7.0/24", "162.159.8.0/24", "162.159.9.0/24", "162.251.82.0/24", "172.64.112.0/20", "172.64.128.0/20", "172.64.144.0/20", "172.64.16.0/20", "172.64.160.0/20", "172.64.176.0/20", "172.64.192.0/20", "172.64.208.0/20", "172.64.32.0/20", "172.64.36.0/23", "172.64.48.0/20", "172.64.64.0/20", "172.64.96.0/20", "172.65.0.0/19", "172.65.112.0/20", "172.65.128.0/20", "172.65.144.0/20", "172.65.16.0/20", "172.65.160.0/20", "172.65.176.0/20", "172.65.192.0/20", "172.65.208.0/20", "172.65.224.0/20", "172.65.240.0/20", "172.65.32.0/20", "172.65.48.0/20", "172.65.64.0/20", "172.65.80.0/20", "172.65.96.0/20", "172.67.0.0/20", "172.67.112.0/20", "172.67.128.0/20", "172.67.144.0/20", "172.67.16.0/20", "172.67.160.0/20", "172.67.176.0/20", "172.67.192.0/20", "172.67.208.0/20", "172.67.224.0/20", "172.67.240.0/20", "172.67.48.0/20", "172.67.64.0/20", "172.67.80.0/20", "172.68.0.0/22", "172.68.100.0/22", "172.68.104.0/22", "172.68.108.0/22", "172.68.112.0/22", "172.68.116.0/22", "172.68.12.0/22", "172.68.120.0/22", "172.68.124.0/22", "172.68.128.0/22", "172.68.129.0/24", "172.68.131.0/24", "172.68.132.0/22", "172.68.136.0/22", "172.68.140.0/22", "172.68.144.0/22", "172.68.148.0/22", "172.68.152.0/22", "172.68.16.0/20", "172.68.160.0/22", "172.68.164.0/23", "172.68.166.0/23", "172.68.168.0/23", "172.68.170.0/24", "172.68.171.0/24", "172.68.172.0/22", "172.68.176.0/22", "172.68.180.0/22", "172.68.184.0/22", "172.68.188.0/22", "172.68.196.0/22", "172.68.200.0/22", "172.68.204.0/22", "172.68.208.0/22", "172.68.212.0/22", "172.68.216.0/22", "172.68.220.0/22", "172.68.224.0/22", "172.68.228.0/22", "172.68.232.0/22", "172.68.236.0/22", "172.68.24.0/22", "172.68.240.0/22", "172.68.244.0/22", "172.68.248.0/21", "172.68.28.0/24", "172.68.29.0/24", "172.68.30.0/24", "172.68.31.0/24", "172.68.32.0/22", "172.68.36.0/22", "172.68.4.0/22", "172.68.40.0/22", "172.68.44.0/22", "172.68.48.0/22", "172.68.52.0/22", "172.68.56.0/24", "172.68.57.0/24", "172.68.58.0/24", "172.68.59.0/24", "172.68.60.0/22", "172.68.64.0/22", "172.68.68.0/22", "172.68.72.0/22", "172.68.76.0/22", "172.68.8.0/22", "172.68.80.0/22", "172.68.84.0/22", "172.68.88.0/22", "172.68.92.0/22", "172.68.96.0/22", "172.69.0.0/22", "172.69.100.0/22", "172.69.104.0/22", "172.69.108.0/22", "172.69.110.0/24", "172.69.111.0/24", "172.69.112.0/22", "172.69.116.0/22", "172.69.12.0/24", "172.69.120.0/22", "172.69.124.0/22", "172.69.128.0/22", "172.69.132.0/22", "172.69.136.0/22", "172.69.14.0/24", "172.69.140.0/22", "172.69.144.0/22", "172.69.148.0/22", "172.69.15.0/24", "172.69.152.0/22", "172.69.156.0/22", "172.69.157.0/24", "172.69.158.0/24", "172.69.159.0/24", "172.69.16.0/20", "172.69.160.0/22", "172.69.161.0/24", "172.69.162.0/24", "172.69.163.0/24", "172.69.164.0/22", "172.69.168.0/22", "172.69.17.0/24", "172.69.172.0/22", "172.69.176.0/22", "172.69.18.0/24", "172.69.180.0/22", "172.69.184.0/22", "172.69.188.0/22", "172.69.19.0/24", "172.69.192.0/22", "172.69.196.0/22", "172.69.2.0/24", "172.69.20.0/22", "172.69.200.0/22", "172.69.204.0/22", "172.69.208.0/22", "172.69.212.0/22", "172.69.216.0/22", "172.69.217.0/24", "172.69.218.0/24", "172.69.219.0/24", "172.69.220.0/22", "172.69.224.0/22", "172.69.226.0/23", "172.69.227.0/24", "172.69.228.0/22", "172.69.232.0/22", "172.69.234.0/24", "172.69.235.0/24", "172.69.236.0/22", "172.69.24.0/21", "172.69.240.0/22", "172.69.244.0/22", "172.69.246.0/23", "172.69.248.0/22", "172.69.252.0/22", "172.69.253.0/24", "172.69.254.0/24", "172.69.255.0/24", "172.69.3.0/24", "172.69.32.0/22", "172.69.36.0/23", "172.69.38.0/23", "172.69.4.0/22", "172.69.40.0/22", "172.69.44.0/22", "172.69.45.0/24", "172.69.46.0/24", "172.69.47.0/24", "172.69.48.0/22", "172.69.52.0/22", "172.69.56.0/22", "172.69.60.0/22", "172.69.64.0/21", "172.69.72.0/22", "172.69.76.0/22", "172.69.8.0/22", "172.69.80.0/22", "172.69.84.0/22", "172.69.88.0/22", "172.69.92.0/22", "172.69.96.0/22", "173.245.49.0/24", "173.245.52.0/24", "173.245.54.0/24", "173.245.56.0/24", "173.245.58.0/24", "173.245.59.0/24", "173.245.60.0/23", "173.245.62.0/24", "173.245.63.0/24", "185.122.0.0/24", "188.114.100.0/24", "188.114.101.0/24", "188.114.102.0/24", "188.114.103.0/24", "188.114.104.0/24", "188.114.106.0/23", "188.114.108.0/24", "188.114.109.0/24", "188.114.110.0/24", "188.114.111.0/24", "188.114.97.0/24", "188.114.98.0/24", "188.114.99.0/24", "190.93.244.0/22", "197.234.241.0/24", "197.234.242.0/24", "197.234.243.0/24", "198.217.251.0/24", "198.41.129.0/24", "198.41.130.0/24", "198.41.132.0/22", "198.41.136.0/22", "198.41.144.0/22", "198.41.148.0/22", "198.41.152.0/22", "198.41.192.0/19", "198.41.200.0/21", "198.41.208.0/23", "198.41.211.0/24", "198.41.212.0/24", "198.41.213.0/24", "198.41.214.0/23", "198.41.216.0/24", "198.41.220.0/23", "198.41.222.0/24", "198.41.223.0/24", "198.41.224.0/22", "198.41.228.0/22", "198.41.229.0/24", "198.41.230.0/24", "198.41.231.0/24", "198.41.232.0/23", "198.41.233.0/24", "198.41.234.0/24", "198.41.235.0/24", "198.41.236.0/22", "198.41.240.0/24", "198.41.241.0/24", "198.41.242.0/24", "198.41.243.0/24", "198.41.245.0/24", "198.41.246.0/23", "198.41.248.0/23", "198.41.250.0/24", "198.41.252.0/24", "198.41.254.0/24", "198.41.255.0/24", "199.21.96.0/22", "199.27.132.0/24", "23.227.38.0/23", "64.68.192.0/24", "66.235.200.0/24", "8.10.148.0/24", "8.14.199.0/24", "8.14.201.0/24", "8.14.202.0/24", "8.14.203.0/24", "8.14.204.0/24", "8.17.205.0/24", "8.17.206.0/24", "8.17.207.0/24", "8.18.113.0/24", "8.18.194.0/24", "8.18.195.0/24", "8.18.196.0/24", "8.18.50.0/24", "8.19.10.0/24", "8.19.8.0/24", "8.20.100.0/24", "8.20.101.0/24", "8.20.103.0/24", "8.20.122.0/24", "8.20.123.0/24", "8.20.124.0/24", "8.20.125.0/24", "8.20.126.0/24", "8.20.127.0/24", "8.20.253.0/24", "8.21.10.0/24", "8.21.11.0/24", "8.21.13.0/24", "8.21.8.0/24", "8.21.9.0/24", "8.31.160.0/24", "8.31.161.0/24", "8.35.211.0/24", "8.35.57.0/24", "8.35.58.0/24", "8.35.59.0/24", "8.36.216.0/24", "8.36.217.0/24", "8.36.218.0/24", "8.36.219.0/24", "8.36.220.0/24", "8.37.41.0/24", "8.37.43.0/24", "8.38.147.0/24", "8.38.148.0/24", "8.38.149.0/24", "8.38.172.0/24", "8.39.125.0/24", "8.39.126.0/24", "8.39.127.0/24", "8.39.18.0/24", "8.39.201.0/24", "8.39.202.0/24", "8.39.203.0/24", "8.39.204.0/24", "8.39.205.0/24", "8.39.206.0/24", "8.39.207.0/24", "8.39.212.0/24", "8.39.213.0/24", "8.39.214.0/24", "8.39.215.0/24", "8.39.6.0/24", "8.40.107.0/24", "8.40.111.0/24", "8.40.140.0/24", "8.40.26.0/24", "8.40.27.0/24", "8.40.28.0/24", "8.40.29.0/24", "8.40.30.0/24", "8.40.31.0/24", "8.41.36.0/24", "8.41.37.0/24", "8.41.39.0/24", "8.41.5.0/24", "8.41.6.0/24", "8.41.7.0/24", "8.42.161.0/24", "8.42.164.0/24", "8.42.172.0/24", "8.42.245.0/24", "8.42.51.0/24", "8.42.52.0/24", "8.42.54.0/24", "8.42.55.0/24", "8.43.121.0/24", "8.43.122.0/24", "8.43.123.0/24", "8.43.224.0/24", "8.43.225.0/24", "8.43.226.0/24", "8.44.0.0/24", "8.44.1.0/24", "8.44.2.0/24", "8.44.3.0/24", "8.44.58.0/24", "8.44.59.0/24", "8.44.6.0/24", "8.44.60.0/24", "8.44.61.0/24", "8.44.62.0/24", "8.44.63.0/24", "8.45.100.0/24", "8.45.101.0/24", "8.45.102.0/24", "8.45.108.0/24", "8.45.111.0/24", "8.45.144.0/24", "8.45.145.0/24", "8.45.146.0/24", "8.45.147.0/24", "8.45.151.0/24", "8.45.41.0/24", "8.45.42.0/24", "8.45.43.0/24", "8.45.44.0/24", "8.45.45.0/24", "8.45.46.0/24", "8.45.47.0/24", "8.45.97.0/24", "8.46.113.0/24", "8.46.114.0/24", "8.46.115.0/24", "8.46.116.0/24", "8.46.117.0/24", "8.46.118.0/24", "8.46.119.0/24", "8.47.12.0/24", "8.47.13.0/24", "8.47.14.0/24", "8.47.15.0/24", "8.47.69.0/24", "8.47.71.0/24", "8.47.9.0/24", "8.48.130.0/24", "8.48.131.0/24", "8.48.132.0/24", "8.48.133.0/24", "8.48.134.0/24", "8.6.112.0/24", "8.6.144.0/24", "8.6.145.0/24", "8.6.146.0/24", "8.9.230.0/24", "8.9.231.0/24"}
	for i := range cidrs {
		hosts, err := hosts(cidrs[i])
		if err != nil {
			continue
		}

		for _, host := range hosts {
			if host == ip.String() {
				return true
			}
		}
	}
	return false
}
