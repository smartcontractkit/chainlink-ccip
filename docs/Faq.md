# Frequently Asked Questions

### Why are my Pods failing with the `Init:ImagePullBackOff` error after deploying CRIB?

- **Check Kubernetes Events**: Run the `kge` command to list Kubernetes events. This can help you understand the underlying issue.
- **Image Repository Configuration**: If you are using a custom image or an image from a different repository, ensure that you set the environment variable `DEVSPACE_IMAGE` and, if specifying a tag via the `-o ...` argument, that the tag exists.

### I have deployed a CRIB environment on a local Kind cluster, but Pods are in a pending state.

- **Check Kubernetes Events**: Run the `kge` command to list Kubernetes events. This can help you understand the underlying issue.
- **Verify Deployment Profile**: Ensure you run the `devspace deploy` command with the `--profile kind` argument. Otherwise, some Pods may be deployed as StatefulSets. If this happens, it’s best to run `devspace purge` to delete existing deployment and then redeploy using the correct profile.

### I've tried deploying CRIB but encountered issues with building the image.

When attempting to deploy CRIB, you may encounter an error like the following during the image build process:

```console
devspace deploy --profile keystone
info Using namespace 'crib-krr-test'
info Using kube context 'main-stage-cluster-crib'
build:app Execute hook 'pre-image-build-hook' at before:build:app
build:app Build <http://323150190480.dkr.ecr.us-west-2.amazonaws.com/chainlink-devspace:crib-krr-test-1723213730|323150190480.dkr.ecr.us-west-2.amazonaws.com/chainlink-devspace:crib-krr-test-1723213730> with custom command
build:app Execute hook '${SCRIPTS_DIR}/man.sh build-error' at error:build:*
build:app
build:app ###############################
build:app
build:app #         BUILD ERROR         #
build:app
build:app ###############################
build:app
build:app It seems that a build error occurred. Please ensure you are using the latest versions of CRIB,
build:app CCIP, and Chainlink repos, depending on the product you are deploying. If the issue persists,
build:app please reach out to the following Slack channels for assistance:
build:app
build:app     #team-core - for CRIB Core build issues
build:app     #team-ccip - for CRIB CCIP build issues
build:app
build:app If you are unsure whether this is a build or CRIB issue, please reach out to us on the #project-crib Slack channel.
build:app
build_images: build images: error building image <http://323150190480.dkr.ecr.us-west-2.amazonaws.com/chainlink-devspace:crib-krr-test-1723213730|323150190480.dkr.ecr.us-west-2.amazonaws.com/chainlink-devspace:crib-krr-test-1723213730>: error building image: exit status 1
fatal exit status 1
```

First, ensure that you have the latest version of the `develop` branch in the code directory referenced by `$CHAINLINK_CODE_DIR/chainlink`. If you're building a custom image to test your changes, it should be your own branch, not the `develop` branch. To verify the value of the `$CHAINLINK_CODE_DIR` variable, run the following command:

```console
(nix:nix-shell-env) MB-JJKGHY2XRX:core njegos$ devspace list vars | grep CODE
CHAINLINK_CODE_DIR | ../../../
(nix:nix-shell-env) MB-JJKGHY2XRX:core njegos$ ls -l ../../../chainlink
```

If everything is set up correctly, the issue might be due to recent changes in the develop branch. If you're building from this branch, it's possible that a recent merge has caused CRIB to break. If this is the case, please reach out to the [#project-crib](https://chainlink.enterprise.slack.com/archives/C0637K4BBC2) Slack channel for further assistance.

### I've tried deploying CRIB but Docker image build fails with a Hash Sum mismatch error.

This issue can have multiple causes, but we recommend first trying to restart the Docker daemon and checking your internet connection. You may also want to disable [Filter content for Apple devices](https://support.apple.com/en-ca/guide/deployment/dep1129ff8d2/web) if applicable.

For more details and suggestions on how to resolve this issue, refer to the related GitHub [issue](https://github.com/docker/for-mac/issues/7025).
