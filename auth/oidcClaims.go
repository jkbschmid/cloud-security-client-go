// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Cloud Security Client Go contributors
//
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	jwtgo "github.com/dgrijalva/jwt-go/v4"
)

// https://www.iana.org/assignments/jwt/jwt.xhtml#claims
const (
	propKeyID = "kid"
	propAlg   = "alg"
)

type OIDCClaims struct {
	jwtgo.StandardClaims
	UserName   string `json:"user_name,omitempty"`
	GivenName  string `json:"first_name,omitempty"`
	FamilyName string `json:"last_name,omitempty"`
	Email      string `json:"mail,omitempty"`
	ZoneID     string `json:"zone_uuid,omitempty"`
}
