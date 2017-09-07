To deploy k8s clusters in different networks:
# deploy bosh
This is just a regular way using kubo-deployment to deploy bosh:
export kubo_env=~/kubo-env

export kubo_env_name=kubo

export kubo_env_path="${kubo_env}/${kubo_env_name}"

mkdir -p "${kubo_env}"

./bin/generate_env_config "${kubo_env}" ${kubo_env_name} vsphere

Then modify the director.yml, in docs/hack/director.yml is a sample manifest for bosh director.

Note: please make sure the network bosh director is deployed to has access to internet. and also the 'dns_recursor_ip' is the dns server ip. In our dev setup on Nimbus, this dns server ip should be the same nameserver ip in your jumphost's /etc/resolv.conf .

# deploy a K8S cluster
create a file named <cluster-name>-vars.yml

docs/hack/cluster1-vars.yml is a sample manifest to deploy cluster1 to ls-2 (a logical switch created by NSX-T).
Make sure :
 - bosh vm's network has connection to this cluster's network 'ls-2'
 - T1 router is created for ls-2 
 - T1's router port has the sameip specified for 'cluster_internal_gw'
 - T1 is connected to T0
 - T0 has SNAT rule maps 'cluster_internal_cidr' to T0's IP (the port which connects to outside). So that VMs in ls-2 can access internet.

export BOSH_CLIENT=admin

export BOSH_CLIENT_SECRET=$(/usr/local/bin/bosh-cli int ~/kubo-env/kubo/creds.yml --path /admin_password)

/usr/local/bin/bosh-cli -e kubo env

bosh-cli -e kubo login

./bin/deploy_k8s ~/kubo-env/kubo <cluster-name> public

# For speed boost when deploying bosh: 

We noticed that deploying bosh and kubo on Nimbus takes a long time, mostly because of downloading the packages.
A work around is to point the urls to a corporate network location. So I downloaded a few packages to my dbc account.

bosh-deployment/bosh.yml, change the bosh url to 
http://sc-dbc1213.eng.vmware.com/famanda/pks-kubo/bosh-262.3-ubuntu-trusty-3421.9-20170706-183731-831697577-20170706183736.tgz

bosh-deployment/uaa.yml, changed the uaa url to
http://sc-dbc1213.eng.vmware.com/famanda/pks-kubo/uaa-41-ubuntu-trusty-3421.9-20170621-061436-615669844-20170621061441.tgz





