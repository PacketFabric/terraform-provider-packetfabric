#!/bin/sh
# set -x

project=$1
region=$2
google_compute_router_name=$3
google_compute_router_asn=$4

GCLOUD_LOCATION=$(command -v gcloud)
echo "Using gcloud from $GCLOUD_LOCATION"

gcloud --version

echo "running gcloud compute routers describe $google_compute_router_name --project=$project --region=$region --format=json"
google_cloud_router_bgp_peer_name=$(gcloud compute routers describe $google_compute_router_name --project=$project --region=$region --format=json | jq '.bgpPeers[]'.name)
echo "google_cloud_router_bgp_peer_name=$google_cloud_router_bgp_peer_name"

echo "running gcloud compute routers update-bgp-peer $google_compute_router_name --peer-asn=$google_compute_router_asn --peer-name=${google_cloud_router_bgp_peer_name:1:-1} --project=$project --region=$region"
gcloud compute routers update-bgp-peer $google_compute_router_name --peer-asn=$google_compute_router_asn --peer-name=${google_cloud_router_bgp_peer_name:1:-1} --project=$project --region=$region

echo "running gcloud compute routers describe $google_compute_router_name --project=$project --region=$region"
gcloud compute routers describe $google_compute_router_name --project=$project --region=$region
