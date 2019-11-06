package accessor

import (
	"fmt"
	"strings"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/dgrijalva/jwt-go"
)

//go:generate counterfeiter . Access

type Access interface {
	HasToken() bool
	IsAuthenticated() bool
	IsAuthorized(string) bool
	IsAdmin() bool
	IsSystem() bool
	TeamNames() []string
	UserName() string
}

type access struct {
	*jwt.Token
	action            string
	actionRoleMap     ActionRoleMap
	adminClaimKey     string
	adminClaimValues  map[interface{}]bool
	systemClaimKey    string
	systemClaimValues map[interface{}]bool
	teamFactory       db.TeamFactory
}

func NewAccessor(
	token *jwt.Token,
	action string,
	teamFactory db.TeamFactory,
	actionRoleMap ActionRoleMap,
) *access {
	return &access{
		Token:             token,
		action:            action,
		adminClaimKey:     "sub",
		adminClaimValues:  map[interface{}]bool{},
		systemClaimKey:    "sub",
		systemClaimValues: map[interface{}]bool{"CgljbGllbnQtaWQSBmNsaWVudA": true},
		teamFactory:       teamFactory,
		actionRoleMap:     actionRoleMap,
	}
}

func (a *access) HasToken() bool {
	return a.Token != nil
}

func (a *access) IsAuthenticated() bool {
	return a.HasToken() && a.Token.Valid
}

func (a *access) Claims() jwt.MapClaims {
	if a.IsAuthenticated() {
		if claims, ok := a.Token.Claims.(jwt.MapClaims); ok {
			return claims
		}
	}
	return jwt.MapClaims{}
}

func (a *access) IsAuthorized(teamName string) bool {

	if a.IsAdmin() {
		return true
	}

	team, found, err := a.teamFactory.FindTeam(teamName)
	if err != nil {
		return false
	}

	if !found {
		return false
	}

	for _, teamRole := range a.TeamRoles(team.Auth()) {
		if a.hasPermission(teamRole) {
			return true
		}
	}

	return false
}

func (a *access) TeamNames() []string {

	teamNames := []string{}

	teams, err := a.teamFactory.GetTeams()
	if err != nil {
		return teamNames
	}

	for _, team := range teams {
		for _, teamRole := range a.TeamRoles(team.Auth()) {
			if a.hasPermission(teamRole) {
				teamNames = append(teamNames, team.Name())
			}
		}
	}

	return teamNames
}

func (a *access) TeamRoles(auth atc.TeamAuth) []string {

	roles := []string{}

	groups := a.Claims()["groups"]
	connectorID := a.FederatedClaims()["connector_id"]
	userID := a.FederatedClaims()["user_id"]
	userName := a.FederatedClaims()["username"]

	for role, auth := range auth {
		userAuth := auth["users"]
		groupAuth := auth["groups"]

		// backwards compatibility for allow-all-users
		if len(userAuth) == 0 && len(groupAuth) == 0 {
			roles = append(roles, role)
		}

		for _, user := range userAuth {
			if userID != "" {
				if strings.EqualFold(user, fmt.Sprintf("%v:%v", connectorID, userID)) {
					roles = append(roles, role)
				}
			}
			if userName != "" {
				if strings.EqualFold(user, fmt.Sprintf("%v:%v", connectorID, userName)) {
					roles = append(roles, role)
				}
			}
		}

		if claimGroups, ok := groups.([]string); ok {
			for _, group := range groupAuth {
				for _, claimGroup := range claimGroups {
					if claimGroup != "" {
						if strings.EqualFold(group, fmt.Sprintf("%v:%v", connectorID, claimGroup)) {
							roles = append(roles, role)
						}
					}
				}
			}
		}
	}

	return roles
}

func (a *access) hasPermission(role string) bool {
	switch a.actionRoleMap.RoleOfAction(a.action) {
	case "owner":
		return role == "owner"
	case "member":
		return role == "owner" || role == "member"
	case "pipeline-operator":
		return role == "owner" || role == "member" || role == "pipeline-operator"
	case "viewer":
		return role == "owner" || role == "member" || role == "pipeline-operator" || role == "viewer"
	default:
		return false
	}
}

func (a *access) FederatedClaims() map[string]interface{} {
	if claims, ok := a.Claims()["federated_claims"]; ok {
		parsed, ok := claims.(map[string]interface{})
		if ok {
			return parsed
		}
	}

	return map[string]interface{}{}
}

func (a *access) ConnectorID() string {
	var connectorID string
	if connectorClaim, ok := a.FederatedClaims()["connector_id"]; ok {
		connectorID, _ = connectorClaim.(string)
	}
	return connectorID
}

func (a *access) UserID() string {
	var userID string
	if userClaim, ok := a.FederatedClaims()["user_id"]; ok {
		userID, _ = userClaim.(string)
	}
	return userID
}

func (a *access) UserName() string {
	var userName string
	if userClaim, ok := a.FederatedClaims()["user_name"]; ok {
		userName, _ = userClaim.(string)
	}
	return userName
}

func (a *access) IsAdmin() bool {
	claim, ok := a.Claims()[a.adminClaimKey]
	return ok && a.adminClaimValues[claim]
}

func (a *access) IsSystem() bool {
	claim, ok := a.Claims()[a.systemClaimKey]
	return ok && a.systemClaimValues[claim]
}
