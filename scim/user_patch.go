package scim

type (
	UserPatch struct {
		ID                            string                              `json:"id,omitempty"`
		UserName                      string                              `json:"userName,omitempty"`
		Active                        *bool                               `json:"active,omitempty"`
		ExternalID                    *string                             `json:"externalId,omitempty"`
		NickName                      *string                             `json:"nickName,omitempty"`
		ProfileURL                    *string                             `json:"profileUrl,omitempty"`
		DisplayName                   *string                             `json:"displayName,omitempty"`
		UserType                      *string                             `json:"userType,omitempty"`
		Title                         *string                             `json:"title,omitempty"`
		PreferredLanguage             *string                             `json:"preferredLanguage,omitempty"`
		Locale                        *string                             `json:"locale,omitempty"`
		Timezone                      *string                             `json:"timezone,omitempty"`
		Password                      *string                             `json:"password,omitempty"`
		Name                          *NamePatch                          `json:"name,omitempty"`
		Meta                          *Meta                               `json:"meta,omitempty"`
		Emails                        []Email                             `json:"emails,omitempty"`
		Addresses                     *[]Address                          `json:"addresses,omitempty"`
		PhoneNumbers                  *[]PhoneNumber                      `json:"phoneNumbers,omitempty"`
		Roles                         *[]Role                             `json:"roles,omitempty"`
		Photos                        *[]Photo                            `json:"photos,omitempty"`
		Groups                        *[]Group                            `json:"groups,omitempty"`
		Schemas                       []string                            `json:"schemas"`
		EnterpriseUserSchemaExtension *EnterpriseUserSchemaExtensionPatch `json:"urn:scim:schemas:extension:enterprise:1.0,omitempty"`
	}

	EnterpriseUserSchemaExtensionPatch struct {
		EmployeeNumber *string  `json:"employeeNumber,omitempty"`
		CostCenter     *string  `json:"costCenter,omitempty"`
		Organization   *string  `json:"organization,omitempty"`
		Division       *string  `json:"division,omitempty"`
		Department     *string  `json:"department,omitempty"`
		Manager        *Manager `json:"manager,omitempty"`
	}

	NamePatch struct {
		FamilyName      *string `json:"familyName,omitempty"`
		GivenName       *string `json:"givenName,omitempty"`
		HonorificPrefix *string `json:"honorificPrefix,omitempty"`
	}
)
