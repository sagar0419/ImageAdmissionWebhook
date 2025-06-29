package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var allowedregistryPrefix = "registry.docker.io"

func startsWith(image, prefix string) bool {
	return len(image) >= len(prefix) && image[:len(prefix)] == prefix
}

func validateImages(pod corev1.Pod) error {

	for _, container := range pod.Spec.Containers {

		image := container.Image
		// Normalize images that omit registry (e.g., "nginx" → "docker.io/library/nginx")
		if !strings.Contains(image, ".") {
			image = "registry.docker.io/library/" + image
		} else if strings.HasPrefix(image, "docker.io/") {
			image = "registry.docker.io/" + strings.TrimPrefix(image, "docker.io/")
		}

		if !startsWith(image, allowedregistryPrefix) {
			err := fmt.Errorf("container image %s is not from the allowed registry: %s", container.Image, allowedregistryPrefix)
			fmt.Println("Validation failed:", err)
			return err
		}
	}
	return nil
}

func AdmitPods(w http.ResponseWriter, r *http.Request) {
	var admissionReview admissionv1.AdmissionReview
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &admissionReview); err != nil {
		http.Error(w, "Could not parse admission review", http.StatusBadRequest)
		return
	}

	var pod corev1.Pod
	if err := json.Unmarshal(admissionReview.Request.Object.Raw, &pod); err != nil {
		http.Error(w, "Could not parse the object", http.StatusBadRequest)
		return
	}

	var allowed bool
	var resultMsg string
	if err := validateImages(pod); err != nil {
		allowed = false
		resultMsg = err.Error()
	} else {
		allowed = true
		resultMsg = "Container image is valid"
	}

	admissionReview.Response = &admissionv1.AdmissionResponse{
		UID:     admissionReview.Request.UID,
		Allowed: allowed,
		Result: &metav1.Status{
			Message: resultMsg,
		},
	}
	respBytes, _ := json.Marshal(admissionReview)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
