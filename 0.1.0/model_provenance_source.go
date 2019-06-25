/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Source describes the location of the source used for the build.
type ProvenanceSource struct {
	// If provided, the input binary artifacts for the build came from this location.
	ArtifactStorageSourceUri string `json:"artifact_storage_source_uri,omitempty"`
	// Hash(es) of the build source, which can be used to verify that the original source integrity was maintained in the build.  The keys to this map are file paths used as build source and the values contain the hash values for those files.  If the build source came in a single package such as a gzipped tarfile (.tar.gz), the FileHash will be for the single path to that file.
	FileHashes map[string]ProvenanceFileHashes `json:"file_hashes,omitempty"`
	// If provided, the source code used for the build came from this location.
	Context *SourceSourceContext `json:"context,omitempty"`
	// If provided, some of the source code used for the build may be found in these locations, in the case where the source repository had multiple remotes or submodules. This list will not include the context specified in the context field.
	AdditionalContexts []SourceSourceContext `json:"additional_contexts,omitempty"`
}