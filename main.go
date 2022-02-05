package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
	"k8s.io/client-go/kubernetes"
	"secondmate.io/kubernetesapi"
)

func main() {
	dryrun := flag.Bool("dry-run", true, "Should we actually delete the namespace")
	local := flag.Bool("local", false, "local-dev mode")
	labelMatcher := flag.String("label", "secondmate.io/purge=true", "set the label looking for during purge")
	flag.Parse()

	fmt.Printf("--- config ---\n")
	fmt.Printf("dry-run: %t \n", *dryrun)
	fmt.Printf("labelMatcher: %s \n", *labelMatcher)
	fmt.Printf("--- config ---\n\n")
	
	// Setup the Kubernetes Client Connection
	var clientset *kubernetes.Clientset
	if *local {
		clientset = kubernetesapi.LocalConnect()
	} else {
		clientset = kubernetesapi.Connect()
	}

	date := time.Now()

	namespaces := kubernetesapi.GetNameSpaces(clientset, *labelMatcher)
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)
	fmt.Fprint(w, "PURGE\tNamespace\tCreated\tPurgeDate\tNamespaceDeleted\n---\t---\t---\t---\t---\n")

	for _, namespace := range namespaces.Items {
		purgeDateLayout := "2006-01-02 15:04 MST"
		namespacePurgeDate := namespace.Labels["secondmate.io/purge-date"]
		nanespacePurgeHour := namespace.Labels["secondmate.io/purge-hour"]
		nanespacePurgeTZ := namespace.Labels["secondmate.io/purge-tz"]
		rawdate := fmt.Sprintf("%s %s:00 %s",namespacePurgeDate,nanespacePurgeHour,nanespacePurgeTZ)
		// fmt.Printf("\n\n date format %s \n\n", rawdate)
		
		purgeDate, _ := time.Parse(purgeDateLayout, rawdate)
		purge := date.After(purgeDate)
		if purge {
			namespaceDeleted := false
			if !*dryrun {
				namespaceDeleted = kubernetesapi.DeleteNamespace(clientset, namespace.Name)
			}
			fmt.Fprintf(w, "%t\t%s\t%s\t%s\t%t\n", purge, namespace.Name, namespace.CreationTimestamp, purgeDate.String(), namespaceDeleted)
		} else {
			fmt.Fprintf(w, "%t\t%s\t%s\t%s\t%t\n", purge, namespace.Name, namespace.CreationTimestamp, purgeDate.String(), false)
		}
	}
	w.Flush()
}