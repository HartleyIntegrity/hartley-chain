[001/010] Setting up a Hyperledger Fabric Network on AWS

First, you'll need to sign up for an AWS account if you don't have one already. Head over to https://aws.amazon.com/ and create an account.

Once you've logged in to your AWS account, navigate to the AWS Management Console. Search for "Amazon Managed Blockchain" and click on the service to open it.

Click on "Create a new network" in the Amazon Managed Blockchain Dashboard.

Select "Hyperledger Fabric" as the network type, and choose the latest available version.

Provide a name for your network, such as "HartleyPropertyNetwork".

For "Voting policy", choose "Simple majority" and set "Proposal duration" to 24 hours.

Click "Next" to configure the first member. Give your member a name, such as "HartleyOrg".

Create an admin for your member by providing an "Admin username" and "Admin password". Make sure to save these credentials, as you'll need them later.

Click "Next" to configure the network and review the settings.

Click "Create network and first member" to start the network creation process. It may take a while for the network to be ready.