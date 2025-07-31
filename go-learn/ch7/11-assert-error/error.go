// 基于类型断言区别错误类型
package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err)) // "true"
}

// IsNotExist

// func IsNotExist(err error) bool {
// 	return underlyingErrorIs(err, ErrNotExist)
// }

//
// func underlyingErrorIs(err, target error) bool {
// 	// Note that this function is not errors.Is:
// 	// underlyingError only unwraps the specific error-wrapping types
// 	// that it historically did, not all errors implementing Unwrap().
// 	err = underlyingError(err)
// 	if err == target {
// 		return true
// 	}
// 	// To preserve prior behavior, only examine syscall errors.
// 	e, ok := err.(syscallErrorType)
// 	return ok && e.Is(target)
// }

// underlyingError returns the underlying error for known os error types.
// func underlyingError(err error) error {
// 	switch err := err.(type) {
// 	case *PathError:
// 		return err.Err
// 	case *LinkError:
// 		return err.Err
// 	case *SyscallError:
// 		return err.Err
// 	}
// 	return err
// }

// 应该是在这里做了系统差异化处理
// C:/Program Files/Go/src/syscall/syscall_windows.go
// func (e Errno) Is(target error) bool {
// 	switch target {
// 	case oserror.ErrPermission:
// 		return e == ERROR_ACCESS_DENIED ||
// 			e == EACCES ||
// 			e == EPERM
// 	case oserror.ErrExist:
// 		return e == ERROR_ALREADY_EXISTS ||
// 			e == ERROR_DIR_NOT_EMPTY ||
// 			e == ERROR_FILE_EXISTS ||
// 			e == EEXIST ||
// 			e == ENOTEMPTY
// 	case oserror.ErrNotExist:
// 		return e == ERROR_FILE_NOT_FOUND ||
// 			e == _ERROR_BAD_NETPATH ||
// 			e == ERROR_PATH_NOT_FOUND ||
// 			e == ENOENT
// 	case errorspkg.ErrUnsupported:
// 		return e == _ERROR_NOT_SUPPORTED ||
// 			e == _ERROR_CALL_NOT_IMPLEMENTED ||
// 			e == ENOSYS ||
// 			e == ENOTSUP ||
// 			e == EOPNOTSUPP ||
// 			e == EWINDOWS
// 	}
// 	return false
// }
