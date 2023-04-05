A Hyperledger Fabric network using Amazon Web Services (AWS). The network will be composed of multiple nodes, and it will be used to create a property rental website called Hartley, which will store tenancy agreements on the blockchain. I will provide step-by-step instructions for setting up the network, writing smart contracts for tenancy agreements, creating a frontend web application to interact with the network, and building a simple API layer in Go.

Here is a high-level overview of the steps I will take:

Set up an AWS instance to run Hyperledger Fabric

Install Hyperledger Fabric on the AWS instance

Create a new network using Hyperledger Fabric

Write smart contracts for tenancy agreements and deploy them to the network

Create a frontend web application to interact with the network using React

Build a simple API layer in Go to interact with the network

Write unit tests for all the code



[001/100] - Setting up an AWS instance to run Hyperledger Fabric

First, we need to create an AWS account if we don't already have one. We will use the Amazon Elastic Compute Cloud (EC2) service to create an instance to run Hyperledger Fabric.
Once we have an AWS account, we need to navigate to the EC2 Dashboard and click on the "Launch Instance" button.
In the next screen, we need to choose an Amazon Machine Image (AMI) for our instance. We will choose the "Ubuntu Server 20.04 LTS (HVM), SSD Volume Type" AMI.
Next, we need to choose an instance type. We will choose the t2.micro instance type, which is eligible for the free tier.

In the next screen, we need to configure the instance details. We will leave the default options as they are.
In the next screen, we need to add storage to our instance. We will leave the default options as they are.
In the next screen, we need to add tags to our instance. We will leave the default options as they are.
In the next screen, we need to configure the security group for our instance. We will add a new rule to allow SSH access, and another rule to allow HTTP access.
Finally, we need to review our instance details and click on the "Launch" button. We will then create a new key pair to access our instance via SSH.
Once the instance is launched, we can access it via SSH using the key pair we created. We can then proceed with the next step of the tutorial, which is installing Hyperledger Fabric on the instance.