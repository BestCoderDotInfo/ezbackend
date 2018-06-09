---
title: How can I learn blockchain development?
date: '2018-06-09 15:00:00'
comments: true
external-url: null
categories: Blockchain
keywords: Blockchain, Training, Courses, programing, developer
excerpt: Blockchain is a very hot topic right now, and Blockchain developers are in high demand. So it makes sense for a developer to want to dive in and add Blockchain development capabilities in his/her repertoire.
---
Blockchain is a very hot topic right now, and Blockchain developers are in high demand. So it makes sense for a developer to want to dive in and add Blockchain development capabilities in his/her repertoire.

What’s important to understand though, is that Blockchain has come a long way as a technology from the first paper published by Satoshi. Sure, it would be a good thing to start from that paper[1], but it might appear complicated and can be confusing.

Another important point is, Blockchain development is not concentrated to, and around, Bitcoins. You might be a Blockchain developer and still have never worked with Bitcoins. It is entirely possible.

I would try to explain in short, various concepts around which Blockchain development revolves. Then I would also provide a gist of technologies which can be used to become a Blockchain developer.

## Concepts

- **Decentralization** : No one entity should own all the assets or resources in the network. It is important to understand how this is different compared to a distributed network. A distributed network can still be owned by one entity, for instance Netflix. A distributed network is essential for a decentralized network, but it is not necessarily decentralized in it’s own right.
- **Immutability** : The data once written, should always be a part of the network. Any change performed over the data must also be recorded in the network. This can be achieved by creating hash of each transaction/operation performed in the network and add it to the next transaction as a meta data. This chains every new transaction to it’s previous transaction and that is where Blockchain derives it’s name from.
- **Assets/Entities** : The storage network should support storage of data or some form of transferable asset in some form or the other. This can be Bitcoins for Bitcoin network, or Ether[2] or Smart contracts or other forms of data in Ethereum [3]network or Assets in case of BigchainDB[4]. (All these are different kinds of prevalent Blockchain networks or frameworks and I will talk about them in some detail later)
- **Censor resistance** : Although this point is inherent to Decentralization, some Blockchain enthusiasts make a point of mentioning it separately. Basically any single authority should not be able to Censor the flow of data or exchange of information over a Blockchain. But with the advent of private Blockchain technologies, Censor resistance is not a compulsory component of a Blockchain.

To sum this section up, Blockchain is any decentralized network which supports aforementioned characteristics.

Now that we have a basic broad explanation of Blockchain, we can move on to the topic of development.

When Bitcoin was launched, the underlying technology wasn’t much explored. But as the interest grew, experts started looking at the architecture which supported bitcoins. An analysis quickly revealed that this architecture had far reaching potential. To put it in simple words, Bitcoin is to Blockchain, what Email was to Internet. Sure it is one very important application of it, it is not the only possible application. And Blockchain is touted to be Internet 2.0 .

To explore such potential, a lot of new development technologies have come up. Let’s talk about some of them.

- **Ethereum** : Ethereum is not only a Blockchain where one can trade in Ethers. It is also a development framework where one can realize a lot of wonderful ideas and implement them using a Blockchain. The most important aspect of Ethereum Blockchain Development is a concept called Smart Contracts. Smart contracts are basically simple methods or functions which run on the Ethereum Blockchain. They can be imagined as something similar to a Java function for instance. But the difference is that every time the Smart Contracts are executed, the execution takes place over the Blockchain and it is written forever in form of a transaction in it. The transaction can not be deleted.

A lot of Blockchain developers are developing Smart contracts. Smart contracts are written in Solidity[5]. According to the introduction, Solidity is a contract-oriented, high-level language for implementing smart contracts. It was influenced by C++, Python and JavaScript and is designed to target the Ethereum Virtual Machine (EVM). You can learn the development of Solidity by following the tutorial and examples on the mentioned link.
- **BigchainDB** : You might have guessed it by now, that Blockchain suffers from scalability problem. Storing large amount of data on Blockchain is not a great idea. This is the problem that BigchainDb tries to solve. It is another Blockchain related technology which when used stand alone, can provide a very fast (1 million transactions per second) throughput. It can be used as a Blockchain in it’s own right. But it can also be used in conjunction with Ethereum to act as a Blockchain secure database where Ethereum acts as the logic processing part of the stack. The development of BigchainDB can be done using nodeJS or a variety of other programming languages. The entire list can be found in the Driver and Tools section here[6]. A good starting point to start developing using BigchainDB can be found here[7].
- **Hyperledger** : Hyperledger is a Business Blockchain tool developed by Linux foundation and IBM. It provides the capability to create a private permissioned Blockchain. A good point to start learning Blockchain development using Hyperledger would be this[8]

There are a few other technologies too, like IOTA, Lisk, Interledger but they can be handled later on a need to know basis.

My advice would be to stay away from generic explanations about Blockchain and start with one of these technologies and dive into development. Not that those explanations aren’t good, but they don’t serve a lot of purpose when it comes to Blockchain development. The only way to really learn Blockchain development is getting your hands dirty.

##Footnotes
[1] [https://bitcoin.org/bitcoin.pdf](https://bitcoin.org/bitcoin.pdf)
[2] [Smart contract - Wikipedia](https://en.wikipedia.org/wiki/Smart_contract)
[3] [https://www.google.de/url?sa=t&r...](https://www.google.de/url?cad=rja&cd=4&esrc=s&q=&rct=j&sa=t&source=web&uact=8&url=https%3A%2F%2Fwww.ethereum.org%2F&usg=AOvVaw0T5uMT9IXrj-8co72FaCPt&ved=0ahUKEwjEociX_YbYAhWECOwKHW0vBVUQFghdMAM)
[4] [https://www.google.de/url?sa=t&r...](https://www.google.de/url?cad=rja&cd=1&esrc=s&q=&rct=j&sa=t&source=web&uact=8&url=https%3A%2F%2Fwww.bigchaindb.com%2F&usg=AOvVaw3KwkrXrLsrfMnmNZOaFQ-P&ved=0ahUKEwj94J6A_YbYAhXM_KQKHaSgC64QFggnMAA)
[5] [Solidity - Solidity 0.4.20 documentation](https://solidity.readthedocs.io/en/develop/)
[6] [Get started • • BigchainDB](https://www.bigchaindb.com/getstarted/)
[7] [Basic Usage Examples](https://docs.bigchaindb.com/projects/js-driver/en/latest/usage.html)
[8] [Blockchain for Business - An Introduction to Hyperledger Technologies](https://www.edx.org/course/blockchain-business-introduction-linuxfoundationx-lfs171x)