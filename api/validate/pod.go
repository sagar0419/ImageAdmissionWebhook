package validate

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful/v3/log"
	corev1 "k8s.io/api/core/v1"
)

func ValidatePodImages(pod corev1.Pod) error {

	for _, container := range pod.Spec.Containers {

		image := container.Image
		if !strings.Contains(image, "/") {
			image = "registry.docker.io/library/" + image
		} else if strings.HasPrefix(image, "docker.io/") {
			image = "registry.docker.io/" + strings.TrimPrefix(image, "docker.io/")
		}

		if !strings.HasPrefix(image, AllowedImageRegistry) {
			err := fmt.Errorf("container image %s is not from the allowed registry: %s", container.Image, AllowedImageRegistry)
			log.Printf("Validation failed: %v", err)
			return err
		}
	}
	return nil
}

// func AdmitPods(w http.ResponseWriter, r *http.Request) {
// 	var admissionReview admissionv1.AdmissionReview
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Could not read request body", http.StatusBadRequest)
// 		return
// 	}

// 	if err := json.Unmarshal(body, &admissionReview); err != nil {
// 		http.Error(w, "Could not parse admission review", http.StatusBadRequest)
// 		return
// 	}

// 	if admissionReview.Request == nil {
// 		http.Error(w, "Malformed AdmissionReview: missing request", http.StatusBadRequest)
// 		return
// 	}

// 	var pod corev1.Pod
// 	if err := json.Unmarshal(admissionReview.Request.Object.Raw, &pod); err != nil {
// 		http.Error(w, "Could not parse the object", http.StatusBadRequest)
// 		return
// 	}

// 	var allowed bool
// 	var resultMsg string
// 	if err := ValidatePodImages(pod); err != nil {
// 		allowed = false
// 		resultMsg = err.Error()
// 	} else {
// 		allowed = true
// 	}

// 	admissionReview.Response = &admissionv1.AdmissionResponse{
// 		UID:     admissionReview.Request.UID,
// 		Allowed: allowed,
// 		Result: &metav1.Status{
// 			Message: resultMsg,
// 		},
// 	}
// 	respBytes, err := json.Marshal(admissionReview)
// 	if err != nil {
// 		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(respBytes)
// }
