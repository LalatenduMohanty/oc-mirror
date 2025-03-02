package release

const (
	graphBaseImage              = "registry.access.redhat.com/ubi9/ubi:latest"
	graphURL                    = "https://api.openshift.com/api/upgrades_info/graph-data"
	graphArchive                = "cincinnati-graph-data.tar"
	graphPreparationDir         = "graph-preparation"
	graphDataDir                = "/var/lib/cincinnati-graph-data"
	graphDataMountPath          = "/var/lib/cincinnati/graph-data"
	graphImageName              = "openshift/graph-image"
	indexJson                   = "manifest.json"
	operatorImageExtractDir     = "hold-operator"
	workingDir                  = "working-dir"
	dockerProtocol              = "docker://"
	ociProtocol                 = "oci://"
	ociProtocolTrimmed          = "oci:"
	dirProtocol                 = "dir://"
	dirProtocolTrimmed          = "dir:"
	releaseImageDir             = "release-images"
	releaseIndex                = "release-index"
	releaseFiltersDir           = "release-filters"
	operatorImageDir            = "operator-images"
	releaseImageExtractDir      = "hold-release"
	releaseManifests            = "release-manifests"
	imageReferences             = "image-references"
	releaseImageExtractFullPath = releaseManifests + "/" + imageReferences
	blobsDir                    = "blobs/sha256" // TODO blobsDir should not take assumptions about algorithm
	errMsg                      = "[ReleaseImageCollector] %v "
	logFile                     = "logs/release.log"
)
