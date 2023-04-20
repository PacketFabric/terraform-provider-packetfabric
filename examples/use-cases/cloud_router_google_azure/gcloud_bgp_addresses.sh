#!/usr/bin/env bash
set -x

project=$1 
region=$2
google_compute_router_name=$3

GCLOUD=$(command -v gcloud)
echo "Using gcloud from $GCLOUD"

JQ=$(command -v jq)
echo "Using jq from $JQ"

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

# Remove old files
rm -f cloud_router_ip_address.txt
rm -f customer_router_ip_address.txt

echo "running $GCLOUD compute routers describe $google_compute_router_name --project=$project --region=$region --format=json"
output=$($GCLOUD compute routers describe $google_compute_router_name --project=$project --region=$region --format=json)

echo $output

# get ipRange
cloud_router_ip_range=$(echo $output | $JQ '.interfaces[].ipRange')
echo "cloud_router_ip_range=$cloud_router_ip_range"

# get ipAddress
cloud_router_ip_address=$(echo $output | $JQ '.bgpPeers[]'.ipAddress)
echo "cloud_router_ip_address=$cloud_router_ip_address"

# get peerIpAddress
customer_router_ip_address=$(echo $output| $JQ '.bgpPeers[]'.peerIpAddress)
echo "customer_router_ip_address=$customer_router_ip_address"


# Saves BGP IP Addresses to file
subnet="${cloud_router_ip_range#*/}"
subnet="${subnet%\"}"
subnet="${subnet#\"}"

customer_router_ip_address="${customer_router_ip_address%\"}"
customer_router_ip_address="${customer_router_ip_address#\"}"
cloud_router_ip_address="${cloud_router_ip_address%\"}"
cloud_router_ip_address="${cloud_router_ip_address#\"}"

echo "${customer_router_ip_address}/${subnet}" | tr -d '\n' > customer_router_ip_address.txt
echo "${cloud_router_ip_address}/${subnet}" | tr -d '\n' > cloud_router_ip_address.txt

echo "cat cloud_router_ip_address.txt"
cat cloud_router_ip_address.txt
echo "cat customer_router_ip_address.txt"
cat customer_router_ip_address.txt