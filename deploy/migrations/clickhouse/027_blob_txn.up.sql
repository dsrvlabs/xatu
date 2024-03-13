ALTER TABLE mempool_transaction_local on cluster '{cluster}'
ADD COLUMN gas_tip_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas,
ADD COLUMN gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas_tip_cap,
ADD COLUMN blob_gas Nullable(UInt64) Codec(ZSTD(1)) AFTER call_data_size,
ADD COLUMN blob_gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER blob_gas,
ADD COLUMN blob_hashes Array(String) Codec(ZSTD(1)) AFTER blob_gas_fee_cap,
ADD COLUMN blob_sidecars_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_hashes,
ADD COLUMN blob_sidecars_empty_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_sidecars_size,
COMMENT COLUMN gas_tip_cap 'The priority fee (tip) the user has set for the transaction',
COMMENT COLUMN gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_gas 'The maximum gas provided for the blob transaction execution',
COMMENT COLUMN blob_gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_hashes 'The hashes of the blob commitments for blob transactions',
COMMENT COLUMN blob_sidecars_size 'The total size of the sidecars for blob transactions in bytes',
COMMENT COLUMN blob_sidecars_empty_size 'The total empty size of the sidecars for blob transactions in bytes';

ALTER TABLE mempool_transaction on cluster '{cluster}'
ADD COLUMN gas_tip_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas,
ADD COLUMN gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas_tip_cap,
ADD COLUMN blob_gas Nullable(UInt64) Codec(ZSTD(1)) AFTER call_data_size,
ADD COLUMN blob_gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER blob_gas,
ADD COLUMN blob_hashes Array(String) Codec(ZSTD(1)) AFTER blob_gas_fee_cap,
ADD COLUMN blob_sidecars_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_hashes,
ADD COLUMN blob_sidecars_empty_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_sidecars_size,
COMMENT COLUMN gas_tip_cap 'The priority fee (tip) the user has set for the transaction',
COMMENT COLUMN gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_gas 'The maximum gas provided for the blob transaction execution',
COMMENT COLUMN blob_gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_hashes 'The hashes of the blob commitments for blob transactions',
COMMENT COLUMN blob_sidecars_size 'The total size of the sidecars for blob transactions in bytes',
COMMENT COLUMN blob_sidecars_empty_size 'The total empty size of the sidecars for blob transactions in bytes';

ALTER TABLE canonical_beacon_block_execution_transaction_local on cluster '{cluster}'
ADD COLUMN gas_tip_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas,
ADD COLUMN gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas_tip_cap,
ADD COLUMN blob_gas Nullable(UInt64) Codec(ZSTD(1)) AFTER call_data_size,
ADD COLUMN blob_gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER blob_gas,
ADD COLUMN blob_hashes Array(String) Codec(ZSTD(1)) AFTER blob_gas_fee_cap,
ADD COLUMN blob_sidecars_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_hashes,
ADD COLUMN blob_sidecars_empty_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_sidecars_size,
COMMENT COLUMN gas_tip_cap 'The priority fee (tip) the user has set for the transaction',
COMMENT COLUMN gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_gas 'The maximum gas provided for the blob transaction execution',
COMMENT COLUMN blob_gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_hashes 'The hashes of the blob commitments for blob transactions',
COMMENT COLUMN blob_sidecars_size 'The total size of the sidecars for blob transactions in bytes',
COMMENT COLUMN blob_sidecars_empty_size 'The total empty size of the sidecars for blob transactions in bytes';

ALTER TABLE canonical_beacon_block_execution_transaction on cluster '{cluster}'
ADD COLUMN gas_tip_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas,
ADD COLUMN gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER gas_tip_cap,
ADD COLUMN blob_gas Nullable(UInt64) Codec(ZSTD(1)) AFTER call_data_size,
ADD COLUMN blob_gas_fee_cap Nullable(UInt128) Codec(ZSTD(1)) AFTER blob_gas,
ADD COLUMN blob_hashes Array(String) Codec(ZSTD(1)) AFTER blob_gas_fee_cap,
ADD COLUMN blob_sidecars_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_hashes,
ADD COLUMN blob_sidecars_empty_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_sidecars_size,
COMMENT COLUMN gas_tip_cap 'The priority fee (tip) the user has set for the transaction',
COMMENT COLUMN gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_gas 'The maximum gas provided for the blob transaction execution',
COMMENT COLUMN blob_gas_fee_cap 'The max fee the user has set for the transaction',
COMMENT COLUMN blob_hashes 'The hashes of the blob commitments for blob transactions',
COMMENT COLUMN blob_sidecars_size 'The total size of the sidecars for blob transactions in bytes',
COMMENT COLUMN blob_sidecars_empty_size 'The total empty size of the sidecars for blob transactions in bytes';

ALTER TABLE canonical_beacon_blob_sidecar_local on cluster '{cluster}'
ADD COLUMN blob_empty_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_size,
COMMENT COLUMN blob_empty_size 'The total empty size of the blob in bytes';

ALTER TABLE canonical_beacon_blob_sidecar on cluster '{cluster}'
ADD COLUMN blob_empty_size Nullable(UInt32) Codec(ZSTD(1)) AFTER blob_size,
COMMENT COLUMN blob_empty_size 'The total empty size of the blob in bytes';

ALTER TABLE block_native_mempool_transaction_local on cluster '{cluster}'
COMMENT COLUMN unique_key 'Unique key for the row',
COMMENT COLUMN updated_date_time 'When this row was last updated',
COMMENT COLUMN detecttime 'Timestamp that the transaction was detected in mempool',
COMMENT COLUMN hash 'Unique identifier hash for a given transaction',
COMMENT COLUMN status 'Status of the transaction',
COMMENT COLUMN region 'The geographic region for the node that detected the transaction',
COMMENT COLUMN reorg 'If there was a reorg, refers to the blockhash of the reorg',
COMMENT COLUMN replace 'If the transaction was replaced (speedup/cancel), the transaction hash of the replacement',
COMMENT COLUMN curblocknumber 'The block number the event was detected in',
COMMENT COLUMN failurereason 'If a transaction failed, this field provides contextual information',
COMMENT COLUMN blockspending 'If a transaction was finalized (confirmed, failed), this refers to the number of blocks that the transaction was waiting to get on-chain',
COMMENT COLUMN timepending 'If a transaction was finalized (confirmed, failed), this refers to the time in milliseconds that the transaction was waiting to get on-chain',
COMMENT COLUMN nonce 'A unique number which counts the number of transactions sent from a given address',
COMMENT COLUMN gas 'The maximum number of gas units allowed for the transaction',
COMMENT COLUMN gasprice 'The price offered to the miner/validator per unit of gas. Denominated in wei',
COMMENT COLUMN value 'The amount of ETH transferred or sent to contract. Denominated in wei',
COMMENT COLUMN toaddress 'The destination of a given transaction',
COMMENT COLUMN fromaddress 'The source/initiator of a given transaction',
COMMENT COLUMN datasize 'The size of the call data of the transaction in bytes',
COMMENT COLUMN data4bytes 'The first 4 bytes of the call data of the transaction',
COMMENT COLUMN network 'The specific Ethereum network used',
COMMENT COLUMN type '"Post EIP-1559, this indicates how the gas parameters are submitted to the network: - type 0 - legacy - type 1 - usage of access lists according to EIP-2930 - type 2 - using maxpriorityfeepergas and maxfeepergas"',
COMMENT COLUMN maxpriorityfeepergas 'The maximum value for a tip offered to the miner/validator per unit of gas. The actual tip paid can be lower if (maxfee - basefee) < maxpriorityfee. Denominated in wei',
COMMENT COLUMN maxfeepergas 'The maximum value for the transaction fee (including basefee and tip) offered to the miner/validator per unit of gas. Denominated in wei',
COMMENT COLUMN basefeepergas 'The fee per unit of gas paid and burned for the curblocknumber. This fee is algorithmically determined. Denominated in wei',
COMMENT COLUMN dropreason 'If the transaction was dropped from the mempool, this describes the contextual reason for the drop',
COMMENT COLUMN rejectionreason 'If the transaction was rejected from the mempool, this describes the contextual reason for the rejection',
COMMENT COLUMN stuck 'A transaction was detected in the queued area of the mempool and is not eligible for inclusion in a block',
COMMENT COLUMN gasused 'If the transaction was published on-chain, this value indicates the amount of gas that was actually consumed. Denominated in wei';

ALTER TABLE block_native_mempool_transaction on cluster '{cluster}'
COMMENT COLUMN unique_key 'Unique key for the row',
COMMENT COLUMN updated_date_time 'When this row was last updated',
COMMENT COLUMN detecttime 'Timestamp that the transaction was detected in mempool',
COMMENT COLUMN hash 'Unique identifier hash for a given transaction',
COMMENT COLUMN status 'Status of the transaction',
COMMENT COLUMN region 'The geographic region for the node that detected the transaction',
COMMENT COLUMN reorg 'If there was a reorg, refers to the blockhash of the reorg',
COMMENT COLUMN replace 'If the transaction was replaced (speedup/cancel), the transaction hash of the replacement',
COMMENT COLUMN curblocknumber 'The block number the event was detected in',
COMMENT COLUMN failurereason 'If a transaction failed, this field provides contextual information',
COMMENT COLUMN blockspending 'If a transaction was finalized (confirmed, failed), this refers to the number of blocks that the transaction was waiting to get on-chain',
COMMENT COLUMN timepending 'If a transaction was finalized (confirmed, failed), this refers to the time in milliseconds that the transaction was waiting to get on-chain',
COMMENT COLUMN nonce 'A unique number which counts the number of transactions sent from a given address',
COMMENT COLUMN gas 'The maximum number of gas units allowed for the transaction',
COMMENT COLUMN gasprice 'The price offered to the miner/validator per unit of gas. Denominated in wei',
COMMENT COLUMN value 'The amount of ETH transferred or sent to contract. Denominated in wei',
COMMENT COLUMN toaddress 'The destination of a given transaction',
COMMENT COLUMN fromaddress 'The source/initiator of a given transaction',
COMMENT COLUMN datasize 'The size of the call data of the transaction in bytes',
COMMENT COLUMN data4bytes 'The first 4 bytes of the call data of the transaction',
COMMENT COLUMN network 'The specific Ethereum network used',
COMMENT COLUMN type '"Post EIP-1559, this indicates how the gas parameters are submitted to the network: - type 0 - legacy - type 1 - usage of access lists according to EIP-2930 - type 2 - using maxpriorityfeepergas and maxfeepergas"',
COMMENT COLUMN maxpriorityfeepergas 'The maximum value for a tip offered to the miner/validator per unit of gas. The actual tip paid can be lower if (maxfee - basefee) < maxpriorityfee. Denominated in wei',
COMMENT COLUMN maxfeepergas 'The maximum value for the transaction fee (including basefee and tip) offered to the miner/validator per unit of gas. Denominated in wei',
COMMENT COLUMN basefeepergas 'The fee per unit of gas paid and burned for the curblocknumber. This fee is algorithmically determined. Denominated in wei',
COMMENT COLUMN dropreason 'If the transaction was dropped from the mempool, this describes the contextual reason for the drop',
COMMENT COLUMN rejectionreason 'If the transaction was rejected from the mempool, this describes the contextual reason for the rejection',
COMMENT COLUMN stuck 'A transaction was detected in the queued area of the mempool and is not eligible for inclusion in a block',
COMMENT COLUMN gasused 'If the transaction was published on-chain, this value indicates the amount of gas that was actually consumed. Denominated in wei';

ALTER TABLE canonical_beacon_proposer_duty_local on cluster '{cluster}'
COMMENT COLUMN unique_key 'Unique key for the row',
COMMENT COLUMN updated_date_time 'When this row was last updated',
COMMENT COLUMN event_date_time 'When the client fetched the proposer duty information from a beacon node',
COMMENT COLUMN slot 'The slot number for which the proposer duty is assigned',
COMMENT COLUMN slot_start_date_time 'The wall clock time when the slot started',
COMMENT COLUMN epoch 'The epoch number containing the slot',
COMMENT COLUMN epoch_start_date_time 'The wall clock time when the epoch started',
COMMENT COLUMN proposer_validator_index 'The validator index of the proposer for the slot',
COMMENT COLUMN proposer_pubkey 'The public key of the validator proposer',
COMMENT COLUMN meta_client_name 'Name of the client that generated the event',
COMMENT COLUMN meta_client_id 'Unique Session ID of the client that generated the event. This changes every time the client is restarted.',
COMMENT COLUMN meta_client_version 'Version of the client that generated the event',
COMMENT COLUMN meta_client_implementation 'Implementation of the client that generated the event',
COMMENT COLUMN meta_client_os 'Operating system of the client that generated the event',
COMMENT COLUMN meta_client_ip 'IP address of the client that generated the event',
COMMENT COLUMN meta_client_geo_city 'City of the client that generated the event',
COMMENT COLUMN meta_client_geo_country 'Country of the client that generated the event',
COMMENT COLUMN meta_client_geo_country_code 'Country code of the client that generated the event',
COMMENT COLUMN meta_client_geo_continent_code 'Continent code of the client that generated the event',
COMMENT COLUMN meta_client_geo_longitude 'Longitude of the client that generated the event',
COMMENT COLUMN meta_client_geo_latitude 'Latitude of the client that generated the event',
COMMENT COLUMN meta_client_geo_autonomous_system_number 'Autonomous system number of the client that generated the event',
COMMENT COLUMN meta_client_geo_autonomous_system_organization 'Autonomous system organization of the client that generated the event',
COMMENT COLUMN meta_network_id 'Ethereum network ID',
COMMENT COLUMN meta_network_name 'Ethereum network name',
COMMENT COLUMN meta_consensus_version 'Ethereum consensus client version that generated the event',
COMMENT COLUMN meta_consensus_version_major 'Ethereum consensus client major version that generated the event',
COMMENT COLUMN meta_consensus_version_minor 'Ethereum consensus client minor version that generated the event',
COMMENT COLUMN meta_consensus_version_patch 'Ethereum consensus client patch version that generated the event',
COMMENT COLUMN meta_consensus_implementation 'Ethereum consensus client implementation that generated the event',
COMMENT COLUMN meta_labels 'Labels associated with the even';

ALTER TABLE canonical_beacon_proposer_duty on cluster '{cluster}'
COMMENT COLUMN unique_key 'Unique key for the row',
COMMENT COLUMN updated_date_time 'When this row was last updated',
COMMENT COLUMN event_date_time 'When the client fetched the proposer duty information from a beacon node',
COMMENT COLUMN slot 'The slot number for which the proposer duty is assigned',
COMMENT COLUMN slot_start_date_time 'The wall clock time when the slot started',
COMMENT COLUMN epoch 'The epoch number containing the slot',
COMMENT COLUMN epoch_start_date_time 'The wall clock time when the epoch started',
COMMENT COLUMN proposer_validator_index 'The validator index of the proposer for the slot',
COMMENT COLUMN proposer_pubkey 'The public key of the validator proposer',
COMMENT COLUMN meta_client_name 'Name of the client that generated the event',
COMMENT COLUMN meta_client_id 'Unique Session ID of the client that generated the event. This changes every time the client is restarted.',
COMMENT COLUMN meta_client_version 'Version of the client that generated the event',
COMMENT COLUMN meta_client_implementation 'Implementation of the client that generated the event',
COMMENT COLUMN meta_client_os 'Operating system of the client that generated the event',
COMMENT COLUMN meta_client_ip 'IP address of the client that generated the event',
COMMENT COLUMN meta_client_geo_city 'City of the client that generated the event',
COMMENT COLUMN meta_client_geo_country 'Country of the client that generated the event',
COMMENT COLUMN meta_client_geo_country_code 'Country code of the client that generated the event',
COMMENT COLUMN meta_client_geo_continent_code 'Continent code of the client that generated the event',
COMMENT COLUMN meta_client_geo_longitude 'Longitude of the client that generated the event',
COMMENT COLUMN meta_client_geo_latitude 'Latitude of the client that generated the event',
COMMENT COLUMN meta_client_geo_autonomous_system_number 'Autonomous system number of the client that generated the event',
COMMENT COLUMN meta_client_geo_autonomous_system_organization 'Autonomous system organization of the client that generated the event',
COMMENT COLUMN meta_network_id 'Ethereum network ID',
COMMENT COLUMN meta_network_name 'Ethereum network name',
COMMENT COLUMN meta_consensus_version 'Ethereum consensus client version that generated the event',
COMMENT COLUMN meta_consensus_version_major 'Ethereum consensus client major version that generated the event',
COMMENT COLUMN meta_consensus_version_minor 'Ethereum consensus client minor version that generated the event',
COMMENT COLUMN meta_consensus_version_patch 'Ethereum consensus client patch version that generated the event',
COMMENT COLUMN meta_consensus_implementation 'Ethereum consensus client implementation that generated the event',
COMMENT COLUMN meta_labels 'Labels associated with the even';