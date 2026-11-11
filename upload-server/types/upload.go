package types

const (
	// ClientHeader is the header used to identify the client
	ClientHeader = "x-nb-client"
	// ClientHeaderValue is the value of the ClientHeader
	ClientHeaderValue = "netbird"
	// GetURLPath is the path for the GetURL request
	GetURLPath = "/upload-url"

	// DefaultBundleURL is the KilatMesh debug-bundle upload endpoint of
	// PT Kilat Cloud Indonesia. Placeholder on the company domain: it must
	// be provisioned (or overridden via --upload-bundle-url) before the
	// bundle upload feature is used; school debug bundles must never be
	// sent to third-party upload services by default.
	DefaultBundleURL = "https://upload.debug.kilat-cloud.com" + GetURLPath
)

// GetURLResponse is the response for the GetURL request
type GetURLResponse struct {
	URL string
	Key string
}
