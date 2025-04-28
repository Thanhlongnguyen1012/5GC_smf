package errors

import (
	"net/http"

	"github.com/Thanhlongnguyen1012/5GC_smf/internal/models"
)

var (
	DnnDeniedError = models.ProblemDetails{
		Title:  "DNN Denied",
		Status: http.StatusForbidden,
		Detail: "The subscriber does not have the necessary subscription to access the DNN",
		Cause:  "DNN_DENIED",
	}
	DnnNotSupported = models.ProblemDetails{
		Title:         "DNN Not Supported",
		Status:        http.StatusForbidden,
		Detail:        "The DNN is not supported by the SMF. ",
		Cause:         "DNN_NOT_SUPPORTED",
		InvalidParams: nil,
	}
	InsufficientResourceSliceDnn = models.ProblemDetails{
		Title:         "DNN Resource insufficient",
		Status:        http.StatusInternalServerError,
		Detail:        "The request cannot be provided due to insufficient resources for the specific slice and DNN.",
		Cause:         "INSUFFICIENT_RESOURCES_SLICE_DNN ",
		InvalidParams: nil,
	}
	IpAllocError = models.ProblemDetails{
		Title:         "IP Allocation Error",
		Status:        http.StatusInternalServerError,
		Detail:        "The request cannot be provided due to insufficient resources for the IP allocation.",
		Cause:         "INSUFFICIENT_RESOURCES",
		InvalidParams: nil,
	}
	SubscriptionDataFetchError = models.ProblemDetails{
		Title:         "Subscription Data Fetch error",
		Status:        http.StatusInternalServerError,
		Detail:        "The request cannot be provided due to failure in fetching subscription data.",
		Cause:         "REQUEST_REJECTED",
		InvalidParams: nil,
	}
	SubscriptionDataLenError = models.ProblemDetails{
		Title:         "Subscription Data Fetch error",
		Status:        http.StatusInternalServerError,
		Detail:        "The request cannot be provided due to not receiving any subscription data.  ",
		Cause:         "REQUEST_REJECTED",
		InvalidParams: nil,
	}
	UPFDataPathError = models.ProblemDetails{
		Title:         "UPF Data Path Failure",
		Status:        http.StatusInternalServerError,
		Detail:        "The request cannot be provided due to failure in fetching UPF data path.",
		Cause:         "REQUEST_REJECTED",
		InvalidParams: nil,
	}
	AMFDiscoveryFailure = models.ProblemDetails{
		Title:         "AMF Discovery Failure",
		Status:        http.StatusInternalServerError,
		Detail:        "The request cannot be provided due to failure in AMF discovery .",
		Cause:         "REQUEST_REJECTED",
		InvalidParams: nil,
	}
	PduSessionTypeNotSupported = models.ProblemDetails{
		Title:         "PduSession Type Not Supported",
		Status:        http.StatusForbidden,
		Detail:        "Unstructured PDU Type is not Supported.",
		Cause:         "REQUEST_REJECTED",
		InvalidParams: nil,
	}
)
