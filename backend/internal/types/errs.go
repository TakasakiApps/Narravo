package types

type (
	RuntimeException    any
	FileSystemException any
	IOException         any

	ArgValueNotValidException   any
	ServerAbortedException      any
	ServerMiddleWareException   any
	ServerUnauthorizedException any
	ServerBadRequestException   any
	ServerNotModifiedException  any

	RequestDataExpiredException any

	DecryptFailedException any

	JsonMarshalFailedException   any
	JsonUnmarshalFailedException any

	DataBaseConnectFailedException   any
	DataBaseOperationFailedException any

	JWTSignFailedException  any
	JWTParseFailedException any

	TencentCloudApiKeyOrSecretInvalidException any
	TTSModelNotFoundException                  any
	TencentCloudSDKErrorException              any
)
