#!/bin/bash

version=">= 0.3.2"

if [[ $1 = "cleanup" ]]; then
    echo -e "\nDelete *state* .*lock* .terraform secret.tfvars secret.json .DS_Store cloud_router_ip_address.txt customer_router_ip_address.txt"
    find . -name ".terraform" -type d -exec rm -rf "{}" \;
    find . -name ".DS_Store" -type f -delete
    find . -name ".*lock*" -type f -delete
    find . -name "*state*" -type f -delete
    find . -name secret.tfvars -type f -delete
    find . -name secret.json -type f -delete
    find . -name cloud_router_ip_address.txt -type f -delete
    find . -name customer_router_ip_address.txt -type f -delete
fi

if [[ $1 = "local" ]]; then
  echo -e "\nSwitch to terraform.local/PacketFabric/packetfabric ~> 0.0.0\n"
  sed -i '' -e 's#PacketFabric/packetfabric#terraform.local/PacketFabric/packetfabric#g' ./use-cases/*/main.tf
  sed -i '' -e 's#PacketFabric/packetfabric#terraform.local/PacketFabric/packetfabric#g' ./use-cases/*/provider.tf
  sed -i '' -e 's#PacketFabric/packetfabric#terraform.local/PacketFabric/packetfabric#g' ./provider/provider.tf
  sed -i '' -e 's#PacketFabric/packetfabric#terraform.local/PacketFabric/packetfabric#g' ./main.tf
  sed -i '' -e "s#$version#~> 0.0.0#g" ./use-cases/*/main.tf
  sed -i '' -e "s#$version#~> 0.0.0#g" ./use-cases/*/provider.tf
  sed -i '' -e "s#$version#~> 0.0.0#g" ./provider/provider.tf
  sed -i '' -e "s#$version#~> 0.0.0#g" ./main.tf
  sed -i '' -e "s#127.0.0.1#117.109.121.202#g" ./use-cases/*/variables.tf
fi

if [[ $1 = "remote" ]]; then
  echo -e "\nSwitch to PacketFabric/packetfabric $version\n"
  sed -i '' -e 's#terraform.local/PacketFabric/packetfabric#PacketFabric/packetfabric#g' ./use-cases/*/main.tf
  sed -i '' -e 's#terraform.local/PacketFabric/packetfabric#PacketFabric/packetfabric#g' ./use-cases/*/provider.tf
  sed -i '' -e 's#terraform.local/PacketFabric/packetfabric#PacketFabric/packetfabric#g' ./provider/provider.tf
  sed -i '' -e 's#terraform.local/PacketFabric/packetfabric#PacketFabric/packetfabric#g' ./main.tf
  sed -i '' -e "s#~> 0.0.0#$version#g" ./use-cases/*/main.tf
  sed -i '' -e "s#~> 0.0.0#$version#g" ./use-cases/*/provider.tf
  sed -i '' -e "s#~> 0.0.0#$version#g" ./provider/provider.tf
  sed -i '' -e "s#~> 0.0.0#$version#g" ./main.tf
  sed -i '' -e "s#117.109.121.202#127.0.0.1#g" ./use-cases/*/variables.tf
fi

if [[ $1 = "dev" ]]; then
  echo -e "\nSwitch to api.dev.packetfabric.net\n"
  sed -i '' -e "s#api.packetfabric.com#api.dev.packetfabric.net#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#api.packetfabric.com#api.dev.packetfabric.net#g" ./variables.tf
  sed -i '' -e "s#SFO1#LAB1#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#SFO1#LAB1#g" ./variables.tf
  sed -i '' -e "s#SFO6#LAB6#g" ./variables.tf
  sed -i '' -e "s#New York#Denver Test#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#New York#Denver Test#g" ./variables.tf
  sed -i '' -e 's#default = "PacketFabric"#default = "Packet Fabric Test"#g' ./use-cases/*/variables.tf
  sed -i '' -e 's#default = "PacketFabric"#default = "Packet Fabric Test"#g' ./variables.tf
  sed -i '' -e "s#PF-AP-WDC1-1726464#PF-AP-LAB5-2756010#g" ./variables.tf
  sed -i '' -e "s#PDB-ROJ-9Y0K#ROM-57Z-XA0R#g" ./variables.tf
  sed -i '' -e "s#ATL#LON#g" ./variables.tf
  sed -i '' -e 's#default = "A"#default = "Y"#g' ./variables.tf
  PF-AP-ATL1-1744189
fi

if [[ $1 = "prod" ]]; then
  echo -e "\nSwitch to api.packetfabric.com\n"
  sed -i '' -e "s#api.dev.packetfabric.net#api.packetfabric.com#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#api.dev.packetfabric.net#api.packetfabric.com#g" ./variables.tf
  sed -i '' -e "s#LAB1#SFO1#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#LAB1#SFO1#g" ./variables.tf
  sed -i '' -e "s#LAB6#SFO6#g" ./variables.tf
  sed -i '' -e "s#Denver Test#New York#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#Denver Test#New York#g" ./variables.tf
  sed -i '' -e 's#default = "Packet Fabric Test"#default = "PacketFabric"#g' ./use-cases/*/variables.tf
  sed -i '' -e 's#default = "Packet Fabric Test"#default = "PacketFabric"#g' ./variables.tf
  sed -i '' -e "s#PF-AP-LAB5-2756010#PF-AP-WDC1-1726464#g" ./variables.tf
  sed -i '' -e "s#ROM-57Z-XA0R#PDB-ROJ-9Y0K#g" ./variables.tf
  sed -i '' -e "s#LON#ATL#g" ./variables.tf
  sed -i '' -e 's#default = "Y"#default = "A"#g' ./variables.tf
fi

echo -e "\nCheck provider settings in examples:"
echo
grep -A 1 "PacketFabric/packetfabric" ./use-cases/*/main.tf
echo
grep -A 1 "PacketFabric/packetfabric" ./use-cases/*/provider.tf
echo
grep -A 1 "PacketFabric/packetfabric" ./provider/provider.tf
echo
grep -A 1 "PacketFabric/packetfabric" ./main.tf
echo
echo -e "\nNumber of variables with api.packetfabric.com: $(grep "api.packetfabric.com" ./use-cases/*/variables.tf | wc -l)"
echo -e "Number of variables with api.dev.packetfabric.net: $(grep "api.dev.packetfabric.net" ./use-cases/*/variables.tf | wc -l)"

echo

echo -e "\nFiles to cleanup:"
find . -name ".terraform" -type d
find . -name ".DS_Store" -type f
find . -name ".*lock*" -type f
find . -name "*state*" -type f
find . -name secret.tfvars -type f
find . -name secret.json -type f
find . -name cloud_router_ip_address.txt -type f
find . -name customer_router_ip_address.txt -type f

echo -e "\nEmpty files:"
find . -empty

echo -e "\nPacketFabric Terraform Provider Remote version set to \"$version\""

echo -e "\nOptions:"
echo -e "\t./$(basename $0) [cleanup]: delete files to cleanup"
echo -e "\t./$(basename $0) [local]: switch to locally built PacketFabric provider"
echo -e "\t./$(basename $0) [remote]: switch to PacketFabric provider on the Terraform registry (using \"$version\")"
echo -e "\t./$(basename $0) [dev]: switch to PacketFabric dev endpoint and variables"
echo -e "\t./$(basename $0) [prod]: switch to PacketFabric prod endpoint and variables\n"