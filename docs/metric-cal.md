# Metrics calculation

### Validator Monitoring dashboard:

- **Validator Identity**

  Validator public Key: Node public key of the validator as base-58 encoded public key 

  Validator vote key: Vote Key of the validator as base-58 encoded pubKey

- **Validator Information**

    Solana node version: Returns the current solana versions running on the node, result feild is `solana-core` from `getVersion` method

    Solana node health: Checking wether solana is running active on this port or not, if it is active node status will be  marked as **UP** or else **DOWN** , result got from method `getHealth`.

    IP Address: Gossip network address for the node, result is got from the feild of `gossip` from the method `getClusterNodes`.
    
    Vote Account: If the valiadtor is a non deliquent, Epochvoteaccount is true and active stake is non-zero then vote account is marked as **yes** or else marked as **no**, epochvoteaccount and activestake calculated from method `getVoteAccounts`.
    
    Validator Active stake: The stake, delegated to this vote account and active in this epoch calculated by `activatedStake` feild from method `getVoteAccounts`.
        
    Commision: Validator's vote account commision, percentage (0-100) of rewards payout owed to the vote account, result feild is `commission` from the method `getVoteAccounts`.

- **validator Health**

    Network Epoch: Current network's  epoch height, calculated by calling the method `getEpochInfo`, the result presented in the feild of `epoch`.

    Validator Epoch: Current validator's epoch height,  calculated by calling the method `getEpochInfo`, result feild is `epoch`.

    Epoch Difference: Epoch Difference is calulated by subtracting validator's epoch height from network's epoch height.
   
     Validator status: If the valiadtor is a non deliquent and Epochvoteaccount is true and active stake is non-zero then validator is **voting** or else marked as **jailed**, epochvoteaccount and activestake calculated byb calling method `getVoteAccounts`.

- **Validator Performance**

    Block Height - Network: The latest block height of network, result field is `blockHeight` got from `getEpochInfo` method.
    
    Block Height - Validator: The latest block height committed by the validator, result field is `blockHeight` got from `getEpochInfo` method.
    
    Height Difference: Calculated by subtracting the vaidator's block height from network's block height

    Solana Blocktime: Difference between current and previous block time, Calculate previous and current block time by using the method `getBlockTIme`.
    
    Vote Height - Network: The latest vote height of network, calculated by calling method `getVoteAccounts` the result field is `LastVote`. 
    
    Vote Height - Validator: The latest vote height committed by validator, calculated by calling method `getVoteAccounts` the result field is `LastVote`. 

    Vote Height Difference: Calculated by subratcting validator's vote height from network's vote height

    Account Balance: Account balance of the validator, result got from method `getBalance`.
    
    Current slot height: Returns the current slot the node is processing, calculated by calling method `getSlot`.
    
    Confirmed Blocktime - Network: Calculated from `getConfirmedBlock` takes slot height as parameter and returns estimated production time of confirmed block of Network.

    Confirmed Blocktime - Validator: Calculated from  method `getConfirmedBlock` takes slot height as parameter and returns estimated production time of confirmed block of Validator. 
    
    Block Time Difference: Calculated by subtracting confirmed blocktime of network from confirmed blocktime of validator.

    Current Epoch - vote credits: Total current epoch vote credits of validator's vote account,  calculated from method `getVoteAccounts`, result field `epochCredits` has array of vote credits, result is sum of all the current epoch vote credits.

    Previous Epoch - Vote credits: Total previous epoch vote credits of validator's vote account,  calculated from method `getVoteAccounts`, result field `epochCredits` has array of vote credits, result is sum of all previous epoch vote credits.

    Total valid slots: number of leader valid slots per leader, If the block present in Confirmed Blocks it counted as valid slot. Confirmed Blocks calculated from method `getConfirmedBlocks`.
    
    Total skipped slots: number of leader skipped slots per leader, If the block is not presented in Confirmed Blocks it counted as skipped slot, Confirmed Blocks calculated from method `getConfirmedBlocks`.

- **Validator Details**

   solana slot leader: Current slot leader address as base-58 enocoded string, result got from method `getSlotLeader`.
   
   Transaction Count: Total number of transactions presented in ledger, Calculated from method `getTransactionCount`.

   Validator last voted: Most recent slot voted on by this vote account, got result field `LastVote` from method `getVoteAccounts`.
   
   Solana confirmed slot height: Current slot height, result feild is `AbsoluteSlot` from method `getEpochInfo`
    
   Current Active Validators: calculates from the method `getVoteAccounts` returns array of current and delinquent validators, current active validators are the sum of current active validators, i.e validators who are voting.
    
   Delinquent validators: calculates from the method `getVoteAccounts`,  returns array of current and delinquent validators, delinquent validators are the sum of delinquent validators presents, i.e validators who are jailed.

   Confirmed epoch last slot: Is calucated by adding first slot of the epoch and no.of slots in the epoch.

   Validator Root slot: Root slot per validator