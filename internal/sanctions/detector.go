package sanctions

import (
	"log"
	"sync"
)

// Detector is responsible for detecting if a given address is sanctioned.
type Detector struct {
	SanctionedAddresses map[string]struct{} // We can replace the storage to Database for scalability.
	mu                  sync.RWMutex
}

// NewDetector creates a new Detector instance with an initial list of sanctioned addresses.
func NewDetector(initialAddresses []string) *Detector {
	log.Println("Initializing sanctions detector")
	detector := &Detector{
		SanctionedAddresses: make(map[string]struct{}),
	}

	for _, addr := range initialAddresses {
		detector.AddAddress(addr)
	}

	log.Printf("Sanctions detector initialized with %d addresses\n", len(initialAddresses))

	return detector
}

// AddAddress adds a new address to the sanctioned list.
func (d *Detector) AddAddress(address string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.SanctionedAddresses[address] = struct{}{}
	log.Printf("Added sanctioned address: %s\n", address)
}

// RemoveAddress removes an address from the sanctioned list.
func (d *Detector) RemoveAddress(address string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.SanctionedAddresses, address)
	log.Printf("Removed sanctioned address: %s\n", address)
}

// IsSanctioned checks if a given address is sanctioned.
func (d *Detector) IsSanctioned(address string) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	_, exists := d.SanctionedAddresses[address]
	if exists {
		log.Printf("Address %s is sanctioned\n", address)
	} else {
		log.Printf("Address %s is not sanctioned\n", address)
	}
	return exists
}
