#!/bin/bash
set -x

project=$1
region=$2
google_compute_router_name=$3
google_compute_router_asn=$4

GCLOUD=$(command -v gcloud)
echo "Using gcloud from $GCLOUD"

JQ=$(command -v jq)
echo "Using gcloud from $JQ"

if ! command -v gcloud --version &> /dev/null
then
    echo "gcloud --version could not be found"
    # dummy values to avoid errors in terraform
    echo "127.0.0.1/29" > cloud_router_ip_address.txt
    echo "127.0.0.2/29" >  customer_router_ip_address.txt
    exit
fi

if ! command -v jq --version &> /dev/null
then
    echo "jq --version could not be found"
    # dummy values to avoid errors in terraform
    echo "127.0.0.1/29" > cloud_router_ip_address.txt
    echo "127.0.0.2/29" >  customer_router_ip_address.txt
    exit
fi

$GCLOUD --version
$JQ --version

echo "running $GCLOUD compute routers describe $google_compute_router_name --project=$project --region=$region --format=json"
google_cloud_router_bgp_peer_name=$($GCLOUD compute routers describe $google_compute_router_name --project=$project --region=$region --format=json | $JQ '.bgpPeers[]'.name)
echo "google_cloud_router_bgp_peer_name=$google_cloud_router_bgp_peer_name"

echo "running $GCLOUD compute routers update-bgp-peer $google_compute_router_name --peer-asn=$google_compute_router_asn --peer-name=${google_cloud_router_bgp_peer_name:1:-1} --project=$project --region=$region"
$GCLOUD compute routers update-bgp-peer $google_compute_router_name --peer-asn=$google_compute_router_asn --peer-name=${google_cloud_router_bgp_peer_name:1:-1} --project=$project --region=$region

echo "running $GCLOUD compute routers describe $google_compute_router_name --project=$project --region=$region"
$GCLOUD compute routers describe $google_compute_router_name --project=$project --region=$region
