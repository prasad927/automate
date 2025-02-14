// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/compliance/reporting/reporting.proto

package reporting

import policy "github.com/chef/automate/api/external/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListReports", "compliance:reporting:reports", "compliance:reports:list", "POST", "/api/v0/compliance/reporting/reports", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListReportIds", "compliance:reporting:report-ids", "compliance:reportids:list", "POST", "/api/v0/compliance/reporting/report-ids", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListControlItems", "compliance:reporting:control", "compliance:controlItems:list", "POST", "/api/v0/compliance/reporting/controls", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ControlItemRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "text":
					return m.Text
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListControlInfo", "compliance:reporting:reportcontrols:{id}", "compliance:reportcontrols:list", "POST", "/api/v0/compliance/reporting/reportcontrols/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ReadReport", "compliance:reporting:reports:{id}", "compliance:reports:get", "POST", "/api/v0/compliance/reporting/reports/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ReadNodeHeader", "compliance:reporting:nodeheader:{id}", "compliance:nodeheader:get", "POST", "/api/v0/compliance/reporting/nodeheader/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListSuggestions", "compliance:reporting:suggestions", "compliance:reportSuggestions:list", "POST", "/api/v0/compliance/reporting/suggestions", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*SuggestionRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "type":
					return m.Type
				case "text":
					return m.Text
				case "type_key":
					return m.TypeKey
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListProfiles", "compliance:reporting:profiles", "compliance:reportProfiles:list", "POST", "/api/v0/compliance/reporting/profiles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ExportReportManager", "compliance:reporting:reports", "compliance:reports:list", "POST", "/api/v0/compliance/reporting/reportmanager/export", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ReadNode", "compliance:reporting:nodes:{id}", "compliance:reportNodes:get", "GET", "/api/v0/compliance/reporting/nodes/id/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Id); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/ListNodes", "compliance:reporting:nodes", "compliance:reportNodes:list", "POST", "/api/v0/compliance/reporting/nodes/search", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*Query); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "type":
					return m.Type
				case "sort":
					return m.Sort
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/api/v0/compliance/reporting/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.compliance.reporting.v1.ReportingService/LicenseUsageNodes", "compliance:reporting:licenseusage", "compliance:reportingLicenseUsage:list", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
