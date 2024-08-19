# OCR3 Commit Plugin

## Context
The purpose of the OCR3 Commit Plugin is to write reports to a configured destination chain. These reports
contain metadata of cross-chain messages, from a set of source chains, that can be executed on the destination chain.

## Commit Plugin Design

The plugin is implemented as a state machine, and moves from state to state each round. There are 3 states:
1. SelectingIntervalsForReport
    - Determine intervals to be included in the next report
2. BuildingReport
    - Build a report from the intervals determined in the previous round
3. WaitingForReportTransmission
    - Check if the maximum committed sequence numbers on the dest chain have changed since generating the most
      recent report, i.e. check if the report has been committed.
    - If the maximum committed sequence numbers have changed (i.e. the report has been committed) or the maximum
      number of check attempts have been exhausted, move to the SelectingIntervalsForReport state and generate a new
      report.
    - If the maximum committed sequence numbers have _not_ changed (i.e. the report is still in-flight) and the
      maximum number of check attempts are not been exhausted, move to the WaitingForReportTransmission state in order
      to check again.

This approach leads to a clear separation of concerns and addresses the complications that can arise if a report
is not successfully transmitted (as we explicitly only continue once we know the previous report has been committed).
In this design, full messages are no longer in the observations, only merkle roots and intervals are. This reduces the
size of observations, which reduces bandwidth and improves performance.

This is the state machine diagram. States are in boxes, outcomes are within arrows.

              Start
                |
                V
    -------------------------------
    | SelectingIntervalsForReport | <---------|
    -------------------------------           |
                |                             |
        ReportIntervalsSelected               |
                |                             |
                V                             |
        ------------------                    |
        | BuildingReport | -- ReportEmpty --->|
        ------------------                    |
                |                     ReportTransmitted
         ReportGenerated                     or
                |                    ReportNotTransmitted
                V                             |
    --------------------------------          |
    | WaitingForReportTransmission | -------->|
    --------------------------------
            |           ^
            |           |
        ReportNotYetTransmitted
