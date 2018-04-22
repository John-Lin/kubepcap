package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"k8s.io/api/core/v1"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func generatePodMessage(pod *v1.Pod) string {
	return fmt.Sprintf("Pod: %s, Phase: %s\n", pod.GetName(), pod.Status.Phase)
}

func generateContainerMessage(pod *v1.Pod) (containersMsg []string) {
	for i, _ := range pod.Status.ContainerStatuses {
		if pod.Status.ContainerStatuses[i].State.Waiting != nil {
			containersMsg = append(containersMsg, fmt.Sprintf("- Container[%d]: %s, State: Waiting. Reason: %s\n",
				i,
				pod.Status.ContainerStatuses[i].Name,
				pod.Status.ContainerStatuses[i].State.Waiting.Reason,
			))
		} else if pod.Status.ContainerStatuses[i].State.Running != nil {
			containersMsg = append(containersMsg, fmt.Sprintf("- Container[%d]: %s, State: Running. Started At: %s\n",
				i,
				pod.Status.ContainerStatuses[i].Name,
				pod.Status.ContainerStatuses[i].State.Running.StartedAt.String(),
			))
		} else if pod.Status.ContainerStatuses[i].State.Terminated != nil {
			containersMsg = append(containersMsg, fmt.Sprintf("- Container[%d]: %s, State: Terminated. Exit Code: %d, Reason: %s\n",
				i,
				pod.Status.ContainerStatuses[i].Name,
				pod.Status.ContainerStatuses[i].State.Terminated.ExitCode,
				pod.Status.ContainerStatuses[i].State.Terminated.Reason,
			))
		} else {
			containersMsg = append(containersMsg, fmt.Sprintf("- Undefined container state"))
		}

	}
	return
}

func render(podMsg string, containersMsg []string) {
	var reRunning = regexp.MustCompile(`Pod:\s(.*),\sPhase:\sRunning`)
	var rePending = regexp.MustCompile(`Pod:\s(.*),\sPhase:\sPending`)
	var reFailed = regexp.MustCompile(`Pod:\s(.*),\sPhase:\sFailed`)
	var reSucceeded = regexp.MustCompile(`Pod:\s(.*),\sPhase:\sSucceeded`)

	if reRunning.MatchString(podMsg) {
		log.Infof(podMsg + "\t" + strings.Join(containersMsg[:], "\t"))
	} else if rePending.MatchString(podMsg) {
		log.Warnf(podMsg + "\t" + strings.Join(containersMsg[:], "\t"))
	} else if reFailed.MatchString(podMsg) {
		log.Errorf(podMsg + "\t" + strings.Join(containersMsg[:], "\t"))
	} else if reSucceeded.MatchString(podMsg) {
		log.Infof(podMsg + "\t" + strings.Join(containersMsg[:], "\t"))
	} else {
		log.Infof(podMsg)
	}

}
