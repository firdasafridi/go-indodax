// Package indodax provide a library for accesing Indodax API (see
// https://indodax.com/downloads/BITCOINCOID-API-DOCUMENTATION.pdf for HTTP API documentation).
//
// Indodax provide public and private APIs.
// The public APIs can be accessed directly by creating new client with empty
// token and secret parameters.
// The private APIs can only be accessed by using token and secret keys (API
// credential).
//
// An API credential can be retrieved manually by logging in into your
// Indodax Exchange account (https://indodax.com/market) and open the
// "Trade API" menu section or https://indodax.com/trade_api.
//
// Please keep these credentials safe and do not reveal to any external party.
//
// Beside passing the token and secret to NewClient or Authenticate, this
// library also read token and secret values from environment variables
// "INDODAX_KEY" for token and "INDODAX_SECRET" for secret.
//
package indodax
