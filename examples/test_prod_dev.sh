#!/bin/bash

if [[ $1 = "cleanup" ]]; then
    echo -e "\nDelete *state* .*lock* .terraform secret.tfvars secret.json .DS_Store cloud_router_ip_address.txt customer_router_ip_address.txt source_env_var.sh"
    find . -name ".terraform" -type d -exec rm -rf "{}" \;
    find . -name ".DS_Store" -type f -delete
    find . -name ".*lock*" -type f -delete
    find . -name "*state*" -type f -delete
    find . -name secret.tfvars -type f -delete
    find . -name secret.json -type f -delete
    find . -name source_env_var.sh -type f -delete
    find . -name cloud_router_ip_address.txt -type f -delete
    find . -name customer_router_ip_address.txt -type f -delete
fi

if [[ $1 = "dev" ]]; then
  echo -e "\nSwitch to api.dev.packetfabric.net variables\n"
  sed -i '' -e 's#default = "PacketFabric"#default = "Packet Fabric Test"#g' ./use-cases/*/variables.tf # Azure Cloud Provider
  sed -i '' -e 's#default = "PacketFabric"#default = "Packet Fabric Test"#g' ./variables.tf # Azure Cloud Provider
  sed -i '' -e "s#New York#Denver Test#g" ./use-cases/*/variables.tf # Azure Cloud location
  sed -i '' -e "s#New York#Denver Test#g" ./variables.tf # Azure Cloud location
  sed -i '' -e "s#PF-AP-WDC1-1726464#PF-AP-LAB5-2756010#g" ./variables.tf # Port - Demo B to Romain Corp
  sed -i '' -e "s#PF-BC-RNO-CHI-1729807-PF#PF-BC-GOG-LON-2796821-PF#g" ./variables.tf # Virtual - Demo B to Romain Corp
  sed -i '' -e "s#SFO1#LAB1#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#SFO1#LAB1#g" ./variables.tf 
  sed -i '' -e "s#SFO6#LAB6#g" ./variables.tf
  sed -i '' -e 's#default = "A"#default = "Y"#g' ./variables.tf
  sed -i '' -e "s#PD-WUY-9VB0#ROM-57Z-XA0R#g" ./variables.tf # Marketplace - Demo A to Romain Corp
  sed -i '' -e "s#IXW-XRH-K2VX#PI-QOS-7H3M#g" ./variables.tf # IX - IX-Denver to	PacketFabric - IX
  sed -i '' -e "s#HOU#LON#g" ./variables.tf # Marketplace - Demo A to Romain Corp
  sed -i '' -e "s#DEN#GOG#g" ./variables.tf # IX - IX-Denver to	PacketFabric - IX
fi

if [[ $1 = "prod" ]]; then
  echo -e "\nSwitch to api.packetfabric.com variables\n"
  sed -i '' -e 's#default = "Packet Fabric Test"#default = "PacketFabric"#g' ./use-cases/*/variables.tf # Azure Cloud Provider
  sed -i '' -e 's#default = "Packet Fabric Test"#default = "PacketFabric"#g' ./variables.tf # Azure Cloud Provider
  sed -i '' -e "s#Denver Test#New York#g" ./use-cases/*/variables.tf # Azure Cloud location
  sed -i '' -e "s#Denver Test#New York#g" ./variables.tf # Azure Cloud location
  sed -i '' -e "s#PF-AP-LAB5-2756010#PF-AP-WDC1-1726464#g" ./variables.tf # Port - Romain Corp to Demo B
  sed -i '' -e "s#PF-BC-GOG-LON-2796821-PF#PF-BC-RNO-CHI-1729807-PF#g" ./variables.tf # Virtual - Demo B to Romain Corp
  sed -i '' -e "s#LAB1#SFO1#g" ./use-cases/*/variables.tf
  sed -i '' -e "s#LAB1#SFO1#g" ./variables.tf
  sed -i '' -e "s#LAB6#SFO6#g" ./variables.tf
  sed -i '' -e 's#default = "Y"#default = "A"#g' ./variables.tf
  sed -i '' -e "s#ROM-57Z-XA0R#PD-WUY-9VB0#g" ./variables.tf # Marketplace - Romain Corp to Demo A
  sed -i '' -e "s#PI-QOS-7H3M#IXW-XRH-K2VX#g" ./variables.tf # IX - PacketFabric - IX to IX-Denver
  sed -i '' -e "s#LON#HOU#g" ./variables.tf # Marketplace - Demo A to Romain Corp
  sed -i '' -e "s#GOG#DEN#g" ./variables.tf # IX - IX-Denver to	PacketFabric - IX
fi

prod_dev=$(grep "Packet Fabric Test" ./use-cases/*/variables.tf | wc -l)

if [[ "$prod_dev" -eq "0" ]]; then
   echo -e "\nvariables.tf set for PacketFabric dev examples."
else
   echo -e "\nvariables.tf set for PacketFabric prod examples."
fi

echo -e "\nEmpty files:"
find . -empty

echo -e "\nFiles to cleanup:"
find . -name ".terraform" -type d
find . -name ".DS_Store" -type f
find . -name ".*lock*" -type f
find . -name "*state*" -type f
find . -name secret.tfvars -type f
find . -name secret.json -type f
find . -name source_env_var.sh -type f
find . -name cloud_router_ip_address.txt -type f
find . -name customer_router_ip_address.txt -type f

echo -e "\nOptions:"
echo -e "\t./$(basename $0) [dev]: switch from prod to dev"
echo -e "\t./$(basename $0) [prod]: switch from dev to prod"
echo -e "\t./$(basename $0) [cleanup]: delete .terraform, lock, state, secret, etc..."