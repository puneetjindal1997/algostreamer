// Copyright (C) 2022 AlgoNode Org.
//
// algostreamer is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// algostreamer is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with algostreamer.  If not, see <https://www.gnu.org/licenses/>.

package simple

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"log"
	"strconv"

	"github.com/algonode/algostreamer/internal/algod"
	"github.com/algonode/algostreamer/rabbitmq"
	"github.com/algonode/algostreamer/redis"
	"github.com/algorand/go-algorand/protocol"

	"github.com/algorand/go-codec/codec"
)

func handleBlockStdOut(b *algod.BlockWrap, redis redis.Database) error {
	var result map[string]interface{}
	var output []byte
	enc := codec.NewEncoderBytes(&output, protocol.JSONStrictHandle)
	err := enc.Encode(b)
	if err != nil {
		log.Println("testtttttttttttttttttttttttttttt")
		return err
	}
	// if b.Block.Round == blockNumber {
	for _, txn := range b.Block.Payset {
		log.Println("txn.Txn.ApplicationID", txn.Txn.ApplicationID)
		resp, err := redis.Get(strconv.Itoa(int(txn.Txn.ApplicationID)))
		if err != nil {
			// error handling
		}
		if resp.(string) == "true" {
			switch txn.Txn.OnCompletion {
			case 1:
				appId := txn.Txn.ApplicationID
				txnSender := txn.Txn.Sender
				result = map[string]interface{}{
					"application_id": appId,
					"address":        txnSender,
				}
				rabbitmq.Send("add_new_user_with_address", result)
			default:
				log.Println("some default value")
			}
			if len(txn.ApplyData.EvalDelta.Logs) > 0 {
				log.Println("----> FOUNDDD  ", txn.ApplyData.EvalDelta.Logs)
				switch txn.ApplyData.EvalDelta.Logs[0] {
				case "newOrder":
					// order id
					text := txn.ApplyData.EvalDelta.Logs[1]
					encodedText := base64.StdEncoding.EncodeToString([]byte(text))
					orderIdByte, err := base64.StdEncoding.DecodeString(encodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					orderId := binary.BigEndian.Uint64(orderIdByte)
					// price
					priceStr := txn.ApplyData.EvalDelta.Logs[3]
					encodeTextPriceStr := base64.StdEncoding.EncodeToString([]byte(priceStr))
					priceInByte, err := base64.StdEncoding.DecodeString(encodeTextPriceStr)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					price := binary.BigEndian.Uint64(priceInByte)
					// amount
					amountStr := txn.ApplyData.EvalDelta.Logs[4]
					encodeTextAmountStr := base64.StdEncoding.EncodeToString([]byte(amountStr))
					amountInByte, err := base64.StdEncoding.DecodeString(encodeTextAmountStr)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					amount := binary.BigEndian.Uint64(amountInByte)
					// pair
					pairStr := txn.ApplyData.EvalDelta.Logs[2]
					encodeTextPairStr := base64.StdEncoding.EncodeToString([]byte(pairStr))
					pairInByte, err := base64.StdEncoding.DecodeString(encodeTextPairStr)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					pair := string(pairInByte)
					// order_type
					orderTypeStr := txn.ApplyData.EvalDelta.Logs[5]
					orderType := orderTypeStr
					// slot
					slotText := txn.ApplyData.EvalDelta.Logs[7]
					encodedSlotText := base64.StdEncoding.EncodeToString([]byte(slotText))
					slotByte, err := base64.StdEncoding.DecodeString(encodedSlotText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					slot := binary.BigEndian.Uint64(slotByte)
					// base_coin_avaliable
					baseCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[8]
					encodedBaseCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(baseCoinAvaliableText))
					baseCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedBaseCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					baseCoinAvaliable := binary.BigEndian.Uint64(baseCoinAvaliableByte)
					// base_coin_locked
					baseCoinLockedText := txn.ApplyData.EvalDelta.Logs[9]
					encodedBaseCoinLockedText := base64.StdEncoding.EncodeToString([]byte(baseCoinLockedText))
					baseCoinLockedByte, err := base64.StdEncoding.DecodeString(encodedBaseCoinLockedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					baseCoinLocked := binary.BigEndian.Uint64(baseCoinLockedByte)
					// price_coin_locked
					priceCoinLockedText := txn.ApplyData.EvalDelta.Logs[11]
					encodedPriceCoinLockedText := base64.StdEncoding.EncodeToString([]byte(priceCoinLockedText))
					priceCoinLockedByte, err := base64.StdEncoding.DecodeString(encodedPriceCoinLockedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					priceCoinLocked := binary.BigEndian.Uint64(priceCoinLockedByte)
					// price_coin_avaliable
					priceCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[10]
					encodedPriceCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(priceCoinAvaliableText))
					priceCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedPriceCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					priceCoinAvaliable := binary.BigEndian.Uint64(priceCoinAvaliableByte)
					// sender
					txnSender := txn.Txn.Sender
					// application id
					appId := txn.Txn.ApplicationID
					result = map[string]interface{}{
						"application_id":       appId,
						"orderId":              orderId,
						"units":                amount,
						"price":                price,
						"pair":                 pair,
						"orderType":            orderType,
						"slot":                 slot,
						"base_coin_available":  baseCoinAvaliable,
						"base_coin_locked":     baseCoinLocked,
						"price_coin_locked":    priceCoinLocked,
						"price_coin_available": priceCoinAvaliable,
						"address":              txnSender,
					}
					rabbitmq.Send("new_order_queue_manage", result)
				case "cancelOrder":
					// order_id
					text := txn.ApplyData.EvalDelta.Logs[1]
					encodedText := base64.StdEncoding.EncodeToString([]byte(text))
					orderIdByte, err := base64.StdEncoding.DecodeString(encodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					orderId := binary.BigEndian.Uint64(orderIdByte)
					// sender
					txnSender := txn.Txn.Sender
					// base_coin_avaliable
					baseCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[2]
					encodedBaseCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(baseCoinAvaliableText))
					baseCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedBaseCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					baseCoinAvaliable := binary.BigEndian.Uint64(baseCoinAvaliableByte)
					// base_coin_locked
					baseCoinLockedText := txn.ApplyData.EvalDelta.Logs[3]
					encodedBaseCoinLockedText := base64.StdEncoding.EncodeToString([]byte(baseCoinLockedText))
					baseCoinLockedByte, err := base64.StdEncoding.DecodeString(encodedBaseCoinLockedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					baseCoinLocked := binary.BigEndian.Uint64(baseCoinLockedByte)
					// price_coin_locked
					priceCoinLockedText := txn.ApplyData.EvalDelta.Logs[5]
					encodedPriceCoinLockedText := base64.StdEncoding.EncodeToString([]byte(priceCoinLockedText))
					priceCoinLockedByte, err := base64.StdEncoding.DecodeString(encodedPriceCoinLockedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					priceCoinLocked := binary.BigEndian.Uint64(priceCoinLockedByte)
					// price_coin_avaliable
					priceCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[4]
					encodedPriceCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(priceCoinAvaliableText))
					priceCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedPriceCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					priceCoinAvaliable := binary.BigEndian.Uint64(priceCoinAvaliableByte)
					// application_id
					appId := txn.Txn.ApplicationID
					result = map[string]interface{}{
						"application_id":       appId,
						"orderId":              orderId,
						"base_coin_available":  baseCoinAvaliable,
						"base_coin_locked":     baseCoinLocked,
						"price_coin_locked":    priceCoinLocked,
						"price_coin_available": priceCoinAvaliable,
						"address":              txnSender,
					}
					rabbitmq.Send("Algo_Cancel_Order", result)
				case "match_order":
					// trade_id
					tradeIdText := txn.ApplyData.EvalDelta.Logs[1]
					tradeIdEncodedText := base64.StdEncoding.EncodeToString([]byte(tradeIdText))
					tradeIdByte, err := base64.StdEncoding.DecodeString(tradeIdEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					tradeId := binary.BigEndian.Uint64(tradeIdByte)
					// application id
					appId := txn.Txn.ApplicationID
					// buyer order id
					buyerOrderIdText := txn.ApplyData.EvalDelta.Logs[2]
					buyerOrderIdEncodedText := base64.StdEncoding.EncodeToString([]byte(buyerOrderIdText))
					buyerOrderIdByte, err := base64.StdEncoding.DecodeString(buyerOrderIdEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					buyerOrderId := binary.BigEndian.Uint64(buyerOrderIdByte)
					// seller order id
					sellerOrderIdText := txn.ApplyData.EvalDelta.Logs[3]
					sellerOrderIdEncodedText := base64.StdEncoding.EncodeToString([]byte(sellerOrderIdText))
					sellerOrderIdByte, err := base64.StdEncoding.DecodeString(sellerOrderIdEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					sellerOrderId := binary.BigEndian.Uint64(sellerOrderIdByte)
					// trade price
					tradePriceText := txn.ApplyData.EvalDelta.Logs[4]
					tradePriceEncodedText := base64.StdEncoding.EncodeToString([]byte(tradePriceText))
					tradePriceByte, err := base64.StdEncoding.DecodeString(tradePriceEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					tradePrice := binary.BigEndian.Uint64(tradePriceByte)
					// trade amount
					tradeAmountText := txn.ApplyData.EvalDelta.Logs[5]
					tradeAmountEncodedText := base64.StdEncoding.EncodeToString([]byte(tradeAmountText))
					tradeAmountByte, err := base64.StdEncoding.DecodeString(tradeAmountEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					tradeAmount := binary.BigEndian.Uint64(tradeAmountByte)
					// buyer address
					encodeBuyerAddress := txn.ApplyData.EvalDelta.Logs[6]
					buyerAd := base64.StdEncoding.EncodeToString([]byte(encodeBuyerAddress))
					buyerAddress := buyerAd
					// seller address
					encodingSellerAddress := txn.ApplyData.EvalDelta.Logs[11]
					sellerAd := base64.StdEncoding.EncodeToString([]byte(encodingSellerAddress))
					sellerAddress := sellerAd
					// buyer price coin avaliable
					buyerPriceCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[7]
					buyerPriceCoinAvaliableEncodedText := base64.StdEncoding.EncodeToString([]byte(buyerPriceCoinAvaliableText))
					buyerPriceCoinAvaliableByte, err := base64.StdEncoding.DecodeString(buyerPriceCoinAvaliableEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					buyerPriceCoinAvaliable := binary.BigEndian.Uint64(buyerPriceCoinAvaliableByte)
					// buyer price coin locked
					buyerPriceCoinLockedText := txn.ApplyData.EvalDelta.Logs[8]
					buyerPriceCoinLockedEncodedText := base64.StdEncoding.EncodeToString([]byte(buyerPriceCoinLockedText))
					buyerPriceCoinLockedByte, err := base64.StdEncoding.DecodeString(buyerPriceCoinLockedEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					buyerPriceCoinLocked := binary.BigEndian.Uint64(buyerPriceCoinLockedByte)
					// buyer base coin avaliable
					buyerBaseCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[9]
					buyerBaseCoinAvaliableEncodedText := base64.StdEncoding.EncodeToString([]byte(buyerBaseCoinAvaliableText))
					buyerBaseCoinAvaliableByte, err := base64.StdEncoding.DecodeString(buyerBaseCoinAvaliableEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					buyerBaseCoinAvaliable := binary.BigEndian.Uint64(buyerBaseCoinAvaliableByte)
					// buyer base coin locked
					buyerBaseCoinLockedText := txn.ApplyData.EvalDelta.Logs[10]
					buyerBaseCoinLockedEncodedText := base64.StdEncoding.EncodeToString([]byte(buyerBaseCoinLockedText))
					buyerBaseCoinLockedByte, err := base64.StdEncoding.DecodeString(buyerBaseCoinLockedEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					buyerBaseCoinLocked := binary.BigEndian.Uint64(buyerBaseCoinLockedByte)
					// seller price coin avaliable
					sellerPriceCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[12]
					sellerPriceCoinAvaliableEncodedText := base64.StdEncoding.EncodeToString([]byte(sellerPriceCoinAvaliableText))
					sellerPriceCoinAvaliableByte, err := base64.StdEncoding.DecodeString(sellerPriceCoinAvaliableEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					sellerPriceCoinAvaliable := binary.BigEndian.Uint64(sellerPriceCoinAvaliableByte)
					// seller price coin locked
					sellerPriceCoinLockedText := txn.ApplyData.EvalDelta.Logs[13]
					sellerPriceCoinLockedEncodedText := base64.StdEncoding.EncodeToString([]byte(sellerPriceCoinLockedText))
					sellerPriceCoinLockedByte, err := base64.StdEncoding.DecodeString(sellerPriceCoinLockedEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					sellerPriceCoinLocked := binary.BigEndian.Uint64(sellerPriceCoinLockedByte)
					// seller base coin avaliable
					sellerBaseCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[14]
					sellerBaseCoinAvaliableEncodedText := base64.StdEncoding.EncodeToString([]byte(sellerBaseCoinAvaliableText))
					sellerBaseCoinAvaliableByte, err := base64.StdEncoding.DecodeString(sellerBaseCoinAvaliableEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					sellerBaseCoinAvaliable := binary.BigEndian.Uint64(sellerBaseCoinAvaliableByte)
					// seller base coin locked
					sellerBaseCoinLockedText := txn.ApplyData.EvalDelta.Logs[15]
					sellerBaseCoinLockedEncodedText := base64.StdEncoding.EncodeToString([]byte(sellerBaseCoinLockedText))
					sellerBaseCoinLockedByte, err := base64.StdEncoding.DecodeString(sellerBaseCoinLockedEncodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					sellerBaseCoinLocked := binary.BigEndian.Uint64(sellerBaseCoinLockedByte)
					// result
					result = map[string]interface{}{
						"trade_id":               tradeId,
						"app_id":                 appId,
						"b_order_id":             buyerOrderId,
						"s_order_id":             sellerOrderId,
						"trade_price":            tradePrice,
						"trade_amount":           tradeAmount,
						"buyer_address":          buyerAddress,
						"seller_address":         sellerAddress,
						"b_price_coin_available": buyerPriceCoinAvaliable,
						"b_price_coin_locked":    buyerPriceCoinLocked,
						"b_base_coin_available":  buyerBaseCoinAvaliable,
						"b_base_coin_locked":     buyerBaseCoinLocked,
						"s_price_coin_available": sellerPriceCoinAvaliable,
						"s_price_coin_locked":    sellerPriceCoinLocked,
						"s_base_coin_available":  sellerBaseCoinAvaliable,
						"s_base_coin_locked":     sellerBaseCoinLocked,
					}
					rabbitmq.Send("trade_settle", result)
				case "self_trade":
					// application_id
					appId := txn.Txn.ApplicationID
					// order_id
					text := txn.ApplyData.EvalDelta.Logs[1]
					encodedText := base64.StdEncoding.EncodeToString([]byte(text))
					orderIdByte, err := base64.StdEncoding.DecodeString(encodedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					orderId := binary.BigEndian.Uint64(orderIdByte)
					// self address
					encodeSelfAddress := txn.ApplyData.EvalDelta.Logs[2]
					sEnc := base64.StdEncoding.EncodeToString([]byte(encodeSelfAddress))
					selfAddress := sEnc
					// base_coin_avaliable
					baseCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[3]
					encodedBaseCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(baseCoinAvaliableText))
					baseCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedBaseCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					baseCoinAvaliable := binary.BigEndian.Uint64(baseCoinAvaliableByte)
					// base_coin_locked
					baseCoinLockedText := txn.ApplyData.EvalDelta.Logs[4]
					encodedBaseCoinLockedText := base64.StdEncoding.EncodeToString([]byte(baseCoinLockedText))
					baseCoinLockedByte, err := base64.StdEncoding.DecodeString(encodedBaseCoinLockedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					baseCoinLocked := binary.BigEndian.Uint64(baseCoinLockedByte)
					// price_coin_locked
					priceCoinLockedText := txn.ApplyData.EvalDelta.Logs[6]
					encodedPriceCoinLockedText := base64.StdEncoding.EncodeToString([]byte(priceCoinLockedText))
					priceCoinLockedByte, err := base64.StdEncoding.DecodeString(encodedPriceCoinLockedText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					priceCoinLocked := binary.BigEndian.Uint64(priceCoinLockedByte)
					// price_coin_avaliable
					priceCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[5]
					encodedPriceCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(priceCoinAvaliableText))
					priceCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedPriceCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					priceCoinAvaliable := binary.BigEndian.Uint64(priceCoinAvaliableByte)
					result = map[string]interface{}{
						"app_id":               appId,
						"address":              selfAddress,
						"order_id":             orderId,
						"base_coin_available":  baseCoinAvaliable,
						"base_coin_locked":     baseCoinLocked,
						"price_coin_locked":    priceCoinLocked,
						"price_coin_available": priceCoinAvaliable,
					}
					rabbitmq.Send("process_self_trade", result)
				case "withdraw":
					// pair
					pairStr := txn.ApplyData.EvalDelta.Logs[1]
					encodeTextPairStr := base64.StdEncoding.EncodeToString([]byte(pairStr))
					pairInByte, err := base64.StdEncoding.DecodeString(encodeTextPairStr)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					pair := string(pairInByte)
					// asset_id
					assetIdText := txn.ApplyData.EvalDelta.Logs[2]
					encodedAssetIdText := base64.StdEncoding.EncodeToString([]byte(assetIdText))
					assetIdByte, err := base64.StdEncoding.DecodeString(encodedAssetIdText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					assetId := binary.BigEndian.Uint64(assetIdByte)
					// price_coin_avaliable
					priceCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[4]
					encodedPriceCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(priceCoinAvaliableText))
					priceCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedPriceCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					priceCoinAvaliable := binary.BigEndian.Uint64(priceCoinAvaliableByte)
					// base_coin_avaliable
					baseCoinAvaliableText := txn.ApplyData.EvalDelta.Logs[3]
					encodedBaseCoinAvaliableText := base64.StdEncoding.EncodeToString([]byte(baseCoinAvaliableText))
					baseCoinAvaliableByte, err := base64.StdEncoding.DecodeString(encodedBaseCoinAvaliableText)
					if err != nil {
						log.Println("errrrrrrr", err)
					}
					baseCoinAvaliable := binary.BigEndian.Uint64(baseCoinAvaliableByte)
					// application id
					appId := txn.Txn.ApplicationID
					// sender
					txnSender := txn.Txn.Sender
					result = map[string]interface{}{
						"app_id":               appId,
						"asset_id":             assetId,
						"pair":                 pair,
						"base_coin_available":  baseCoinAvaliable,
						"price_coin_available": priceCoinAvaliable,
						"address":              txnSender,
					}
					rabbitmq.Send("Algo_Withdraw", result)
				case "closeout":
					// application id
					appId := txn.Txn.ApplicationID
					// sender
					txnSender := txn.Txn.Sender
					result = map[string]interface{}{
						"app_id":  appId,
						"address": txnSender,
					}
					rabbitmq.Send("process_closeout", result)
				default:
					log.Println("No case matched")
				}
			}
		}
	}
	return nil
}

func SimplePusher(ctx context.Context, blocks chan *algod.BlockWrap, status chan *algod.Status, redis redis.Database) error {

	go func() {
		for {
			select {
			case <-status:
				//noop
			case b := <-blocks:
				handleBlockStdOut(b, redis)
			case <-ctx.Done():
			}

		}
	}()
	return nil
}
