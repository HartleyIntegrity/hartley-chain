import { Gateway, Wallets } from "fabric-network";
import FabricCAServices from "fabric-ca-client";
import path from "path";

const ccpPath = path.resolve(
  __dirname,
  "..",
  "..",
  "blockchain",
  "network",
  "connection.json"
);

const enrollAdmin = async () => {
  const caURL = "https://localhost:7054";
  const ca = new FabricCAServices(caURL);

  const walletPath = path.join(process.cwd(), "wallet");
  const wallet = await Wallets.newFileSystemWallet(walletPath);
  console.log(`Wallet path: ${walletPath}`);

  const identity = await wallet.get("admin");
  if (identity) {
    console.log(
      "An identity for the admin user already exists in the wallet"
    );
    return;
  }

  const enrollment = await ca.enroll({
    enrollmentID: "admin",
    enrollmentSecret: "adminpw",
  });
  const x509Identity = {
    credentials: {
      certificate: enrollment.certificate,
      privateKey: enrollment.key.toBytes(),
    },
    mspId: "Org1MSP",
    type: "X.509",
  };
  await wallet.put("admin", x509Identity);
  console.log(
    "Successfully enrolled admin user and imported it into the wallet"
  );
};

const registerUser = async (username) => {
  const caURL = "https://localhost:7054";
  const ca = new FabricCAServices(caURL);

  const walletPath = path.join(process.cwd(), "wallet");
  const wallet = await Wallets.newFileSystemWallet(walletPath);
  console.log(`Wallet path: ${walletPath}`);

  const identity = await wallet.get(username);
  if (identity) {
    console.log(`An identity for the user ${username} already exists in the wallet`);
    return;
  }

  const adminIdentity = await wallet.get("admin");
  if (!adminIdentity) {
    console.log("Admin identity not found in wallet");
    return;
  }

  const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
  const adminUser = await provider.getUserContext(adminIdentity, "admin");

  const secret = await ca.register({
    affiliation: "org1.department1",
    enrollmentID: username,
    role: "client",
  }, adminUser);
  const enrollment = await ca.enroll({
    enrollmentID: username,
    enrollmentSecret: secret,
  });
  const x509Identity = {
    credentials: {
      certificate: enrollment.certificate,
      privateKey: enrollment.key.toBytes(),
    },
    mspId: "Org1MSP",
    type: "X.509",
  };
  await wallet.put(username, x509Identity);
  console.log(`Successfully registered and enrolled user ${username} and imported it into the wallet`);
};

const getTenancies = async () => {
    const gateway = new Gateway();
    await gateway.connect(ccpPath, {
      wallet: await Wallets.newFileSystemWallet(path.join(process.cwd(), "wallet")),
      identity: "user1",
      discovery: { enabled: true, asLocalhost: true },
    });
  
    const network = await gateway.getNetwork("mychannel");
    const contract = network.getContract("hartley");
  
    const result = await contract.evaluateTransaction("queryAllTenancies");
    return JSON.parse(result.toString());
  };
  
  const createTenancy = async (tenancy) => {
    const gateway = new Gateway();
    await gateway.connect(ccpPath, {
      wallet: await Wallets.newFileSystemWallet(path.join(process.cwd(), "wallet")),
      identity: "user1",
      discovery: { enabled: true, asLocalhost: true },
    });
  
    const network = await gateway.getNetwork("mychannel");
    const contract = network.getContract("hartley");
  
    await contract.submitTransaction(
      "createTenancy",
      tenancy.id,
      tenancy.propertyId,
      tenancy.tenantName,
      tenancy.startDate,
      tenancy.endDate,
      tenancy.rentAmount
    );
  };

export { getTenancies, createTenancy }
  