package relay

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pokt-network/pocket-core/logs"
	"github.com/pokt-network/pocket-core/rpc/shared"
	"github.com/pokt-network/pocket-core/service"
)

// "Relay" handles the localhost:<relay-port>/v1/relaycall.
func Relay(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	relay := &service.Relay{}
	if err := shared.PopModel(w, r, ps, relay); err != nil {
		logs.NewLog(err.Error(), logs.ErrorLevel, logs.JSONLogFormat)
		shared.WriteErrorResponse(w, 400, err.Error())
		return
	}
	if relay.Blockchain == "" || relay.NetworkID == "" || relay.DevID == "" || relay.Version == "" || relay.Data == "" {
		shared.WriteErrorResponse(w, 400, "The request was not properly formatted")
		return
	}
	response, err := service.RouteRelay(*relay)
	if err != nil {
		logs.NewLog(err.Error(), logs.ErrorLevel, logs.JSONLogFormat)
		// error is embedded in the response
	}
	shared.WriteJSONResponse(w, response) // relay the response
}

// "RelayInfo" handles a get request to localhost:<relay-port>/v1/relay call.
// And provides the developers with an in-client reference to the API call
func RelayInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	info := shared.InfoStruct(r, "Relay", service.Relay{}, "Response from hosted chain")
	shared.WriteInfoResponse(w, info)
}
