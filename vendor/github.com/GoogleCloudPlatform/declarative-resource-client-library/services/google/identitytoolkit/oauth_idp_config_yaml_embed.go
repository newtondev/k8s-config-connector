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
// gen_go_data -package identitytoolkit -var YAML_oauth_idp_config blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/oauth_idp_config.yaml

package identitytoolkit

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/identitytoolkit/oauth_idp_config.yaml
var YAML_oauth_idp_config = []byte("info:\n  title: IdentityToolkit/OAuthIdpConfig\n  description: The IdentityToolkit OAuthIdpConfig resource\n  x-dcl-struct-name: OAuthIdpConfig\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a OAuthIdpConfig\n    parameters:\n    - name: oAuthIdpConfig\n      required: true\n      description: A full instance of a OAuthIdpConfig\n  apply:\n    description: The function used to apply information about a OAuthIdpConfig\n    parameters:\n    - name: oAuthIdpConfig\n      required: true\n      description: A full instance of a OAuthIdpConfig\n  delete:\n    description: The function used to delete a OAuthIdpConfig\n    parameters:\n    - name: oAuthIdpConfig\n      required: true\n      description: A full instance of a OAuthIdpConfig\n  deleteAll:\n    description: The function used to delete all OAuthIdpConfig\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many OAuthIdpConfig\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    OAuthIdpConfig:\n      title: OAuthIdpConfig\n      x-dcl-id: projects/{{project}}/oauthIdpConfigs/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - project\n      properties:\n        clientId:\n          type: string\n          x-dcl-go-name: ClientId\n          description: The client id of an OAuth client.\n        clientSecret:\n          type: string\n          x-dcl-go-name: ClientSecret\n          description: The client secret of the OAuth client, to enable OIDC code\n            flow.\n          x-dcl-sensitive: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: The config's display name set by developers.\n        enabled:\n          type: boolean\n          x-dcl-go-name: Enabled\n          description: True if allows the user to sign in with the provider.\n        issuer:\n          type: string\n          x-dcl-go-name: Issuer\n          description: For OIDC Idps, the issuer identifier.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of the Config resource\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        responseType:\n          type: object\n          x-dcl-go-name: ResponseType\n          x-dcl-go-type: OAuthIdpConfigResponseType\n          description: 'The multiple response type to request for in the OAuth authorization\n            flow. This can possibly be a combination of set bits (e.g.: {id\\_token,\n            token}).'\n          x-dcl-server-default: true\n          properties:\n            code:\n              type: boolean\n              x-dcl-go-name: Code\n              description: If true, authorization code is returned from IdP's authorization\n                endpoint.\n            idToken:\n              type: boolean\n              x-dcl-go-name: IdToken\n              description: If true, ID token is returned from IdP's authorization\n                endpoint.\n            token:\n              type: boolean\n              x-dcl-go-name: Token\n              description: If true, access token is returned from IdP's authorization\n                endpoint.\n")

// 3677 bytes
// MD5: 953182f1bcf3a3bce2b05c9dbd5c9a66
