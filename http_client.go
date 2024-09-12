//go:build ignore

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package main

import (
	"io"
	"log"
	"net/http"
	"net/netip"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun/netstack"
)

func main() {
	tun, tnet, err := netstack.CreateNetTUN(
		[]netip.Addr{netip.MustParseAddr("172.16.0.2")},
		[]netip.Addr{netip.MustParseAddr("8.8.8.8")},
		1420)
	if err != nil {
		log.Panic(err)
	}
	dev := device.NewDevice(tun, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, ""))
	err = dev.IpcSet(`private_key=20345982adfd56854bd7c960d6730bad46bb666660132a44cf503c83231c1559
public_key=6e65ce0be17517110c17d77288ad87e7fd5252dcc7d09b95a39d61db03df832a
endpoint=162.159.192.1:2408
allowed_ip=0.0.0.0/0
`)
	err = dev.Up()
	if err != nil {
		log.Panic(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			DialContext: tnet.DialContext,
		},
	}
	resp, err := client.Get("https://ip.sb/cdn-cgi/trace")
	if err != nil {
		log.Panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(body))
}
