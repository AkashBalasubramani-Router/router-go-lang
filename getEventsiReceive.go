// package main

// import (
// 	"context"
// 	"fmt"
// 	"math/big"
// 	"strconv"

// 	"github.com/NethermindEth/juno/core/felt"
// 	"github.com/NethermindEth/starknet.go/rpc"
// 	"github.com/NethermindEth/starknet.go/utils"
// 	ethrpc "github.com/ethereum/go-ethereum/rpc"
// )

// func QueryIsendEvents(startblock uint64, endblock uint64) {
// 	base := "https://starknet-goerli.g.alchemy.com/v2/bbDTAS9Qs68Po3ssBBPVm378Pye6oS2A"
// 	fmt.Println("Starting simpeCall example")
// 	c, err := ethrpc.DialContext(context.Background(), base)
// 	if err != nil {
// 		fmt.Println("Failed to connect to the client, did you specify the url in the .env.mainnet?")
// 		panic(err)
// 	}
// 	clientv02 := rpc.NewProvider(c)
// 	fmt.Println("Established connection with the client")

// 	// startblock_felt := utils.Uint64ToFelt(startblock)
// 	// endblock_felt := utils.Uint64ToFelt(endblock)

// 	// blockhash1, _ := utils.HexToFelt("0x3f56b1eda98fe5f1baf347a4dd62d9870133a9b417812f104724bcda3decfd9")
// 	// blockhash2, _ := utils.HexToFelt("0x1f91401f309e87f5e8113a5d022ce84e45ddcc434f6c7bd21e66d6715b7204c")
// 	// blockhash3, _ := utils.HexToFelt("0x2b388197c31599e3a774c6d8a62a498884b5ec211b10827bfd484bd27aa2a40")
// 	contractAddress, _ := utils.HexToFelt("0x029d74077978f8fd0f680fb17fc439dc38b54b750ad45ca78bb9823e2d5c1c67")
// 	// iSendEvent, _ := utils.HexToFelt("0x20ff74db85f0d9700fcdb580fe4f2dfefd06b4ecde7b8a1753c0f142bda6402")
// 	iReceiveEvent, _ := utils.HexToFelt("0x3d3a6fa480ecc176ffb470cfbadabc9c568e614a3ff5877bd8c1bcf1ba7c332")
// 	// valsetUpdated, _ := utils.HexToFelt("0x206930eb6d782c3657f345b7cd4fa1c9f14dd743d26cc26f252875bc7b9d275")

// 	eventInput := rpc.EventsInput{
// 		EventFilter: rpc.EventFilter{
// 			FromBlock: rpc.BlockID{
// 				Number: &startblock,
// 			},
// 			ToBlock: rpc.BlockID{
// 				Number: &endblock,
// 			},
// 			Address: contractAddress,
// 			Keys: [][]*felt.Felt{[]*felt.Felt{
// 				iReceiveEvent,
// 			}},
// 		},
// 		ResultPageRequest: rpc.ResultPageRequest{
// 			ChunkSize: 1,
// 		},
// 	}

// 	events, err := clientv02.Events(context.Background(), eventInput)
// 	if err != nil {
// 		fmt.Println("Unsuccessful")
// 		panic(err)
// 	}

// 	for _, emittedEvent := range events.Events {
// 		dataCount := len(emittedEvent.Event.Data)
// 		fmt.Println("DataCount :", dataCount)
// 		var requestIdentifier, eventNonce *big.Int
// 		var execData []byte
// 		var srcChainId, destChainId, relayerRouterAddress, requestSender string
// 		var execStatus bool
// 		// , destChainId,
// 		if dataCount >= 12 {
// 			fmt.Println("Event Data:")
// 			transaction := rpc.InvokeTransactionReceipt(rpc.CommonTransactionReceipt{
// 				TransactionHash: emittedEvent.TransactionHash,
// 				Status:          rpc.TransactionAcceptedOnL2,
// 				Events: []rpc.Event{{
// 					FromAddress: contractAddress,
// 				}},
// 			})
// 			fees := transaction.ActualFee
// 			transactionhash := transaction.TransactionHash
// 			println("Fees --->", fees)
// 			println("Fees --->", transactionhash.String())
// 			println("Fees --->", emittedEvent.TransactionHash.String())
// 			for i := 0; i < 3; i += 2 {
// 				lowPart := new(big.Int)
// 				lowPart.SetString(emittedEvent.Event.Data[i].String(), 0)

// 				highPart := new(big.Int)
// 				highPart.SetString(emittedEvent.Event.Data[i+1].String(), 0)

// 				combinedBigInt := new(big.Int).Mul(highPart, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
// 				combinedBigInt.Add(combinedBigInt, lowPart)

// 				switch i {
// 				case 0:
// 					requestIdentifier = combinedBigInt
// 				case 2:
// 					eventNonce = combinedBigInt
// 				}

// 			}

// 			stringVariable4 := emittedEvent.Event.Data[4].String()
// 			stringVariable5 := emittedEvent.Event.Data[5].String()

// 			stringVariable4Int, err := strconv.Atoi(stringVariable4[2:]) // Removing "0x" prefix
// 			if err != nil {
// 				fmt.Println("Error converting hex for stringVariable7:", err)
// 			}

// 			srcChainId = strconv.Itoa(stringVariable4Int)

// 			stringVariable5Int, err := strconv.Atoi(stringVariable5[2:]) // Removing "0x" prefix
// 			if err != nil {
// 				fmt.Println("Error converting hex for stringVariable7:", err)
// 			}

// 			destChainId = strconv.Itoa(stringVariable5Int)

// 			relayerRouterAddress = emittedEvent.Event.Data[6].String()
// 			requestSender = emittedEvent.Event.Data[7].String()

// 			arraySize1 := new(big.Int)
// 			arraySize1.SetString(emittedEvent.Event.Data[8].String(), 0)

// 			// Construct arrayValue1 based on arraySize1
// 			arrayValue1 := []byte{}
// 			for j := 0; j < int(arraySize1.Int64()); j++ {
// 				value := emittedEvent.Event.Data[9+j].String()
// 				arrayValue1 = append(arrayValue1, value...)
// 			}

// 			execData = arrayValue1

// 			stringVariable10 := emittedEvent.Event.Data[8+int(arraySize1.Int64())+1].String()

// 			intValue, err := strconv.ParseInt(stringVariable10, 0, 0)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				return
// 			}

// 			execStatus = intValue != 0

// 			fmt.Println("Boolean value:", execStatus)

// 			fmt.Println("requestIdentifier:", requestIdentifier)
// 			fmt.Println("eventNonce:", eventNonce)
// 			fmt.Println("srcChainId:", srcChainId)
// 			fmt.Println("destChainId:", destChainId)
// 			fmt.Println("relayerRouterAddress:", relayerRouterAddress)
// 			fmt.Println("requestSender:", requestSender)
// 			fmt.Println("execData:", execData)

// 		}
// 	}
// }

// func main() {
// 	var startblock, endblock uint64
// 	startblock = 851131
// 	endblock = 851136

// 	QueryIsendEvents(startblock, endblock)
// }
