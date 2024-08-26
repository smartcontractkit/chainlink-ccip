# Logging

Application logs are automatically collected and sent to the centralized logging stack, where they are accessible in [Grafana Loki](https://grafana.ops.prod.cldev.sh/explore) alongside logs from other applications (including containers) running in EKS. This setup works out of the box, requiring no additional configuration. Loki is a powerful log aggregation system designed to integrate seamlessly with Grafana, enabling you to search, explore, and visualize logs in conjunction with metrics and traces.

To access the logs:

1. **Navigate to the Explore tab in Grafana**:

   - This is where you can run queries against your logs in real-time.
   - Select Loki as the data source in the Explore view.

2. **Query and Analyze Logs**:
   - Use LogQL, Loki’s query language, to filter and search your logs.
   - You can visualize the results, making it easier to identify trends, and create separate dashboards if needed.

For more details, please check the [LOKI documentation](https://smartcontract-it.atlassian.net/wiki/spaces/OBS/pages/676855845/Loki+Logs+Cheat+Sheet).

Note: For CRIB environments deployed on the local Kind cluster, logs are not currently shipped, as support for Kind is in the alpha phase.
