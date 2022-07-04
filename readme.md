# V2 docs
- https://aws.github.io/aws-sdk-go-v2/docs/
- https://pkg.go.dev/github.com/aws/aws-sdk-go-v2#section-readme

# Error Handle

- 文件: https://aws.github.io/aws-sdk-go-v2/docs/handling-errors/
- 定義: https://github.com/aws/smithy-go/blob/main/errors.go

## Service Client Errors
- OperationError provides contextual information about the service name and operation that is associated with an underlying error.
- This information can be useful for applications that perform `batches of operations` to one or more services, with a centralized error handling mechanism.

```
import "log"
import "github.com/aws/smithy-go"

// ...

if err != nil {
	var oe *smithy.OperationError
	if errors.As(err, &oe) {
		log.Printf("failed to call service: %s, operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap())
    }
    return
}
```

## 判斷指定 error

```
import "log"
import "github.com/aws/aws-sdk-go-v2/service/s3/types"

// ...

if err != nil {
	var bne *types.BucketAlreadyExists
	if errors.As(err, &bne) {
		log.Println("error:", bne)
    }
    return
}
```

## APIError
- This interface can be used to handle both modeled or un-modeled service error responses.
- Additionally, this type provides indication of whether the fault of the error was due to the client or server `if known`.

```
import "log"
import "github.com/aws/smithy-go"

// ...

if err != nil {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		log.Printf("code: %s, message: %s, fault: %s", ae.ErrorCode(), ae.ErrorMessage(), ae.ErrorFault().String())
    }
    return
}
```