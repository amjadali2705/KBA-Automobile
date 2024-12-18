package main

type Config struct {
	CertPath     string `json:"certPath"`
	KeyDirectory string `json:"keyPath"`
	TLSCertPath  string `json:"tlsCertPath"`
	PeerEndpoint string `json:"peerEndpoint"`
	GatewayPeer  string `json:"gatewayPeer"`
	MSPID        string `json:"mspID"`
}

var profile = map[string]Config{

	"manufacturer": {
		CertPath:     "../Automobile-network/organizations/peerOrganizations/manufacturer.auto.com/users/User1@manufacturer.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Automobile-network/organizations/peerOrganizations/manufacturer.auto.com/users/User1@manufacturer.auto.com/msp/keystore/",
		TLSCertPath:  "../Automobile-network/organizations/peerOrganizations/manufacturer.auto.com/peers/peer0.manufacturer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.manufacturer.auto.com",
		MSPID:        "ManufacturerMSP",
	},

	"dealer": {
		CertPath:     "../Automobile-network/organizations/peerOrganizations/dealer.auto.com/users/User1@dealer.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Automobile-network/organizations/peerOrganizations/dealer.auto.com/users/User1@dealer.auto.com/msp/keystore/",
		TLSCertPath:  "../Automobile-network/organizations/peerOrganizations/dealer.auto.com/peers/peer0.dealer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:9051",
		GatewayPeer:  "peer0.dealer.auto.com",
		MSPID:        "DealerMSP",
	},

	"mvd": {
		CertPath:     "../Automobile-network/organizations/peerOrganizations/mvd.auto.com/users/User1@mvd.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Automobile-network/organizations/peerOrganizations/mvd.auto.com/users/User1@mvd.auto.com/msp/keystore/",
		TLSCertPath:  "../Automobile-network/organizations/peerOrganizations/mvd.auto.com/peers/peer0.mvd.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:11051",
		GatewayPeer:  "peer0.mvd.auto.com",
		MSPID:        "MvdMSP",
	},

	"manufacturer2": {
		CertPath:     "../Automobile-network/organizations/peerOrganizations/manufacturer.auto.com/users/User2@manufacturer.auto.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Automobile-network/organizations/peerOrganizations/manufacturer.auto.com/users/User2@manufacturer.auto.com/msp/keystore/",
		TLSCertPath:  "../Automobile-network/organizations/peerOrganizations/manufacturer.auto.com/peers/peer0.manufacturer.auto.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.manufacturer.auto.com",
		MSPID:        "ManufacturerMSP",
	},
}