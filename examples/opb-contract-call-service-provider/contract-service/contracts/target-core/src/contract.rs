use cosmwasm_std::{
    to_binary, Deps,Binary, HumanAddr, DepsMut, Env, HandleResponse, InitResponse, MessageInfo,CosmosMsg, WasmMsg, StdResult
};

use crate::error::ContractError;
use crate::msg::{HandleMsg, InitMsg,QueryMsg};
use crate::state::REQUESTS;

// Note, you can use StdResult in some functions where you do not
// make use of the custom errors
pub fn init(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: InitMsg,
) -> Result<InitResponse, ContractError> {
    Ok(InitResponse::default())
}

// And declare a custom Error variant for the ones where you will want to make use of it
pub fn handle(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: HandleMsg,
) -> Result<HandleResponse, ContractError> {
    match msg {
        HandleMsg::CallService{request_id, endpoint_address, call_data} 
        => call_service(deps, request_id, endpoint_address, call_data),
    }
}

pub fn call_service(deps: DepsMut, request_id:String, endpoint_address: HumanAddr, call_data: Binary) -> Result<HandleResponse, ContractError> {
    let executed = REQUESTS.load(deps.storage, &request_id);
    if executed.is_ok() {
        Err(ContractError::Unauthorized{})
    }else{
        let msg = to_binary(&call_data)?;
        let msgs = vec![CosmosMsg::Wasm(WasmMsg::Execute {
            contract_addr: endpoint_address,
            msg: msg,
            send: vec![],
        })];

        REQUESTS.save(deps.storage, &request_id, &true)?;

        Ok(HandleResponse {
            messages: msgs,
            data: None,
            attributes: vec![],
        })
    }
}

pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    to_binary("no query function")
}