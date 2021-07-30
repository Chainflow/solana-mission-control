# Metrics calculation

### Validator Monitoring dashboard:

- **Validator Identity**

    Vote Key: Vote key of the validator.

    Pub Key: Public key of the validator.

- **Validator Information**

    Node Version: Returns the current solana version running on the node,considered result feild is `solana-core` from `getVersion` method.

    Node Health: Checking whether the node is running or not. Will get the result from the method `getHealth` and consider the result accordingly. If the result field is `ok` then it marked as *UP* or else *DOWN*.

    IP Address: Gossip network address of the node, considered result from the method `getClusterNodes` of field `gossip`.
    
    Vote Account: If the validator is non-deliquent, Epoch vote account is true and active stake is non-zero then it marked as **yes** or else **no**, epoch vote account status and active stake calculated from the method `getVoteAccounts`.
    
    Validator Active stake: The stake, delegated to the vote account and active in current epoch will be calculated from `activatedStake` feild from the method `getVoteAccounts`.
        
    Commission: Validator's vote account commission, percentage (0-100) of rewards payout owed to the vote account, result feild is `commission` from the method `getVoteAccounts`.

- **Validator Health**

    Network Epoch: Current epoch of network, calculated by calling the method `getEpochInfo` and considered result is `epoch` field.

    Validator Epoch: Current validator's epoch height, calculated by calling the method `getEpochInfo`, result field is `epoch`.

    Epoch Difference: Epoch Difference is calulated by subtracting validator's epoch from network's epoch.
   
    Validator status: If the valiadtor is a non-deliquent and Epochvoteaccount is true and active stake is non-zero then validator is **voting**, otherwise considered as **jailed**. Will be getting epochvoteaccount and activestake by calling the method `getVoteAccounts`.

    Solana Current slot height: Returns the current slot the node is processing, calculated by calling method `getSlot`.

    Block Height - Network: The latest block height of network, considered result field is `blockHeight` which we can get from the method `getEpochInfo`.

    Block Height - Validator: The latest block height committed by the validator, result field is `blockHeight` got from `getEpochInfo` method.
    
    Block Height Difference: Calculated by subtracting the vaidator's block height from network's block height.

    Vote Height - Network: The latest vote height of network, we can get this by calling method `getVoteAccounts` the result field is `LastVote`.
    
    Vote Height - Validator: The latest vote height of the validator, calculated by calling method `getVoteAccounts` the result field is `LastVote`.

    Vote Height Difference: Calculated by subratcting validator's vote height from network's vote height.

    Skip Rate-Validator: Skip rate of the validator which will be calculated by calling solana client binary `solana validators`. From the set of validators will compare the configured public key and if it matches then will store the data.

    Skip Rate-Network: Skip rate of the network will be calculated by getting considering all the skip rate of the the validators in the network. Will get the data by executing the command `solana validators`.

    Skip Rate-Difference: Calculated by subtracting validator's skip rate from network's skip rate.

- **Validator Performance**

    Current Epoch - vote credits: Total vote credits for current epoch of validator's vote account, calculated from method `getVoteAccounts`, result field is `epochCredits` which has array of vote credits, result is sum of all the current epoch vote credits.

    Previous Epoch - Vote credits: Total vote credits for previous epoch of validator's vote account, calculated from method `getVoteAccounts`, considered field is `epochCredits`, which is a array of vote credits, result is sum of all previous epoch vote credits.
    
    IdentityAccount Balance: Account balance of the validator, we can get the result by calling the method `getBalance`.
        
    Confirmed Blocktime - Network: Calculated from `getConfirmedBlock` takes slot height as parameter and returns estimated production time of confirmed block of network.

    Confirmed Blocktime - Validator: Calculated from method `getConfirmedBlock` takes slot height as parameter and returns estimated production time of confirmed block of validator.

    Confirmed Epoch Last Slot - Validator: Is calucated by adding validator first slot of the epoch and number of slots in the epoch.

    Confirmed Epoch Last Slot - Network: Is calucated by adding first slot of the network epoch and number of slots in the epoch.

    Transaction Count: Total number of transactions in a ledger, calculated from method `getTransactionCount`.

    Vote Account Balance: Vote account balance of the validator, result got from method `getBalance`.

- **Recent Block Production-Current Epoch Metrics Info**
   
   Leader Slots - Validator: Leader slots of a validator in current epoch, considered
    result field is `LeaderSlots` from the method `BlockProduction`.

   Total Slots - Current Epoch: Total slots in current epoch, got result from field `TotalSlots`, from the method `BlockProduction`.

   Blocks Produced - Validator: Blocks produced of a validator in current epoch, considered result field is `BlocksProduced` from the method `BlockProduction`.

   Total Blocks Produced - Current Epoch: Total blocks produced in current epoch, considered result field is `TotalBlocksProduced` from the method `BlockProduction`.

   Skipped Slots - Validator: Skipped slots of a validator in current epoch, considered result field is `SkippedSlots` from the method `BlockProduction`.

   Total Skipped Slots - Current Epoch: Total skipped slots in current epoch, considered result field is `TotalSlotsSkipped` from the method `BlockProduction`.

- **Extra Information**

   Solana Slot Leader: Leader of the current slot, result got from the method `getSlotLeader`.

   Current Active Validators: Calculated from the method `getVoteAccounts`, which returns array of current and delinquent validators. From that considered current active validators as sum of the active validators, i.e validators who are voting.
    
   Delinquent Validators: Calculated from the method `getVoteAccounts`, which returns array of current and delinquent validators, delinquent validators are the number of vaidators i.e., who are not voting.
  
   Validator Last Voted: Most recent slot voted by the validator, considered result field is `LastVote` from the method `getVoteAccounts`.

   Solana Confirmed Slot Height: Current slot height,considered result feild is `AbsoluteSlot` from the method `getEpochInfo`.
   
   Validator Root slot: Root slot of the validator, which we can get from the method `getVoteAccounts`.
