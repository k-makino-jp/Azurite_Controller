package azuritectl

func isSuccessfulStatusCode(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
