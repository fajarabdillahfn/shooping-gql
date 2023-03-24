package v1

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
)

var qtyLeft = map[string]int{}

func (u *useCase) Checkout(ctx context.Context, productsBought map[string]int) (*model.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*2000)
	defer cancel()

	var (
		googleHomeSku = "120P90"
		alexaSku      = "A304SD"
		macBookSku    = "43N23P"
		cart          = model.Cart{
			TotalPrice: 0,
		}
	)

	for sku, qty := range productsBought {
		ctxSku := context.WithValue(ctx, "sku", sku)
		product, err := u.ShopRepo.GetBySku(ctxSku)
		if err != nil {
			return nil, err
		}

		if product.Quantity < qty {
			return nil, fmt.Errorf("INVALID QUANTITY: we don't have enough %s product", product.Name)
		}

		checkoutProduct := model.CheckoutProduct{
			Sku:      product.Sku,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: qty,
		}

		switch {
		// PROMOTION 2: Buy 3 Google Homes for the price of 2
		case sku == googleHomeSku && qty >= 3:
			disc := math.Floor(3 / float64(qty))
			checkoutProduct.TotalPrice = product.Price * (float64(qty) - disc)

		// PROMOTION 3: Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers
		case sku == alexaSku && qty > 3:
			checkoutProduct.TotalPrice = (product.Price * float64(qty)) * 0.9

		default:
			checkoutProduct.TotalPrice = product.Price * float64(qty)
		}

		cart.Products = append(cart.Products, &checkoutProduct)
		cart.TotalPrice += checkoutProduct.TotalPrice

		qtyLeft[sku] = product.Quantity - qty
	}

	// PROMOTION 1: Each sale of a MacBook Pro comes with a free Raspberry Pi B
	_, buyMacbookPro := productsBought[macBookSku]
	if buyMacbookPro {
		err := u.calculatePromotion3(ctx, productsBought, &cart)
		if err != nil {
			return nil, err
		}
	}

	for sku, qty := range qtyLeft {
		ctxSku := context.WithValue(ctx, "sku", sku)

		err := u.ShopRepo.UpdateQuantity(ctxSku, uint(qty))
		if err != nil {
			return nil, err
		}
	}

	return &cart, nil
}

// PROMOTION 1: Each sale of a MacBook Pro comes with a free Raspberry Pi B
func (u *useCase) calculatePromotion3(ctx context.Context, productsBought map[string]int, cart *model.Cart) error {
	var (
		rasPiSku   = "234234"
		macBookSku = "43N23P"
	)

	qtyRasPi, buyRasPi := productsBought[rasPiSku]
	qtyMacBook, _ := productsBought[macBookSku]

	var (
		newQtyRasPi int
		newPrice    float64
	)

	ctxSku := context.WithValue(ctx, "sku", rasPiSku)
	rasPi, err := u.ShopRepo.GetBySku(ctxSku)
	if err != nil {
		return err
	}

	if qtyMacBook > qtyRasPi {
		if rasPi.Quantity > qtyMacBook {
			newQtyRasPi = qtyMacBook
		} else {
			newQtyRasPi = rasPi.Quantity
		}
	} else {
		newQtyRasPi = qtyRasPi
	}

	if newQtyRasPi-qtyMacBook < 0 {
		newPrice = 0
	} else {
		newPrice = float64(newQtyRasPi-qtyMacBook) * rasPi.Price
	}

	if buyRasPi {
		for i, productCart := range cart.Products {
			if productCart.Sku == rasPiSku {
				cart.Products[i].Quantity = newQtyRasPi
				cart.Products[i].TotalPrice = newPrice
				break
			}
		}
		minQty := newQtyRasPi - qtyMacBook
		if minQty < 0 {
			minQty = 0
		} else if minQty == 0 {
			minQty = qtyRasPi
		}
		cart.TotalPrice -= float64(minQty) * rasPi.Price
	} else if newQtyRasPi > 0{
		rasPiProduct := model.CheckoutProduct{
			Sku:        rasPi.Sku,
			Name:       rasPi.Name,
			Price:      rasPi.Price,
			Quantity:   newQtyRasPi,
			TotalPrice: newPrice,
		}

		cart.Products = append(cart.Products, &rasPiProduct)
	}

	qtyLeft[rasPiSku] = rasPi.Quantity - newQtyRasPi

	return nil
}
