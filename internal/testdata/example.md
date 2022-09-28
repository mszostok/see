kubectl alpha events -n foo


# Update pod 'foo' with the annotation 'description' and the value 'my frontend'
# If the same annotation is set multiple times, only the last value will be applied
kubectl annotate pods foo description='my frontend'

# Update a pod identified by type and name in "pod.json"
kubectl annotate -f pod.json description='my frontend'

# Update pod 'foo' with the annotation 'description' and the value 'my frontend running nginx', overwriting any existing value
kubectl annotate --overwrite pods foo description='my frontend running nginx'

# Update all pods in the namespace
kubectl annotate pods --all description='my frontend running nginx'

  # Update pod 'foo' only if the resource is unchanged from version 1
  kubectl annotate pods foo description='my frontend running nginx' --resource-version=1

#Update pod 'foo' by removing an annotation named 'description' if it exists
#Does not require the --overwrite flag
kubectl annotate pods foo description-## SEE ALSO

