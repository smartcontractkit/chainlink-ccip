# Frequently Asked Questions

### Why are my Pods failing with the `Init:ImagePullBackOff` error after deploying Crib?

- **Check Kubernetes Events**: Run the `kge` command to list Kubernetes events. This can help you understand the underlying issue.
- **Image Repository Configuration**: If you are using a custom image or an image from a different repository, ensure that you set the environment variable `DEVSPACE_IMAGE`.

### I have deployed a CRIB environment on a local Kind cluster, but Pods are in a pending state.

- **Check Kubernetes Events**: Run the `kge` command to list Kubernetes events. This can help you understand the underlying issue.
- **Verify Deployment Profile**: Ensure you run the `devspace deploy` command with the `--profile kind` argument. Otherwise, some Pods may be deployed as StatefulSets. If this happens, it’s best to run `devspace purge` to delete existing deployment and then redeploy using the correct profile.
