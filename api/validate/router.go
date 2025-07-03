package validate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/emicklei/go-restful/v3/log"
	admissionv1 "k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AllowedImageRegistry = os.Getenv("IMAGE_REGISTRY")

func init() {
	if AllowedImageRegistry == "" {
		log.Printf("IMAGE_REGISTRY environment variable is not set")
		os.Exit(1)
	}
}

func AdmissionRouter(w http.ResponseWriter, r *http.Request) {

	var admissionReview admissionv1.AdmissionReview
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read the request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &admissionReview); err != nil {
		http.Error(w, "Unable to parse admission review", http.StatusBadRequest)
		return
	}

	if admissionReview.Request == nil {
		http.Error(w, "Malformed AdmissionReview: missing request", http.StatusBadRequest)
		return
	}
	var allowed bool
	var resultMsg string
	kind := admissionReview.Request.Kind.Kind
	fmt.Println("kind is ", kind)

	switch kind {

	case "Pod":
		var pod corev1.Pod
		if err := json.Unmarshal(admissionReview.Request.Object.Raw, &pod); err != nil {
			http.Error(w, "Could not parse the pod request", http.StatusBadRequest)
		}

		if err := ValidatePodImages(pod); err != nil {
			allowed = false
			resultMsg = err.Error()
		} else {
			allowed = true
		}

		admissionReview.Response = &admissionv1.AdmissionResponse{
			UID:     admissionReview.Request.UID,
			Allowed: allowed,
			Result: &metav1.Status{
				Message: resultMsg,
			},
		}

		respBytes, err := json.Marshal(admissionReview)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)

	case "Deployment":
		var deploy appsv1.Deployment
		if err := json.Unmarshal(admissionReview.Request.Object.Raw, &deploy); err != nil {
			http.Error(w, "Could not parse the object", http.StatusBadRequest)
			return
		}

		if err := validateDeployment(&deploy); err != nil {
			allowed = false
			resultMsg = err.Error()
		} else {
			allowed = true
		}
		admissionReview.Response = &admissionv1.AdmissionResponse{
			UID:     admissionReview.Request.UID,
			Allowed: allowed,
			Result: &metav1.Status{
				Message: resultMsg,
			},
		}

		respBytes, err := json.Marshal(admissionReview)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)

	default:
		http.Error(w, fmt.Sprintf("Unsupported kind: %s", kind), http.StatusBadRequest)
		return
	}
}
