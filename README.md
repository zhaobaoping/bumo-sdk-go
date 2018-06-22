# Bumo Go SDK

## Introduction
Go developers can easily operate Bumo blockchain via the Bumo Go SDK. And you can complete the installation of by downloading the go package in a few minutes.

1. [doc](/doc) is the usage documentations for the Bumo Go SDK.
2. [src](/src)  is the source code for the Bumo Go SDK.
3. [test](/test)  is the test  code for the Bumo Go SDK.

## Environment

go 1.10.1 or above.

## How to use?

#### Import packages

```
import (

	"github.com/bumoproject/bumo-sdk-go/src/bumo"
)
```

#### Initializes the structure
>Initialize Error and BumoSdk structures

```
   var Err bumo.Error
   var bumosdk bumo.BumoSdk
```

#### New a connection

```
  	bumosdk.Newbumo(url)
```
#### Generate an inactive account
>Generate an Account by calling CreateInactive, for example：

```
    newPublicKey, newPrivateKey, newAddress, Err := bumosdk.Account.CreateInactive()
```

#### Initiate the TX
> Functions such as creating an active account, issuing assets, transferring assets, and sending BU can be completed by the following four steps.

1. Create operation
   
By invoking the corresponding method, the action for the specified function is created, and the method for issuing assets is invoked, such as an example：
   
```
    operation, Err := bumosdk.Account.Asset.Issue(sourceAddress, issueAddress, code, amount)
```


2.  Create transaction

  Build transaction and set up information such as gasPrice, feeLimit
> Note: gasPrice and feeLimit unit is MO, and 1 BU = 10 ^ 8 MO
   
   For example：
   
```
    //Need to pass in the cost parameter gasPrice、 feeLimit
    transaction, Err := bumosdk.CreateTransactionWithFee(address, nonce, gasPrice, feeLimit, issueData)
```
3.  Signature

   Transaction data is signed
   
   For example：
   
```
    signTransaction, publicKey, Err := bumosdk.SignTransaction(transaction, sourcePrivateKey)
```
4. Submit a transaction

Submit transaction and wait for block network confirmation results. Generally, the confirmation time is 10 seconds and the confirmation timeout is 500 seconds。
 
```
    submitTransaction, Err := bumosdk.SubmitTransaction(transaction, signTransaction, publicKey)
```



#### Query
Call the corresponding interface to bumo, for example: Get account information
```
    addressInfo, Err := bumosdk.Account.GetInfo(address)
```


## Example project
Bumo Go SDK provides rich examples for developers' reference

[Sample document entry](doc/bumo-sdk-go.md "")

