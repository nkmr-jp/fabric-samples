package fabric

import "path/filepath"

var (
	// connection profile path
	ccpPath = filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
	// credentials path
	// See: https://hyperledger-fabric.readthedocs.io/en/release-2.2/membership/membership.html#msp-structure
	// See: https://hyperledger-fabric.readthedocs.io/en/release-2.2/test_network.html#bring-up-the-network-with-certificate-authorities
	credPath = filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)
	// certificate pem
	certPath = filepath.Join(credPath, "signcerts", "cert.pem")
	// private key dir
	keyDir = filepath.Join(credPath, "keystore")
)
