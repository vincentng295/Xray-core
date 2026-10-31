package minimal

import (
	// Core dispatcher & proxyman
	_ "github.com/xtls/xray-core/app/dispatcher"
	_ "github.com/xtls/xray-core/app/proxyman/inbound"
	_ "github.com/xtls/xray-core/app/proxyman/outbound"

	// Tun and Freedom
	_ "github.com/xtls/xray-core/proxy/tun"
	_ "github.com/xtls/xray-core/proxy/freedom"

	// minimal transports
	_ "github.com/xtls/xray-core/transport/internet/tcp"
	_ "github.com/xtls/xray-core/transport/internet/udp"
)