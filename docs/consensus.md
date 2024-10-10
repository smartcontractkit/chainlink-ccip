# Consensus and Sub-consensus

Consensus is achieved when a required number of observations are made for a
particular data point. The CCIP plugins are built using the OffChain Reporting
(OCR) framework to manage collecting observations from a permissioned set of
nodes. The required threshold of observations is derived from the value `F`,
for strict observations we would require `2 * F + 1` observations. In other
cases `F + 1` observations are allowed. The number of nodes participating is
derived from `3 * F + 1`.

## Role DON

The CCIP plugins introduce the "Role DON" protocol built as a gossip layer
ontop of OCR, this allows for multiple distinct groups of observations. For
CCIP the roles assigned to a node indicate which chains they should be
observing. It is no longer a requirement for all nodes to support all chains.

### FRoleDON and FChain

Each role has its own `F` value. In order to distinguish between them there are
two new terms:

* **FRoleDON**: the OCR level F which determines observation thresholds for
  the OCR instance, previously `F`.
* **FChain**: the role level F which determines the observation threshold for
  each role.

## Configuration

Because OCR is a permissioned network, roles are assigned during configuration.
The plugins must ensure observations for each chain are only made by the
designated nodes.

Each role also has its own value for `FChain`. This value determines the number
of nodes we should configure as `3 * FChain + 1` along with the strict
strict observation threshold as `2 * FChain + 1`.
