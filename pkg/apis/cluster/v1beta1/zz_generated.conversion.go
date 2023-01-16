//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Alipay Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	url "net/url"
	unsafe "unsafe"

	cluster "code.alipay.com/ant-iac/karbour/pkg/apis/cluster"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClusterAccess)(nil), (*cluster.ClusterAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterAccess_To_cluster_ClusterAccess(a.(*ClusterAccess), b.(*cluster.ClusterAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ClusterAccess)(nil), (*ClusterAccess)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ClusterAccess_To_v1beta1_ClusterAccess(a.(*cluster.ClusterAccess), b.(*ClusterAccess), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterAccessCredential)(nil), (*cluster.ClusterAccessCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterAccessCredential_To_cluster_ClusterAccessCredential(a.(*ClusterAccessCredential), b.(*cluster.ClusterAccessCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ClusterAccessCredential)(nil), (*ClusterAccessCredential)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ClusterAccessCredential_To_v1beta1_ClusterAccessCredential(a.(*cluster.ClusterAccessCredential), b.(*ClusterAccessCredential), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterExtension)(nil), (*cluster.ClusterExtension)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterExtension_To_cluster_ClusterExtension(a.(*ClusterExtension), b.(*cluster.ClusterExtension), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ClusterExtension)(nil), (*ClusterExtension)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ClusterExtension_To_v1beta1_ClusterExtension(a.(*cluster.ClusterExtension), b.(*ClusterExtension), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterExtensionList)(nil), (*cluster.ClusterExtensionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterExtensionList_To_cluster_ClusterExtensionList(a.(*ClusterExtensionList), b.(*cluster.ClusterExtensionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ClusterExtensionList)(nil), (*ClusterExtensionList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ClusterExtensionList_To_v1beta1_ClusterExtensionList(a.(*cluster.ClusterExtensionList), b.(*ClusterExtensionList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterExtensionProxyOptions)(nil), (*cluster.ClusterExtensionProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterExtensionProxyOptions_To_cluster_ClusterExtensionProxyOptions(a.(*ClusterExtensionProxyOptions), b.(*cluster.ClusterExtensionProxyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ClusterExtensionProxyOptions)(nil), (*ClusterExtensionProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ClusterExtensionProxyOptions_To_v1beta1_ClusterExtensionProxyOptions(a.(*cluster.ClusterExtensionProxyOptions), b.(*ClusterExtensionProxyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterExtensionSpec)(nil), (*cluster.ClusterExtensionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterExtensionSpec_To_cluster_ClusterExtensionSpec(a.(*ClusterExtensionSpec), b.(*cluster.ClusterExtensionSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ClusterExtensionSpec)(nil), (*ClusterExtensionSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ClusterExtensionSpec_To_v1beta1_ClusterExtensionSpec(a.(*cluster.ClusterExtensionSpec), b.(*ClusterExtensionSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterExtensionStatus)(nil), (*cluster.ClusterExtensionStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterExtensionStatus_To_cluster_ClusterExtensionStatus(a.(*ClusterExtensionStatus), b.(*cluster.ClusterExtensionStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.ClusterExtensionStatus)(nil), (*ClusterExtensionStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_ClusterExtensionStatus_To_v1beta1_ClusterExtensionStatus(a.(*cluster.ClusterExtensionStatus), b.(*ClusterExtensionStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*X509)(nil), (*cluster.X509)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_X509_To_cluster_X509(a.(*X509), b.(*cluster.X509), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*cluster.X509)(nil), (*X509)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_cluster_X509_To_v1beta1_X509(a.(*cluster.X509), b.(*X509), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*url.Values)(nil), (*ClusterExtensionProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_url_Values_To_v1beta1_ClusterExtensionProxyOptions(a.(*url.Values), b.(*ClusterExtensionProxyOptions), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_ClusterAccess_To_cluster_ClusterAccess(in *ClusterAccess, out *cluster.ClusterAccess, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.Insecure = (*bool)(unsafe.Pointer(in.Insecure))
	out.Credential = (*cluster.ClusterAccessCredential)(unsafe.Pointer(in.Credential))
	return nil
}

// Convert_v1beta1_ClusterAccess_To_cluster_ClusterAccess is an autogenerated conversion function.
func Convert_v1beta1_ClusterAccess_To_cluster_ClusterAccess(in *ClusterAccess, out *cluster.ClusterAccess, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterAccess_To_cluster_ClusterAccess(in, out, s)
}

func autoConvert_cluster_ClusterAccess_To_v1beta1_ClusterAccess(in *cluster.ClusterAccess, out *ClusterAccess, s conversion.Scope) error {
	out.Endpoint = in.Endpoint
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.Insecure = (*bool)(unsafe.Pointer(in.Insecure))
	out.Credential = (*ClusterAccessCredential)(unsafe.Pointer(in.Credential))
	return nil
}

// Convert_cluster_ClusterAccess_To_v1beta1_ClusterAccess is an autogenerated conversion function.
func Convert_cluster_ClusterAccess_To_v1beta1_ClusterAccess(in *cluster.ClusterAccess, out *ClusterAccess, s conversion.Scope) error {
	return autoConvert_cluster_ClusterAccess_To_v1beta1_ClusterAccess(in, out, s)
}

func autoConvert_v1beta1_ClusterAccessCredential_To_cluster_ClusterAccessCredential(in *ClusterAccessCredential, out *cluster.ClusterAccessCredential, s conversion.Scope) error {
	out.Type = cluster.CredentialType(in.Type)
	out.ServiceAccountToken = in.ServiceAccountToken
	out.X509 = (*cluster.X509)(unsafe.Pointer(in.X509))
	return nil
}

// Convert_v1beta1_ClusterAccessCredential_To_cluster_ClusterAccessCredential is an autogenerated conversion function.
func Convert_v1beta1_ClusterAccessCredential_To_cluster_ClusterAccessCredential(in *ClusterAccessCredential, out *cluster.ClusterAccessCredential, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterAccessCredential_To_cluster_ClusterAccessCredential(in, out, s)
}

func autoConvert_cluster_ClusterAccessCredential_To_v1beta1_ClusterAccessCredential(in *cluster.ClusterAccessCredential, out *ClusterAccessCredential, s conversion.Scope) error {
	out.Type = CredentialType(in.Type)
	out.ServiceAccountToken = in.ServiceAccountToken
	out.X509 = (*X509)(unsafe.Pointer(in.X509))
	return nil
}

// Convert_cluster_ClusterAccessCredential_To_v1beta1_ClusterAccessCredential is an autogenerated conversion function.
func Convert_cluster_ClusterAccessCredential_To_v1beta1_ClusterAccessCredential(in *cluster.ClusterAccessCredential, out *ClusterAccessCredential, s conversion.Scope) error {
	return autoConvert_cluster_ClusterAccessCredential_To_v1beta1_ClusterAccessCredential(in, out, s)
}

func autoConvert_v1beta1_ClusterExtension_To_cluster_ClusterExtension(in *ClusterExtension, out *cluster.ClusterExtension, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ClusterExtensionSpec_To_cluster_ClusterExtensionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ClusterExtensionStatus_To_cluster_ClusterExtensionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_ClusterExtension_To_cluster_ClusterExtension is an autogenerated conversion function.
func Convert_v1beta1_ClusterExtension_To_cluster_ClusterExtension(in *ClusterExtension, out *cluster.ClusterExtension, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterExtension_To_cluster_ClusterExtension(in, out, s)
}

func autoConvert_cluster_ClusterExtension_To_v1beta1_ClusterExtension(in *cluster.ClusterExtension, out *ClusterExtension, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cluster_ClusterExtensionSpec_To_v1beta1_ClusterExtensionSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cluster_ClusterExtensionStatus_To_v1beta1_ClusterExtensionStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cluster_ClusterExtension_To_v1beta1_ClusterExtension is an autogenerated conversion function.
func Convert_cluster_ClusterExtension_To_v1beta1_ClusterExtension(in *cluster.ClusterExtension, out *ClusterExtension, s conversion.Scope) error {
	return autoConvert_cluster_ClusterExtension_To_v1beta1_ClusterExtension(in, out, s)
}

func autoConvert_v1beta1_ClusterExtensionList_To_cluster_ClusterExtensionList(in *ClusterExtensionList, out *cluster.ClusterExtensionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.ClusterExtension)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_ClusterExtensionList_To_cluster_ClusterExtensionList is an autogenerated conversion function.
func Convert_v1beta1_ClusterExtensionList_To_cluster_ClusterExtensionList(in *ClusterExtensionList, out *cluster.ClusterExtensionList, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterExtensionList_To_cluster_ClusterExtensionList(in, out, s)
}

func autoConvert_cluster_ClusterExtensionList_To_v1beta1_ClusterExtensionList(in *cluster.ClusterExtensionList, out *ClusterExtensionList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterExtension)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_ClusterExtensionList_To_v1beta1_ClusterExtensionList is an autogenerated conversion function.
func Convert_cluster_ClusterExtensionList_To_v1beta1_ClusterExtensionList(in *cluster.ClusterExtensionList, out *ClusterExtensionList, s conversion.Scope) error {
	return autoConvert_cluster_ClusterExtensionList_To_v1beta1_ClusterExtensionList(in, out, s)
}

func autoConvert_v1beta1_ClusterExtensionProxyOptions_To_cluster_ClusterExtensionProxyOptions(in *ClusterExtensionProxyOptions, out *cluster.ClusterExtensionProxyOptions, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_v1beta1_ClusterExtensionProxyOptions_To_cluster_ClusterExtensionProxyOptions is an autogenerated conversion function.
func Convert_v1beta1_ClusterExtensionProxyOptions_To_cluster_ClusterExtensionProxyOptions(in *ClusterExtensionProxyOptions, out *cluster.ClusterExtensionProxyOptions, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterExtensionProxyOptions_To_cluster_ClusterExtensionProxyOptions(in, out, s)
}

func autoConvert_cluster_ClusterExtensionProxyOptions_To_v1beta1_ClusterExtensionProxyOptions(in *cluster.ClusterExtensionProxyOptions, out *ClusterExtensionProxyOptions, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_cluster_ClusterExtensionProxyOptions_To_v1beta1_ClusterExtensionProxyOptions is an autogenerated conversion function.
func Convert_cluster_ClusterExtensionProxyOptions_To_v1beta1_ClusterExtensionProxyOptions(in *cluster.ClusterExtensionProxyOptions, out *ClusterExtensionProxyOptions, s conversion.Scope) error {
	return autoConvert_cluster_ClusterExtensionProxyOptions_To_v1beta1_ClusterExtensionProxyOptions(in, out, s)
}

func autoConvert_url_Values_To_v1beta1_ClusterExtensionProxyOptions(in *url.Values, out *ClusterExtensionProxyOptions, s conversion.Scope) error {
	// WARNING: Field TypeMeta does not have json tag, skipping.

	if values, ok := map[string][]string(*in)["path"]; ok && len(values) > 0 {
		if err := runtime.Convert_Slice_string_To_string(&values, &out.Path, s); err != nil {
			return err
		}
	} else {
		out.Path = ""
	}
	return nil
}

// Convert_url_Values_To_v1beta1_ClusterExtensionProxyOptions is an autogenerated conversion function.
func Convert_url_Values_To_v1beta1_ClusterExtensionProxyOptions(in *url.Values, out *ClusterExtensionProxyOptions, s conversion.Scope) error {
	return autoConvert_url_Values_To_v1beta1_ClusterExtensionProxyOptions(in, out, s)
}

func autoConvert_v1beta1_ClusterExtensionSpec_To_cluster_ClusterExtensionSpec(in *ClusterExtensionSpec, out *cluster.ClusterExtensionSpec, s conversion.Scope) error {
	out.Provider = in.Provider
	if err := Convert_v1beta1_ClusterAccess_To_cluster_ClusterAccess(&in.Access, &out.Access, s); err != nil {
		return err
	}
	out.Finalized = (*bool)(unsafe.Pointer(in.Finalized))
	return nil
}

// Convert_v1beta1_ClusterExtensionSpec_To_cluster_ClusterExtensionSpec is an autogenerated conversion function.
func Convert_v1beta1_ClusterExtensionSpec_To_cluster_ClusterExtensionSpec(in *ClusterExtensionSpec, out *cluster.ClusterExtensionSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterExtensionSpec_To_cluster_ClusterExtensionSpec(in, out, s)
}

func autoConvert_cluster_ClusterExtensionSpec_To_v1beta1_ClusterExtensionSpec(in *cluster.ClusterExtensionSpec, out *ClusterExtensionSpec, s conversion.Scope) error {
	out.Provider = in.Provider
	if err := Convert_cluster_ClusterAccess_To_v1beta1_ClusterAccess(&in.Access, &out.Access, s); err != nil {
		return err
	}
	out.Finalized = (*bool)(unsafe.Pointer(in.Finalized))
	return nil
}

// Convert_cluster_ClusterExtensionSpec_To_v1beta1_ClusterExtensionSpec is an autogenerated conversion function.
func Convert_cluster_ClusterExtensionSpec_To_v1beta1_ClusterExtensionSpec(in *cluster.ClusterExtensionSpec, out *ClusterExtensionSpec, s conversion.Scope) error {
	return autoConvert_cluster_ClusterExtensionSpec_To_v1beta1_ClusterExtensionSpec(in, out, s)
}

func autoConvert_v1beta1_ClusterExtensionStatus_To_cluster_ClusterExtensionStatus(in *ClusterExtensionStatus, out *cluster.ClusterExtensionStatus, s conversion.Scope) error {
	out.Healthy = in.Healthy
	return nil
}

// Convert_v1beta1_ClusterExtensionStatus_To_cluster_ClusterExtensionStatus is an autogenerated conversion function.
func Convert_v1beta1_ClusterExtensionStatus_To_cluster_ClusterExtensionStatus(in *ClusterExtensionStatus, out *cluster.ClusterExtensionStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_ClusterExtensionStatus_To_cluster_ClusterExtensionStatus(in, out, s)
}

func autoConvert_cluster_ClusterExtensionStatus_To_v1beta1_ClusterExtensionStatus(in *cluster.ClusterExtensionStatus, out *ClusterExtensionStatus, s conversion.Scope) error {
	out.Healthy = in.Healthy
	return nil
}

// Convert_cluster_ClusterExtensionStatus_To_v1beta1_ClusterExtensionStatus is an autogenerated conversion function.
func Convert_cluster_ClusterExtensionStatus_To_v1beta1_ClusterExtensionStatus(in *cluster.ClusterExtensionStatus, out *ClusterExtensionStatus, s conversion.Scope) error {
	return autoConvert_cluster_ClusterExtensionStatus_To_v1beta1_ClusterExtensionStatus(in, out, s)
}

func autoConvert_v1beta1_X509_To_cluster_X509(in *X509, out *cluster.X509, s conversion.Scope) error {
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	out.PrivateKey = *(*[]byte)(unsafe.Pointer(&in.PrivateKey))
	return nil
}

// Convert_v1beta1_X509_To_cluster_X509 is an autogenerated conversion function.
func Convert_v1beta1_X509_To_cluster_X509(in *X509, out *cluster.X509, s conversion.Scope) error {
	return autoConvert_v1beta1_X509_To_cluster_X509(in, out, s)
}

func autoConvert_cluster_X509_To_v1beta1_X509(in *cluster.X509, out *X509, s conversion.Scope) error {
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	out.PrivateKey = *(*[]byte)(unsafe.Pointer(&in.PrivateKey))
	return nil
}

// Convert_cluster_X509_To_v1beta1_X509 is an autogenerated conversion function.
func Convert_cluster_X509_To_v1beta1_X509(in *cluster.X509, out *X509, s conversion.Scope) error {
	return autoConvert_cluster_X509_To_v1beta1_X509(in, out, s)
}
