package main

import (
	"github.com/Nik-U/ringsig"
	"github.com/Nik-U/ringsig/shacham"
)

func main() {
	// Generate new scheme parameters
	scheme, _ := shacham.New(2)

	// In the real world, these parameters would be saved using scheme.WriteTo
	// and then loaded by the clients using shacham.Load.

	// Two clients generate key pairs
	alice := scheme.KeyGen()
	bob := scheme.KeyGen()

	// We will sign over the ring of the two users. In general, higher level
	// communication protocols will somehow specify the ring used to sign the
	// message (either explicitly or implicitly).
	ring := []ringsig.PublicKey{alice.Public, bob.Public}

	// Alice signs a message intended for Bob
	sig, _ := scheme.Sign("message", ring, alice)

	// Bob verifies the signature
	if scheme.Verify("message", sig, ring) {
		// Both Alice and Bob are now convinced that Alice signed the message.
		// However, nobody else can be convinced of this cryptographically.
		fmt.Println("this is alice message")
	}
}