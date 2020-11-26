package fabric

import "path/filepath"

var (
	ccpPath = filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
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
	certPath = filepath.Join(credPath, "signcerts", "cert.pem")
	keyDir   = filepath.Join(credPath, "keystore")
)
