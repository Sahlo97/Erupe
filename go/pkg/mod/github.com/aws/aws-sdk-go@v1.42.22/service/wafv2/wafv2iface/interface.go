// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package wafv2iface provides an interface to enable mocking the AWS WAFV2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package wafv2iface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/wafv2"
)

// WAFV2API provides an interface to enable mocking the
// wafv2.WAFV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS WAFV2.
//    func myFunc(svc wafv2iface.WAFV2API) bool {
//        // Make svc.AssociateWebACL request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := wafv2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockWAFV2Client struct {
//        wafv2iface.WAFV2API
//    }
//    func (m *mockWAFV2Client) AssociateWebACL(input *wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockWAFV2Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type WAFV2API interface {
	AssociateWebACL(*wafv2.AssociateWebACLInput) (*wafv2.AssociateWebACLOutput, error)
	AssociateWebACLWithContext(aws.Context, *wafv2.AssociateWebACLInput, ...request.Option) (*wafv2.AssociateWebACLOutput, error)
	AssociateWebACLRequest(*wafv2.AssociateWebACLInput) (*request.Request, *wafv2.AssociateWebACLOutput)

	CheckCapacity(*wafv2.CheckCapacityInput) (*wafv2.CheckCapacityOutput, error)
	CheckCapacityWithContext(aws.Context, *wafv2.CheckCapacityInput, ...request.Option) (*wafv2.CheckCapacityOutput, error)
	CheckCapacityRequest(*wafv2.CheckCapacityInput) (*request.Request, *wafv2.CheckCapacityOutput)

	CreateIPSet(*wafv2.CreateIPSetInput) (*wafv2.CreateIPSetOutput, error)
	CreateIPSetWithContext(aws.Context, *wafv2.CreateIPSetInput, ...request.Option) (*wafv2.CreateIPSetOutput, error)
	CreateIPSetRequest(*wafv2.CreateIPSetInput) (*request.Request, *wafv2.CreateIPSetOutput)

	CreateRegexPatternSet(*wafv2.CreateRegexPatternSetInput) (*wafv2.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetWithContext(aws.Context, *wafv2.CreateRegexPatternSetInput, ...request.Option) (*wafv2.CreateRegexPatternSetOutput, error)
	CreateRegexPatternSetRequest(*wafv2.CreateRegexPatternSetInput) (*request.Request, *wafv2.CreateRegexPatternSetOutput)

	CreateRuleGroup(*wafv2.CreateRuleGroupInput) (*wafv2.CreateRuleGroupOutput, error)
	CreateRuleGroupWithContext(aws.Context, *wafv2.CreateRuleGroupInput, ...request.Option) (*wafv2.CreateRuleGroupOutput, error)
	CreateRuleGroupRequest(*wafv2.CreateRuleGroupInput) (*request.Request, *wafv2.CreateRuleGroupOutput)

	CreateWebACL(*wafv2.CreateWebACLInput) (*wafv2.CreateWebACLOutput, error)
	CreateWebACLWithContext(aws.Context, *wafv2.CreateWebACLInput, ...request.Option) (*wafv2.CreateWebACLOutput, error)
	CreateWebACLRequest(*wafv2.CreateWebACLInput) (*request.Request, *wafv2.CreateWebACLOutput)

	DeleteFirewallManagerRuleGroups(*wafv2.DeleteFirewallManagerRuleGroupsInput) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error)
	DeleteFirewallManagerRuleGroupsWithContext(aws.Context, *wafv2.DeleteFirewallManagerRuleGroupsInput, ...request.Option) (*wafv2.DeleteFirewallManagerRuleGroupsOutput, error)
	DeleteFirewallManagerRuleGroupsRequest(*wafv2.DeleteFirewallManagerRuleGroupsInput) (*request.Request, *wafv2.DeleteFirewallManagerRuleGroupsOutput)

	DeleteIPSet(*wafv2.DeleteIPSetInput) (*wafv2.DeleteIPSetOutput, error)
	DeleteIPSetWithContext(aws.Context, *wafv2.DeleteIPSetInput, ...request.Option) (*wafv2.DeleteIPSetOutput, error)
	DeleteIPSetRequest(*wafv2.DeleteIPSetInput) (*request.Request, *wafv2.DeleteIPSetOutput)

	DeleteLoggingConfiguration(*wafv2.DeleteLoggingConfigurationInput) (*wafv2.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationWithContext(aws.Context, *wafv2.DeleteLoggingConfigurationInput, ...request.Option) (*wafv2.DeleteLoggingConfigurationOutput, error)
	DeleteLoggingConfigurationRequest(*wafv2.DeleteLoggingConfigurationInput) (*request.Request, *wafv2.DeleteLoggingConfigurationOutput)

	DeletePermissionPolicy(*wafv2.DeletePermissionPolicyInput) (*wafv2.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyWithContext(aws.Context, *wafv2.DeletePermissionPolicyInput, ...request.Option) (*wafv2.DeletePermissionPolicyOutput, error)
	DeletePermissionPolicyRequest(*wafv2.DeletePermissionPolicyInput) (*request.Request, *wafv2.DeletePermissionPolicyOutput)

	DeleteRegexPatternSet(*wafv2.DeleteRegexPatternSetInput) (*wafv2.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetWithContext(aws.Context, *wafv2.DeleteRegexPatternSetInput, ...request.Option) (*wafv2.DeleteRegexPatternSetOutput, error)
	DeleteRegexPatternSetRequest(*wafv2.DeleteRegexPatternSetInput) (*request.Request, *wafv2.DeleteRegexPatternSetOutput)

	DeleteRuleGroup(*wafv2.DeleteRuleGroupInput) (*wafv2.DeleteRuleGroupOutput, error)
	DeleteRuleGroupWithContext(aws.Context, *wafv2.DeleteRuleGroupInput, ...request.Option) (*wafv2.DeleteRuleGroupOutput, error)
	DeleteRuleGroupRequest(*wafv2.DeleteRuleGroupInput) (*request.Request, *wafv2.DeleteRuleGroupOutput)

	DeleteWebACL(*wafv2.DeleteWebACLInput) (*wafv2.DeleteWebACLOutput, error)
	DeleteWebACLWithContext(aws.Context, *wafv2.DeleteWebACLInput, ...request.Option) (*wafv2.DeleteWebACLOutput, error)
	DeleteWebACLRequest(*wafv2.DeleteWebACLInput) (*request.Request, *wafv2.DeleteWebACLOutput)

	DescribeManagedRuleGroup(*wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error)
	DescribeManagedRuleGroupWithContext(aws.Context, *wafv2.DescribeManagedRuleGroupInput, ...request.Option) (*wafv2.DescribeManagedRuleGroupOutput, error)
	DescribeManagedRuleGroupRequest(*wafv2.DescribeManagedRuleGroupInput) (*request.Request, *wafv2.DescribeManagedRuleGroupOutput)

	DisassociateWebACL(*wafv2.DisassociateWebACLInput) (*wafv2.DisassociateWebACLOutput, error)
	DisassociateWebACLWithContext(aws.Context, *wafv2.DisassociateWebACLInput, ...request.Option) (*wafv2.DisassociateWebACLOutput, error)
	DisassociateWebACLRequest(*wafv2.DisassociateWebACLInput) (*request.Request, *wafv2.DisassociateWebACLOutput)

	GetIPSet(*wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error)
	GetIPSetWithContext(aws.Context, *wafv2.GetIPSetInput, ...request.Option) (*wafv2.GetIPSetOutput, error)
	GetIPSetRequest(*wafv2.GetIPSetInput) (*request.Request, *wafv2.GetIPSetOutput)

	GetLoggingConfiguration(*wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationWithContext(aws.Context, *wafv2.GetLoggingConfigurationInput, ...request.Option) (*wafv2.GetLoggingConfigurationOutput, error)
	GetLoggingConfigurationRequest(*wafv2.GetLoggingConfigurationInput) (*request.Request, *wafv2.GetLoggingConfigurationOutput)

	GetManagedRuleSet(*wafv2.GetManagedRuleSetInput) (*wafv2.GetManagedRuleSetOutput, error)
	GetManagedRuleSetWithContext(aws.Context, *wafv2.GetManagedRuleSetInput, ...request.Option) (*wafv2.GetManagedRuleSetOutput, error)
	GetManagedRuleSetRequest(*wafv2.GetManagedRuleSetInput) (*request.Request, *wafv2.GetManagedRuleSetOutput)

	GetPermissionPolicy(*wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error)
	GetPermissionPolicyWithContext(aws.Context, *wafv2.GetPermissionPolicyInput, ...request.Option) (*wafv2.GetPermissionPolicyOutput, error)
	GetPermissionPolicyRequest(*wafv2.GetPermissionPolicyInput) (*request.Request, *wafv2.GetPermissionPolicyOutput)

	GetRateBasedStatementManagedKeys(*wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error)
	GetRateBasedStatementManagedKeysWithContext(aws.Context, *wafv2.GetRateBasedStatementManagedKeysInput, ...request.Option) (*wafv2.GetRateBasedStatementManagedKeysOutput, error)
	GetRateBasedStatementManagedKeysRequest(*wafv2.GetRateBasedStatementManagedKeysInput) (*request.Request, *wafv2.GetRateBasedStatementManagedKeysOutput)

	GetRegexPatternSet(*wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error)
	GetRegexPatternSetWithContext(aws.Context, *wafv2.GetRegexPatternSetInput, ...request.Option) (*wafv2.GetRegexPatternSetOutput, error)
	GetRegexPatternSetRequest(*wafv2.GetRegexPatternSetInput) (*request.Request, *wafv2.GetRegexPatternSetOutput)

	GetRuleGroup(*wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error)
	GetRuleGroupWithContext(aws.Context, *wafv2.GetRuleGroupInput, ...request.Option) (*wafv2.GetRuleGroupOutput, error)
	GetRuleGroupRequest(*wafv2.GetRuleGroupInput) (*request.Request, *wafv2.GetRuleGroupOutput)

	GetSampledRequests(*wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error)
	GetSampledRequestsWithContext(aws.Context, *wafv2.GetSampledRequestsInput, ...request.Option) (*wafv2.GetSampledRequestsOutput, error)
	GetSampledRequestsRequest(*wafv2.GetSampledRequestsInput) (*request.Request, *wafv2.GetSampledRequestsOutput)

	GetWebACL(*wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error)
	GetWebACLWithContext(aws.Context, *wafv2.GetWebACLInput, ...request.Option) (*wafv2.GetWebACLOutput, error)
	GetWebACLRequest(*wafv2.GetWebACLInput) (*request.Request, *wafv2.GetWebACLOutput)

	GetWebACLForResource(*wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error)
	GetWebACLForResourceWithContext(aws.Context, *wafv2.GetWebACLForResourceInput, ...request.Option) (*wafv2.GetWebACLForResourceOutput, error)
	GetWebACLForResourceRequest(*wafv2.GetWebACLForResourceInput) (*request.Request, *wafv2.GetWebACLForResourceOutput)

	ListAvailableManagedRuleGroupVersions(*wafv2.ListAvailableManagedRuleGroupVersionsInput) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error)
	ListAvailableManagedRuleGroupVersionsWithContext(aws.Context, *wafv2.ListAvailableManagedRuleGroupVersionsInput, ...request.Option) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error)
	ListAvailableManagedRuleGroupVersionsRequest(*wafv2.ListAvailableManagedRuleGroupVersionsInput) (*request.Request, *wafv2.ListAvailableManagedRuleGroupVersionsOutput)

	ListAvailableManagedRuleGroups(*wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
	ListAvailableManagedRuleGroupsWithContext(aws.Context, *wafv2.ListAvailableManagedRuleGroupsInput, ...request.Option) (*wafv2.ListAvailableManagedRuleGroupsOutput, error)
	ListAvailableManagedRuleGroupsRequest(*wafv2.ListAvailableManagedRuleGroupsInput) (*request.Request, *wafv2.ListAvailableManagedRuleGroupsOutput)

	ListIPSets(*wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error)
	ListIPSetsWithContext(aws.Context, *wafv2.ListIPSetsInput, ...request.Option) (*wafv2.ListIPSetsOutput, error)
	ListIPSetsRequest(*wafv2.ListIPSetsInput) (*request.Request, *wafv2.ListIPSetsOutput)

	ListLoggingConfigurations(*wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsWithContext(aws.Context, *wafv2.ListLoggingConfigurationsInput, ...request.Option) (*wafv2.ListLoggingConfigurationsOutput, error)
	ListLoggingConfigurationsRequest(*wafv2.ListLoggingConfigurationsInput) (*request.Request, *wafv2.ListLoggingConfigurationsOutput)

	ListManagedRuleSets(*wafv2.ListManagedRuleSetsInput) (*wafv2.ListManagedRuleSetsOutput, error)
	ListManagedRuleSetsWithContext(aws.Context, *wafv2.ListManagedRuleSetsInput, ...request.Option) (*wafv2.ListManagedRuleSetsOutput, error)
	ListManagedRuleSetsRequest(*wafv2.ListManagedRuleSetsInput) (*request.Request, *wafv2.ListManagedRuleSetsOutput)

	ListRegexPatternSets(*wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsWithContext(aws.Context, *wafv2.ListRegexPatternSetsInput, ...request.Option) (*wafv2.ListRegexPatternSetsOutput, error)
	ListRegexPatternSetsRequest(*wafv2.ListRegexPatternSetsInput) (*request.Request, *wafv2.ListRegexPatternSetsOutput)

	ListResourcesForWebACL(*wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error)
	ListResourcesForWebACLWithContext(aws.Context, *wafv2.ListResourcesForWebACLInput, ...request.Option) (*wafv2.ListResourcesForWebACLOutput, error)
	ListResourcesForWebACLRequest(*wafv2.ListResourcesForWebACLInput) (*request.Request, *wafv2.ListResourcesForWebACLOutput)

	ListRuleGroups(*wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error)
	ListRuleGroupsWithContext(aws.Context, *wafv2.ListRuleGroupsInput, ...request.Option) (*wafv2.ListRuleGroupsOutput, error)
	ListRuleGroupsRequest(*wafv2.ListRuleGroupsInput) (*request.Request, *wafv2.ListRuleGroupsOutput)

	ListTagsForResource(*wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *wafv2.ListTagsForResourceInput, ...request.Option) (*wafv2.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*wafv2.ListTagsForResourceInput) (*request.Request, *wafv2.ListTagsForResourceOutput)

	ListWebACLs(*wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error)
	ListWebACLsWithContext(aws.Context, *wafv2.ListWebACLsInput, ...request.Option) (*wafv2.ListWebACLsOutput, error)
	ListWebACLsRequest(*wafv2.ListWebACLsInput) (*request.Request, *wafv2.ListWebACLsOutput)

	PutLoggingConfiguration(*wafv2.PutLoggingConfigurationInput) (*wafv2.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationWithContext(aws.Context, *wafv2.PutLoggingConfigurationInput, ...request.Option) (*wafv2.PutLoggingConfigurationOutput, error)
	PutLoggingConfigurationRequest(*wafv2.PutLoggingConfigurationInput) (*request.Request, *wafv2.PutLoggingConfigurationOutput)

	PutManagedRuleSetVersions(*wafv2.PutManagedRuleSetVersionsInput) (*wafv2.PutManagedRuleSetVersionsOutput, error)
	PutManagedRuleSetVersionsWithContext(aws.Context, *wafv2.PutManagedRuleSetVersionsInput, ...request.Option) (*wafv2.PutManagedRuleSetVersionsOutput, error)
	PutManagedRuleSetVersionsRequest(*wafv2.PutManagedRuleSetVersionsInput) (*request.Request, *wafv2.PutManagedRuleSetVersionsOutput)

	PutPermissionPolicy(*wafv2.PutPermissionPolicyInput) (*wafv2.PutPermissionPolicyOutput, error)
	PutPermissionPolicyWithContext(aws.Context, *wafv2.PutPermissionPolicyInput, ...request.Option) (*wafv2.PutPermissionPolicyOutput, error)
	PutPermissionPolicyRequest(*wafv2.PutPermissionPolicyInput) (*request.Request, *wafv2.PutPermissionPolicyOutput)

	TagResource(*wafv2.TagResourceInput) (*wafv2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *wafv2.TagResourceInput, ...request.Option) (*wafv2.TagResourceOutput, error)
	TagResourceRequest(*wafv2.TagResourceInput) (*request.Request, *wafv2.TagResourceOutput)

	UntagResource(*wafv2.UntagResourceInput) (*wafv2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *wafv2.UntagResourceInput, ...request.Option) (*wafv2.UntagResourceOutput, error)
	UntagResourceRequest(*wafv2.UntagResourceInput) (*request.Request, *wafv2.UntagResourceOutput)

	UpdateIPSet(*wafv2.UpdateIPSetInput) (*wafv2.UpdateIPSetOutput, error)
	UpdateIPSetWithContext(aws.Context, *wafv2.UpdateIPSetInput, ...request.Option) (*wafv2.UpdateIPSetOutput, error)
	UpdateIPSetRequest(*wafv2.UpdateIPSetInput) (*request.Request, *wafv2.UpdateIPSetOutput)

	UpdateManagedRuleSetVersionExpiryDate(*wafv2.UpdateManagedRuleSetVersionExpiryDateInput) (*wafv2.UpdateManagedRuleSetVersionExpiryDateOutput, error)
	UpdateManagedRuleSetVersionExpiryDateWithContext(aws.Context, *wafv2.UpdateManagedRuleSetVersionExpiryDateInput, ...request.Option) (*wafv2.UpdateManagedRuleSetVersionExpiryDateOutput, error)
	UpdateManagedRuleSetVersionExpiryDateRequest(*wafv2.UpdateManagedRuleSetVersionExpiryDateInput) (*request.Request, *wafv2.UpdateManagedRuleSetVersionExpiryDateOutput)

	UpdateRegexPatternSet(*wafv2.UpdateRegexPatternSetInput) (*wafv2.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetWithContext(aws.Context, *wafv2.UpdateRegexPatternSetInput, ...request.Option) (*wafv2.UpdateRegexPatternSetOutput, error)
	UpdateRegexPatternSetRequest(*wafv2.UpdateRegexPatternSetInput) (*request.Request, *wafv2.UpdateRegexPatternSetOutput)

	UpdateRuleGroup(*wafv2.UpdateRuleGroupInput) (*wafv2.UpdateRuleGroupOutput, error)
	UpdateRuleGroupWithContext(aws.Context, *wafv2.UpdateRuleGroupInput, ...request.Option) (*wafv2.UpdateRuleGroupOutput, error)
	UpdateRuleGroupRequest(*wafv2.UpdateRuleGroupInput) (*request.Request, *wafv2.UpdateRuleGroupOutput)

	UpdateWebACL(*wafv2.UpdateWebACLInput) (*wafv2.UpdateWebACLOutput, error)
	UpdateWebACLWithContext(aws.Context, *wafv2.UpdateWebACLInput, ...request.Option) (*wafv2.UpdateWebACLOutput, error)
	UpdateWebACLRequest(*wafv2.UpdateWebACLInput) (*request.Request, *wafv2.UpdateWebACLOutput)
}

var _ WAFV2API = (*wafv2.WAFV2)(nil)
