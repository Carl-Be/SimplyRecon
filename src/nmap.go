package main

import (
	"log"
	"os/exec"
	"sync"
)

var TcpOutfile string = "nmap_tcp_all"
var UdpOutfile string = "nmap_udp_top_100"
var LootLoc string = "../loot/nmap"

func WrapTcpScan(wg *sync.WaitGroup) {
	nmap_tcp := exec.Command("nmap", "-sV", "-sC", "-p-", "-oA", TcpOutfile, "0.0.0.0", "--min-rate", "1250")
	nmap_tcp.Dir = LootLoc

	if err := nmap_tcp.Run(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func WrapUdpScan(wg *sync.WaitGroup) {
	nmap_udp := exec.Command("nmap", "-sU", "--top-port", "100", "-oA", UdpOutfile, "0.0.0.0", "--min-rate", "1250")
	nmap_udp.Dir = LootLoc

	if err := nmap_udp.Run(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func WrapXsltproc(wg *sync.WaitGroup, targetFile string) {
	xsltproc := exec.Command("xsltproc", targetFile+".xml", "-o", targetFile+".html")
	xsltproc.Dir = LootLoc

	if err := xsltproc.Run(); err != nil {
		log.Fatal(err)
	}

	wg.Done()
}

func StartNmapScanning(wg *sync.WaitGroup) {

	var childWg sync.WaitGroup
	childWg.Add(2)
	go WrapTcpScan(&childWg)
	go WrapUdpScan(&childWg)
	childWg.Wait()

	childWg.Add(2)
	go WrapXsltproc(&childWg, TcpOutfile)
	go WrapXsltproc(&childWg, UdpOutfile)
	childWg.Wait()

	wg.Done()
}
