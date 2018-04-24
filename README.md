# cluster-operator

# Development Deployment

## Initial (One-Time) Setup

  * Install required packages:
    * Fedora: `sudo dnf install golang make docker ansible`
    * Mac OSX:
      * [Go](https://golang.org/doc/install#osx) 
      * [Ansible](http://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#latest-releases-via-pip)
    * NOTE: Ansible 2.5+ is required to run our deployment playbooks.
  * Change docker to allow insecure pulls (required for `oc cluster up`):
    * Edit `/etc/sysconfig/docker`
    * Change `OPTIONS=` to include `--insecure-registry 172.30.0.0/16`
  * Enable and Start docker:
    * `sudo systemctl enable docker`
    * `sudo systemctl start docker`
  * Install the OpenShift and Kubernetes Python clients:
    * `sudo pip install kubernetes openshift`
  * Clone this repo to `$HOME/go/src/github.com/openshift/cluster-operator`
  * Get cfssl:
    * `go get -u github.com/cloudflare/cfssl/cmd/...`
  * Get the oc client binary
    * Fedora: Download a recent oc client binary from `origin/releases` (doesn't have to be 3.10):
      * https://github.com/openshift/origin/releases
      * Alternatively, you can also compile `oc` from source.
      * Note: It is recommended to put the `oc` binary somewhere in your path.
    * Mac OSX: [Minishift](https://github.com/minishift/minishift/releases) is the recommended development environment
  * Start an OpenShift cluster:
    * Fedora: `oc cluster up --image="docker.io/openshift/origin"`
    * Mac OSX: Follow the Minishift [Getting Started Guide](https://docs.openshift.org/latest/minishift/getting-started/index.html)
    * Note: Startup output will contain the URL to the web console for your openshift cluster, save this for later
  * Login to the OpenShift cluster as admin:
    * `oc login -u system:admin`
  * Grant admin rights to login to the web console
    * `oc adm policy add-cluster-role-to-user cluster-admin admin`
  * Ensure the following files are available on your local machine:
    * `$HOME/.aws/credentials` - your AWS credentials
    * `$HOME/.ssh/libra.pem` - the SSH private key to use for AWS


## Deploy / Re-deploy Cluster Operator

  * Compile the Go code and create the Cluster Operator images (both Go and Ansible):
    * Mac OSX only: `eval $(minishift docker-env)`
    * `make images`
  * Idempotently deploy cluster operator to the OpenShift cluster.
    * `ansible-playbook contrib/ansible/deploy-playbook.yaml`
  * If your image changed, but the kubernetes config did not (which is usually the case), you should pods appropriately:
    * `oc delete pod -l app=cluster-operator-controller-manager`
    * Or if you would rather delete all pods including the apiserver (which seldom changes) and our etcd (which would delete your stored clusters): `oc delete pod --all -n openshift-cluster-operator`

## Creating a Sample Cluster

WARNING: Running the script below creates a number of resources in your active project. By default `oc cluster up` creates a project named `myproject` and sets it as active.

  * `contrib/examples/create-cluster.sh`

When using the `create-cluster.sh` script, provisioning on AWS is disabled by default.
To enable it, you must either set the USE_REAL_AWS variable or specify a
real cluster-operator-ansible image to use in the ANSIBLE_IMAGE variable.

  * `USE_REAL_AWS=1 contrib/examples/create-cluster.sh`
  * `ANSIBLE_IMAGE=cluster-operator-ansible:canary contrib/examples/create-cluster.sh`

You can then check the provisioning status of your cluster by running `oc describe cluster <your-cluster-name>`

## Developing With OpenShift Ansible

The Cluster Operator uses its own Ansible image which layers our playbooks and roles on top of the upstream (https://github.com/openshift/openshift-ansible)[OpenShift Ansible] images. Typically our Ansible changes only require work in this repo. See the `build/cluster-operator-ansible` directory for the Dockerfile and playbooks we layer in.

To build the *cluster-operator-ansible* image you can just run `make images` normally.

**WARNING**: This image is built using OpenShift Ansible v3.10. This can be adjusted by specifying the CO_ANSIBLE_URL and CO_ANSIBLE_BRANCH environment variables to use a different branch/repository for the base openshift-ansible image.

You can run cluster-operator-ansible playbooks standalone by creating an inventory like:

```
[OSEv3:children]
masters
nodes
etcd

[OSEv3:vars]
ansible_become=true
ansible_ssh_user=centos
openshift_deployment_type=origin
openshift_release="3.10"
oreg_url=openshift/origin-${component}:v3.10.0
openshift_aws_ami=ami-833d37f9

[masters]

[etcd]

[nodes]
```

You can then run ansible with the above inventory file and your cluster ID:

`ansible-playbook -i ec2-hosts build/cluster-operator-ansible/playbooks/cluster-operator/node-config-daemonset.yml -e openshift_aws_clusterid=dgoodwin-cluster`

