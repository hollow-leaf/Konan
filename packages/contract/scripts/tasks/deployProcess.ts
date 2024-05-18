import { task } from "hardhat/config";
import { readFileSync, writeFileSync } from "../helpers/pathHelper";
task("deploy:contract", "Deploy contract")
  .addParam("contract")
  .setAction(async ({ contract }, hre) => {
    await hre.run("compile");
    const [signer]: any = await hre.ethers.getSigners();
    const contractFactory = await hre.ethers.getContractFactory(contract);
    // if you mint in constructor, you need to add value in deploy function
    const deployContract: any = await contractFactory.connect(signer).deploy();
    console.log(`TestToken.sol deployed to ${deployContract.address}`);

    const address = {
      main: deployContract.address,
    };
    const addressData = JSON.stringify(address);
    writeFileSync(
      `scripts/address/${hre.network.name}/`,
      "mainContract.json",
      addressData,
    );

    await deployContract.deployed();
  });

task("deploy:BlackStone", "Deploy BlackStone")
  .addFlag("verify", "Validate contract after deploy")
  .setAction(async ({ verify }, hre) => {
    await hre.run("compile");
    const [signer]: any = await hre.ethers.getSigners();
    const feeData = await hre.ethers.provider.getFeeData();
    const balance = await hre.ethers.provider.getBalance(signer.address);
    console.log(`balance: ${balance}`);
    console.log(`maxPriorityFeePerGas: ${feeData.maxPriorityFeePerGas}`);
    const blackStoneFactory = await hre.ethers.getContractFactory(
      "contracts/BlackStone.sol:BlackStone",
    );
    const subscriptionId =
      "109523572682734995777696152140023272905430657577104923148923374208498494773437";
    const maxBatchSize_ = 10;
    const drawtypes_ = [9,9,5];
    const collectionSize_ = 500;
    const amountForDevs_ = 100;
    const blackStoneDeployContract: any = await blackStoneFactory
      .connect(signer)
      .deploy(
        subscriptionId,
        maxBatchSize_,
        drawtypes_,
        collectionSize_,
        amountForDevs_, 
        {
        maxPriorityFeePerGas: feeData.maxPriorityFeePerGas,
        maxFeePerGas: feeData.maxFeePerGas,
        // gasLimit: 6000000, // optional: for some weird infra network
      });
    console.log(`BlackStone.sol deployed to ${blackStoneDeployContract.address}`);
    const setProperty0 = await blackStoneDeployContract.setDrawProperty(0, [15,15,15,15,10,10,10,5,5])
    console.log(`setProperty0: ${setProperty0.hash}`);
    const setProperty1 = await blackStoneDeployContract.setDrawProperty(1, [20,20,20,10,10,10,5,5,0])
    console.log(`setProperty1: ${setProperty1.hash}`);
    const setProperty2 = await blackStoneDeployContract.setDrawProperty(2, [40,30,20,10,0])
    console.log(`setProperty2: ${setProperty2.hash}`);
    const setBaseURI = await blackStoneDeployContract.setBaseURI("https://blackstone.com/")
    console.log(`setBaseURI: ${setBaseURI.hash}`);
    const devMint = await blackStoneDeployContract.devMint(10);
    console.log(`devMint: ${devMint.hash}`);
    const token1 = await blackStoneDeployContract.tokenURI(0);
    console.log(`tokenURI: ${token1}`);
    const address = {
      main: blackStoneDeployContract.address,
    };
    const addressData = JSON.stringify(address);
    writeFileSync(
      `scripts/address/${hre.network.name}/`,
      "BlackStone.json",
      addressData,
    );

    await blackStoneDeployContract.deployed();

    if (verify) {
      console.log("verifying contract...");
      await blackStoneDeployContract.deployTransaction.wait(3);
      try {
        await hre.run("verify:verify", {
          address: blackStoneDeployContract.address,
          constructorArguments: [
            subscriptionId,
            maxBatchSize_,
            drawtypes_,
            collectionSize_,
            amountForDevs_
          ],
          contract: "contracts/BlackStone.sol:BlackStone",
        });
      } catch (e) {
        console.log(e);
      }
    }
  });
