package raydium

import "github.com/gagliardetto/solana-go"

var RAYDIUM_PROGRAM_ID = solana.MustPublicKeyFromBase58("CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK")
var POOLSTATE_WSOLUSDC_025 = solana.MustPublicKeyFromBase58("3tD34VtprDSkYCnATtQLCiVgTkECU3d12KtjupeR6N2X")
var METAPLEX = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
var COMPUTE_BUDGET = solana.MustPublicKeyFromBase58("ComputeBudget111111111111111111111111111111")

var AMM_CONFIG_SEED = []byte("amm_config")
var POOL_SEED = []byte("pool")
var POOL_VAULT_SEED = []byte("pool_vault")
var POOL_REWARD_VAULT_SEED = []byte("pool_reward_vault")
var POSITION_SEED = []byte("position")
var TICK_ARRAY_SEED = []byte("tick_array")
var OPERATION_SEED = []byte("operation")
