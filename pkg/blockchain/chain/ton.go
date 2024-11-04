package chain

// hd wallet產不成功，先擱置
//import (
//	"crypto/ed25519"
//	"github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/xssnick/tonutils-go/ton/wallet"
//	"mouse/pkg/lib/hdwallet"
//)
//
//type tonChain struct{}
//
//func newTonChain() *tonChain {
//	return &tonChain{}
//}
//
//// 跟btc一樣，有很多版本的地址跟變數，每個錢包使用的版本或變數不同
//var _tonAddrConfigMap = map[string]_tonAddrConfig{
//	"default": {
//		versionCfg: wallet.ConfigV5R1Final{
//			NetworkGlobalID: wallet.MainnetGlobalID,
//		},
//		subWallet: wallet.DefaultSubwallet,
//	},
//	"safepal": {
//		versionCfg: wallet.V3R2,
//		subWallet:  wallet.DefaultSubwallet,
//	},
//	"bitget_v4": {
//		versionCfg: wallet.V4R2,
//		subWallet:  wallet.DefaultSubwallet,
//	},
//	"bitget_v5": {
//		versionCfg: wallet.ConfigV5R1Final{
//			NetworkGlobalID: wallet.MainnetGlobalID,
//		},
//		subWallet: uint32(0),
//	},
//}
//
//type _tonAddrConfig struct {
//	versionCfg wallet.VersionConfig
//	subWallet  uint32
//}
//
//// GenAddr 产生地址
//func (s *tonChain) GenAddr() (string, string, error) {
//	// 生成新的密钥对
//	pubKey, privateKey, err := ed25519.GenerateKey(nil)
//	if err != nil {
//		return "", "", err
//	}
//
//	addrConfig := _tonAddrConfigMap["default"]
//	address, err := wallet.AddressFromPubKey(pubKey, addrConfig.versionCfg, addrConfig.subWallet)
//	if err != nil {
//		return "", "", err
//	}
//
//	// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
//	return address.Bounce(false).String(), hexutil.Encode(privateKey.Seed()), nil
//}
//
//const (
//	_Iterations   = 100000
//	_Salt         = "TON default seed"
//	_BasicSalt    = "TON seed version"
//	_PasswordSalt = "TON fast seed version"
//)
//
//var _tonDerivationPath = hdwallet.MustParseDerivationPath("m/44'/607'/0'/0/0")
//
//// GenHdAddr 产生Hd wallet 地址，返回的key為mnemonic
//func (s *tonChain) GenHdAddr() (string, string, error) {
//	// 生成新的密钥对
//	//privateKey, err := crypto.GenerateKey()
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//mnemonic, err := bip39.NewMnemonic(crypto.FromECDSA(privateKey))
//	//if err != nil {
//	//	return "", "", err
//	//}
//
//	mnemonic := "group danger moment hen erase trash mixture hockey glow lady nothing bring still result today decrease high nerve develop giggle monkey pattern nurse ignore"
//
//	hdwalletD, err := hdwallet.NewFromMnemonic(mnemonic, true)
//	if err != nil {
//		return "", "", err
//	}
//	key := hdwalletD.MasterKey
//	for _, n := range _tonDerivationPath {
//		key, err = key.Derive(n)
//		if err != nil {
//			return "", "", err
//		}
//	}
//
//	privateKeyEc, err := key.ECPrivKey()
//	if err != nil {
//		return "", "", err
//	}
//	sum := privateKeyEc.Serialize()
//
//	//seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//deriveKey, err := signer.DeriveForPath("m/44'/607'/0'/0'/0'", seed)
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//seed = deriveKey.Key
//
//	//hash := hmac.New(sha512.New, nil)
//	//_, err = hash.Write(seed)
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//sum := hash.Sum(nil)
//	//sum := seed
//
//	priKey := ed25519.NewKeyFromSeed(sum[:32])
//	pubKey := priKey.Public().(ed25519.PublicKey)
//
//	addrConfig := _tonAddrConfigMap["bitget_v4"]
//	address, err := wallet.AddressFromPubKey(pubKey, addrConfig.versionCfg, addrConfig.subWallet)
//	if err != nil {
//		return "", "", err
//	}
//
//	// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
//	return address.Bounce(false).String(), mnemonic, nil
//
//	//hash := hmac.New(sha512.New, []byte("ed25519 seed"))
//	//_, err = hash.Write(seed)
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//sum := hash.Sum(nil)
//	//seed = sum[:32]
//
//	//d, err := signer.DeriveForPath("m/44'/607'/0'/0'/0'", seed)
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//seed = d.Key
//
//	//password := ""
//	//mac := hmac.New(sha512.New, []byte(seed))
//	//mac.Write([]byte(password))
//	//hash := mac.Sum(nil)
//	//priKey := ed25519.NewKeyFromSeed(seed[:])
//	//pubKey := priKey.Public().(ed25519.PublicKey)
//
//	//k := pbkdf2.Key(seed, nil, _Iterations, 32, sha512.New)
//	//key := ed25519.NewKeyFromSeed(seed)
//	//addrConfig := _tonAddrConfigMap["bitget_v5"]
//	//address, err := wallet.AddressFromPubKey(key.Public().(ed25519.PublicKey), addrConfig.versionCfg, addrConfig.subWallet)
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//return address.Bounce(false).String(), mnemonic, nil
//
//	//if len(seed) < 12 {
//	//	return nil, fmt.Errorf("seed should have at least 12 words")
//	//}
//	//for _, s := range seed {
//	//	if !words[s] {
//	//		return nil, fmt.Errorf("unknown word '%s' in seed", s)
//	//	}
//	//}
//
//	//seed := mnemonic
//	//password := ""
//	//mac := hmac.New(sha512.New, []byte(seed))
//	//mac.Write([]byte(password))
//	//hash := mac.Sum(nil)
//	//
//	////p := pbkdf2.Key(hash, []byte(_BasicSalt), _Iterations/256, 1, sha512.New)
//	////if p[0] != 0 {
//	////	return "", "", errors.New("invalid seed")
//	////}
//	//
//	//k := pbkdf2.Key(hash, []byte(_Salt), _Iterations, 32, sha512.New)
//	//key := ed25519.NewKeyFromSeed(k)
//	//
//	//addrConfig := _tonAddrConfigMap["bitget_v4"]
//	//address, err := wallet.AddressFromPubKey(key.Public().(ed25519.PublicKey), addrConfig.versionCfg, addrConfig.subWallet)
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//return address.Bounce(false).String(), mnemonic, nil
//
//	//
//	//wallet.FromSeed()
//	//
//	//priKey := ed25519.NewKeyFromSeed(key.Key)
//	//pubKey := priKey.Public().(ed25519.PublicKey)
//	//
//	//addrConfig := _tonAddrConfigMap["bitget_v4"]
//	//address, err := wallet.AddressFromPubKey(pubKey, addrConfig.versionCfg, addrConfig.subWallet)
//	//if err != nil {
//	//	return "", "", err
//	//}
//	//
//	//// Bounce(false) => UQ 開頭，Bounce(true) => EQ 開頭
//	//return address.Bounce(false).String(), mnemonic, nil
//	//
//	//suiSigner, err := signer.NewSignertWithMnemonic(mnemonic)
//	//if err != nil {
//	//	return "", "", eris.Wrap(err, "failed to generate signer")
//	//}
//	//
//	//return suiSigner.Address, mnemonic, nil
//}
