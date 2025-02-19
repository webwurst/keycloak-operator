// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationExecutionInfo) DeepCopyInto(out *AuthenticationExecutionInfo) {
	*out = *in
	if in.RequirementChoices != nil {
		in, out := &in.RequirementChoices, &out.RequirementChoices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationExecutionInfo.
func (in *AuthenticationExecutionInfo) DeepCopy() *AuthenticationExecutionInfo {
	if in == nil {
		return nil
	}
	out := new(AuthenticationExecutionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticatorConfig) DeepCopyInto(out *AuthenticatorConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticatorConfig.
func (in *AuthenticatorConfig) DeepCopy() *AuthenticatorConfig {
	if in == nil {
		return nil
	}
	out := new(AuthenticatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericStatus) DeepCopyInto(out *GenericStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericStatus.
func (in *GenericStatus) DeepCopy() *GenericStatus {
	if in == nil {
		return nil
	}
	out := new(GenericStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keycloak) DeepCopyInto(out *Keycloak) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keycloak.
func (in *Keycloak) DeepCopy() *Keycloak {
	if in == nil {
		return nil
	}
	out := new(Keycloak)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Keycloak) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakApiClient) DeepCopyInto(out *KeycloakApiClient) {
	*out = *in
	if in.DefaultRoles != nil {
		in, out := &in.DefaultRoles, &out.DefaultRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RedirectUris != nil {
		in, out := &in.RedirectUris, &out.RedirectUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WebOrigins != nil {
		in, out := &in.WebOrigins, &out.WebOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProtocolMappers != nil {
		in, out := &in.ProtocolMappers, &out.ProtocolMappers
		*out = make([]KeycloakProtocolMapper, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakApiClient.
func (in *KeycloakApiClient) DeepCopy() *KeycloakApiClient {
	if in == nil {
		return nil
	}
	out := new(KeycloakApiClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakApiPasswordReset) DeepCopyInto(out *KeycloakApiPasswordReset) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakApiPasswordReset.
func (in *KeycloakApiPasswordReset) DeepCopy() *KeycloakApiPasswordReset {
	if in == nil {
		return nil
	}
	out := new(KeycloakApiPasswordReset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakApiRealm) DeepCopyInto(out *KeycloakApiRealm) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]*KeycloakUser, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakUser)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Clients != nil {
		in, out := &in.Clients, &out.Clients
		*out = make([]*KeycloakClient, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakClient)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.IdentityProviders != nil {
		in, out := &in.IdentityProviders, &out.IdentityProviders
		*out = make([]*KeycloakIdentityProvider, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KeycloakIdentityProvider)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.EventsListeners != nil {
		in, out := &in.EventsListeners, &out.EventsListeners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakApiRealm.
func (in *KeycloakApiRealm) DeepCopy() *KeycloakApiRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakApiRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakApiUser) DeepCopyInto(out *KeycloakApiUser) {
	*out = *in
	if in.RealmRoles != nil {
		in, out := &in.RealmRoles, &out.RealmRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClientRoles != nil {
		in, out := &in.ClientRoles, &out.ClientRoles
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.RequiredActions != nil {
		in, out := &in.RequiredActions, &out.RequiredActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakApiUser.
func (in *KeycloakApiUser) DeepCopy() *KeycloakApiUser {
	if in == nil {
		return nil
	}
	out := new(KeycloakApiUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakBackup) DeepCopyInto(out *KeycloakBackup) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakBackup.
func (in *KeycloakBackup) DeepCopy() *KeycloakBackup {
	if in == nil {
		return nil
	}
	out := new(KeycloakBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClient) DeepCopyInto(out *KeycloakClient) {
	*out = *in
	if in.KeycloakApiClient != nil {
		in, out := &in.KeycloakApiClient, &out.KeycloakApiClient
		*out = new(KeycloakApiClient)
		(*in).DeepCopyInto(*out)
	}
	if in.OutputSecret != nil {
		in, out := &in.OutputSecret, &out.OutputSecret
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClient.
func (in *KeycloakClient) DeepCopy() *KeycloakClient {
	if in == nil {
		return nil
	}
	out := new(KeycloakClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakClientPair) DeepCopyInto(out *KeycloakClientPair) {
	*out = *in
	if in.KcClient != nil {
		in, out := &in.KcClient, &out.KcClient
		*out = new(KeycloakClient)
		(*in).DeepCopyInto(*out)
	}
	if in.SpecClient != nil {
		in, out := &in.SpecClient, &out.SpecClient
		*out = new(KeycloakClient)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakClientPair.
func (in *KeycloakClientPair) DeepCopy() *KeycloakClientPair {
	if in == nil {
		return nil
	}
	out := new(KeycloakClientPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakIdentityProvider) DeepCopyInto(out *KeycloakIdentityProvider) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakIdentityProvider.
func (in *KeycloakIdentityProvider) DeepCopy() *KeycloakIdentityProvider {
	if in == nil {
		return nil
	}
	out := new(KeycloakIdentityProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakIdentityProviderPair) DeepCopyInto(out *KeycloakIdentityProviderPair) {
	*out = *in
	if in.KcIdentityProvider != nil {
		in, out := &in.KcIdentityProvider, &out.KcIdentityProvider
		*out = new(KeycloakIdentityProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.SpecIdentityProvider != nil {
		in, out := &in.SpecIdentityProvider, &out.SpecIdentityProvider
		*out = new(KeycloakIdentityProvider)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakIdentityProviderPair.
func (in *KeycloakIdentityProviderPair) DeepCopy() *KeycloakIdentityProviderPair {
	if in == nil {
		return nil
	}
	out := new(KeycloakIdentityProviderPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakList) DeepCopyInto(out *KeycloakList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Keycloak, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakList.
func (in *KeycloakList) DeepCopy() *KeycloakList {
	if in == nil {
		return nil
	}
	out := new(KeycloakList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakProtocolMapper) DeepCopyInto(out *KeycloakProtocolMapper) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakProtocolMapper.
func (in *KeycloakProtocolMapper) DeepCopy() *KeycloakProtocolMapper {
	if in == nil {
		return nil
	}
	out := new(KeycloakProtocolMapper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealm) DeepCopyInto(out *KeycloakRealm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealm.
func (in *KeycloakRealm) DeepCopy() *KeycloakRealm {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmList) DeepCopyInto(out *KeycloakRealmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeycloakRealm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmList.
func (in *KeycloakRealmList) DeepCopy() *KeycloakRealmList {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeycloakRealmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmSpec) DeepCopyInto(out *KeycloakRealmSpec) {
	*out = *in
	if in.KeycloakApiRealm != nil {
		in, out := &in.KeycloakApiRealm, &out.KeycloakApiRealm
		*out = new(KeycloakApiRealm)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmSpec.
func (in *KeycloakRealmSpec) DeepCopy() *KeycloakRealmSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakRealmStatus) DeepCopyInto(out *KeycloakRealmStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakRealmStatus.
func (in *KeycloakRealmStatus) DeepCopy() *KeycloakRealmStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakRealmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakSpec) DeepCopyInto(out *KeycloakSpec) {
	*out = *in
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]KeycloakBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakSpec.
func (in *KeycloakSpec) DeepCopy() *KeycloakSpec {
	if in == nil {
		return nil
	}
	out := new(KeycloakSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakStatus) DeepCopyInto(out *KeycloakStatus) {
	*out = *in
	out.GenericStatus = in.GenericStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakStatus.
func (in *KeycloakStatus) DeepCopy() *KeycloakStatus {
	if in == nil {
		return nil
	}
	out := new(KeycloakStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUser) DeepCopyInto(out *KeycloakUser) {
	*out = *in
	if in.KeycloakApiUser != nil {
		in, out := &in.KeycloakApiUser, &out.KeycloakApiUser
		*out = new(KeycloakApiUser)
		(*in).DeepCopyInto(*out)
	}
	if in.OutputSecret != nil {
		in, out := &in.OutputSecret, &out.OutputSecret
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUser.
func (in *KeycloakUser) DeepCopy() *KeycloakUser {
	if in == nil {
		return nil
	}
	out := new(KeycloakUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUserClientRole) DeepCopyInto(out *KeycloakUserClientRole) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUserClientRole.
func (in *KeycloakUserClientRole) DeepCopy() *KeycloakUserClientRole {
	if in == nil {
		return nil
	}
	out := new(KeycloakUserClientRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakUserPair) DeepCopyInto(out *KeycloakUserPair) {
	*out = *in
	if in.KcUser != nil {
		in, out := &in.KcUser, &out.KcUser
		*out = new(KeycloakUser)
		(*in).DeepCopyInto(*out)
	}
	if in.SpecUser != nil {
		in, out := &in.SpecUser, &out.SpecUser
		*out = new(KeycloakUser)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakUserPair.
func (in *KeycloakUserPair) DeepCopy() *KeycloakUserPair {
	if in == nil {
		return nil
	}
	out := new(KeycloakUserPair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenResponse) DeepCopyInto(out *TokenResponse) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenResponse.
func (in *TokenResponse) DeepCopy() *TokenResponse {
	if in == nil {
		return nil
	}
	out := new(TokenResponse)
	in.DeepCopyInto(out)
	return out
}
