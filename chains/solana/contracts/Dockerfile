ARG ANCHOR_CLI

FROM backpackapp/build:${ANCHOR_CLI}

COPY . /contracts

ENV CCIP_BUILD_GIT_HASH=dummyvaluewhichis40characterslongexactly
RUN cd /contracts && anchor build

# only keep downloaded artifacts in /root/.cargo cached
RUN rm -rf /contracts
