ARG ANCHOR_CLI

FROM backpackapp/build:${ANCHOR_CLI}

COPY . /contracts

ENV CCIP_BUILD_GIT_HASH
RUN [ ! -z "${CCIP_BUILD_GIT_HASH}" ] || { echo "CCIP_BUILD_GIT_HASH env var cannot be empty"; exit 1; }
RUN cd /contracts && anchor build

# only keep downloaded artifacts in /root/.cargo cached
RUN rm -rf /contracts
