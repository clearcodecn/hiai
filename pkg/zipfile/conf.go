package zipfile

var conf = `{
  "logEnabled": false,
  "loglevel": "warning",
  "indexId": "4653921950098877425",
  "muxEnabled": false,
  "sysProxyType": 1,
  "enableStatistics": false,
  "keepOlderDedupl": false,
  "statisticsFreshRate": 1,
  "remoteDNS": null,
  "domainStrategy4Freedom": null,
  "defAllowInsecure": false,
  "domainStrategy": "IPIfNonMatch",
  "domainMatcher": null,
  "routingIndex": 0,
  "enableRoutingAdvanced": true,
  "ignoreGeoUpdateCore": false,
  "systemProxyExceptions": null,
  "systemProxyAdvancedProtocol": null,
  "autoUpdateInterval": 0,
  "autoUpdateSubInterval": 0,
  "checkPreReleaseUpdate": false,
  "enableSecurityProtocolTls13": false,
  "trayMenuServersLimit": 30,
  "inbound": [
    {
      "localPort": 10808,
      "protocol": "socks",
      "udpEnabled": true,
      "sniffingEnabled": true,
      "allowLANConn": false,
      "user": null,
      "pass": null
    }
  ],
  "vmess": [
    {
      "indexId": "4653921950098877425",
      "configType": 1,
      "configVersion": 2,
      "sort": 0,
      "address": "proxy.x.hi-ai.top",
      "port": 443,
      "id": "%s",
      "alterId": 0,
      "security": "auto",
      "network": "ws",
      "remarks": "proxy.x.hi-ai.top",
      "headerType": "none",
      "requestHost": "proxy.x.hi-ai.top",
      "path": "/puppet",
      "streamSecurity": "tls",
      "allowInsecure": "False",
      "testResult": "",
      "subid": "",
      "flow": "",
      "sni": "",
      "alpn": [],
      "groupId": "",
      "coreType": null,
      "preSocksPort": 0,
      "fingerprint": null
    }
  ],
  "kcpItem": {
    "mtu": 1350,
    "tti": 50,
    "uplinkCapacity": 12,
    "downlinkCapacity": 100,
    "congestion": false,
    "readBufferSize": 2,
    "writeBufferSize": 2
  },
  "subItem": [
    {
      "id": "5396092270645152360",
      "remarks": "remarks",
      "url": "https://hneko.xyz/api/v1/client/subscribe?token=6faea4ba221dc802f87e75f0b75c024e",
      "enabled": true,
      "userAgent": "",
      "groupId": ""
    }
  ],
  "uiItem": {
    "enableAutoAdjustMainLvColWidth": true,
    "mainLocation": "476, 204",
    "mainSize": "968, 632",
    "mainLvColWidth": {
      "def": 25,
      "configType": 182,
      "remarks": 283,
      "address": 319,
      "port": 47,
      "security": 60,
      "network": 60,
      "streamSecurity": 72,
      "subRemarks": 61,
      "testResult": 188
    }
  },
  "routings": [
    {
      "remarks": "绕过大陆(Whitelist)",
      "url": "",
      "rules": [
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "direct",
          "ip": null,
          "domain": [
            "domain:example-example.com",
            "domain:example-example2.com"
          ],
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "block",
          "ip": null,
          "domain": [
            "geosite:category-ads-all"
          ],
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "direct",
          "ip": null,
          "domain": [
            "geosite:cn"
          ],
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "direct",
          "ip": [
            "geoip:private",
            "geoip:cn"
          ],
          "domain": null,
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": "0-65535",
          "inboundTag": null,
          "outboundTag": "proxy",
          "ip": null,
          "domain": null,
          "protocol": null,
          "enabled": true
        }
      ],
      "enabled": true,
      "locked": false,
      "customIcon": null
    },
    {
      "remarks": "黑名单(Blacklist)",
      "url": "",
      "rules": [
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "direct",
          "ip": null,
          "domain": null,
          "protocol": [
            "bittorrent"
          ],
          "enabled": true
        },
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "block",
          "ip": null,
          "domain": [
            "geosite:category-ads-all"
          ],
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "proxy",
          "ip": [
            "geoip:telegram"
          ],
          "domain": [
            "geosite:gfw",
            "geosite:greatfire",
            "geosite:tld-!cn"
          ],
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": "0-65535",
          "inboundTag": null,
          "outboundTag": "direct",
          "ip": null,
          "domain": null,
          "protocol": null,
          "enabled": true
        }
      ],
      "enabled": true,
      "locked": false,
      "customIcon": null
    },
    {
      "remarks": "全局(Global)",
      "url": "",
      "rules": [
        {
          "type": null,
          "port": "0-65535",
          "inboundTag": null,
          "outboundTag": "proxy",
          "ip": null,
          "domain": null,
          "protocol": null,
          "enabled": true
        }
      ],
      "enabled": true,
      "locked": false,
      "customIcon": null
    },
    {
      "remarks": "locked",
      "url": "",
      "rules": [
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "proxy",
          "ip": null,
          "domain": [
            "geosite:google"
          ],
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "direct",
          "ip": null,
          "domain": [
            "domain:example-example.com",
            "domain:example-example2.com"
          ],
          "protocol": null,
          "enabled": true
        },
        {
          "type": null,
          "port": null,
          "inboundTag": null,
          "outboundTag": "block",
          "ip": null,
          "domain": [
            "geosite:category-ads-all"
          ],
          "protocol": null,
          "enabled": true
        }
      ],
      "enabled": true,
      "locked": true,
      "customIcon": null
    }
  ],
  "constItem": {
    "speedTestUrl": "http://cachefly.cachefly.net/10mb.test",
    "speedPingTestUrl": "https://www.google.com/generate_204",
    "defIEProxyExceptions": "localhost;127.*;10.*;172.16.*;172.17.*;172.18.*;172.19.*;172.20.*;172.21.*;172.22.*;172.23.*;172.24.*;172.25.*;172.26.*;172.27.*;172.28.*;172.29.*;172.30.*;172.31.*;192.168.*"
  },
  "globalHotkeys": null,
  "groupItem": [],
  "coreTypeItem": [
    {
      "configType": 1,
      "coreType": 2
    },
    {
      "configType": 2,
      "coreType": 2
    },
    {
      "configType": 3,
      "coreType": 2
    },
    {
      "configType": 4,
      "coreType": 2
    },
    {
      "configType": 5,
      "coreType": 2
    },
    {
      "configType": 6,
      "coreType": 2
    }
  ]
}`
