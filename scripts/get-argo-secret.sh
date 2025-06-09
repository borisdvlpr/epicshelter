#!/bin/bash
set -eo pipefail

ARGO_SECRET=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)
echo "ArgoCD credentials"
echo "------------------------"
echo "URL:      http://localhost:8080"
echo "Username: admin"
echo "Password: $ARGO_SECRET"
echo "------------------------"
