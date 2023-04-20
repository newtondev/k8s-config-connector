// Copyright 2023 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package iam -var YAML_workforce_pool_provider blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/workforce_pool_provider.yaml

package iam

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/iam/workforce_pool_provider.yaml
var YAML_workforce_pool_provider = []byte("info:\n  title: Iam/WorkforcePoolProvider\n  description: The Iam WorkforcePoolProvider resource\n  x-dcl-struct-name: WorkforcePoolProvider\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a WorkforcePoolProvider\n    parameters:\n    - name: workforcePoolProvider\n      required: true\n      description: A full instance of a WorkforcePoolProvider\n  apply:\n    description: The function used to apply information about a WorkforcePoolProvider\n    parameters:\n    - name: workforcePoolProvider\n      required: true\n      description: A full instance of a WorkforcePoolProvider\n  delete:\n    description: The function used to delete a WorkforcePoolProvider\n    parameters:\n    - name: workforcePoolProvider\n      required: true\n      description: A full instance of a WorkforcePoolProvider\n  deleteAll:\n    description: The function used to delete all WorkforcePoolProvider\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: workforcePool\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many WorkforcePoolProvider\n    parameters:\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: workforcePool\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    WorkforcePoolProvider:\n      title: WorkforcePoolProvider\n      x-dcl-id: locations/{{location}}/workforcePools/{{workforce_pool}}/providers/{{name}}\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - attributeMapping\n      - location\n      - workforcePool\n      properties:\n        attributeCondition:\n          type: string\n          x-dcl-go-name: AttributeCondition\n          description: 'A [Common Expression Language](https://opensource.google/projects/cel)\n            expression, in plain text, to restrict what otherwise valid authentication\n            credentials issued by the provider should not be accepted. The expression\n            must output a boolean representing whether to allow the federation. The\n            following keywords may be referenced in the expressions: * `assertion`:\n            JSON representing the authentication credential issued by the provider.\n            * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`.\n            `google.profile_photo` and `google.display_name` are not supported. *\n            `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`.\n            The maximum length of the attribute condition expression is 4096 characters.\n            If unspecified, all valid authentication credentials will be accepted.\n            The following example shows how to only allow credentials with a mapped\n            `google.groups` value of `admins`: ``` \"''admins'' in google.groups\" ```'\n        attributeMapping:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: AttributeMapping\n          description: 'Required. Maps attributes from the authentication credentials\n            issued by an external identity provider to Google Cloud attributes, such\n            as `subject` and `segment`. Each key must be a string specifying the Google\n            Cloud IAM attribute to map to. The following keys are supported: * `google.subject`:\n            The principal IAM is authenticating. You can reference this value in IAM\n            bindings. This is also the subject that appears in Cloud Logging logs.\n            This is a required field and the mapped subject cannot exceed 127 bytes.\n            * `google.groups`: Groups the authenticating user belongs to. You can\n            grant groups access to resources using an IAM `principalSet` binding;\n            access applies to all members of the group. * `google.display_name`: The\n            name of the authenticated user. This is an optional field and the mapped\n            display name cannot exceed 100 bytes. If not set, `google.subject` will\n            be displayed instead. This attribute cannot be referenced in IAM bindings.\n            * `google.profile_photo`: The URL that specifies the authenticated user''s\n            thumbnail photo. This is an optional field. When set, the image will be\n            visible as the user''s profile picture. If not set, a generic user icon\n            will be displayed instead. This attribute cannot be referenced in IAM\n            bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`,\n            where {custom_attribute} is the name of the custom attribute to be mapped.\n            You can define a maximum of 50 custom attributes. The maximum length of\n            a mapped attribute key is 100 characters, and the key may only contain\n            the characters [a-z0-9_]. You can reference these attributes in IAM policies\n            to define fine-grained access for a workforce pool to Google Cloud resources.\n            For example:'\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: A user-specified description of the provider. Cannot exceed\n            256 characters.\n        disabled:\n          type: boolean\n          x-dcl-go-name: Disabled\n          description: Whether the provider is disabled. You cannot use a disabled\n            provider to exchange tokens. However, existing tokens still grant access.\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: A user-specified display name for the provider. Cannot exceed\n            32 characters.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'Output only. The resource name of the provider. Format: `locations/{location}/workforcePools/{workforce_pool_id}/providers/{provider_id}`'\n          x-kubernetes-immutable: true\n        oidc:\n          type: object\n          x-dcl-go-name: Oidc\n          x-dcl-go-type: WorkforcePoolProviderOidc\n          description: An OpenId Connect 1.0 identity provider configuration.\n          x-dcl-conflicts:\n          - saml\n          required:\n          - issuerUri\n          - clientId\n          - webSsoConfig\n          properties:\n            clientId:\n              type: string\n              x-dcl-go-name: ClientId\n              description: Required. The client ID. Must match the audience claim\n                of the JWT issued by the identity provider.\n            issuerUri:\n              type: string\n              x-dcl-go-name: IssuerUri\n              description: Required. The OIDC issuer URI. Must be a valid URI using\n                the 'https' scheme.\n            webSsoConfig:\n              type: object\n              x-dcl-go-name: WebSsoConfig\n              x-dcl-go-type: WorkforcePoolProviderOidcWebSsoConfig\n              description: Required. Configuration for web single sign-on for the\n                OIDC provider. Here, web sign-in refers to console sign-in and gcloud\n                sign-in through the browser.\n              required:\n              - responseType\n              - assertionClaimsBehavior\n              properties:\n                assertionClaimsBehavior:\n                  type: string\n                  x-dcl-go-name: AssertionClaimsBehavior\n                  x-dcl-go-type: WorkforcePoolProviderOidcWebSsoConfigAssertionClaimsBehaviorEnum\n                  description: 'Required. The behavior for how OIDC Claims are included\n                    in the `assertion` object used for attribute mapping and attribute\n                    condition. Possible values: ASSERTION_CLAIMS_BEHAVIOR_UNSPECIFIED,\n                    ONLY_ID_TOKEN_CLAIMS'\n                  enum:\n                  - ASSERTION_CLAIMS_BEHAVIOR_UNSPECIFIED\n                  - ONLY_ID_TOKEN_CLAIMS\n                responseType:\n                  type: string\n                  x-dcl-go-name: ResponseType\n                  x-dcl-go-type: WorkforcePoolProviderOidcWebSsoConfigResponseTypeEnum\n                  description: 'Required. The Response Type to request for in the\n                    OIDC Authorization Request for web sign-in. Possible values: RESPONSE_TYPE_UNSPECIFIED,\n                    ID_TOKEN'\n                  enum:\n                  - RESPONSE_TYPE_UNSPECIFIED\n                  - ID_TOKEN\n        saml:\n          type: object\n          x-dcl-go-name: Saml\n          x-dcl-go-type: WorkforcePoolProviderSaml\n          description: A SAML identity provider configuration.\n          x-dcl-conflicts:\n          - oidc\n          required:\n          - idpMetadataXml\n          properties:\n            idpMetadataXml:\n              type: string\n              x-dcl-go-name: IdpMetadataXml\n              description: 'Required. SAML Identity provider configuration metadata\n                xml doc. The xml document should comply with [SAML 2.0 specification](https://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf).\n                The max size of the acceptable xml document will be bounded to 128k\n                characters. The metadata xml document should satisfy the following\n                constraints: 1) Must contain an Identity Provider Entity ID. 2) Must\n                contain at least one non-expired signing key certificate. 3) For each\n                signing key: a) Valid from should be no more than 7 days from now.\n                b) Valid to should be no more than 10 years in the future. 4) Up to\n                3 IdP signing keys are allowed in the metadata xml. When updating\n                the provider''s metadata xml, at least one non-expired signing key\n                must overlap with the existing metadata. This requirement is skipped\n                if there are no non-expired signing keys present in the existing metadata.'\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: WorkforcePoolProviderStateEnum\n          readOnly: true\n          description: 'Output only. The state of the provider. Possible values: STATE_UNSPECIFIED,\n            ACTIVE, DELETED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - ACTIVE\n          - DELETED\n        workforcePool:\n          type: string\n          x-dcl-go-name: WorkforcePool\n          description: The workforce_pool for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Iam/WorkforcePool\n            field: name\n            parent: true\n")

// 10921 bytes
// MD5: 9140a5eecf53c89dfe8e4e53cfe9a40d
