output: {
	apiVersion: "apps.kruise.io/v1alpha1"
	kind:       "StatefulSet"
	spec: {
		selector: matchLabels: {
			"app.oam.dev/component": context.name
		}

		if parameter["replicas"] != _|_ {
			replicas: parameter.replicas
		}

		template: {
			metadata: labels: {
				"app.oam.dev/component": context.name
			}
			spec: {
				if parameter["rollingUpdateStrategy"] != _|_ && parameter["rollingUpdateStrategy"]["inPlaceUpdateStrategy"] != _|_ {
					readinessGates: [{
						conditionType: "InPlaceUpdateReady"
					}]
				}
				containers: [{
					name:  context.name
					image: parameter.image

					if parameter["cmd"] != _|_ {
						command: parameter.cmd
					}

					if parameter["env"] != _|_ {
						env: parameter.env
					}

					if context["config"] != _|_ {
						env: context.config
					}

					ports: [{
						containerPort: parameter.port
					}]

					if parameter["cpu"] != _|_ {
						resources: {
							limits:
								cpu: parameter.cpu
							requests:
								cpu: parameter.cpu
						}
					}
				}]
		}
		}

		if parameter["orderedPodManagementPolicy"] != _|_ && parameter.orderedPodManagementPolicy {
			podManagementPolicy: "OrderedReady"
		}
		if parameter["parallelPodManagementPolicy"] != _|_ {
			podManagementPolicy: "Parallel"
		}

		updateStrategy: {
			type: "RollingUpdate"
			rollingUpdate: {
				if parameter["parallelPodManagementPolicy"] != _|_ && parameter["parallelPodManagementPolicy"]["maxUnavailable"] != _|_ {
					maxUnavailable: parameter.parallelPodManagementPolicy.maxUnavailable
				}

				if parameter["rollingUpdateStrategy"]["podUpdatePolicy"] != _|_ {
					podUpdatePolicy: parameter.rollingUpdateStrategy.podUpdatePolicy
				}

				if parameter["rollingUpdateStrategy"]["inPlaceUpdateStrategy"] != _|_ {
					inPlaceUpdateStrategy: parameter.rollingUpdateStrategy.inPlaceUpdateStrategy
				}
			}

		}
	}
}
parameter: {
	// +usage=Which image would you like to use for your service
	// +short=i
	image: string

	// +usage=Commands to run in the container
	cmd?: [...string]

	// +usage=Traffic port
	// +short=p
	port: *80 | int

	// +usage=Define arguments by using environment variables
	env?: [...{
		// +usage=Environment variable name
		name: string
		// +usage=The value of the environment variable
		value?: string
		// +usage=Specifies a source the value of this var should come from
		valueFrom?: {
			// +usage=Selects a key of a secret in the pod's namespace
			secretKeyRef: {
				// +usage=The name of the secret in the pod's namespace to select from
				name: string
				// +usage=The key of the secret to select from. Must be a valid secret key
				key: string
			}
		}
	}]

	// +usage=Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	cpu?: string

	// +usage=Pod replicas
	replicas?: int | *1

	// +usage=Pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. Conflicts with `parallelPodManagementPolicy`
	orderedPodManagementPolicy: *false | bool

	// +usage=Pods are crated in parallel to match the desired scale without waiting, and on scale down will delete all pods at once. Conflicts with `orderedPodManagementPolicy`
	parallelPodManagementPolicy?: {
		//+usage=The maximum number of pods that can be unavailable during the update. `40` means `40%` of desired pods
		maxUnavailable?: int
	}

	//+usage=indicates how pods should be updated, allows in (`ReCreate`, `InPlaceIfPossible`, `InPlaceOnly`)
	podUpdatePolicy?: *"ReCreate" | string

	//+usage=Strategy for in-place update
	inPlaceUpdateStrategy?: {
		//+usage=The timespan between set Pod status to not-ready and update images in Pod spec when in-place update a Pod
		gracePeriodSeconds?: int
	}
}
