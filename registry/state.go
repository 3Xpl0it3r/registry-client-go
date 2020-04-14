package registry


const (
	BLOB_UNKONWN_MESSAGE = "blob unkonow to registry:\t%s"
	BLOB_UNKONWN = "BLOB_UNKONWN"

	BLOB_UPLOAD_INVALID_MESSAGE = "blob upload invalid:\t%s"
	BLOB_UPLOAD_INVALID = "BLOB_UPLOAD_INVALID"

	BLOB_UPLOAD_UNKNOWN_MESSAGE = "blob upload unknown to registry\t%s"
	BLOB_UPLOAD_UNKNOWN= "BLOB_UPLOAD_UNKNOWN"

	DIGEST_INVALID_MESSAGE = "provided digest did not match uploaded content\t%s"
	DIGEST_INVALID = "DIGEST_INVALID"

	MANIFESTS_BLOB_UNKNOWN_MESSAGE = "blob unknown to registry:\t%s"
	MANIFESTS_BLOB_UNKNOWN = "MANIFESTS_BLOB_UNKNOWN"

	MANIFEST_INVALID_MESSAGE = "manifest invalid:\t%s"
	MANIFEST_INVALID = "MANIFEST_INVALID"

	MANIFEST_UNKNOWN_MESSAGE = "manifest unknown:\t%s"
	MANIFEST_UNKNOWN = "MANIFEST_UNKNOWN"

	MANIFEST_UNVERIFIED_MESSAGE = "manifest failed signature verification:\t%s"
	MANIFEST_UNVERIFIED = "MANIFEST_UNVERIFIED"

	NAME_INVALID = "NAME_INVALID"
	NAME_INVALID_MESSAGE = "invalid repository name:\t%s"

	NAME_UNKONOWN_MESSAGE = "repository name not known to registry:\t%s"
	NAME_UNKONOWN = "NAME_UNKONOWN"

	SIZE_INVALID_MESSAGE = "provided length did not match content lentth:\t%s"
	SIZE_INVALID = "SIZE_INVALID"

	TAG_INVALID_MESSAGE = "manifest tag did not match URI:\t%s"
	TAG_INVALID = "TAG_INVALID"

	UNAUTHORIZED_MESSAGE = "authentication required:\t%s"
	UNAUTHORIZED = "UNAUTHORIZED"

	DENIED_MESSAGE = "request access to the resource is denied:\t%s"
	DENIED = "DENIED"

	UNSUPPORTED_MESSAGE = "the operation is unsupported:\t%s"
	UNSUPPORTED = "UNSUPPORTED"



	// custom define extra error
	UNKNOWN = "UNKNOWN"
	UNKONWN_MESSAGE = "cannot get standard registry error:\t%s"

	INTERNAL_SERVER_ERROR_MESSAGE = "internal server error:\t%s"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"

	CLIENT_ERROR_MESSAGE = "client error:\t%s"
	CLIENT_ERROR = "CLIENT_ERROR"


	SUCCESS = "SUCCESS"
	SUCCESS_MESSAGE = "successfully:\t%s"
)
