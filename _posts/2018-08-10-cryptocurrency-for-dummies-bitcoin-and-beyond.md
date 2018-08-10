---
layout: post
author: Toptal
image: assets/2018-08-10/cover.png
featured: false
hidden: false
title: "Cryptocurrency for Dummies: Bitcoin and Beyond"
date: 2018-08-10 09:00
comments: true
external-url:
categories: Cryptocurrency
keywords: Bitcoin, Cryptocurrency, Cryptography
excerpt: Bitcoin created a lot of buzz on the Internet. It was ridiculed, it was attacked, and eventually it was accepted and became a part of our lives. However, Bitcoin is not alone. At this moment, there are over 700 AltCoin implementations, which use similar principles and various cryptocurrency algorithms.
---

Bitcoin created a lot of buzz on the Internet. It was ridiculed, it was attacked, and eventually it was accepted and became a part of our lives. However, [Bitcoin](https://www.toptal.com/bitcoin) is not alone. At this moment, there are [over 700 AltCoin](http://mapofcoins.com/) implementations, which use similar principles and various cryptocurrency algorithms.

So, what do you need to create something like Bitcoin?

Without trying to understand your personal motivation for creating a decentralized, anonymous system for exchanging money/information (but still hoping that it is in scope of moral and legal activities), let’s first break down the basic requirements for our new payment system:

- All transactions should be made over the Internet
- We do not want to have a central authority that will process transactions
- Users should be anonymous and identified only by their virtual identity
- A single user can have as many virtual identities as he or she likes
- Value supply (new virtual bills) must be added in a controlled way

## Decentralized Information Sharing Over Internet

Fulfilling the first two requirements from our list, removing a central authority for information exchange over the Internet, is already possible. What you need is a peer-to-peer (P2P) network.

Information sharing in P2P networks is similar to information sharing among friends and family. If you share information with at least one member of the network, eventually this information will reach every other member of the network. The only difference is that in digital networks this information will not be altered in any way.

![](/assets/2018-08-10/image1.png){:height="100%" width="100%"}

You have probably heard of BitTorrent, one of the most popular P2P file sharing (content delivery) systems. Another popular application for P2P sharing is *Skype*, as well as other chat systems.

Bottom line is that you can implement or use one of the existing open-source P2P protocols to support your new cryptocurrency, which we’ll call *Topcoin*.

## Hashing Algorithm

To understand digital identities, we need to understand how **cryptographic hashing works**. Hashing is the process of mapping digital data of any arbitrary size to data of a fixed size. In simpler words, hashing is a process of taking some information that is readable and making something that makes no sense at all.

>You can compare hashing to getting answers from politicians. Information you provide to them is clear and understandable, while the output they provide looks like random stream of words.

![](/assets/2018-08-10/image2.png){:height="100%" width="100%"}

There are a few requirements that a good hashing algorithm needs:

- Output length of hashing algorithm must be fixed (a good value is 256 bytes)
- Even the smallest change in input data must produce significant difference in output
- Same input will always produce same output
- There must be no way to reverse the output value to calculate the input
- Calculating the HASH value should not be compute intensive and should be fast

If you take a look at the simple statistics, we will have a limited (but huge) number of possible HASH values, simply because our HASH length is limited. However, our hashing algorithm (let’s name it Politician256) should be reliable enough that it only produces duplicate hash values for different inputs about as frequently as a monkey in a zoo manages to correctly type Hamlet on a typewriter!

*If you think Hamlet is just a name or a word, please stop reading now, or read about the [Infinite Monkey Theorem](https://en.wikipedia.org/wiki/Infinite_monkey_theorem).*

## Digital Signature

When signing a paper, all you need to do is append your signature to the text of a document. A digital signature is similar: you just need to append your personal data to the document you are signing.

If you understand that the hashing algorithm adheres to the rule where **even the smallest change in input data must produce significant difference in output**, then it is obvious that the HASH value created for the original document will be different from the HASH value created for the document with the appended signature.

A combination of the original document and the HASH value produced for the document with your personal data appended is a **digitally signed document**.

And this is how we get to your **virtual identity**, which is defined as the data you appended to the document before you created that HASH value.

Next, you need to make sure that your signature cannot be copied, and no one can execute any transaction on your behalf. The best way to make sure that your signature is secured, is to keep it yourself, and provide a different method for someone else to validate the signed document. Again, we can fall back on technology and algorithms that are readily available. What we need to use is [public-key cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography) also known as **asymmetric cryptography**.

To make this work, you need to create a **private key** and a **public key**. These two keys will be in some kind of mathematical correlation and will depend on each other. The algorithm that you will use to make these keys will assure that each private key will have a different public key. As their names suggest, a private key is information that you will keep just for yourself, while a public key is information that you will share.

If you use your private key (your identity) and original document as input values for the **signing algorithm** to create a HASH value, assuming you kept your key secret, you can be sure that no one else can produce the same HASH value for that document.

![](/assets/2018-08-10/image3.png){:height="100%" width="100%"}

If anyone needs to validate your signature, he or she will use the original document, the HASH value you produced, and your public key as inputs for the **signature verifying algorithm** to verify that these values match.

![](/assets/2018-08-10/image4.png){:height="100%" width="100%"}

## How to send Bitcoin/Money

Assuming that you have implemented P2P communication, mechanisms for creating digital identities (private and public keys), and provided ways for users to sign documents using their private keys, you are ready to start sending information to your peers.

Since we do not have a central authority that will validate how much money you have, the system will have to ask you about it every time, and then check if you lied or not. So, your transaction record might contain the following information:

- I have 100 Topcoins
- I want to send 10 coins to my pharmacist for the medication (you would include your pharmacists public key here)
- I want to give one coin as transaction fee to the system (we will come back to this later)
- I want to keep the remaining 89 coins

The only thing left to do is digitally sign the transaction record with your private key and transmit the transaction record to your peers in the network. At that point, everyone will receive the information that someone (your virtual identity) is sending money to someone else (your pharmacist’s virtual identity).

Your job is done. However, your medication will not be paid for until the whole network agrees that you really did have 100 coins, and therefore could execute this transaction. Only after your transaction is validated will your pharmacist get the funds and send you the medication.

## Cryptocurrency Miners: A New Breed of Agent

Miners are known to be very hard working people who are, in my opinion, heavily underpaid. In the digital world of cryptocurrency, miners play a very similar role, except in this case, they do the computationally-intensive work instead of digging piles of dirt. Unlike real miners, some cryptocurrency miners earned a small fortune over the past five years, but many others lost a fortune on this risky endeavour.

Miners are the core component of the system and their main purpose is to confirm the validity of each and every transaction requested by users.

In order to confirm the validity of your transaction (or a combination of several transactions requested by a few other users), miners will do two things.

First, they will rely on the fact that “everyone knows everything,” meaning that every transaction executed in the system is copied and available to any peer in the network. They will look into the history of your transactions to verify that you actually had 100 coins to begin with. Once your account balance is confirmed, they will generate a specific HASH value. **This hash value must have a specific format; it must start with certain number of zeros**.

There are two inputs for calculating this HASH value:

- Transaction record data
- Miner’s proof-of-work

Considering that even the smallest change in input data must produce a significant difference in output HASH value, miners have a very difficult task. They need to find a specific value for a proof-of-work variable that will produce a HASH beginning with zeros. If your system requires a minimum of 40 zeros in each validated transaction, the miner will need to calculate approximately 2^40 different HASH values in order to find the right proof-of-work.

Once a miner finds the proper value for proof-of-work, he or she is entitled to a transaction fee (the single coin you were willing to pay), which can be added as part of the validated transaction. Every validated transaction is transmitted to peers in the network and stored in a specific database format known as the Blockchain.

But what happens if the number of miners goes up, and their hardware becomes much more efficient? Bitcoin used to be mined on CPUs, then GPUs and FPGAs, but ultimately miners started designing their own ASIC chips, which were vastly more powerful than these early solutions. As the hash rate goes up, so does the mining difficulty, thus ensuring equilibrium. When more hashing power is introduced into the network, the difficulty goes up and vice versa; if many miners decide to pull the plug because their operation is no longer profitable, difficulty is readjusted to match the new hash rate.

## Blockchain for Dummies: The Global Cryptocurrency Ledger

The blockchain contains the history of all transactions performed in the system. Every validated transaction, or batch of transactions, becomes another ring in the chain.

So, the Bitcoin blockchain is, essentially, a public ledger where transactions are listed in a chronological order.

The first ring in the Bitcoin blockchain is called the **Genesis Block**
To read more about how the blockchain works, I suggest reading [Blockchain Technology Explained: Powering Bitcoin](https://www.toptal.com/bitcoin/blockchain-technology-powering-bitcoin), by Nermin Hajdarbegovic.

There is no limit to how many miners may be active in your system. This means that it is possible for two or more miners to validate the same transaction. If this happens, the system will check the total effort each miner invested in validating the transaction by simply counting zeros. The miner that invested more effort (found more leading zeros) will prevail and his or her block will be accepted.

>This article has been written by **DEMIR SELMANOVIC** - *LEAD TECHNICAL EDITOR @ TOPTAL*. You can read original post at [here](https://www.toptal.com/bitcoin/cryptocurrency-for-dummies-bitcoin-and-beyond).