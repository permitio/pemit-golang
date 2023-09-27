package enforcement

import "strings"

const (
	DefaultTenant  = "default"
	DefaultTimeout = 30
	AllowKey       = "allow"
)

const (
	reqMethod           = "POST"
	reqContentTypeKey   = "Content-Type"
	reqContentTypeValue = "application/json"
	reqAuthKey          = "Authorization"
)

type packageName string
type sidecarPath string

const (
	mainPolicyPackage       packageName = "permit.root"
	bulkPolicyPackage       packageName = "permit.bulk"
	allTenantsPolicyPackage packageName = "permit.all_tenants"
	userPermissionsPackage  packageName = "permit.user_permissions"
)

const (
	mainPolicy       sidecarPath = "/allowed"
	bulkPolicy       sidecarPath = "/allowed/bulk"
	allTenantsPolicy sidecarPath = "/allowed/all-tenants"
	userPermissions  sidecarPath = "/user-permissions"
)

type checkOperationConfig struct {
	sidecarPath sidecarPath
	opaPath     string
}

var policyMap = map[packageName]checkOperationConfig{
	mainPolicyPackage: {
		sidecarPath: mainPolicy,
		opaPath:     strings.Replace(string(mainPolicyPackage), ".", "/", -1),
	},
	bulkPolicyPackage: {
		sidecarPath: bulkPolicy,
		opaPath:     strings.Replace(string(bulkPolicyPackage), ".", "/", -1),
	},
	allTenantsPolicyPackage: {
		sidecarPath: allTenantsPolicy,
		opaPath:     strings.Replace(string(allTenantsPolicyPackage), ".", "/", -1),
	},
	userPermissionsPackage: {
		sidecarPath: userPermissions,
		opaPath:     strings.Replace(string(userPermissionsPackage), ".", "/", -1),
	},
}
