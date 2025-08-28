package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil/pg"
)

func CreateDB(ctx context.Context, lggr logger.Logger, dbURL string) (*sqlx.DB, error) {
	dbConfig := pg.DBConfig{
		IdleInTxSessionTimeout: 10 * time.Second,
		LockTimeout:            10 * time.Second,
		MaxOpenConns:           100,
		MaxIdleConns:           50,
		EnableTracing:          false,
	}

	db, err := dbConfig.New(ctx, dbURL, pg.DriverPostgres)
	if err != nil {
		lggr.Fatal("Error connecting to database", err)
	}

	lggr.Info("Connected to database")
	initDB(ctx, lggr, db)

	return db, nil
}

func initDB(ctx context.Context, lggr logger.Logger, db *sqlx.DB) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		lggr.Panicf("Error starting transaction: %v", err)
	}

	_, err = tx.Exec(`
	CREATE SCHEMA evm;

CREATE TYPE evm.tx_attempts_state AS ENUM (
    'in_progress',
    'insufficient_eth',
    'broadcast'
);

CREATE TYPE evm.txes_state AS ENUM (
    'unstarted',
    'in_progress',
    'fatal_error',
    'unconfirmed',
    'confirmed_missing_receipt',
    'confirmed',
    'finalized'
);

CREATE FUNCTION evm.f_log_poller_filter_hash(name text, evm_chain_id numeric, address bytea, event bytea, topic2 bytea, topic3 bytea, topic4 bytea) RETURNS bigint
    LANGUAGE sql IMMUTABLE COST 25 PARALLEL SAFE
    AS $_$SELECT hashtextextended(textin(record_out(($1,$2,$3,$4,$5,$6,$7))), 0)$_$;

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE evm.key_states (
    id integer NOT NULL,
    address bytea NOT NULL,
    disabled boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    evm_chain_id numeric(78,0) NOT NULL,
    CONSTRAINT chk_address_length CHECK ((octet_length(address) = 20))
);

CREATE SEQUENCE evm.eth_key_states_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.eth_key_states_id_seq OWNED BY evm.key_states.id;

CREATE TABLE evm.receipts (
    id bigint NOT NULL,
    tx_hash bytea NOT NULL,
    block_hash bytea NOT NULL,
    block_number bigint NOT NULL,
    transaction_index bigint NOT NULL,
    receipt jsonb NOT NULL,
    created_at timestamp with time zone NOT NULL,
    CONSTRAINT chk_hash_length CHECK (((octet_length(tx_hash) = 32) AND (octet_length(block_hash) = 32)))
);

CREATE SEQUENCE evm.eth_receipts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.eth_receipts_id_seq OWNED BY evm.receipts.id;

CREATE TABLE evm.tx_attempts (
    id bigint NOT NULL,
    eth_tx_id bigint NOT NULL,
    gas_price numeric(78,0),
    signed_raw_tx bytea NOT NULL,
    hash bytea NOT NULL,
    broadcast_before_block_num bigint,
    state evm.tx_attempts_state NOT NULL,
    created_at timestamp with time zone NOT NULL,
    chain_specific_gas_limit bigint NOT NULL,
    tx_type smallint DEFAULT 0 NOT NULL,
    gas_tip_cap numeric(78,0),
    gas_fee_cap numeric(78,0),
    is_purge_attempt boolean DEFAULT false NOT NULL,
    CONSTRAINT chk_cannot_broadcast_before_block_zero CHECK (((broadcast_before_block_num IS NULL) OR (broadcast_before_block_num > 0))),
    CONSTRAINT chk_chain_specific_gas_limit_not_zero CHECK ((chain_specific_gas_limit > 0)),
    CONSTRAINT chk_eth_tx_attempts_fsm CHECK ((((state = ANY (ARRAY['in_progress'::evm.tx_attempts_state, 'insufficient_eth'::evm.tx_attempts_state])) AND (broadcast_before_block_num IS NULL)) OR (state = 'broadcast'::evm.tx_attempts_state))),
    CONSTRAINT chk_hash_length CHECK ((octet_length(hash) = 32)),
    CONSTRAINT chk_legacy_or_dynamic CHECK ((((tx_type = 0) AND (gas_price IS NOT NULL) AND (gas_tip_cap IS NULL) AND (gas_fee_cap IS NULL)) OR ((tx_type = 2) AND (gas_price IS NULL) AND (gas_tip_cap IS NOT NULL) AND (gas_fee_cap IS NOT NULL)))),
    CONSTRAINT chk_sanity_fee_cap_tip_cap CHECK (((gas_tip_cap IS NULL) OR (gas_fee_cap IS NULL) OR (gas_tip_cap <= gas_fee_cap))),
    CONSTRAINT chk_signed_raw_tx_present CHECK ((octet_length(signed_raw_tx) > 0)),
    CONSTRAINT chk_tx_type_is_byte CHECK (((tx_type >= 0) AND (tx_type <= 255)))
);

CREATE SEQUENCE evm.eth_tx_attempts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.eth_tx_attempts_id_seq OWNED BY evm.tx_attempts.id;

CREATE TABLE evm.txes (
    id bigint NOT NULL,
    nonce bigint,
    from_address bytea NOT NULL,
    to_address bytea NOT NULL,
    encoded_payload bytea NOT NULL,
    value numeric(78,0) NOT NULL,
    gas_limit bigint NOT NULL,
    error text,
    broadcast_at timestamp with time zone,
    created_at timestamp with time zone NOT NULL,
    meta jsonb,
    subject uuid,
    pipeline_task_run_id uuid,
    min_confirmations integer,
    evm_chain_id numeric(78,0) NOT NULL,
    transmit_checker jsonb,
    initial_broadcast_at timestamp with time zone,
    idempotency_key character varying(2000),
    signal_callback boolean DEFAULT false,
    callback_completed boolean DEFAULT false,
    state evm.txes_state DEFAULT 'unstarted'::evm.txes_state NOT NULL,
    CONSTRAINT chk_broadcast_at_is_sane CHECK ((broadcast_at > '2019-01-01 00:00:00+00'::timestamp with time zone)),
    CONSTRAINT chk_error_cannot_be_empty CHECK (((error IS NULL) OR (length(error) > 0))),
    CONSTRAINT chk_from_address_length CHECK ((octet_length(from_address) = 20)),
    CONSTRAINT chk_to_address_length CHECK ((octet_length(to_address) = 20))
);

CREATE SEQUENCE evm.eth_txes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.eth_txes_id_seq OWNED BY evm.txes.id;

CREATE TABLE evm.forwarders (
    id bigint NOT NULL,
    address bytea NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    evm_chain_id numeric(78,0) NOT NULL,
    CONSTRAINT chk_address_length CHECK ((octet_length(address) = 20))
);

CREATE SEQUENCE evm.evm_forwarders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.evm_forwarders_id_seq OWNED BY evm.forwarders.id;

CREATE TABLE evm.log_poller_filters (
    id bigint NOT NULL,
    name text NOT NULL,
    address bytea NOT NULL,
    event bytea NOT NULL,
    evm_chain_id numeric(78,0),
    created_at timestamp with time zone NOT NULL,
    retention bigint DEFAULT 0,
    topic2 bytea,
    topic3 bytea,
    topic4 bytea,
    max_logs_kept bigint DEFAULT 0 NOT NULL,
    logs_per_block bigint DEFAULT 0 NOT NULL,
    is_legacy_name boolean DEFAULT false,
    CONSTRAINT evm_log_poller_filters_address_check CHECK ((octet_length(address) = 20)),
    CONSTRAINT evm_log_poller_filters_event_check CHECK ((octet_length(event) = 32)),
    CONSTRAINT evm_log_poller_filters_name_check CHECK ((length(name) > 0)),
    CONSTRAINT log_poller_filters_topic2_check CHECK ((octet_length(topic2) = 32)),
    CONSTRAINT log_poller_filters_topic3_check CHECK ((octet_length(topic3) = 32)),
    CONSTRAINT log_poller_filters_topic4_check CHECK ((octet_length(topic4) = 32))
);

CREATE SEQUENCE evm.evm_log_poller_filters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.evm_log_poller_filters_id_seq OWNED BY evm.log_poller_filters.id;

CREATE TABLE evm.upkeep_states (
    id integer NOT NULL,
    work_id text NOT NULL,
    evm_chain_id numeric(20,0) NOT NULL,
    upkeep_id numeric(78,0) NOT NULL,
    completion_state smallint NOT NULL,
    ineligibility_reason smallint NOT NULL,
    block_number bigint NOT NULL,
    inserted_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT work_id_len_chk CHECK (((length(work_id) > 0) AND (length(work_id) < 255)))
);

CREATE SEQUENCE evm.evm_upkeep_states_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.evm_upkeep_states_id_seq OWNED BY evm.upkeep_states.id;

CREATE TABLE evm.heads (
    id bigint NOT NULL,
    hash bytea NOT NULL,
    number bigint NOT NULL,
    parent_hash bytea NOT NULL,
    created_at timestamp with time zone NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    l1_block_number bigint,
    evm_chain_id numeric(78,0) NOT NULL,
    base_fee_per_gas numeric(78,0),
    CONSTRAINT chk_hash_size CHECK ((octet_length(hash) = 32)),
    CONSTRAINT chk_parent_hash_size CHECK ((octet_length(parent_hash) = 32))
);

CREATE SEQUENCE evm.heads_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE evm.heads_id_seq OWNED BY evm.heads.id;

CREATE TABLE evm.log_poller_blocks (
    evm_chain_id numeric(78,0) NOT NULL,
    block_hash bytea NOT NULL,
    block_number bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    block_timestamp timestamp with time zone NOT NULL,
    finalized_block_number bigint DEFAULT 0 NOT NULL,
    id bigint NOT NULL,
    safe_block_number bigint DEFAULT 0 NOT NULL,
    CONSTRAINT log_poller_blocks_block_number_check CHECK ((block_number > 0)),
    CONSTRAINT log_poller_blocks_finalized_block_number_check CHECK ((finalized_block_number >= 0)),
    CONSTRAINT log_poller_blocks_safe_block_number_check CHECK ((safe_block_number >= 0))
);

ALTER TABLE evm.log_poller_blocks ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME evm.log_poller_blocks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE evm.logs (
    evm_chain_id numeric(78,0) NOT NULL,
    log_index bigint NOT NULL,
    block_hash bytea NOT NULL,
    block_number bigint NOT NULL,
    address bytea NOT NULL,
    event_sig bytea NOT NULL,
    topics bytea[] NOT NULL,
    tx_hash bytea NOT NULL,
    data bytea NOT NULL,
    created_at timestamp with time zone NOT NULL,
    block_timestamp timestamp with time zone NOT NULL,
    id bigint NOT NULL,
    CONSTRAINT logs_block_number_check CHECK ((block_number > 0))
);

ALTER TABLE evm.logs ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME evm.logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

ALTER TABLE ONLY evm.forwarders ALTER COLUMN id SET DEFAULT nextval('evm.evm_forwarders_id_seq'::regclass);

ALTER TABLE ONLY evm.heads ALTER COLUMN id SET DEFAULT nextval('evm.heads_id_seq'::regclass);

ALTER TABLE ONLY evm.key_states ALTER COLUMN id SET DEFAULT nextval('evm.eth_key_states_id_seq'::regclass);

ALTER TABLE ONLY evm.log_poller_filters ALTER COLUMN id SET DEFAULT nextval('evm.evm_log_poller_filters_id_seq'::regclass);

ALTER TABLE ONLY evm.receipts ALTER COLUMN id SET DEFAULT nextval('evm.eth_receipts_id_seq'::regclass);

ALTER TABLE ONLY evm.tx_attempts ALTER COLUMN id SET DEFAULT nextval('evm.eth_tx_attempts_id_seq'::regclass);

ALTER TABLE ONLY evm.txes ALTER COLUMN id SET DEFAULT nextval('evm.eth_txes_id_seq'::regclass);

ALTER TABLE ONLY evm.upkeep_states ALTER COLUMN id SET DEFAULT nextval('evm.evm_upkeep_states_id_seq'::regclass);

ALTER TABLE ONLY evm.log_poller_blocks
    ADD CONSTRAINT block_hash_uniq UNIQUE (evm_chain_id, block_hash);

ALTER TABLE evm.txes
    ADD CONSTRAINT chk_eth_txes_fsm CHECK ((((state = 'unstarted'::evm.txes_state) AND (nonce IS NULL) AND (error IS NULL) AND (broadcast_at IS NULL) AND (initial_broadcast_at IS NULL)) OR ((state = 'in_progress'::evm.txes_state) AND (nonce IS NOT NULL) AND (error IS NULL) AND (broadcast_at IS NULL) AND (initial_broadcast_at IS NULL)) OR ((state = 'fatal_error'::evm.txes_state) AND (error IS NOT NULL)) OR ((state = 'unconfirmed'::evm.txes_state) AND (nonce IS NOT NULL) AND (error IS NULL) AND (broadcast_at IS NOT NULL) AND (initial_broadcast_at IS NOT NULL)) OR ((state = 'confirmed'::evm.txes_state) AND (nonce IS NOT NULL) AND (error IS NULL) AND (broadcast_at IS NOT NULL) AND (initial_broadcast_at IS NOT NULL)) OR ((state = 'confirmed_missing_receipt'::evm.txes_state) AND (nonce IS NOT NULL) AND (error IS NULL) AND (broadcast_at IS NOT NULL) AND (initial_broadcast_at IS NOT NULL)) OR ((state = 'finalized'::evm.txes_state) AND (nonce IS NOT NULL) AND (error IS NULL) AND (broadcast_at IS NOT NULL) AND (initial_broadcast_at IS NOT NULL)))) NOT VALID;

ALTER TABLE ONLY evm.key_states
    ADD CONSTRAINT eth_key_states_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.receipts
    ADD CONSTRAINT eth_receipts_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.tx_attempts
    ADD CONSTRAINT eth_tx_attempts_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.txes
    ADD CONSTRAINT eth_txes_idempotency_key_key UNIQUE (idempotency_key);

ALTER TABLE ONLY evm.txes
    ADD CONSTRAINT eth_txes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.forwarders
    ADD CONSTRAINT evm_forwarders_address_key UNIQUE (address);

ALTER TABLE ONLY evm.forwarders
    ADD CONSTRAINT evm_forwarders_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.log_poller_filters
    ADD CONSTRAINT evm_log_poller_filters_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.upkeep_states
    ADD CONSTRAINT evm_upkeep_states_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.heads
    ADD CONSTRAINT heads_pkey1 PRIMARY KEY (id);

ALTER TABLE ONLY evm.log_poller_blocks
    ADD CONSTRAINT log_poller_blocks_pkey PRIMARY KEY (id);

ALTER TABLE ONLY evm.logs
    ADD CONSTRAINT logs_pkey PRIMARY KEY (id);

CREATE INDEX evm_logs_by_timestamp ON evm.logs USING btree (evm_chain_id, address, event_sig, block_timestamp, block_number);

CREATE INDEX evm_logs_idx ON evm.logs USING btree (evm_chain_id, block_number, address, event_sig);

CREATE INDEX evm_logs_idx_data_word_five ON evm.logs USING btree (address, event_sig, evm_chain_id, "substring"(data, 129, 32));

CREATE INDEX evm_logs_idx_data_word_four ON evm.logs USING btree ("substring"(data, 97, 32));

CREATE INDEX evm_logs_idx_data_word_one ON evm.logs USING btree ("substring"(data, 1, 32));

CREATE INDEX evm_logs_idx_data_word_three ON evm.logs USING btree ("substring"(data, 65, 32));

CREATE INDEX evm_logs_idx_data_word_two ON evm.logs USING btree ("substring"(data, 33, 32));

CREATE INDEX evm_logs_idx_topic_four ON evm.logs USING btree ((topics[4]));

CREATE INDEX evm_logs_idx_topic_three ON evm.logs USING btree ((topics[3]));

CREATE INDEX evm_logs_idx_topic_two ON evm.logs USING btree ((topics[2]));

CREATE INDEX evm_logs_idx_tx_hash ON evm.logs USING btree (tx_hash);

CREATE INDEX idx_eth_receipts_block_number ON evm.receipts USING btree (block_number);

CREATE INDEX idx_eth_receipts_created_at ON evm.receipts USING brin (created_at);

CREATE UNIQUE INDEX idx_eth_receipts_unique ON evm.receipts USING btree (tx_hash, block_hash);

CREATE INDEX idx_eth_tx_attempts_broadcast_before_block_num ON evm.tx_attempts USING btree (broadcast_before_block_num);

CREATE INDEX idx_eth_tx_attempts_created_at ON evm.tx_attempts USING brin (created_at);

CREATE UNIQUE INDEX idx_eth_tx_attempts_hash ON evm.tx_attempts USING btree (hash);

CREATE INDEX idx_eth_tx_attempts_unbroadcast ON evm.tx_attempts USING btree (state) WHERE (state <> 'broadcast'::evm.tx_attempts_state);

CREATE UNIQUE INDEX idx_eth_tx_attempts_unique_gas_prices ON evm.tx_attempts USING btree (eth_tx_id, gas_price);

CREATE INDEX idx_eth_txes_broadcast_at ON evm.txes USING brin (broadcast_at);

CREATE INDEX idx_eth_txes_created_at ON evm.txes USING brin (created_at);

CREATE INDEX idx_eth_txes_from_address ON evm.txes USING btree (from_address);

CREATE INDEX idx_eth_txes_initial_broadcast_at ON evm.txes USING brin (initial_broadcast_at);

CREATE INDEX idx_eth_txes_min_unconfirmed_nonce_for_key_evm_chain_id ON evm.txes USING btree (evm_chain_id, from_address, nonce) WHERE (state = 'unconfirmed'::evm.txes_state);

CREATE UNIQUE INDEX idx_eth_txes_nonce_from_address_per_evm_chain_id ON evm.txes USING btree (evm_chain_id, from_address, nonce);

CREATE UNIQUE INDEX idx_eth_txes_pipeline_run_task_id ON evm.txes USING btree (pipeline_task_run_id) WHERE (pipeline_task_run_id IS NOT NULL);

CREATE INDEX idx_eth_txes_state_from_address_evm_chain_id ON evm.txes USING btree (evm_chain_id, from_address, state) WHERE ((state <> 'confirmed'::evm.txes_state) AND (state <> 'finalized'::evm.txes_state));

CREATE INDEX idx_eth_txes_unstarted_subject_id_evm_chain_id ON evm.txes USING btree (evm_chain_id, subject, id) WHERE ((subject IS NOT NULL) AND (state = 'unstarted'::evm.txes_state));

CREATE INDEX idx_evm_key_states_address ON evm.key_states USING btree (address);

CREATE UNIQUE INDEX idx_evm_key_states_evm_chain_id_address ON evm.key_states USING btree (evm_chain_id, address);

CREATE INDEX idx_evm_logs_ccip_exec_state_change_read ON evm.logs USING btree (address, evm_chain_id, (topics[2]), (topics[3]), block_number, log_index, tx_hash) WHERE (event_sig = '\x05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b'::bytea);

CREATE INDEX idx_evm_logs_ccip_message_sent_read_latest ON evm.logs USING btree (address, evm_chain_id, (topics[2]), block_number DESC) INCLUDE (topics, event_sig, block_number, log_index, tx_hash) WHERE (event_sig = '\x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32'::bytea);

CREATE INDEX idx_evm_logs_ccip_message_sent_read_seq ON evm.logs USING btree (address, evm_chain_id, (topics[2]), (topics[3]), block_number) INCLUDE (log_index, block_hash, event_sig, topics, tx_hash, created_at, block_timestamp) WHERE (event_sig = '\x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32'::bytea);

CREATE INDEX idx_evm_upkeep_state_added_at_chain_id ON evm.upkeep_states USING btree (evm_chain_id, inserted_at);

CREATE UNIQUE INDEX idx_evm_upkeep_state_chainid_workid ON evm.upkeep_states USING btree (evm_chain_id, work_id);

CREATE INDEX idx_forwarders_created_at ON evm.forwarders USING brin (created_at);

CREATE INDEX idx_forwarders_evm_address ON evm.forwarders USING btree (address);

CREATE INDEX idx_forwarders_evm_chain_id ON evm.forwarders USING btree (evm_chain_id);

CREATE INDEX idx_forwarders_updated_at ON evm.forwarders USING brin (updated_at);

CREATE UNIQUE INDEX idx_heads_evm_chain_id_hash ON evm.heads USING btree (evm_chain_id, hash);

CREATE INDEX idx_heads_evm_chain_id_number ON evm.heads USING btree (evm_chain_id, number);

CREATE UNIQUE INDEX idx_log_poller_blocks_chain_block ON evm.log_poller_blocks USING btree (evm_chain_id, block_number DESC);

CREATE INDEX idx_logs_chain_address_event_block_logindex ON evm.logs USING btree (evm_chain_id, address, event_sig, block_number);

CREATE UNIQUE INDEX idx_logs_chain_block_logindex ON evm.logs USING btree (evm_chain_id, block_number, log_index);

CREATE UNIQUE INDEX idx_only_one_in_progress_tx_per_account_id_per_evm_chain_id ON evm.txes USING btree (evm_chain_id, from_address) WHERE (state = 'in_progress'::evm.txes_state);

CREATE UNIQUE INDEX idx_only_one_unbroadcast_attempt_per_eth_tx ON evm.tx_attempts USING btree (eth_tx_id) WHERE (state <> 'broadcast'::evm.tx_attempts_state);

CREATE INDEX idx_receipts_tx_hash ON evm.receipts USING btree (tx_hash);

CREATE INDEX idx_receipts_tx_hash_id ON evm.receipts USING btree (tx_hash, id);

CREATE INDEX idx_tx_attempts_eth_tx_id_hash ON evm.tx_attempts USING btree (eth_tx_id, hash);

CREATE INDEX idx_tx_attempts_eth_tx_id_state_hash ON evm.tx_attempts USING btree (eth_tx_id, state, hash);

CREATE INDEX idx_txes_evm_chain_id_state_nonce ON evm.txes USING btree (evm_chain_id, state, nonce);

CREATE UNIQUE INDEX log_poller_filters_hash_key ON evm.log_poller_filters USING btree (evm.f_log_poller_filter_hash(name, evm_chain_id, address, event, topic2, topic3, topic4));

ALTER TABLE ONLY evm.receipts
    ADD CONSTRAINT eth_receipts_tx_hash_fkey FOREIGN KEY (tx_hash) REFERENCES evm.tx_attempts(hash) ON DELETE CASCADE;

ALTER TABLE ONLY evm.tx_attempts
    ADD CONSTRAINT eth_tx_attempts_eth_tx_id_fkey FOREIGN KEY (eth_tx_id) REFERENCES evm.txes(id) ON DELETE CASCADE;
	`)

	if err != nil {
		// Migrations are already applied, so this is expected.
		lggr.Info("Migrations already applied. Database initialised", "err", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		lggr.Panicf("Error committing transaction: %v", err)
	}

	lggr.Info("Database initialized")
}
