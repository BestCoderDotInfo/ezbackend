---
layout: post
author: derek
image: assets/images/smart-contract.png
featured: false
hidden: false
title: How to write and deploy first Smart Contract to testnet - part 1
date: 2018-08-06 14:00
comments: true
external-url:
categories: Blockchain
keywords: blockchain, smart contract, dapp
excerpt: In this article, I will show you how to write and deploy simple smart contract to testnet.
---

First you need to understand some basic concepts:

- [ https://en.wikipedia.org/wiki/Blockchain](https://en.wikipedia.org/wiki/Blockchain)

- [https://en.wikipedia.org/wiki/Ethereum](https://en.wikipedia.org/wiki/Ethereum)

- [https://en.wikipedia.org/wiki/Smart_contract](https://en.wikipedia.org/wiki/Smart_contract)


Next:

- **Ethereum Virtual Machine (EVM)**: Imagine it like a super computer to run all of the SC (Smart contract).Like its name, it is not a virtual and physical computers). Read more: [https://nulltx.com/what-is-the-ethereum-virtual-machine](https://nulltx.com/what-is-the-ethereum-virtual-machine/).

- What is **gas** ? In the main unit is gas EVM measurement to charge for every transaction with SC.Each calculation happens EVM all need gas.The contract, the more complex the need more gas. We have the recipe: **Fee for transaction  = Total gas used * gas price** . Read more: [https://ethereum.stackexchange.com/questions/3/what-is-meant-by-the-term-gas](https://ethereum.stackexchange.com/questions/3/what-is-meant-by-the-term-gas)

## 1. Write a Smart Contract

### 1.1 Setup develop environment:

- Install node js: [http://blog.teamtreehouse.com/install-node-js-npm-mac](http://blog.teamtreehouse.com/install-node-js-npm-mac).


- Install Truffle. It look like a framework allow write Ethereum Smart Contract : [https://github.com/trufflesuite/truffle](https://github.com/trufflesuite/truffle) . Run  `npm install -g truffle `.

- Install local blockchain with [Ganache](https://truffleframework.com/ganache).
Ganache define 10 ETH addresses to test . Note: URL of RPC look like:

![alt text](https://s3-ap-southeast-1.amazonaws.com/kipalog.com/k0x9rtpa4u_1_5cApmJQCnFBpYRJ_47emIg.png)

### 1.2 Begin write SC

First, create your project's folder. At here, I call it is **first_contract**. `mkdir first_contract`  then `cd first_contract` . Run shell command: `truffle init `.
![alt text](https://s3-ap-southeast-1.amazonaws.com/kipalog.com/iigk7whrwl_image5.png)

- **contracts** - this is where your smart contracts will go. You’ll notice that there is a `Migrations.sol` file already. This is needed in order to run migrations (to deploy your smart contract) later.

- **migrations** - this is where you will write your migrations. Migrations are used to deploy your Smart Contracts.

- **test** - mocha style tests for the contracts

- Smart Contract has been write by **Solidity** programing language. Read more: [http://solidity.readthedocs.io/en/v0.4.24](http://solidity.readthedocs.io/en/v0.4.24/)

At this article, I have a simple SC **contracts/Fcontracts.sol** :

```solidity
pragma solidity ^0.4.24;

contract Fcontracts {

  mapping (address => uint) fcontracts;

  function updateFcontracts(uint fcontract) public {
    fcontracts[msg.sender] = getFcontracts(msg.sender) + fcontract;
  }

  function getFcontracts(address addr) public view returns(uint) {
    return fcontracts[addr];
  }

}

```

Explain:

```
pragma solidity ^0.4.24;
```

We write SC with solidity version 0.4.24.

```
contract Fcontracts
```

Everything related to **Fcontracts** goes inside this contract. Essentially, a contract in solidity is the collection of functions and state (code and data) sitting at an address on the Ethereum blockchain.

```
 mapping (address => uint) fcontracts;
```
In Solidity, a mapping is referred to a hash table, which consists of key types and value type pairs. We define a mapping like any other variable type. Here, we're creating a mapping, which accepts first the key type (in our case, it will be an address type), and the value type will be our **unit** that we created above, then we're referring to this mapping as instructors.

```solidity
function updateFcontracts(uint fcontract) public {
    fcontracts[msg.sender] = getFcontracts(msg.sender) + fcontract;
}

function getFcontracts(address addr) public view returns(uint) {
    return fcontracts[addr];
}
```

In this SC have 2 public methods: **getter** và **setter**. This is very simple :stuck_out_tongue: .

We need test to ensure it work fine. Create **test/FconFcontracts.js**. We will use Mocha testing framework to implement for test:

```javascript
// web3 is a global variable, injected by Truffle.js
const BigNumber = web3.BigNumber

// artifacts is a global variable, injected by Truffle.js
const Fcontracts = artifacts.require("./Fcontracts.sol")

require('chai')
  .use(require('chai-as-promised'))
  .use(require('chai-bignumber')(BigNumber))
  .should()

contract('Fcontracts', function(walletAddresses) {
  let me = walletAddresses[0]
  let contract

  beforeEach(async function () {
    contract = await Fcontracts.new()
  })

  it('should create contract', async function () {
    contract.should.exist

    const fcontracts = await contract.getFcontracts(me)
    fcontracts.should.be.bignumber.equal(new BigNumber(0))
  })

  it('should updateFcontracts and getFcontracts correctly', async function () {
    // initially i have 0 shares
    let fcontracts = await contract.getFcontracts(me)
    fcontracts.should.be.bignumber.equal(new BigNumber(0))

    await contract.updateFcontracts(1, { from: me })

    fcontracts = await contract.getFcontracts(me)
    fcontracts.should.be.bignumber.equal(new BigNumber(1))
  })

})
```
At here, I have 2 methods **getFcontracts** and **updateFcontracts**. Run `truffle test` to ensure test passed.

![alt text](https://s3-ap-southeast-1.amazonaws.com/kipalog.com/hvmkuo2d5p_image4.png)

Then, we will deploy SC to testnet. At here is develop environment.

### 1.3 Deploy SC to local blockchain

Open app Ganache, then setting **truffle** to deploy local blockchain.

As you see, this project have 2 fileds: **truffle.js** and **truffle-config.js** . Cause I develop with macOS, so I will remove **truffle-config.js**. Then config looks like:

```javascript
module.exports = {
  networks: {
    development: {
      host: 'localhost',
      port: 7545,
      network_id: '*' // Match any network id
    }
  }
}
````

**7545** this is Ganache RPC port.

We need write deploy migration. Create file: **migrations/2_deploy_contracts.js** :

```javascript
var Fcontracts = artifacts.require('Fcontracts')

module.exports = function (deployer, network, accounts) {
  deployer.deploy(Fcontracts)
}
```

Run command to deploy SC:

```
truffle migrate --network development
```

Deployed console screen:

![alt text](https://s3-ap-southeast-1.amazonaws.com/kipalog.com/yr6vl2qr60_Screen%20Shot%202018-08-02%20at%2011.57.38%20AM.png)

After we has been successfuly deployed, it will have transactions:
![alt text](https://s3-ap-southeast-1.amazonaws.com/kipalog.com/tdz3fzdivm_1.png)

In the part 2, I will show you how to deploy Smart contract to real Testnet such as: Rinkeby, Ropsten, Kovan. Thanks for reading!
