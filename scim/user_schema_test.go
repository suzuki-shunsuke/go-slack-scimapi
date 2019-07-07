package scim

import (
	"context"
	"testing"

	"gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/require"
)

func TestClient_GetUserSchema(t *testing.T) {
	defer gock.Off()

	data := []struct {
		statusCode int
		body       string
		isError    bool
		exp        *Schema
	}{
		{
			statusCode: 200,
			isError:    false,
			exp: &Schema{
				ID:          "urn:scim:schemas:core:1.0:User",
				Name:        "User",
				Description: "Core User",
				Schema: []string{
					"urn:scim:schemas:core:1.0",
					"urn:scim:schemas:extension:enterprise:1.0",
				},
				Endpoint: "/Users",
				Attributes: []Attribute{
					{
						Name:        "id",
						Type:        "string",
						Description: "Unique identifier for the SCIM resource as defined by the Service Provider. Each representation of the resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider’s entire set of resources. It MUST be a stable, non-reassignable identifier that does not change when the same resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. REQUIRED.",

						Schema:   "urn:scim:schemas:core:1.0",
						ReadOnly: true,
						Required: true,
					},
					{
						Name:        "userName",
						Type:        "string",
						Description: "Unique identifier for the User, typically used by the user to directly authenticate to the service provider. Often displayed to the user as their unique identifier within the system (as opposed to id or externalId, which are generally opaque and not user-friendly identifiers). Each User MUST include a non-empty userName value. This identifier MUST be unique across the Service Consumer's entire set of Users. REQUIRED.",

						Schema:   "urn:scim:schemas:core:1.0",
						Required: true,
					},
					{
						Name:        "nickName",
						Type:        "string",
						Description: "The casual way to address the user in real life, e.g. \"Bob\" or \"Bobby\" instead of \"Robert\". This attribute SHOULD NOT be used to represent a User's username (e.g. bjensen or mpepperidge).",

						Schema: "urn:scim:schemas:core:1.0",
					},
					{
						Name:        "name",
						Type:        "complex",
						Description: "The components of the user’s real name. Providers MAY return just the full name as a single string in the formatted sub-attribute, or they MAY return just the individual component attributes using the other sub-attributes, or they MAY return both. If both variants are returned, they SHOULD be describing the same name, with the formatted name indicating how the component attributes should be combined.",

						Schema: "urn:scim:schemas:core:1.0",
						SubAttributes: []Attribute{
							{
								Name:        "familyName",
								Type:        "string",
								Description: "The family name of the User, or Last Name in most Western languages (e.g. Jensen given the full name Ms. Barbara J Jensen, III.).",
							}, {
								Name:        "givenName",
								Type:        "string",
								Description: "The given name of the User, or First Name in most Western languages (e.g. Barbara given the full name Ms. Barbara J Jensen, III.).",
							}, {
								Name:        "honorificPrefix",
								Type:        "string",
								Description: "The honorific prefix(es) of the User, or Title in most Western languages (e.g. Ms. given the full name Ms. Barbara J Jensen, III.).",
							},
						},
					},
					{
						Name:                          "emails",
						Type:                          "complex",
						MultiValued:                   true,
						MultiValuedAttributeChildName: "email",
						Description:                   "E-mail addresses for the user. The value SHOULD be canonicalized by the Service Provider, e.g. bjensen@example.com instead of bjensen@EXAMPLE.COM. Canonical Type values of work, home, and other.",

						Schema:   "urn:scim:schemas:core:1.0",
						Required: true,
						SubAttributes: []Attribute{
							{
								Name:        "value",
								Type:        "string",
								Description: "E-mail addresses for the user. The value SHOULD be canonicalized by the Service Provider, e.g. bjensen@example.com instead of bjensen@EXAMPLE.COM. Canonical Type values of work, home, and other.",
								Required:    true,
							}, {
								Name:            "type",
								Type:            "string",
								Description:     "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
								CanonicalValues: []string{"work", "home", "other"},
							}, {
								Name:        "primary",
								Type:        "boolean",
								Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
							},
						},
					},
					{
						Name:                          "photos",
						Type:                          "complex",
						MultiValued:                   true,
						MultiValuedAttributeChildName: "photo",
						Description:                   "URL of a photo of the User. The value SHOULD be a canonicalized URL, and MUST point to an image file (e.g. a GIF, JPEG, or PNG image file) rather than to a web page containing an image. Service Providers MAY return the same image at different sizes, though it is recognized that no standard for describing images of various sizes currently exists. Note that this attribute SHOULD NOT be used to send down arbitrary photos taken by this User, but specifically profile photos of the User suitable for display when describing the User. Instead of the standard Canonical Values for type, this attribute defines the following Canonical Values to represent popular photo sizes: photo, thumbnail.",

						Schema: "urn:scim:schemas:core:1.0",
						SubAttributes: []Attribute{
							{
								Name:        "value",
								Type:        "string",
								Description: "URL of a photo of the User. The value SHOULD be a canonicalized URL, and MUST point to an image file (e.g. a GIF, JPEG, or PNG image file) rather than to a web page containing an image. Service Providers MAY return the same image at different sizes, though it is recognized that no standard for describing images of various sizes currently exists. Note that this attribute SHOULD NOT be used to send down arbitrary photos taken by this User, but specifically profile photos of the User suitable for display when describing the User. Instead of the standard Canonical Values for type, this attribute defines the following Canonical Values to represent popular photo sizes: photo, thumbnail.",
								Required:    true,
							}, {
								Name:            "type",
								Type:            "string",
								Description:     "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
								CanonicalValues: []string{"photo", "thumbnail"},
							}, {
								Name:        "primary",
								Type:        "boolean",
								Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
							},
						},
					}, {
						Name:                          "groups",
						Type:                          "complex",
						MultiValued:                   true,
						MultiValuedAttributeChildName: "group",
						Description:                   "A list of groups that the user belongs to, either through direct membership, or dynamically calculated. The values are meant to enable expression of common group or role based access control models, although no explicit authorization model is defined. It is intended that the semantics of group membership and any behavior or authorization granted as a result of membership are defined by the Service Provider. The Canonical types \"direct\" and \"indirect\" are defined to describe how the group membership was derived. Direct group membership indicates the User is directly associated with the group and SHOULD indicate that Consumers may modify membership through the Group Resource. Indirect membership indicates User membership is transitive or dynamic and implies that Consumers cannot modify indirect group membership through the Group resource but MAY modify direct group membership through the Group resource which MAY influence indirect memberships. If the SCIM Service Provider exposes a Group resource, the value MUST be the \"id\" attribute of the corresponding Group resources to which the user belongs. Since this attribute is read-only, group membership changes MUST be applied via the Group Resource. READ-ONLY.",

						Schema:   "urn:scim:schemas:core:1.0",
						ReadOnly: true,
						SubAttributes: []Attribute{
							{
								Name:        "value",
								Type:        "string",
								Description: "The \"id\" of a SCIM Group resource. REQUIRED.",
								ReadOnly:    true,
								Required:    true,
							}, {
								Name:        "display",
								Type:        "string",
								Description: "A human readable name, primarily used for display purposes. READ-ONLY.",
								ReadOnly:    true,
							},
						},
					},
					{
						Name:        "active",
						Type:        "string",
						Description: "A Boolean value indicating the User's administrative status. The definitive meaning of this attribute is determined by the Service Provider though a value of true infers the User is, for example, able to login while a value of false implies the User's account has been suspended.",

						Schema: "urn:scim:schemas:core:1.0",
					}, {
						Name:                          "addresses",
						Type:                          "complex",
						MultiValued:                   true,
						MultiValuedAttributeChildName: "address",
						Description:                   "A physical mailing address for this User. Canonical Type Values of work, home, and other. The value attribute is a complex type with the following sub-attributes. All Sub-Attributes are OPTIONAL.",

						Schema: "urn:scim:schemas:core:1.0",
						SubAttributes: []Attribute{
							{
								Name:        "streetAddress",
								Type:        "string",
								Description: "The full street address component, which may include house number, street name, P.O. box, and multi-line extended street address information. This attribute MAY contain newlines.",
							}, {
								Name:        "locality",
								Type:        "string",
								Description: "The city or locality component.",
							}, {
								Name:        "region",
								Type:        "string",
								Description: "The state or region component.",
							}, {
								Name:        "postalCode",
								Type:        "string",
								Description: "The zipcode or postal code component.",
							}, {
								Name:        "country",
								Type:        "string",
								Description: "The country name component. When specified the value MUST be in ISO 3166-1 alpha 2 \"short\" code format; e.g., the United States and Sweden are \"US\" and \"SE\", respectively.",
							}, {
								Name:            "type",
								Type:            "string",
								Description:     "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
								CanonicalValues: []string{"work", "home", "other"},
							}, {
								Name:        "primary",
								Type:        "boolean",
								Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
							},
						},
					}, {
						Name:                          "phoneNumbers",
						Type:                          "complex",
						MultiValued:                   true,
						MultiValuedAttributeChildName: "phoneNumber",
						Description:                   "Phone numbers for the User. The value SHOULD be canonicalized by the Service Provider according to format in RFC3966 e.g. 'tel:+1-201-555-0123'. Canonical Type values of work, home, mobile, fax, pager and other.",

						Schema: "urn:scim:schemas:core:1.0",
						SubAttributes: []Attribute{
							{
								Name:        "value",
								Type:        "string",
								Description: "Phone numbers for the User. The value SHOULD be canonicalized by the Service Provider according to format in RFC3966 e.g. 'tel:+1-201-555-0123'. Canonical Type values of work, home, mobile, fax, pager and other.",
								Required:    true,
							}, {
								Name:            "type",
								Type:            "string",
								Description:     "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
								CanonicalValues: []string{"work", "home", "mobile", "pager", "fax", "other"},
							}, {
								Name:        "primary",
								Type:        "boolean",
								Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
							},
						},
					}, {
						Name:        "displayName",
						Type:        "string",
						Description: "The name of the User, suitable for display to end-users. Each User returned MAY include a non-empty displayName value. The name SHOULD be the full name of the User being described if known (e.g. Babs Jensen or Ms. Barbara J Jensen, III), but MAY be a username or handle, if that is all that is available (e.g. bjensen). The value provided SHOULD be the primary textual label by which this User is normally displayed by the Service Provider when presenting it to end-users.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:        "profileUrl",
						Type:        "string",
						Description: "A fully qualified URL to a page representing the User's online profile.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:        "userType",
						Type:        "string",
						Description: "Used to identify the organization to user relationship. Typical values used might be \"Contractor\", \"Employee\", \"Intern\", \"Temp\", \"External\", and \"Unknown\" but any value may be used.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:        "title",
						Type:        "string",
						Description: "The user’s title, such as “Vice President”.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:        "preferredLanguage",
						Type:        "string",
						Description: "Indicates the User's preferred written or spoken language. Generally used for selecting a localized User interface. Valid values are concatenation of the ISO 639-1 two letter language code, an underscore, and the ISO 3166-1 2 letter country code; e.g., 'en_US' specifies the language English and country US.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:        "locale",
						Type:        "string",
						Description: "Used to indicate the User's default location for purposes of localizing items such as currency, date time format, numerical representations, etc. A locale value is a concatenation of the ISO 639-1 two letter language code, an underscore, and the ISO 3166-1 2 letter country code; e.g., 'en_US' specifies the language English and country US.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:        "timezone",
						Type:        "string",
						Description: "The User's time zone in the \"Olson\" timezone database format; e.g.,'America/Los_Angeles'.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:        "password",
						Type:        "string",
						Description: "The User's clear text password. This attribute is intended to be used as a means to specify an initial password when creating a new User or to reset an existing User's password. No accepted standards exist to convey password policies, hence Consumers should expect Service Providers to reject password values. This value MUST never be returned by a Service Provider in any form.",
						Schema:      "urn:scim:schemas:core:1.0",
					}, {
						Name:                          "roles",
						Type:                          "complex",
						MultiValued:                   true,
						MultiValuedAttributeChildName: "role",
						Description:                   "A list of roles for the User that collectively represent who the User is; e.g., \"Student\", \"Faculty\". No vocabulary or syntax is specified though it is expected that a role value is a String or label representing a collection of entitlements. This value has NO canonical types.",
						Schema:                        "urn:scim:schemas:core:1.0",
						SubAttributes: []Attribute{
							{
								Name:        "value",
								Type:        "string",
								Description: "A list of roles for the User that collectively represent who the User is; e.g., \"Student\", \"Faculty\". No vocabulary or syntax is specified though it is expected that a role value is a String or label representing a collection of entitlements. This value has NO canonical types.",
							}, {
								Name:        "type",
								Type:        "string",
								Description: "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
							}, {
								Name:        "primary",
								Type:        "boolean",
								Description: "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
							},
						},
					}, {
						Name:        "urn:scim:schemas:extension:enterprise:1.0",
						Type:        "complex",
						MultiValued: true,
						Description: "The following SCIM extension defines attributes commonly used in representing users that belong to, or act on behalf of a business or enterprise. The enterprise user extension is identified using the following URI: 'urn:scim:schemas:extension:enterprise:1.0'.",
						Schema:      "urn:scim:schemas:extension:enterprise:1.0",
						SubAttributes: []Attribute{
							{
								Name:        "employeeNumber",
								Type:        "string",
								Description: "Numeric or alphanumeric identifier assigned to a person, typically based on order of hire or association with an organization.",
							}, {
								Name:        "costCenter",
								Type:        "string",
								Description: "Identifies the name of a cost center.",
							}, {
								Name:        "organization",
								Type:        "string",
								Description: "Identifies the name of an organization.",
							}, {
								Name:        "division",
								Type:        "string",
								Description: "Identifies the name of a division.",
							}, {
								Name:        "department",
								Type:        "string",
								Description: "Identifies the name of a department.",
							}, {
								Name:        "manager",
								Type:        "complex",
								Description: "The User's manager. A complex type that optionally allows Service Providers to represent organizational hierarchy by referencing the \"id\" attribute of another User.",
								SubAttributes: []Attribute{
									{
										Name:        "managerId",
										Type:        "string",
										Description: "The id of the SCIM resource representing the User's manager. REQUIRED.",
									},
								},
							},
						},
					},
				},
			},
			body: `{
  "id": "urn:scim:schemas:core:1.0:User",
  "name": "User",
  "description": "Core User",
  "schema": [
    "urn:scim:schemas:core:1.0",
    "urn:scim:schemas:extension:enterprise:1.0"
  ],
  "endpoint": "/Users",
  "attributes": [
    {
      "name": "id",
      "type": "string",
      "multiValued": false,
      "description": "Unique identifier for the SCIM resource as defined by the Service Provider. Each representation of the resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider’s entire set of resources. It MUST be a stable, non-reassignable identifier that does not change when the same resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. REQUIRED.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": true,
      "required": true,
      "caseExact": false
    },
    {
      "name": "userName",
      "type": "string",
      "multiValued": false,
      "description": "Unique identifier for the User, typically used by the user to directly authenticate to the service provider. Often displayed to the user as their unique identifier within the system (as opposed to id or externalId, which are generally opaque and not user-friendly identifiers). Each User MUST include a non-empty userName value. This identifier MUST be unique across the Service Consumer's entire set of Users. REQUIRED.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": true,
      "caseExact": false
    },
    {
      "name": "nickName",
      "type": "string",
      "multiValued": false,
      "description": "The casual way to address the user in real life, e.g. \"Bob\" or \"Bobby\" instead of \"Robert\". This attribute SHOULD NOT be used to represent a User's username (e.g. bjensen or mpepperidge).",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "name",
      "type": "complex",
      "multiValued": false,
      "description": "The components of the user’s real name. Providers MAY return just the full name as a single string in the formatted sub-attribute, or they MAY return just the individual component attributes using the other sub-attributes, or they MAY return both. If both variants are returned, they SHOULD be describing the same name, with the formatted name indicating how the component attributes should be combined.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "familyName",
          "type": "string",
          "multiValued": false,
          "description": "The family name of the User, or Last Name in most Western languages (e.g. Jensen given the full name Ms. Barbara J Jensen, III.).",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "givenName",
          "type": "string",
          "multiValued": false,
          "description": "The given name of the User, or First Name in most Western languages (e.g. Barbara given the full name Ms. Barbara J Jensen, III.).",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "honorificPrefix",
          "type": "string",
          "multiValued": false,
          "description": "The honorific prefix(es) of the User, or Title in most Western languages (e.g. Ms. given the full name Ms. Barbara J Jensen, III.).",
          "readOnly": false,
          "required": false,
          "caseExact": false
        }
      ]
    },
    {
      "name": "emails",
      "type": "complex",
      "multiValued": true,
      "multiValuedAttributeChildName": "email",
      "description": "E-mail addresses for the user. The value SHOULD be canonicalized by the Service Provider, e.g. bjensen@example.com instead of bjensen@EXAMPLE.COM. Canonical Type values of work, home, and other.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": true,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "value",
          "type": "string",
          "multiValued": false,
          "description": "E-mail addresses for the user. The value SHOULD be canonicalized by the Service Provider, e.g. bjensen@example.com instead of bjensen@EXAMPLE.COM. Canonical Type values of work, home, and other.",
          "readOnly": false,
          "required": true,
          "caseExact": false
        },
        {
          "name": "type",
          "type": "string",
          "multiValued": false,
          "description": "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
          "readOnly": false,
          "required": false,
          "caseExact": false,
          "canonicalValues": [
            "work",
            "home",
            "other"
          ]
        },
        {
          "name": "primary",
          "type": "boolean",
          "multiValued": false,
          "description": "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        }
      ]
    },
    {
      "name": "photos",
      "type": "complex",
      "multiValued": true,
      "multiValuedAttributeChildName": "photo",
      "description": "URL of a photo of the User. The value SHOULD be a canonicalized URL, and MUST point to an image file (e.g. a GIF, JPEG, or PNG image file) rather than to a web page containing an image. Service Providers MAY return the same image at different sizes, though it is recognized that no standard for describing images of various sizes currently exists. Note that this attribute SHOULD NOT be used to send down arbitrary photos taken by this User, but specifically profile photos of the User suitable for display when describing the User. Instead of the standard Canonical Values for type, this attribute defines the following Canonical Values to represent popular photo sizes: photo, thumbnail.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "value",
          "type": "string",
          "multiValued": false,
          "description": "URL of a photo of the User. The value SHOULD be a canonicalized URL, and MUST point to an image file (e.g. a GIF, JPEG, or PNG image file) rather than to a web page containing an image. Service Providers MAY return the same image at different sizes, though it is recognized that no standard for describing images of various sizes currently exists. Note that this attribute SHOULD NOT be used to send down arbitrary photos taken by this User, but specifically profile photos of the User suitable for display when describing the User. Instead of the standard Canonical Values for type, this attribute defines the following Canonical Values to represent popular photo sizes: photo, thumbnail.",
          "readOnly": false,
          "required": true,
          "caseExact": false
        },
        {
          "name": "type",
          "type": "string",
          "multiValued": false,
          "description": "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
          "readOnly": false,
          "required": false,
          "caseExact": false,
          "canonicalValues": [
            "photo",
            "thumbnail"
          ]
        },
        {
          "name": "primary",
          "type": "boolean",
          "multiValued": false,
          "description": "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        }
      ]
    },
    {
      "name": "groups",
      "type": "complex",
      "multiValued": true,
      "multiValuedAttributeChildName": "group",
      "description": "A list of groups that the user belongs to, either through direct membership, or dynamically calculated. The values are meant to enable expression of common group or role based access control models, although no explicit authorization model is defined. It is intended that the semantics of group membership and any behavior or authorization granted as a result of membership are defined by the Service Provider. The Canonical types \"direct\" and \"indirect\" are defined to describe how the group membership was derived. Direct group membership indicates the User is directly associated with the group and SHOULD indicate that Consumers may modify membership through the Group Resource. Indirect membership indicates User membership is transitive or dynamic and implies that Consumers cannot modify indirect group membership through the Group resource but MAY modify direct group membership through the Group resource which MAY influence indirect memberships. If the SCIM Service Provider exposes a Group resource, the value MUST be the \"id\" attribute of the corresponding Group resources to which the user belongs. Since this attribute is read-only, group membership changes MUST be applied via the Group Resource. READ-ONLY.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": true,
      "required": false,
      "caseExact": false,
      "subattributes": [
        {
          "name": "value",
          "type": "string",
          "multiValued": false,
          "description": "The \"id\" of a SCIM Group resource. REQUIRED.",
          "readOnly": true,
          "required": true,
          "caseExact": false
        },
        {
          "name": "display",
          "type": "string",
          "multiValued": false,
          "description": "A human readable name, primarily used for display purposes. READ-ONLY.",
          "readOnly": true,
          "required": false,
          "caseExact": false
        }
      ]
    },
    {
      "name": "active",
      "type": "string",
      "multiValued": false,
      "description": "A Boolean value indicating the User's administrative status. The definitive meaning of this attribute is determined by the Service Provider though a value of true infers the User is, for example, able to login while a value of false implies the User's account has been suspended.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "addresses",
      "type": "complex",
      "multiValued": true,
      "multiValuedAttributeChildName": "address",
      "description": "A physical mailing address for this User. Canonical Type Values of work, home, and other. The value attribute is a complex type with the following sub-attributes. All Sub-Attributes are OPTIONAL.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "streetAddress",
          "type": "string",
          "multiValued": false,
          "description": "The full street address component, which may include house number, street name, P.O. box, and multi-line extended street address information. This attribute MAY contain newlines.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "locality",
          "type": "string",
          "multiValued": false,
          "description": "The city or locality component.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "region",
          "type": "string",
          "multiValued": false,
          "description": "The state or region component.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "postalCode",
          "type": "string",
          "multiValued": false,
          "description": "The zipcode or postal code component.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "country",
          "type": "string",
          "multiValued": false,
          "description": "The country name component. When specified the value MUST be in ISO 3166-1 alpha 2 \"short\" code format; e.g., the United States and Sweden are \"US\" and \"SE\", respectively.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "type",
          "type": "string",
          "multiValued": false,
          "description": "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
          "readOnly": false,
          "required": false,
          "caseExact": false,
          "canonicalValues": [
            "work",
            "home",
            "other"
          ]
        },
        {
          "name": "primary",
          "type": "boolean",
          "multiValued": false,
          "description": "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        }
      ]
    },
    {
      "name": "phoneNumbers",
      "type": "complex",
      "multiValued": true,
      "multiValuedAttributeChildName": "phoneNumber",
      "description": "Phone numbers for the User. The value SHOULD be canonicalized by the Service Provider according to format in RFC3966 e.g. 'tel:+1-201-555-0123'. Canonical Type values of work, home, mobile, fax, pager and other.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "value",
          "type": "string",
          "multiValued": false,
          "description": "Phone numbers for the User. The value SHOULD be canonicalized by the Service Provider according to format in RFC3966 e.g. 'tel:+1-201-555-0123'. Canonical Type values of work, home, mobile, fax, pager and other.",
          "readOnly": false,
          "required": true,
          "caseExact": false
        },
        {
          "name": "type",
          "type": "string",
          "multiValued": false,
          "description": "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
          "readOnly": false,
          "required": false,
          "caseExact": false,
          "canonicalValues": [
            "work",
            "home",
            "mobile",
            "pager",
            "fax",
            "other"
          ]
        },
        {
          "name": "primary",
          "type": "boolean",
          "multiValued": false,
          "description": "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        }
      ]
    },
    {
      "name": "displayName",
      "type": "string",
      "multiValued": false,
      "description": "The name of the User, suitable for display to end-users. Each User returned MAY include a non-empty displayName value. The name SHOULD be the full name of the User being described if known (e.g. Babs Jensen or Ms. Barbara J Jensen, III), but MAY be a username or handle, if that is all that is available (e.g. bjensen). The value provided SHOULD be the primary textual label by which this User is normally displayed by the Service Provider when presenting it to end-users.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "profileUrl",
      "type": "string",
      "multiValued": false,
      "description": "A fully qualified URL to a page representing the User's online profile.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "userType",
      "type": "string",
      "multiValued": false,
      "description": "Used to identify the organization to user relationship. Typical values used might be \"Contractor\", \"Employee\", \"Intern\", \"Temp\", \"External\", and \"Unknown\" but any value may be used.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "title",
      "type": "string",
      "multiValued": false,
      "description": "The user’s title, such as “Vice President”.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "preferredLanguage",
      "type": "string",
      "multiValued": false,
      "description": "Indicates the User's preferred written or spoken language. Generally used for selecting a localized User interface. Valid values are concatenation of the ISO 639-1 two letter language code, an underscore, and the ISO 3166-1 2 letter country code; e.g., 'en_US' specifies the language English and country US.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "locale",
      "type": "string",
      "multiValued": false,
      "description": "Used to indicate the User's default location for purposes of localizing items such as currency, date time format, numerical representations, etc. A locale value is a concatenation of the ISO 639-1 two letter language code, an underscore, and the ISO 3166-1 2 letter country code; e.g., 'en_US' specifies the language English and country US.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "timezone",
      "type": "string",
      "multiValued": false,
      "description": "The User's time zone in the \"Olson\" timezone database format; e.g.,'America/Los_Angeles'.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "password",
      "type": "string",
      "multiValued": false,
      "description": "The User's clear text password. This attribute is intended to be used as a means to specify an initial password when creating a new User or to reset an existing User's password. No accepted standards exist to convey password policies, hence Consumers should expect Service Providers to reject password values. This value MUST never be returned by a Service Provider in any form.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false
    },
    {
      "name": "roles",
      "type": "complex",
      "multiValued": true,
      "multiValuedAttributeChildName": "role",
      "description": "A list of roles for the User that collectively represent who the User is; e.g., \"Student\", \"Faculty\". No vocabulary or syntax is specified though it is expected that a role value is a String or label representing a collection of entitlements. This value has NO canonical types.",
      "schema": "urn:scim:schemas:core:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "value",
          "type": "string",
          "multiValued": false,
          "description": "A list of roles for the User that collectively represent who the User is; e.g., \"Student\", \"Faculty\". No vocabulary or syntax is specified though it is expected that a role value is a String or label representing a collection of entitlements. This value has NO canonical types.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "type",
          "type": "string",
          "multiValued": false,
          "description": "A label indicating the attribute’s function; e.g., 'work' or 'home'.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "primary",
          "type": "boolean",
          "multiValued": false,
          "description": "A Boolean value indicating the 'primary' or preferred attribute value for this attribute, e.g. the preferred mailing address or primary e-mail address. The primary attribute value 'true' MUST appear no more than once.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        }
      ]
    },
    {
      "name": "urn:scim:schemas:extension:enterprise:1.0",
      "type": "complex",
      "multiValued": true,
      "description": "The following SCIM extension defines attributes commonly used in representing users that belong to, or act on behalf of a business or enterprise. The enterprise user extension is identified using the following URI: 'urn:scim:schemas:extension:enterprise:1.0'.",
      "schema": "urn:scim:schemas:extension:enterprise:1.0",
      "readOnly": false,
      "required": false,
      "caseExact": false,
      "subAttributes": [
        {
          "name": "employeeNumber",
          "type": "string",
          "multiValued": false,
          "description": "Numeric or alphanumeric identifier assigned to a person, typically based on order of hire or association with an organization.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "costCenter",
          "type": "string",
          "multiValued": false,
          "description": "Identifies the name of a cost center.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "organization",
          "type": "string",
          "multiValued": false,
          "description": "Identifies the name of an organization.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "division",
          "type": "string",
          "multiValued": false,
          "description": "Identifies the name of a division.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "department",
          "type": "string",
          "multiValued": false,
          "description": "Identifies the name of a department.",
          "readOnly": false,
          "required": false,
          "caseExact": false
        },
        {
          "name": "manager",
          "type": "complex",
          "multiValued": false,
          "description": "The User's manager. A complex type that optionally allows Service Providers to represent organizational hierarchy by referencing the \"id\" attribute of another User.",
          "readOnly": false,
          "required": false,
          "caseExact": false,
          "subAttributes": {
            "name": "managerId",
            "type": "string",
            "multiValued": false,
            "description": "The id of the SCIM resource representing the User's manager. REQUIRED.",
            "readOnly": false,
            "required": false,
            "caseExact": false
          }
        }
      ]
    }
  ]
}`,
		},
	}

	ctx := context.Background()
	client := NewClient("XXX")
	for _, d := range data {
		gock.New("https://api.slack.com").
			Get("/scim/v1/Schemas/Users").
			MatchType("json").Reply(d.statusCode).
			BodyString(d.body)
		schema, resp, err := client.GetUserSchema(ctx)
		if d.isError {
			require.NotNil(t, err)
			return
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, schema)
		require.Equal(t, d.statusCode, resp.StatusCode)
		if d.exp != nil {
			require.Equal(t, d.exp, schema)
		}
	}
}
