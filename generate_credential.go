package utilities

import "fmt"

const (
	// RoleNameUser const for the name of the user role.
	RoleNameUser int = 1
	// RoleNameModerator const for the name of the moderator role.
	RoleNameModerator int = 2
	// RoleNameAdmin const for the name of the admin role.
	RoleNameAdmin int = 3
)

// GenerateCredential func for generating credential:
// <object>:<isOwn?>:<action>
func GenerateCredential(object, action string, isOwnCredential bool) string {
	credential := fmt.Sprintf("%s:%s", object, action)
	if isOwnCredential {
		// If this credential is own to user.
		credential = fmt.Sprintf("%s:own:%s", object, action)
	}
	return credential
}

// GenerateCredentialsByRole func for generating slice of the credentials,
// by given role number.
func GenerateCredentialsByRole(role int) ([]string, error) {
	var credentials []string
	switch role {
	case RoleNameAdmin:
		credentials = []string{
			// Any object:
			GenerateCredential("projects", "create", false),
			GenerateCredential("projects", "update", false),
			GenerateCredential("projects", "delete", false),
			GenerateCredential("tasks", "create", false),
			GenerateCredential("tasks", "update", false),
			GenerateCredential("tasks", "delete", false),
			GenerateCredential("answers", "create", false),
			GenerateCredential("answers", "update", false),
			GenerateCredential("answers", "delete", false),
			GenerateCredential("user_password", "update", false),
			GenerateCredential("user_attrs", "update", false),
			// Is own object:
			GenerateCredential("projects", "update", true),
			GenerateCredential("projects", "delete", true),
			GenerateCredential("tasks", "update", true),
			GenerateCredential("tasks", "delete", true),
			GenerateCredential("answers", "update", true),
			GenerateCredential("answers", "delete", true),
			GenerateCredential("user_password", "update", true),
			GenerateCredential("user_attrs", "update", true),
		}
	case RoleNameModerator:
		credentials = []string{
			// Any object:
			GenerateCredential("projects", "create", false),
			GenerateCredential("projects", "update", false),
			GenerateCredential("tasks", "create", false),
			GenerateCredential("tasks", "update", false),
			GenerateCredential("answers", "create", false),
			GenerateCredential("answers", "update", false),
			GenerateCredential("user_attrs", "update", false),
			// Is own object:
			GenerateCredential("projects", "update", true),
			GenerateCredential("projects", "delete", true),
			GenerateCredential("tasks", "update", true),
			GenerateCredential("tasks", "delete", true),
			GenerateCredential("answers", "update", true),
			GenerateCredential("answers", "delete", true),
			GenerateCredential("user_password", "update", true),
			GenerateCredential("user_attrs", "update", true),
		}
	case RoleNameUser:
		credentials = []string{
			// Any object:
			GenerateCredential("projects", "create", false),
			GenerateCredential("tasks", "create", false),
			GenerateCredential("answers", "create", false),
			// Is own object:
			GenerateCredential("projects", "update", true),
			GenerateCredential("projects", "delete", true),
			GenerateCredential("tasks", "update", true),
			GenerateCredential("tasks", "delete", true),
			GenerateCredential("answers", "update", true),
			GenerateCredential("user_password", "update", true),
			GenerateCredential("user_attrs", "update", true),
		}
	default:
		return credentials, fmt.Errorf(GenerateErrorMessage(400, "role", fmt.Sprint(role)))
	}
	return credentials, nil
}
