package pkg

import "errors"

// error message conventions
var (
	// ErrUserNotFound error user does not exist
	ErrUserNotFound = errors.New("User not found")

	// ErrMenteeNotFound error mentee does not exist
	ErrMenteeNotFound = errors.New("Mentee not found")

	// ErrMentorNotFound error mentor does not exist
	ErrMentorNotFound = errors.New("Mentor not found")

	// ErrCategoryNotFound error category does not exist
	ErrCategoryNotFound = errors.New("Category not found")

	// ErrCourseNotFound error course does not exist
	ErrCourseNotFound = errors.New("Course not found")

	// ErrModuleNotFound error module does not exist
	ErrModuleNotFound = errors.New("Module not found")

	// ErrMaterialNotFound error material does not exist
	ErrMaterialNotFound = errors.New("Material not found")

	// ErrMaterialAssetNotFound error material asset does not exist
	ErrMaterialAssetNotFound = errors.New("Material asset not found")

	// ErrAssignmentNotFound error assignment does not exist
	ErrAssignmentNotFound = errors.New("Assignment not found")

	// ErrEmailAlreadyExist error email already exist
	ErrEmailAlreadyExist = errors.New("Email already exist")

	// ErrPasswordLengthInvalid error invalid password length
	ErrPasswordLengthInvalid = errors.New("Password must be min. 6 characters long")

	// ErrPasswordNotMatch error both password not match
	ErrPasswordNotMatch = errors.New("Both password not matched")

	// ErrOTPExpired error otp expired
	ErrOTPExpired = errors.New("OTP has been expired")

	// ErrOTPNotMatch error OTP not match with server
	ErrOTPNotMatch = errors.New("OTP not match")

	// ErrAccessForbidden error access forbidden
	ErrAccessForbidden = errors.New("Access forbidden")

	// ErrUserUnauthorized error user unauthorized
	ErrUserUnauthorized = errors.New("User unauthorized")

	// ErrInvalidRequest error invalid request body
	ErrInvalidRequest = errors.New("Invalid request body")

	// ErrInvalidJWTPayload error invalid JWT payloads
	ErrInvalidJWTPayload = errors.New("Invalid JWT payloads")

	// error unsupported file upload
	ErrUnsupportedAssignmentFile = errors.New("Unsupported assignment file. Acceptable file format: .pdf")

	// error invalid token header
	ErrInvalidTokenHeader = errors.New("Invalid token header")

	// error unsupported video file
	ErrUnsupportedVideoFile = errors.New("Unsupported video file. Acceptable file format: .mp4 or .mkv")

	// error unsupported image file
	ErrUnsupportedImageFile = errors.New("Unsupported image file. Acceptable file format: .jpeg, .jpg, and .png")

	// error record not found (cannot specify the error)
	ErrRecordNotFound = errors.New("Record not found")

	// ErrInternalServerError error internal server error
	ErrInternalServerError = errors.New("Internal server error")
)
