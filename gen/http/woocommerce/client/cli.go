// Code generated by goa v3.6.2, DO NOT EDIT.
//
// woocommerce HTTP client CLI support package
//
// Command:
// $ goa gen goa/design -o ./

package client

import (
	"encoding/json"
	"fmt"
	woocommerce "goa/gen/woocommerce"
	"strconv"
)

// BuildReturnWoocommercePayload builds the payload for the woocommerce
// return_woocommerce endpoint from CLI flags.
func BuildReturnWoocommercePayload(woocommerceReturnWoocommerceBody string) (*woocommerce.WoocommerceReturnArgs, error) {
	var err error
	var body ReturnWoocommerceRequestBody
	{
		err = json.Unmarshal([]byte(woocommerceReturnWoocommerceBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"success\": \"0\",\n      \"user_id\": \"1\"\n   }'")
		}
	}
	v := &woocommerce.WoocommerceReturnArgs{
		UserID:  body.UserID,
		Success: body.Success,
	}

	return v, nil
}

// BuildCallbackWoocommercePayload builds the payload for the woocommerce
// callback_woocommerce endpoint from CLI flags.
func BuildCallbackWoocommercePayload(woocommerceCallbackWoocommerceBody string) (*woocommerce.WoocommerceCallbackArgs, error) {
	var err error
	var body CallbackWoocommerceRequestBody
	{
		err = json.Unmarshal([]byte(woocommerceCallbackWoocommerceBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"consumer_key\": \"Officiis quaerat aperiam aut ipsa esse culpa.\",\n      \"consumer_secret\": \"Aliquam doloribus quo consequatur.\",\n      \"key_id\": 8991172639607242822,\n      \"key_permissions\": \"Nisi deleniti placeat ut exercitationem est mollitia.\",\n      \"user_id\": \"Aperiam itaque impedit praesentium velit harum dolor.\"\n   }'")
		}
	}
	v := &woocommerce.WoocommerceCallbackArgs{
		KeyID:          body.KeyID,
		UserID:         body.UserID,
		ConsumerKey:    body.ConsumerKey,
		ConsumerSecret: body.ConsumerSecret,
		KeyPermissions: body.KeyPermissions,
	}

	return v, nil
}

// BuildRetrieveOrdersPayload builds the payload for the woocommerce
// retrieve_orders endpoint from CLI flags.
func BuildRetrieveOrdersPayload(woocommerceRetrieveOrdersStoreID string, woocommerceRetrieveOrdersPlatformRefIds string, woocommerceRetrieveOrdersAuthorization string, woocommerceRetrieveOrdersToken string) (*woocommerce.GetWoocommerce, error) {
	var err error
	var storeID *int32
	{
		if woocommerceRetrieveOrdersStoreID != "" {
			var v int64
			v, err = strconv.ParseInt(woocommerceRetrieveOrdersStoreID, 10, 32)
			val := int32(v)
			storeID = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for storeID, must be INT32")
			}
		}
	}
	var platformRefIds []string
	{
		if woocommerceRetrieveOrdersPlatformRefIds != "" {
			err = json.Unmarshal([]byte(woocommerceRetrieveOrdersPlatformRefIds), &platformRefIds)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for platformRefIds, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"Tempore optio.\",\n      \"Non est illo.\"\n   ]'")
			}
		}
	}
	var authorization *string
	{
		if woocommerceRetrieveOrdersAuthorization != "" {
			authorization = &woocommerceRetrieveOrdersAuthorization
		}
	}
	var token *string
	{
		if woocommerceRetrieveOrdersToken != "" {
			token = &woocommerceRetrieveOrdersToken
		}
	}
	v := &woocommerce.GetWoocommerce{}
	v.StoreID = storeID
	v.PlatformRefIds = platformRefIds
	v.Authorization = authorization
	v.Token = token

	return v, nil
}

// BuildRetrieveProductsPayload builds the payload for the woocommerce
// retrieve_products endpoint from CLI flags.
func BuildRetrieveProductsPayload(woocommerceRetrieveProductsStoreID string, woocommerceRetrieveProductsPlatformRefIds string, woocommerceRetrieveProductsAuthorization string, woocommerceRetrieveProductsToken string) (*woocommerce.GetWoocommerce, error) {
	var err error
	var storeID *int32
	{
		if woocommerceRetrieveProductsStoreID != "" {
			var v int64
			v, err = strconv.ParseInt(woocommerceRetrieveProductsStoreID, 10, 32)
			val := int32(v)
			storeID = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for storeID, must be INT32")
			}
		}
	}
	var platformRefIds []string
	{
		if woocommerceRetrieveProductsPlatformRefIds != "" {
			err = json.Unmarshal([]byte(woocommerceRetrieveProductsPlatformRefIds), &platformRefIds)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for platformRefIds, \nerror: %s, \nexample of valid JSON:\n%s", err, "'[\n      \"Perspiciatis ut dolore corrupti occaecati accusamus.\",\n      \"Quae eos velit.\",\n      \"Iure voluptate expedita non consectetur.\",\n      \"Enim molestiae voluptas ipsa fuga.\"\n   ]'")
			}
		}
	}
	var authorization *string
	{
		if woocommerceRetrieveProductsAuthorization != "" {
			authorization = &woocommerceRetrieveProductsAuthorization
		}
	}
	var token *string
	{
		if woocommerceRetrieveProductsToken != "" {
			token = &woocommerceRetrieveProductsToken
		}
	}
	v := &woocommerce.GetWoocommerce{}
	v.StoreID = storeID
	v.PlatformRefIds = platformRefIds
	v.Authorization = authorization
	v.Token = token

	return v, nil
}