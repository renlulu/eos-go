package generated

//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treeset" AccountNameSet(common.AccountName,common.CompareName,false)
//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treemap" AccountNameUint32Map(common.AccountName,uint32,common.CompareName,false)
//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treemap" AccountNameUint64Map(common.AccountName,uint64,common.CompareName,false)
//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treeset" AccountDeltaSet(common.AccountDelta,common.CompareAccountDelta,false)
//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treeset" AccountDeltaSet(common.AccountDelta,common.CompareAccountDelta,false)
//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treeset" NamePairSet(common.NamePair,common.CompareNamePair,false)


//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treeset" PermissionLevelSet(common.PermissionLevel,common.ComparePermissionLevel,false)
//go:generate gotemplate -outfmt "gen_%v" "github.com/eosspark/eos-go/common/container/treeset" PublicKeySet(ecc.PublicKey,ecc.ComparePubKey,false)

//go:generate go build
