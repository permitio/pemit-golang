package tests

import (
	"context"
	"github.com/permitio/permit-golang/pkg/config"
	"github.com/permitio/permit-golang/pkg/enforcement"
	"github.com/permitio/permit-golang/pkg/errors"
	"github.com/permitio/permit-golang/pkg/models"
	"github.com/permitio/permit-golang/pkg/permit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	logger := zap.NewExample()
	ctx := context.Background()
	const userKey = "test-user3"
	const resourceKey = "document3"
	const roleKey = "editor3"
	permitContext := config.NewPermitContext(config.EnvironmentAPIKeyLevel, "test", "staging")
	permitClient := permit.New(config.NewConfigBuilder("").WithContext(permitContext).WithLogger(logger).Build())

	// Create a user
	newUser := *models.NewUserCreate(userKey)
	newUser.SetFirstName("tesasdt")
	_, err := permitClient.SyncUser(ctx, newUser)
	assert.NoError(t, err)

	// Create a Proxy Config
	mappingRules := []models.MappingRule{}
	action := "read"
	mappingRules = append(mappingRules, models.MappingRule{
		Url:        "https://asdfasdf.com",
		HttpMethod: "delete",
		Resource:   "document3",
		Action:     &action,
	})
	secret := "asdf:asdfasdf"
	proxyConfigCreate := *models.NewProxyConfigCreate(secret, "pxcf", "proxy")
	proxyConfigCreate.SetMappingRules(mappingRules)
	_, err = permitClient.Api.ProxyConfigs.Create(ctx, proxyConfigCreate)
	assert.ErrorContains(t, err, string(errors.ConflictMessage))
	proxyConfigUpdate := models.NewProxyConfigUpdate()
	mappingRules = append(mappingRules, models.MappingRule{
		Url:        "https://mushy.com",
		HttpMethod: "post",
		Resource:   "document3",
		Action:     &action,
	})
	authMechanism := models.BASIC
	proxyConfigUpdate.SetAuthMechanism(authMechanism)
	proxyConfigUpdate.SetSecret(secret)
	proxyConfigUpdate.SetMappingRules(mappingRules)
	_, err = permitClient.Api.ProxyConfigs.Update(ctx, "pxcf", *proxyConfigUpdate)
	// Create a resource
	_, err = permitClient.Api.Resources.Create(ctx, *models.NewResourceCreate(resourceKey, resourceKey, map[string]models.ActionBlockEditable{"read": {}, "write": {}}))
	assert.ErrorContains(t, err, string(errors.ConflictMessage))

	// Create a role
	permissions := []string{resourceKey + ":read", resourceKey + ":write"}
	roleCreate := models.NewRoleCreate(roleKey, roleKey)
	roleCreate.SetPermissions(permissions)
	_, err = permitClient.Api.Roles.Create(ctx, *roleCreate)
	assert.ErrorContains(t, err, string(errors.ConflictMessage))

	// Assign role to user
	_, err = permitClient.Api.Users.AssignRole(ctx, userKey, roleKey, "default")
	assert.ErrorContains(t, err, string(errors.ConflictMessage))

	// Check if user has permission
	time.Sleep(6 * time.Second)

	userCheck := enforcement.UserBuilder(userKey).Build()
	resourceCheck := enforcement.ResourceBuilder(resourceKey).WithTenant("default").Build()
	allowed, err := permitClient.Check(userCheck, "read", resourceCheck)
	assert.NoError(t, err)
	assert.True(t, allowed)

	// Elements
	a, err := permitClient.Elements.LoginAs(ctx, *models.NewUserLoginRequestInput(userKey, "default"))
	assert.NoError(t, err)
	assert.Equal(t, userKey, a.RedirectUrl)
}
