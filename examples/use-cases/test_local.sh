#!/bin/bash

version=">= 0.3.0"

echo
ls -l */*state* */.*lock* */.terraform */secret.tfvars */secret.json

if [[ $1 = "cleanup" ]]; then
    echo -e "\nDelete */*state* */.*lock* */.terraform */secret.tfvars */secret.json"
    rm -fr */*state* */.*lock* */.terraform */secret.tfvars */secret.json
fi

if [[ $1 = "dev" ]]; then
  echo -e "\nSwitch to terraform.local/PacketFabric/packetfabric ~> 0.0.0\n"
  sed -i '' -e 's#PacketFabric/packetfabric#terraform.local/PacketFabric/packetfabric#g' ./*/main.tf
  sed -i '' -e "s#$version#~> 0.0.0#g" ./*/main.tf
  sed -i '' -e "s#127.0.0.1#117.109.121.202#g" ./*/variables.tf
fi

if [[ $1 = "prod" ]]; then
  echo -e "\nSwitch to PacketFabric/packetfabric $version\n"
  sed -i '' -e 's#terraform.local/PacketFabric/packetfabric#PacketFabric/packetfabric#g' ./*/main.tf
  sed -i '' -e "s#~> 0.0.0#$version#g" ./*/main.tf
  sed -i '' -e "s#117.109.121.202#127.0.0.1#g" ./*/variables.tf
fi

echo
grep -A 1 "PacketFabric/packetfabric" ./*/main.tf
echo
grep "127.0.0.1" ./*/variables.tf
echo
grep "117.109.121.202" ./*/variables.tf
echo