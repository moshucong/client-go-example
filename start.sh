#!/bin/bash

export  NAMESPACE=$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace)
/app
