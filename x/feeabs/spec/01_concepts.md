
## Modified Fee Deduct Antehandler

When making a transaction, usually users need to pay fees in the native token, but "fee abstraction" allows them to pay fees in other tokens.

To allow for this, we use modified versions of `MempoolFeeDecorator` and `DeductFeeDecorate`. In these ante handlers, IBC tokens are swapped to the native token before the next fee handler logic is executed.

If a blockchain uses the Fee Abstraction module, it is necessary to replace the MempoolFeeDecorator and `DeductFeeDecorate` with the `FeeAbstrationMempoolFeeDecorator` and `FeeAbstractionDeductFeeDecorate`, respectively.


Example :

```
anteDecorators := []sdk.AnteDecorator{
  ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
  ante.NewRejectExtensionOptionsDecorator(),
  feeabsante.NewFeeAbstrationMempoolFeeDecorator(options.FeeAbskeeper),
  ante.NewValidateBasicDecorator(),
  ante.NewTxTimeoutHeightDecorator(),
  ante.NewValidateMemoDecorator(options.AccountKeeper),
  ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
  feeabsante.NewFeeAbstractionDeductFeeDecorate(options.AccountKeeper, options.BankKeeper, options.FeeAbskeeper, options.FeegrantKeeper),
  // SetPubKeyDecorator must be called before all signature verification decorators
  ante.NewSetPubKeyDecorator(options.AccountKeeper),
  ante.NewValidateSigCountDecorator(options.AccountKeeper),
  ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
  ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
  ante.NewIncrementSequenceDecorator(options.AccountKeeper),
  ibcante.NewAnteDecorator(options.IBCKeeper),
 }

```
