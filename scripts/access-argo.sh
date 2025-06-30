#!/bin/bash
set -eo pipefail

if ! kubectl get namespace argocd &>/dev/null; then
    echo "Error: ArgoCD namespace not found. Please install ArgoCD first."
    exit 1
fi

ARGO_SECRET=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
echo "ArgoCD credentials"
echo "------------------------"
echo "URL:      http://localhost:8080"
echo "Username: admin"
echo "Password: $ARGO_SECRET"
echo "------------------------"

kubectl port-forward svc/argocd-server -n argocd 8080:443
