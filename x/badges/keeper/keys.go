package keeper

import (
	"strconv"
	"strings"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	CollectionKey       = []byte{0x01}
	UserBalanceKey      = []byte{0x02}
	NextCollectionIdKey = []byte{0x03}
	TransferManagerKey  = []byte{0x04}
	ClaimKey            = []byte{0x05}
	NextClaimIdKey      = []byte{0x06}
	UsedClaimDataKey    = []byte{0x07}
	UsedClaimChallengeKey = []byte{0x08}
	UsedClaimAddressKey = []byte{0x09}
	WhitelistIndexKey   = []byte{0x0A}

	Delimiter   = []byte{0xDD}
	Placeholder = []byte{0xFF}

	IDLength = 8

	BalanceKeyDelimiter = "-"
)

// StoreKey is the store key string for nft
const StoreKey = types.ModuleName

type BalanceKeyDetails struct {
	collectionId sdk.Uint
	address   	 string
}

type ClaimKeyDetails struct {
	collectionId sdk.Uint
	claimId   	 sdk.Uint
}
// Helper functions to manipulate the balance keys. These aren't prefixed. They will be after they are passed into the functions further down in this file.

// Creates the balance key from an address and collectionId. Note this is not prefixed yet. It is just performing a delimited string concatenation.
func ConstructBalanceKey(address string, id sdk.Uint) string {
	collection_id_str := id.String()
	address_str := address
	return collection_id_str + BalanceKeyDelimiter + address_str
}

func ConstructClaimKey(collectionId sdk.Uint, claimId sdk.Uint) string {
	collection_id_str := collectionId.String()
	claim_id_str := claimId.String()
	return collection_id_str + BalanceKeyDelimiter + claim_id_str
}

// Creates the used claim data key from an id and data. Note this is not prefixed yet. It is just performing a delimited string concatenation.
func ConstructUsedClaimDataKey(collectionId sdk.Uint, claimId sdk.Uint) string {
	collection_id_str := collectionId.String()
	claim_id_str := claimId.String()
	return collection_id_str + BalanceKeyDelimiter + claim_id_str
}

func ConstructUsedClaimChallengeKey(collectionId sdk.Uint, claimId sdk.Uint, challengeId sdk.Uint, codeLeafIndex sdk.Uint) string {
	collection_id_str := collectionId.String()
	claim_id_str := claimId.String()
	code_leaf_index_str := codeLeafIndex.String()
	challenge_id_str := challengeId.String()
	return collection_id_str + BalanceKeyDelimiter + claim_id_str + BalanceKeyDelimiter + challenge_id_str + BalanceKeyDelimiter + code_leaf_index_str
}

func ConstructUsedWhitelistIndexKey(collectionId sdk.Uint, claimId sdk.Uint, whitelistLeafIndex sdk.Uint) string {
	collection_id_str := collectionId.String()
	claim_id_str := claimId.String()
	whitelist_leaf_index_str := whitelistLeafIndex.String()
	return collection_id_str + BalanceKeyDelimiter + claim_id_str + BalanceKeyDelimiter + whitelist_leaf_index_str
}

func ConstructUsedClaimAddressKey(collectionId sdk.Uint, claimId sdk.Uint, address string) string {
	collection_id_str := collectionId.String()
	claim_id_str := claimId.String()
	return collection_id_str + BalanceKeyDelimiter + claim_id_str + BalanceKeyDelimiter + address
}

// Creates the transfer manager request key from an address and collectionId. Note this is not prefixed yet. It is just performing a delimited string concatenation.
func ConstructTransferManagerRequestKey(collectionId sdk.Uint, address string) string {
	collection_id_str := collectionId.String()
	address_str := address
	return collection_id_str + BalanceKeyDelimiter + address_str + BalanceKeyDelimiter
}

// Helper function to unparse a balance key and get the information from it.
func GetDetailsFromBalanceKey(id string) BalanceKeyDetails {
	result := strings.Split(id, BalanceKeyDelimiter)
	address := result[1]
	collection_id, _ := strconv.ParseUint(result[0], 10, 64)


	return BalanceKeyDetails{
		address:   address,
		collectionId: sdk.NewUint(collection_id),
	}
}

func GetDetailsFromClaimKey(id string) ClaimKeyDetails {
	result := strings.Split(id, BalanceKeyDelimiter)
	collection_id, _ := strconv.ParseUint(result[0], 10, 64)
	claim_id, _ := strconv.ParseUint(result[1], 10, 64)
	

	return ClaimKeyDetails{
		claimId:  sdk.NewUint(claim_id),
		collectionId: sdk.NewUint(collection_id),
	}
}

// Prefixer functions

// collectionStoreKey returns the byte representation of the collection key ([]byte{0x01} + collectionId)
func collectionStoreKey(collectionId sdk.Uint) []byte {
	key := make([]byte, len(CollectionKey)+IDLength)
	copy(key, CollectionKey)
	copy(key[len(CollectionKey):], []byte(collectionId.String()))
	return key
}

// userBalanceStoreKey returns the byte representation of the collection balance store key ([]byte{0x02} + balanceKey)
func userBalanceStoreKey(balanceKey string) []byte {
	key := make([]byte, len(UserBalanceKey)+len(balanceKey))
	copy(key, UserBalanceKey)
	copy(key[len(UserBalanceKey):], []byte(balanceKey))
	return key
}

// usedClaimDataStoreKey returns the byte representation of the used claim data store key ([]byte{0x07} + key)
func usedClaimDataStoreKey(usedClaimDataKey string) []byte {
	key := make([]byte, len(UsedClaimDataKey)+len(usedClaimDataKey))
	copy(key, UsedClaimDataKey)
	copy(key[len(UsedClaimDataKey):], []byte(usedClaimDataKey))
	return key
}

func usedClaimChallengeStoreKey(usedClaimChallengeKey string) []byte {
	key := make([]byte, len(UsedClaimChallengeKey)+len(usedClaimChallengeKey))
	copy(key, UsedClaimChallengeKey)
	copy(key[len(UsedClaimChallengeKey):], []byte(usedClaimChallengeKey))
	return key
}

func usedWhitelistIndexStoreKey(whitelistIndexKey string) []byte {
	key := make([]byte, len(WhitelistIndexKey)+len(whitelistIndexKey))
	copy(key, WhitelistIndexKey)
	copy(key[len(WhitelistIndexKey):], []byte(whitelistIndexKey))
	return key
}

func usedClaimAddressStoreKey(usedClaimAddressKey string) []byte {
	key := make([]byte, len(UsedClaimAddressKey)+len(usedClaimAddressKey))
	copy(key, UsedClaimAddressKey)
	copy(key[len(UsedClaimAddressKey):], []byte(usedClaimAddressKey))
	return key
}

// claimStoreKey returns the byte representation of the claim store key ([]byte{0x05} + claimId)
func claimStoreKey(claimKey string) []byte {
	key := make([]byte, len(ClaimKey)+len(claimKey))
	copy(key, ClaimKey)
	copy(key[len(ClaimKey):], []byte(claimKey))
	return key
}

// managerTransferRequestKey returns the byte representation of the manager transfer store key ([]byte{0x04} + id)
func managerTransferRequestKey(id string) []byte {
	key := make([]byte, len(TransferManagerKey)+len(id))
	copy(key, TransferManagerKey)
	copy(key[len(TransferManagerKey):], []byte(id))
	return key
}

// nextCollectionIdKey returns the byte representation of the next asset id key ([]byte{0x03})
func nextCollectionIdKey() []byte {
	return NextCollectionIdKey
}

// nextClaimIdKey returns the byte representation of the next claim id key ([]byte{0x06})
func nextClaimIdKey() []byte {
	return NextClaimIdKey
}
