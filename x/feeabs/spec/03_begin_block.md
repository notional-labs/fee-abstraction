# Begin Block

At the begin of every block, ``feeabs`` module iterates epochs and runs the following with every epoch:
1. If currentTime < epoch Start time , we stop

2. If epoch counting hasn't started, we check if epoch should start.

```go
    if !(currentBlocktime>currenEpochStartTime.Add(epochInfo.Duration) || !EpochCountingStarted)
```

3. Then check if we should initial epoch start

```go
    epochInfo.EpochCountingStarted = true
	epochInfo.CurrentEpoch = 1
	epochInfo.CurrentEpochStartTime = epochInfo.StartTime
```

4. If not initial epoch start, we execute TWAPQuery and Swap

```go
    executeAllHostChainTWAPQuery()
    executeAllHostChainSwap()
```

5. Emit event `EpochStart`
6. Set epoch info

```go
    setEpochInfo(ctx, epochInfo)
```
