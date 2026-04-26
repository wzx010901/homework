package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

//10734327

func main() {
	//查询指定区块
	//queryBlock()
	//tx()

	interactWithContract()
}

//查找指定区块

func queryBlock() {
	blockNumberFlag := flag.Uint64("number", 0, "block number to query (0 means skip)")
	flag.Parse()

	rpcURL := "https://sepolia.infura.io/v3/7f36dc79c44d46dc8d0ddf9261ac73bf"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	// 最新区块
	latestBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get latest block: %v", err)
	}

	printBlockInfo("Latest Block", latestBlock)

	// 指定区块
	if *blockNumberFlag > 0 {
		num := big.NewInt(0).SetUint64(*blockNumberFlag)
		block, err := fetchBlockWithRetry(ctx, client, num, 3)
		if err != nil {
			log.Fatalf("failed to get block %d: %v", *blockNumberFlag, err)
		}
		printBlockInfo(fmt.Sprintf("Block %d", *blockNumberFlag), block)
	}
}

// fetchBlockWithRetry 带重试机制的区块查询
func fetchBlockWithRetry(ctx context.Context, client *ethclient.Client, blockNumber *big.Int, maxRetries int) (*types.Block, error) {
	var lastErr error
	for i := 0; i < maxRetries; i++ {
		// 每次重试使用新的超时上下文，避免上下文被取消
		reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		block, err := client.BlockByNumber(reqCtx, blockNumber)
		cancel()

		if err == nil {
			return block, nil
		}

		lastErr = err
		if i < maxRetries-1 {
			backoff := time.Duration(i+1) * 500 * time.Millisecond
			log.Printf("[WARN] failed to fetch block %s, retry %d/%d after %v: %v",
				blockNumber.String(), i+1, maxRetries, backoff, err)
			time.Sleep(backoff)
		}
	}
	return nil, fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

// printBlockInfo 打印详细的区块信息
func printBlockInfo(title string, block *types.Block) {
	fmt.Println("======================================")
	fmt.Println(title)
	fmt.Println("======================================")
	fmt.Printf("Block: %+v\n", block)

	// 基本信息
	fmt.Printf("Number       : %d\n", block.Number().Uint64())
	fmt.Printf("Hash         : %s\n", block.Hash().Hex())
	fmt.Printf("Parent Hash  : %s\n", block.ParentHash().Hex())

	// 时间信息
	blockTime := time.Unix(int64(block.Time()), 0)
	fmt.Printf("Timestamp    : %d\n", block.Time())
	fmt.Printf("Time         : %s\n", blockTime.Format(time.RFC3339))
	fmt.Printf("Time (Local) : %s\n", blockTime.Local().Format("2006-01-02 15:04:05 MST"))

	// Gas 信息
	gasUsed := block.GasUsed()
	gasLimit := block.GasLimit()
	gasUsagePercent := float64(gasUsed) / float64(gasLimit) * 100
	fmt.Printf("Gas Used     : %d (%.2f%%)\n", gasUsed, gasUsagePercent)
	fmt.Printf("Gas Limit    : %d\n", gasLimit)

	// 交易信息
	txCount := len(block.Transactions())
	fmt.Printf("Tx Count     : %d\n", txCount)

	// 区块根信息（Merkle 树根）
	fmt.Printf("State Root   : %s\n", block.Root().Hex())
	fmt.Printf("Tx Root      : %s\n", block.TxHash().Hex())
	fmt.Printf("Receipt Root : %s\n", block.ReceiptHash().Hex())

	// 区块大小估算（简化版，实际大小还包括其他字段）
	if txCount > 0 {
		fmt.Printf("\nFirst Tx Hash: %s\n", block.Transactions()[0].Hash().Hex())
		if txCount > 1 {
			fmt.Printf("Last Tx Hash : %s\n", block.Transactions()[txCount-1].Hash().Hex())
		}
	}

	// 难度信息（PoW 相关，PoS 后基本固定）
	fmt.Printf("Difficulty   : %s\n", block.Difficulty().String())

	// 区块奖励相关信息
	coinbase := block.Coinbase()
	if coinbase != (common.Address{}) {
		fmt.Printf("Coinbase     : %s\n", coinbase.Hex())
	}

	fmt.Println("======================================")
	fmt.Println()
}

func tx() {
	toAddrHex := flag.String("to", "0x3afE8f9aC80f08228838Cb2C95e78eD207406E03", "recipient address (required for send mode)")
	amountEth := flag.Float64("amount", 1, "amount in ETH (required for send mode)")
	flag.Parse()

	// 发送交易模式
	if *toAddrHex == "" || *amountEth <= 0 {
		log.Fatal("send mode requires --to and --amount flags")
	}
	sendTransaction(*toAddrHex, *amountEth)

}

// 发送交易
func sendTransaction(toAddrHex string, amountEth float64) {
	rpcURL := "https://sepolia.infura.io/v3/7f36dc79c44d46dc8d0ddf9261ac73bf"

	privKeyHex := ""

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}
	defer client.Close()

	// 解析私钥
	privKey, err := crypto.HexToECDSA(trim0x(privKeyHex))
	if err != nil {
		log.Fatalf("invalid private key: %v", err)
	}

	// 获取发送方地址
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddr := common.HexToAddress(toAddrHex)

	// 获取链 ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}

	// 获取 nonce
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}

	// 获取建议的 Gas 价格（使用 EIP-1559 动态费用）
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalf("failed to get gas tip cap: %v", err)
	}

	// 获取 base fee，计算 fee cap
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get header: %v", err)
	}

	baseFee := header.BaseFee
	if baseFee == nil {
		// 如果不支持 EIP-1559，使用传统 gas price
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			log.Fatalf("failed to get gas price: %v", err)
		}
		baseFee = gasPrice
	}

	// fee cap = base fee * 2 + tip cap（简单策略）
	gasFeeCap := new(big.Int).Add(
		new(big.Int).Mul(baseFee, big.NewInt(2)),
		gasTipCap,
	)

	// 估算 Gas Limit（普通转账固定为 21000）
	gasLimit := uint64(21000)

	// 转换 ETH 金额为 Wei
	// amountEth * 1e18
	amountWei := new(big.Float).Mul(
		big.NewFloat(amountEth),
		big.NewFloat(1e18),
	)
	valueWei, _ := amountWei.Int(nil)

	// 检查余额是否足够
	balance, err := client.BalanceAt(ctx, fromAddr, nil)
	if err != nil {
		log.Fatalf("failed to get balance: %v", err)
	}

	// 计算总费用：value + gasFeeCap * gasLimit
	totalCost := new(big.Int).Add(
		valueWei,
		new(big.Int).Mul(gasFeeCap, big.NewInt(int64(gasLimit))),
	)

	if balance.Cmp(totalCost) < 0 {
		log.Fatalf("insufficient balance: have %s wei, need %s wei", balance.String(), totalCost.String())
	}

	// 构造交易（EIP-1559 动态费用交易）
	txData := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     valueWei,
		Data:      nil,
	}
	tx := types.NewTx(txData)

	// 签名交易
	signer := types.NewLondonSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, privKey)
	if err != nil {
		log.Fatalf("failed to sign transaction: %v", err)
	}

	// 发送交易
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		log.Fatalf("failed to send transaction: %v", err)
	}

	// 输出交易信息
	fmt.Println("=== Transaction Sent ===")
	fmt.Printf("From       : %s\n", fromAddr.Hex())
	fmt.Printf("To         : %s\n", toAddr.Hex())
	fmt.Printf("Value      : %s ETH (%s Wei)\n", fmt.Sprintf("%.6f", amountEth), valueWei.String())
	fmt.Printf("Gas Limit  : %d\n", gasLimit)
	fmt.Printf("Gas Tip Cap: %s Wei\n", gasTipCap.String())
	fmt.Printf("Gas Fee Cap: %s Wei\n", gasFeeCap.String())
	fmt.Printf("Nonce      : %d\n", nonce)
	fmt.Printf("Tx Hash    : %s\n", signedTx.Hash().Hex())
	fmt.Println("\nTransaction is pending. Use --tx flag to query status:")
	fmt.Printf("  go run main.go --tx %s\n", signedTx.Hash().Hex())
}

func printTxBasicInfo(tx *types.Transaction, isPending bool) {
	fmt.Printf("Hash        : %s\n", tx.Hash().Hex())
	fmt.Printf("Nonce       : %d\n", tx.Nonce())
	fmt.Printf("Gas         : %d\n", tx.Gas())
	fmt.Printf("Gas Price   : %s\n", tx.GasPrice().String())
	fmt.Printf("To          : %v\n", tx.To())
	fmt.Printf("Value (Wei) : %s\n", tx.Value().String())
	fmt.Printf("Data Len    : %d bytes\n", len(tx.Data()))
	fmt.Printf("Pending     : %v\n", isPending)
}

func printReceiptInfo(r *types.Receipt) {
	fmt.Printf("Status      : %d\n", r.Status)
	fmt.Printf("BlockNumber : %d\n", r.BlockNumber.Uint64())
	fmt.Printf("BlockHash   : %s\n", r.BlockHash.Hex())
	fmt.Printf("TxIndex     : %d\n", r.TransactionIndex)
	fmt.Printf("Gas Used    : %d\n", r.GasUsed)
	fmt.Printf("Logs        : %d\n", len(r.Logs))
	if len(r.Logs) > 0 {
		fmt.Printf("First Log Address : %s\n", r.Logs[0].Address.Hex())
	}
}

// trim0x 移除十六进制字符串前缀 "0x"
func trim0x(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}

// interactWithContract 与已部署的 SimpleCounter 合约交互
func interactWithContract() {
	// Sepolia 测试网 RPC
	rpcURL := "https://sepolia.infura.io/v3/7f36dc79c44d46dc8d0ddf9261ac73bf"

	// 连接到以太坊节点
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("连接以太坊节点失败: %v", err)
	}
	defer client.Close()

	fmt.Println("✓ 成功连接到 Sepolia 测试网")

	// 已部署的合约地址（需要替换为您实际部署的合约地址）
	contractAddress := common.HexToAddress("0x0c5c010dc3a6e73b13e08963e784d540ec0d6a19")

	// 创建合约实例
	instance, err := NewSimpleCounter(contractAddress, client)
	if err != nil {
		log.Fatalf("创建合约实例失败: %v", err)
	}

	fmt.Println("✓ 合约实例创建成功")
	fmt.Println("======================================")

	// ========== 读取合约状态 ==========
	count, err := instance.GetCount(nil)
	if err != nil {
		log.Fatalf("读取计数失败: %v", err)
	}
	fmt.Printf("当前计数值: %s\n", count.String())

	// ========== 调用合约方法（写入操作）==========
	fmt.Println("\n准备发送交易调用 increment()...")

	// 配置交易参数
	privateKey, err := crypto.HexToECDSA("私钥")
	if err != nil {
		log.Fatalf("私钥解析失败: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥转换失败")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取链 ID（Sepolia = 11155111）
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("获取链 ID 失败: %v", err)
	}

	// 创建授权器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("创建授权器失败: %v", err)
	}

	auth.GasLimit = uint64(300000)
	auth.GasPrice, err = client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("获取 Gas 价格失败: %v", err)
	}

	fmt.Printf("发送方地址: %s\n", fromAddress.Hex())
	fmt.Printf("Gas Limit: %d\n", auth.GasLimit)
	fmt.Printf("Gas Price: %s Wei\n", auth.GasPrice.String())

	// 调用 increment 方法
	tx, err := instance.Increment(auth)
	if err != nil {
		log.Fatalf("调用 increment 失败: %v", err)
	}

	fmt.Printf("\n交易已发送!\n")
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())

	// 等待交易确认
	fmt.Println("\n等待交易确认...")
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("等待交易确认失败: %v", err)
	}

	fmt.Printf("✓ 交易已确认!\n")
	fmt.Printf("区块号: %d\n", receipt.BlockNumber)
	fmt.Printf("Gas 使用: %d\n", receipt.GasUsed)
	fmt.Printf("交易状态: %d (1=成功)\n", receipt.Status)

	// 再次读取计数值，验证更新
	fmt.Println("\n读取更新后的计数值...")
	count, err = instance.GetCount(nil)
	if err != nil {
		log.Fatalf("读取计数失败: %v", err)
	}
	fmt.Printf("当前计数值: %s\n", count.String())
}

//func main() {
//	fmt.Println("======================================")
//	fmt.Println("SimpleCounter 合约交互演示")
//	fmt.Println("======================================\n")
//
//	// 选择执行模式
//	// 1. 部署新合约
//	// deployContract()
//
//	// 2. 与已部署合约交互
//	interactWithContract()
//}
