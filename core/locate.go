package core

// Locators embody the separation of storage mechanism from key.
// An immediate benefit is decoupling from a single specific architecture, but
// ... additionally this allows optimizations like layering based on desired access time
// ... (e.g. memory -> local storage OR S3)

// To locate data (whether in AWS, etc.)
type Locator struct {
	Locate (string)
}
